package entity

// Structure of message
type Message struct {
	ID string `json:"id"`
	Message string `json:"message"`
	CreationDate string `json:"creation-date"`
	UserSenderID string `json:"user-sender-id"`
	UserSenderName string `json:"user-sender-name"`
	UserReceiverID string `json:"user-receiver-id"`
	UserReceiverName string `json:"user-receiver-name"`
}
