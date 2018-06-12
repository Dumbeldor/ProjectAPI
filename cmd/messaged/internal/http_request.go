package internal

import "fmt"

// swagger:parameters messageRequest
type messageRequest struct {
	Message string `json:"message"`
	NameReceiver string `json:"receiver"`
}

// swagger:parameters modifyMessageRequest
type modifyMessageRequest struct {
	Message string `json:"message"`
}

func validateMessage(message string) error {
	if len(message) < 1 {
		return fmt.Errorf("1 characters is the minimum message length")
	}
	if len(message) > 2000 {
		return fmt.Errorf("Please make your message shorter. We've set the limit at 2,000 characters to be courteous to others")
	}

	return nil
}

func (cmr *messageRequest) Validate() error {
	return validateMessage(cmr.Message)
}

func (mmr *modifyMessageRequest) Validate() error {
	return validateMessage(mmr.Message)
}