package internal

import "gitlab.com/projetAPI/ProjetAPI/entity"

// swagger:response messageListResponse
type messageListResponse struct {
	// in: body
	Body struct {
		// Messages
		// required: true
		Messages []entity.Message `json:"messages"`
	}
}
