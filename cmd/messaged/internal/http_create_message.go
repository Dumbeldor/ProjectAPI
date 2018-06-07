package internal

import (
	"github.com/labstack/echo"
	"gitlab.com/projetAPI/easyhttp"
	"net/http"
)

// UserRegister swagger:route POST /v1/message/create message createMessage
//
// Handler to register
//
// Responses:
//    200: MessageResponse
//    400: ErrorResponse
//    409: ErrorResponse
//    500: ErrorResponse
func httpCreateMessage(c echo.Context) error {
	userSess, err := app.ValidateSession(c, sessionReader)
	if err != nil {
		return err
	}
	if userSess == nil {
		return app.Error500String(c, "Error retrieving session")
	}

	var msg easyhttp.MessageResponse
	msg.Body.Message = "Send successful message."
	return c.JSON(http.StatusOK, msg.Body)
}