package mock

import (
	"github.com/op/go-logging"
	"gitlab.com/projetAPI/ProjetAPI/service"
)

// InsertSession use for mock session reader
func InsertSession(sessionWriter service.WriterInterface, log *logging.Logger, userID string, tokenString string) {
	if sessionWriter == nil {
		log.Fatal("sessionWriter nil")
	}
	sessionWriter.Write(userID, tokenString)
}
