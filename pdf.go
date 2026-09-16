// =============================================================================
// (c) 2026 Caldera Technologies Ltd.
// Proprietary and confidential.
// Unauthorized copying or distribution is prohibited.
// =============================================================================

package platformxe

// PdfService handles PDF generation API calls.
type PdfService struct {
	client *Client
}

// GetProcessor returns the PDF processor configuration.
func (s *PdfService) GetProcessor() (*ProcessorConfig, error) {
	var result ProcessorConfig
	err := s.client.doRequestTyped("GET", "/api/v1/pdf/processor", nil, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateProcessor updates the PDF processor configuration.
func (s *PdfService) UpdateProcessor(input map[string]interface{}) (*ProcessorConfig, error) {
	var result ProcessorConfig
	err := s.client.doRequestTyped("PUT", "/api/v1/pdf/processor", input, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// OfferLetter generates an offer letter PDF.
func (s *PdfService) OfferLetter(input map[string]interface{}) (map[string]interface{}, error) {
	return s.client.doRequest("POST", "/api/v1/pdf/offer-letter", input, nil)
}

// PropertyFlyer generates a property flyer PDF.
func (s *PdfService) PropertyFlyer(input map[string]interface{}) (map[string]interface{}, error) {
	return s.client.doRequest("POST", "/api/v1/pdf/property-flyer", input, nil)
}
