package internal

import "fmt"

// swagger:parameters createMessageRequest
type createMessageRequest struct {
	Message string `json:"message"`
	Receiver string `json:"receiver"`
}

func (cmr *createMessageRequest) Validate() error {
	if len(cmr.Message) < 1 {
		return fmt.Errorf("1 characters is the minimum message length")
	}
	if len(cmr.Message) > 2000 {
		return fmt.Errorf("Please make your message shorter. We've set the limit at 2,000 characters to be courteous to others")
	}

	return nil
}
