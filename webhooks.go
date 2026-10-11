// =============================================================================
// (c) 2026 Caldera Technologies Ltd.
// Proprietary and confidential.
// Unauthorized copying or distribution is prohibited.
// =============================================================================

package platformxe

import "fmt"

// WebhooksService handles webhook endpoint management.
type WebhooksService struct {
	client *Client
}

// List returns all webhook endpoints.
func (s *WebhooksService) List() (*WebhookListResponse, error) {
	var result WebhookListResponse
	err := s.client.doRequestTyped("GET", "/api/v1/webhooks", nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new webhook endpoint.
func (s *WebhooksService) Create(input map[string]interface{}) (*Webhook, error) {
	var result Webhook
	err := s.client.doRequestTyped("POST", "/api/v1/webhooks", input, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Get returns a webhook endpoint by ID.
func (s *WebhooksService) Get(id string) (*Webhook, error) {
	var result Webhook
	err := s.client.doRequestTyped("GET", fmt.Sprintf("/api/v1/webhooks/%s", id), nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a webhook endpoint.
func (s *WebhooksService) Update(id string, input map[string]interface{}) (*Webhook, error) {
	var result Webhook
	err := s.client.doRequestTyped("PATCH", fmt.Sprintf("/api/v1/webhooks/%s", id), input, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a webhook endpoint.
func (s *WebhooksService) Delete(id string) (*DeleteResult, error) {
	var result DeleteResult
	err := s.client.doRequestTyped("DELETE", fmt.Sprintf("/api/v1/webhooks/%s", id), nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RotateSecret rotates the signing secret for a webhook endpoint.
func (s *WebhooksService) RotateSecret(id string) (*Webhook, error) {
	var result Webhook
	err := s.client.doRequestTyped("POST", fmt.Sprintf("/api/v1/webhooks/%s/rotate-secret", id), nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Test sends a test event to a webhook endpoint.
func (s *WebhooksService) Test(id string) (*WebhookTestResult, error) {
	var result WebhookTestResult
	err := s.client.doRequestTyped("POST", fmt.Sprintf("/api/v1/webhooks/%s/test", id), nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
