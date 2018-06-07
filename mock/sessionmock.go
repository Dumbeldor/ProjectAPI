package mock

import (
	"gitlab.com/projetAPI/auth"
	"github.com/op/go-logging"
)

const (
	userID      = "a1472317-f1e2-413e-b55d-d412b9f4953f"
	tokenString = "a1472317-f1e2-413e-b55d-d412b9f4953f"
)

// InsertSession use for mock session reader
func InsertSession(sessionWriter auth.WriterInterface, log *logging.Logger) {
	if sessionWriter == nil {
		log.Fatal("sessionWriter nil")
	}
	sessionWriter.Write(userID, tokenString)
}
