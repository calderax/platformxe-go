# PlatformXe Go SDK

Official Go SDK for the [PlatformXe](https://platformxe.com) API — messaging, storage, OCR, PDF, identity, fraud detection, webhooks, workflows, custom events, and authorization (RBAC + ABAC + ReBAC + Federation).

[![Go Reference](https://pkg.go.dev/badge/github.com/calderax/platformxe-go.svg)](https://pkg.go.dev/github.com/calderax/platformxe-go)
[![Go version](https://img.shields.io/badge/go-1.22%2B-blue)](https://go.dev/)

## Install

```bash
go get github.com/calderax/platformxe-go@v1.6.0
```

The SDK targets Go 1.22+ and has no third-party dependencies — only the standard library.

## Quick start

```go
package main

import (
    "fmt"

    px "github.com/calderax/platformxe-go"
)

func main() {
    client := px.NewClient(px.ClientConfig{
        APIKey: "pxk_live_…",
    })

    // Send a transactional email
    _, err := client.Messaging.SendEmail(px.SendEmailInput{
        To:      []string{"user@example.com"},
        Subject: "Welcome",
        HTML:    "<h1>Hello</h1>",
    })
    if err != nil {
        panic(err)
    }

    // Check a permission
    decision, err := client.Permissions.Check("usr_123", "chat/session", "READ")
    if err != nil {
        panic(err)
    }
    fmt.Println(decision.Allowed)

    // Resolve a Nigerian identity
    profile, err := client.Identity.Resolve("BVN", "22012345678", true, "consent_abc")
    if err != nil {
        panic(err)
    }
    fmt.Println(profile)
}
```

## Configuration

```go
client := px.NewClient(px.ClientConfig{
    APIKey:         "pxk_live_…",             // required
    BaseURL:        "https://platformxe.com", // optional — defaults to production
    Timeout:        10,                       // optional — seconds
    Retries:        2,                        // optional — 0 means the default (2); negative = none
    DisableRetries: false,                    // optional — true = exactly one attempt
    FailOpen:       false,                    // optional — a transport failure returns a _failed result
})
```

`BaseURL` accepts the regional production endpoint or a self-hosted PlatformXe instance.

Retries back off from 200 ms, doubling each time. Transport failures, `429`, `502`, `503` and `504` are retried for every method; `500` only for `GET`, `HEAD`, `PUT`, `DELETE` and `OPTIONS`, because PlatformXe answers a handler failure with `500` after side effects may already have landed. A `POST`/`PATCH` `500` and every other non-2xx return an error after one attempt.

## Surface

| Namespace | Domain |
|-----------|--------|
| `client.Messaging` | Transactional email, SMS, WhatsApp click-to-chat |
| `client.Storage` / `client.Documents` | Media + fixed-storage upload, signed URLs |
| `client.Ocr` / `client.Pdf` / `client.Qr` | OCR, PDF generation, QR encoding |
| `client.Identity` | Identity resolution + Nigerian KYC (BVN, NIN, liveness, face match) |
| `client.Fraud` | Fraud Detection — rules, screens, devices, federation, terms |
| `client.Permissions` | RBAC + ABAC + ReBAC + Federation. `Check`, `CheckBatch`, `Resolve`, roles, modules, overrides, policies, relationships |
| `client.Audit` | Decision + mutation audit log query/export |
| `client.Webhooks` | Outbound webhook endpoints |
| `client.Templates` | Content templates |
| `client.Workflows` | Event-driven automations |
| `client.Domains` | Sending domain management |
| `client.Events` | Built-in event ingestion + log + subscriptions |
| `client.Events.Custom` | Tenant-defined custom events (Phase 9A) |
| `client.Events.Custom.Marketplace` | Cross-tenant marketplace (Phase 9C, PRO+) |
| `client.Events.Custom.Federation` | Cross-org event fan-out (Phase 9D + Pattern 3, ENTERPRISE) |
| `client.Threads` | Caldera Threads contextual messaging |
| `client.Exports` | Async data export jobs |
| `client.Usage` | Usage + billing reads |
| `client.Search` | Federated search query |
| `client.Issues` | Tenant-reported issues / support cases |
| `client.Whoami` | API-key resolution + telemetry helpers |
| `px.Register` | Package-level standalone bootstrap helper |

## Pattern 3 — external webhook peers (1.5.0)

Custom Event Federation supports peers addressed by URL + HMAC secret instead of a tenant org id — useful when the receiving system isn't on PlatformXe.

```go
result, err := client.Events.Custom.Federation.AddExternalPeer(
    groupID,
    px.AddEventFederationExternalPeerInput{
        Label:      "Booking.com",
        WebhookURL: "https://booking.example.com/inbound/platformxe",
        Headers:    map[string]string{"Authorization": "Bearer xyz"},
    },
)
if err != nil {
    panic(err)
}
fmt.Println(result.Secret)               // "whsec_…" — store immediately, shown ONCE
fmt.Println(result.Peer.PeerType)        // px.EventFederationPeerTypeExternalWebhook

// Remove later:
_, _ = client.Events.Custom.Federation.RemoveExternalPeer(result.Peer.ID)
```

The receiving endpoint verifies inbound POSTs with `HMAC-SHA256(rawBody, secret)` matched against `X-Caldera-Signature: sha256=<hex>`. See the [federation reference](https://docs.platformxe.com/sdk/federation) for the full wire format.

## Error handling

Every method returns an `error` on failure. Every non-2xx response — a JSON envelope, other JSON, or a non-JSON body such as a gateway's HTML error page — is an `*APIError` carrying `StatusCode`, plus `Code` and `Message` from the body when it has them (otherwise `HTTP_<status>` and the status text). Setting `FailOpen: true` makes a request whose last attempt failed at the transport level return a `{"_failed": true, "error": …}` result with a `nil` error — useful for non-blocking permission checks. It never covers an HTTP error response.

```go
decision, err := client.Permissions.Check(adminID, path, action)
if err != nil {
    var apiErr *px.APIError
    if errors.As(err, &apiErr) {
        // apiErr.Code, apiErr.Message, apiErr.StatusCode
    }
    return err
}
```

## Versioning

| Package | Version |
|---------|---------|
| `github.com/calderax/platformxe-go` (this) | **v1.6.0** *(pending release — tag `sdk-go/v1.6.0` cut on merge)* |
| `@caldera/platformxe-sdk` (TypeScript) | 1.6.0 *(pending release — SDK MR !5)* |
| `@caldera/platformxe-types` (TypeScript shapes) | 3.2.0 |
| `platformxe` (Python) | 1.6.0 *(pending release — tag `sdk-python/v1.6.0` cut on merge)* |
| `calderax/platformxe` (Terraform) | 1.5.1 |

The PlatformXe ecosystem ships **5 published artefacts** that move in lockstep against every API change (the NO-DRIFT policy). The full coverage matrix lives at [docs.platformxe.com/sdk/alignment](https://docs.platformxe.com/sdk/alignment).

**v1.6.0 behaviour changes** (released with the `sdk-go/v1.6.0` tag, cut on merge):

- Any non-2xx response is an `*APIError`. Before, a non-2xx whose body was not a `{"success": false}` envelope (for example a 5xx `{"error": "..."}`) came back as data with a `nil` error, and a non-JSON body surfaced as an unmarshal error.
- A `POST`/`PATCH` `500` is no longer retried, and neither is a 5xx other than `500`/`502`/`503`/`504`. `429` and `502`/`503`/`504` are retried for every method.
- `FailOpen` covers transport failures only: a retried HTTP error status still failing after the last retry returns the `*APIError`.
- `DisableRetries: true` configures a single attempt (`Retries: 0` means the default of 2).

## Distribution

The Go SDK is hosted at `github.com/calderax/platformxe-go`. The source of truth lives in the PlatformXe monorepo at `packages/sdk-go/` and is mirrored to GitHub on every `sdk-go/v*` tag push so `go get …@v<semver>` resolves directly through the Go module proxy.

```bash
go get github.com/calderax/platformxe-go@v1.6.0
```

```go
import px "github.com/calderax/platformxe-go"
```

## Documentation

- **API reference + per-domain SDK guides:** [docs.platformxe.com](https://docs.platformxe.com)
- **Go reference:** [pkg.go.dev/github.com/calderax/platformxe-go](https://pkg.go.dev/github.com/calderax/platformxe-go)
- **OpenAPI 3.1 spec:** `https://api.platformxe.com/api/docs/openapi.json`
- **Status + incidents:** [status.platformxe.com](https://status.platformxe.com)

## License

Proprietary — © 2026 Caldera Technologies Ltd. All rights reserved.
