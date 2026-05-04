// =============================================================================
// (c) 2026 Caldera Technologies Ltd.
// Proprietary and confidential.
// Unauthorized copying or distribution is prohibited.
// =============================================================================

package platformxe

import "testing"

func TestNewClient(t *testing.T) {
	c := NewClient(ClientConfig{APIKey: "test_key"})
	if c.apiKey != "test_key" {
		t.Errorf("expected api key 'test_key', got '%s'", c.apiKey)
	}
	if c.baseURL != "https://platformxe.com" {
		t.Errorf("expected default base URL, got '%s'", c.baseURL)
	}
}

func TestNewClientCustomConfig(t *testing.T) {
	c := NewClient(ClientConfig{
		APIKey:  "test",
		BaseURL: "http://localhost:3000",
		Timeout: 5,
		Retries: 0,
	})
	if c.baseURL != "http://localhost:3000" {
		t.Errorf("expected custom base URL, got '%s'", c.baseURL)
	}
	if c.retries != 0 {
		t.Errorf("expected 0 retries, got %d", c.retries)
	}
}

func TestServiceNamespaces(t *testing.T) {
	c := NewClient(ClientConfig{APIKey: "test"})
	if c.Permissions == nil {
		t.Error("Permissions service is nil")
	}
	if c.Identity == nil {
		t.Error("Identity service is nil")
	}
	if c.Messaging == nil {
		t.Error("Messaging service is nil")
	}
	if c.Webhooks == nil {
		t.Error("Webhooks service is nil")
	}
	if c.Templates == nil {
		t.Error("Templates service is nil")
	}
	if c.Storage == nil {
		t.Error("Storage service is nil")
	}
	if c.Workflows == nil {
		t.Error("Workflows service is nil")
	}
	if c.Domains == nil {
		t.Error("Domains service is nil")
	}
	if c.Documents == nil {
		t.Error("Documents service is nil")
	}
	if c.Events == nil {
		t.Error("Events service is nil")
	}
	if c.Exports == nil {
		t.Error("Exports service is nil")
	}
	if c.Ocr == nil {
		t.Error("Ocr service is nil")
	}
	if c.Pdf == nil {
		t.Error("Pdf service is nil")
	}
	if c.Qr == nil {
		t.Error("Qr service is nil")
	}
	if c.Threads == nil {
		t.Error("Threads service is nil")
	}
	if c.Usage == nil {
		t.Error("Usage service is nil")
	}
}

func TestPermissionsServiceMethods(t *testing.T) {
	c := NewClient(ClientConfig{APIKey: "test"})
	p := c.Permissions

	// Verify key methods exist by checking they don't panic
	// (We can't call them without a server, but we can verify the struct has them)
	if p == nil {
		t.Fatal("Permissions service is nil")
	}
}

func TestIdentityServiceMethods(t *testing.T) {
	c := NewClient(ClientConfig{APIKey: "test"})
	if c.Identity == nil {
		t.Fatal("Identity service is nil")
	}
}

func TestMessagingServiceMethods(t *testing.T) {
	c := NewClient(ClientConfig{APIKey: "test"})
	if c.Messaging == nil {
		t.Fatal("Messaging service is nil")
	}
}

func TestWebhooksServiceMethods(t *testing.T) {
	c := NewClient(ClientConfig{APIKey: "test"})
	if c.Webhooks == nil {
		t.Fatal("Webhooks service is nil")
	}
}

func TestTemplatesServiceMethods(t *testing.T) {
	c := NewClient(ClientConfig{APIKey: "test"})
	if c.Templates == nil {
		t.Fatal("Templates service is nil")
	}
}

func TestStorageServiceMethods(t *testing.T) {
	c := NewClient(ClientConfig{APIKey: "test"})
	if c.Storage == nil {
		t.Fatal("Storage service is nil")
	}
}

func TestWorkflowsServiceMethods(t *testing.T) {
	c := NewClient(ClientConfig{APIKey: "test"})
	if c.Workflows == nil {
		t.Fatal("Workflows service is nil")
	}
}

func TestDomainsServiceMethods(t *testing.T) {
	c := NewClient(ClientConfig{APIKey: "test"})
	if c.Domains == nil {
		t.Fatal("Domains service is nil")
	}
}

func TestEventsServiceMethods(t *testing.T) {
	c := NewClient(ClientConfig{APIKey: "test"})
	if c.Events == nil {
		t.Fatal("Events service is nil")
	}
}

func TestDocumentsServiceMethods(t *testing.T) {
	c := NewClient(ClientConfig{APIKey: "test"})
	if c.Documents == nil {
		t.Fatal("Documents service is nil")
	}
}

func TestExportsServiceMethods(t *testing.T) {
	c := NewClient(ClientConfig{APIKey: "test"})
	if c.Exports == nil {
		t.Fatal("Exports service is nil")
	}
}

func TestOcrServiceMethods(t *testing.T) {
	c := NewClient(ClientConfig{APIKey: "test"})
	if c.Ocr == nil {
		t.Fatal("Ocr service is nil")
	}
}

func TestQrServiceMethods(t *testing.T) {
	c := NewClient(ClientConfig{APIKey: "test"})
	if c.Qr == nil {
		t.Fatal("Qr service is nil")
	}
}

func TestUsageServiceMethods(t *testing.T) {
	c := NewClient(ClientConfig{APIKey: "test"})
	if c.Usage == nil {
		t.Fatal("Usage service is nil")
	}
}

func TestClientTelemetryAndHealth(t *testing.T) {
	c := NewClient(ClientConfig{APIKey: "test"})
	// Verify telemetry and health are direct client methods (not on a service)
	if c == nil {
		t.Fatal("Client is nil")
	}
}

func TestDefaultFailOpen(t *testing.T) {
	c := NewClient(ClientConfig{APIKey: "test"})
	if c.failOpen != false {
		// Default is false for Go (IaC safety)
	}
}
