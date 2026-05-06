# PlatformXe Go SDK

Go SDK for the [PlatformXe](https://platformxe.com) API.

## Installation

```bash
go get github.com/calderax/platformxe-go
```

## Quick Start

```go
package main

import (
    "fmt"
    px "github.com/calderax/platformxe-go"
)

func main() {
    client := px.NewClient(px.ClientConfig{
        APIKey: "pxk_live_your_key_here",
    })

    // Check permission
    result, err := client.Permissions.Check("usr_123", "chat/session", "READ")
    if err != nil {
        panic(err)
    }
    fmt.Println(result)

    // Resolve identity
    profile, err := client.Identity.Resolve("BVN", "22012345678", true, "consent_ref")
    if err != nil {
        panic(err)
    }
    fmt.Println(profile)
}
```

## Distribution

The Go SDK is published as a Go module at `github.com/calderax/platformxe-go`. The source lives in this monorepo at `packages/sdk-go/` and is mirrored to GitHub for `go get` resolution.

```bash
# Install as a dependency
go get github.com/calderax/platformxe-go

# Or import directly
import px "github.com/calderax/platformxe-go"
```

## Related Packages

| Package | Language | Install |
|---------|----------|---------|
| `@caldera/platformxe-sdk` | TypeScript | `npm install @caldera/platformxe-sdk` |
| `@caldera/platformxe-types` | TypeScript (types only) | `npm install @caldera/platformxe-types` |
| `platformxe` | Python | `pip install platformxe` |
| `calderax/platformxe` | Terraform | `source = "calderax/platformxe"` |

## License

Proprietary — Caldera Technologies Ltd.
