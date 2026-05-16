// =============================================================================
// (c) 2026 Caldera Technologies Ltd.
// Proprietary and confidential.
// Unauthorized copying or distribution is prohibited.
// =============================================================================

package platformxe

import "fmt"

// DocumentsService handles fixed-storage documents, folders, and deletion overrides.
type DocumentsService struct {
	client *Client
}

// ListDocuments returns documents with optional filters.
func (s *DocumentsService) ListDocuments(params map[string]string) (*DocumentListResponse, error) {
	var result DocumentListResponse
	err := s.client.doRequestTyped("GET", "/api/v1/storage/fixed/documents", nil, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDocument returns a document by ID.
func (s *DocumentsService) GetDocument(id string) (*Document, error) {
	var result Document
	err := s.client.doRequestTyped("GET", fmt.Sprintf("/api/v1/storage/fixed/documents/%s", id), nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateDocument creates a new fixed-storage document.
func (s *DocumentsService) CreateDocument(input map[string]interface{}) (*Document, error) {
	var result Document
	err := s.client.doRequestTyped("POST", "/api/v1/storage/fixed/documents", input, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateDocument updates a fixed-storage document.
func (s *DocumentsService) UpdateDocument(documentID string, input map[string]interface{}) (*Document, error) {
	var result Document
	err := s.client.doRequestTyped("PATCH", fmt.Sprintf("/api/v1/storage/fixed/documents/%s", documentID), input, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteDocument deletes a fixed-storage document.
func (s *DocumentsService) DeleteDocument(documentID string) (*DeleteResult, error) {
	var result DeleteResult
	err := s.client.doRequestTyped("DELETE", fmt.Sprintf("/api/v1/storage/fixed/documents/%s", documentID), nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListFolders returns all storage folders.
func (s *DocumentsService) ListFolders() (*FolderListResponse, error) {
	var result FolderListResponse
	err := s.client.doRequestTyped("GET", "/api/v1/storage/fixed/folders", nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetFolder returns a folder by ID.
func (s *DocumentsService) GetFolder(id string) (*DocumentFolder, error) {
	var result DocumentFolder
	err := s.client.doRequestTyped("GET", fmt.Sprintf("/api/v1/storage/fixed/folders/%s", id), nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateFolder creates a new storage folder.
func (s *DocumentsService) CreateFolder(input map[string]interface{}) (*DocumentFolder, error) {
	var result DocumentFolder
	err := s.client.doRequestTyped("POST", "/api/v1/storage/fixed/folders", input, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RequestOverride requests a deletion override for a protected document.
func (s *DocumentsService) RequestOverride(input map[string]interface{}) (*DeletionOverride, error) {
	var result DeletionOverride
	err := s.client.doRequestTyped("POST", "/api/v1/storage/fixed/overrides", input, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ProcessOverride approves or rejects a deletion override request.
func (s *DocumentsService) ProcessOverride(overrideID string, input map[string]interface{}) (*DeletionOverride, error) {
	var result DeletionOverride
	err := s.client.doRequestTyped("PATCH", fmt.Sprintf("/api/v1/storage/fixed/overrides/%s", overrideID), input, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
