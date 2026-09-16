// =============================================================================
// (c) 2026 Caldera Technologies Ltd.
// Proprietary and confidential.
// Unauthorized copying or distribution is prohibited.
// =============================================================================
//
// CHANGELOG (2026-09-15):
// NewClient treated an explicit Retries: 0 as unset and defaulted it to 2,
// so retries could never be disabled (TestNewClientCustomConfig failed on
// development too). ClientConfig.DisableRetries is now the explicit switch
// and a negative Retries clamps to 0; omitting both still yields 2. Also
// gofmt alignment of the v1.1.0 service fields.
//
// Also 2026-09-15 (MR !199 round 9): doRequest never checked the HTTP status —
// it only failed when the body said "success": false — so a non-2xx whose
// JSON was not an envelope (for example a 5xx {"error":"..."}) was returned as
// data with a nil error. Any non-2xx is now an *APIError with the status and
// the parsed code/message (the status text for a non-JSON body). 5xx and 429
// are retried on the existing transport backoff and, once retries run out,
// return the last *APIError (FailOpen still covers transport failures only);
// other 4xx return immediately. 2xx handling is unchanged, and the
// success:false error keeps its "[code] message (HTTP n)" text.
//
// Also 2026-09-15 (MR !199 round 10, finding 3): retrying every 5xx repeated
// non-idempotent requests. The server answers a handler throw with 500
// precisely because side effects may have landed, and this client sends no
// idempotency key, so a retried POST could send an email or create a record
// twice. 429 and 502/503/504 (the request did not reach a handler, or was
// refused before it ran) are retried for every method; 500 only for
// GET/HEAD/PUT/DELETE/OPTIONS; a POST/PATCH 500 and any other 5xx return the
// *APIError at once. Released as module v1.6.0.
// =============================================================================

package platformxe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client is the PlatformXe API client.
type Client struct {
	apiKey   string
	baseURL  string
	retries  int
	failOpen bool
	http     *http.Client

	Documents   *DocumentsService
	Domains     *DomainsService
	Events      *EventsService
	Exports     *ExportsService
	Fraud       *FraudService
	Identity    *IdentityService
	Messaging   *MessagingService
	Ocr         *OcrService
	Pdf         *PdfService
	Permissions *PermissionsService
	Qr          *QrService
	Storage     *StorageService
	Templates   *TemplatesService
	Threads     *ThreadsService
	Usage       *UsageService
	Webhooks    *WebhooksService
	Workflows   *WorkflowsService
	// v1.1.0
	Audit  *AuditService
	Issues *IssuesService
	Search *SearchService
	Whoami *WhoamiService
}

// NewClient creates a new PlatformXe API client.
func NewClient(config ClientConfig) *Client {
	if config.BaseURL == "" {
		config.BaseURL = "https://platformxe.com"
	}
	if config.Timeout == 0 {
		config.Timeout = 10
	}
	switch {
	case config.DisableRetries || config.Retries < 0:
		// Explicit "no retries". A negative Retries used to make the request
		// loop run zero attempts and return without ever sending.
		config.Retries = 0
	case config.Retries == 0:
		// Zero value = option omitted → default. Use DisableRetries for zero.
		config.Retries = 2
	}

	c := &Client{
		apiKey:   config.APIKey,
		baseURL:  config.BaseURL,
		retries:  config.Retries,
		failOpen: config.FailOpen,
		http: &http.Client{
			Timeout: time.Duration(config.Timeout) * time.Second,
		},
	}

	c.Documents = &DocumentsService{client: c}
	c.Domains = &DomainsService{client: c}
	c.Events = &EventsService{client: c}
	c.Events.Custom = &CustomEventsService{client: c}
	c.Events.Custom.Marketplace = &MarketplaceService{client: c}
	c.Events.Custom.Federation = &EventFederationService{client: c}
	c.Exports = &ExportsService{client: c}
	c.Fraud = newFraudService(c)
	c.Identity = newIdentityService(c)
	c.Messaging = &MessagingService{client: c}
	c.Ocr = &OcrService{client: c}
	c.Pdf = &PdfService{client: c}
	c.Permissions = &PermissionsService{client: c}
	c.Qr = &QrService{client: c}
	c.Storage = &StorageService{client: c}
	c.Templates = &TemplatesService{client: c}
	c.Threads = &ThreadsService{client: c}
	c.Usage = &UsageService{client: c}
	c.Webhooks = &WebhooksService{client: c}
	c.Workflows = &WorkflowsService{client: c}
	// v1.1.0
	c.Audit = &AuditService{client: c}
	c.Issues = &IssuesService{client: c}
	c.Search = &SearchService{client: c}
	c.Whoami = &WhoamiService{client: c}

	return c
}

// SendTelemetry sends client telemetry metrics as a full batch.
func (c *Client) SendTelemetry(batch TelemetryBatch) (*TelemetryResult, error) {
	var result TelemetryResult
	err := c.doRequestTyped("POST", "/api/v1/telemetry", batch, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// HealthCheck returns the platform health status.
func (c *Client) HealthCheck() (*HealthCheckResult, error) {
	var result HealthCheckResult
	err := c.doRequestTyped("GET", "/api/health", nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Error implements error. The text matches the fmt.Errorf form the client
// returned before *APIError, so callers matching on it keep working.
func (e *APIError) Error() string {
	return fmt.Sprintf("[%s] %s (HTTP %d)", e.Code, e.Message, e.StatusCode)
}

// isRetryableStatus reports whether a non-2xx status is retried on the
// transport backoff. 429 and 502/503/504 are retried for every method: the
// request was rate-limited or never reached a handler. 500 means a handler
// threw and its side effects may have landed, so it is retried only for an
// idempotent method. Every other status is final.
func isRetryableStatus(method string, status int) bool {
	switch status {
	case http.StatusTooManyRequests, http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout:
		return true
	case http.StatusInternalServerError:
		return isIdempotentMethod(method)
	default:
		return false
	}
}

// isIdempotentMethod reports whether repeating a request with this method has
// the same effect as sending it once (RFC 9110 §9.2.2).
func isIdempotentMethod(method string) bool {
	switch strings.ToUpper(method) {
	case http.MethodGet, http.MethodHead, http.MethodPut, http.MethodDelete, http.MethodOptions:
		return true
	default:
		return false
	}
}

// httpStatusError builds the *APIError for a non-2xx response. Code and
// message come from a JSON body when it has them — an {"error":{"code",
// "message"}} object (the envelope included), an {"error":"..."} string, or a
// top-level "message" — otherwise the code is HTTP_<status> and the message is
// the status text (a proxy's HTML error page, for example).
func httpStatusError(status int, body []byte) *APIError {
	apiErr := &APIError{Code: fmt.Sprintf("HTTP_%d", status), Message: http.StatusText(status), StatusCode: status}
	if apiErr.Message == "" {
		apiErr.Message = fmt.Sprintf("HTTP %d", status)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(body, &parsed); err != nil {
		return apiErr
	}
	switch e := parsed["error"].(type) {
	case map[string]interface{}:
		if code, ok := e["code"].(string); ok && code != "" {
			apiErr.Code = code
		}
		if message, ok := e["message"].(string); ok && message != "" {
			apiErr.Message = message
		}
	case string:
		if e != "" {
			apiErr.Message = e
		}
	default:
		if message, ok := parsed["message"].(string); ok && message != "" {
			apiErr.Message = message
		}
	}
	return apiErr
}

// backoff sleeps before the next attempt (200 ms doubled per attempt), and not
// at all after the last one.
func (c *Client) backoff(attempt int) {
	if attempt < c.retries {
		time.Sleep(time.Duration(200*(1<<attempt)) * time.Millisecond)
	}
}

func (c *Client) doRequest(method, path string, body interface{}, params map[string]string) (map[string]interface{}, error) {
	var lastErr error
	// lastErrIsHTTP: the last attempt got a retryable HTTP error status
	// (429/500/502/503/504) rather than a transport failure. FailOpen does not
	// cover it.
	lastErrIsHTTP := false

	for attempt := 0; attempt <= c.retries; attempt++ {
		var reqBody io.Reader
		if body != nil {
			b, err := json.Marshal(body)
			if err != nil {
				return nil, fmt.Errorf("marshal body: %w", err)
			}
			reqBody = bytes.NewReader(b)
		}

		fullURL := c.baseURL + path
		if len(params) > 0 {
			q := url.Values{}
			for k, v := range params {
				q.Set(k, v)
			}
			fullURL += "?" + q.Encode()
		}

		req, err := http.NewRequest(method, fullURL, reqBody)
		if err != nil {
			return nil, fmt.Errorf("create request: %w", err)
		}

		req.Header.Set("x-api-key", c.apiKey)
		req.Header.Set("Accept", "application/json")
		if body != nil {
			req.Header.Set("Content-Type", "application/json")
		}

		resp, err := c.http.Do(req)
		if err != nil {
			lastErr, lastErrIsHTTP = err, false
			c.backoff(attempt)
			continue
		}

		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			lastErr, lastErrIsHTTP = err, false
			continue
		}

		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			apiErr := httpStatusError(resp.StatusCode, respBody)
			if !isRetryableStatus(method, resp.StatusCode) {
				return nil, apiErr
			}
			lastErr, lastErrIsHTTP = apiErr, true
			c.backoff(attempt)
			continue
		}

		// 2xx — unchanged.
		var result map[string]interface{}
		if err := json.Unmarshal(respBody, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response: %w", err)
		}

		if success, ok := result["success"].(bool); ok && !success {
			errMap, _ := result["error"].(map[string]interface{})
			code, _ := errMap["code"].(string)
			message, _ := errMap["message"].(string)
			return nil, &APIError{Code: code, Message: message, StatusCode: resp.StatusCode}
		}

		if data, ok := result["data"].(map[string]interface{}); ok {
			return data, nil
		}
		return result, nil
	}

	if lastErrIsHTTP {
		// A 5xx or 429 on every attempt is an API error, not an outage of the
		// transport: FailOpen does not apply.
		return nil, lastErr
	}
	if c.failOpen {
		return map[string]interface{}{"_failed": true, "error": fmt.Sprintf("%v", lastErr)}, nil
	}
	return nil, fmt.Errorf("request failed after %d attempts: %w", c.retries+1, lastErr)
}

// doRequestTyped performs an HTTP request and unmarshals the response data into the typed result.
func (c *Client) doRequestTyped(method, path string, body interface{}, params map[string]string, result interface{}) error {
	raw, err := c.doRequest(method, path, body, params)
	if err != nil {
		return err
	}
	jsonBytes, err := json.Marshal(raw)
	if err != nil {
		return fmt.Errorf("failed to marshal response: %w", err)
	}
	return json.Unmarshal(jsonBytes, result)
}

// Get performs an HTTP GET request.
func (c *Client) Get(path string, params map[string]string) (map[string]interface{}, error) {
	return c.doRequest("GET", path, nil, params)
}

// Post performs an HTTP POST request.
func (c *Client) Post(path string, body interface{}) (map[string]interface{}, error) {
	return c.doRequest("POST", path, body, nil)
}

// Put performs an HTTP PUT request.
func (c *Client) Put(path string, body interface{}) (map[string]interface{}, error) {
	return c.doRequest("PUT", path, body, nil)
}

// Patch performs an HTTP PATCH request.
func (c *Client) Patch(path string, body interface{}) (map[string]interface{}, error) {
	return c.doRequest("PATCH", path, body, nil)
}

// Delete performs an HTTP DELETE request.
func (c *Client) Delete(path string) (map[string]interface{}, error) {
	return c.doRequest("DELETE", path, nil, nil)
}
