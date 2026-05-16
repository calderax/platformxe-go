// =============================================================================
// (c) 2026 Caldera Technologies Ltd.
// Proprietary and confidential.
// Unauthorized copying or distribution is prohibited.
// =============================================================================

package platformxe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	Audit       *AuditService
	Issues      *IssuesService
	Search      *SearchService
	Whoami      *WhoamiService
}

// NewClient creates a new PlatformXe API client.
func NewClient(config ClientConfig) *Client {
	if config.BaseURL == "" {
		config.BaseURL = "https://platformxe.com"
	}
	if config.Timeout == 0 {
		config.Timeout = 10
	}
	if config.Retries == 0 {
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

func (c *Client) doRequest(method, path string, body interface{}, params map[string]string) (map[string]interface{}, error) {
	var lastErr error

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
			lastErr = err
			if attempt < c.retries {
				time.Sleep(time.Duration(200*(1<<attempt)) * time.Millisecond)
			}
			continue
		}
		defer resp.Body.Close()

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			lastErr = err
			continue
		}

		var result map[string]interface{}
		if err := json.Unmarshal(respBody, &result); err != nil {
			return nil, fmt.Errorf("unmarshal response: %w", err)
		}

		if success, ok := result["success"].(bool); ok && !success {
			errMap, _ := result["error"].(map[string]interface{})
			code, _ := errMap["code"].(string)
			message, _ := errMap["message"].(string)
			return nil, fmt.Errorf("[%s] %s (HTTP %d)", code, message, resp.StatusCode)
		}

		if data, ok := result["data"].(map[string]interface{}); ok {
			return data, nil
		}
		return result, nil
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
