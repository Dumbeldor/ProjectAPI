package mock

import (
	"github.com/op/go-logging"
	"gitlab.com/projetAPI/ProjetAPI/service"
)

const (
	userID      = "a1472317-f1e2-413e-b55d-d412b9f4953f"
	tokenString = "a1472317-f1e2-413e-b55d-d412b9f4953f"
)

// InsertSession use for mock session reader
func InsertSession(sessionWriter service.WriterInterface, log *logging.Logger) {
	if sessionWriter == nil {
		log.Fatal("sessionWriter nil")
	}
	sessionWriter.Write(userID, tokenString)
}
