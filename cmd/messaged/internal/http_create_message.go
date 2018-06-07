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
	var cmreq createMessageRequest
	if !easyhttp.ReadJsonRequest(c.Request().Body, &cmreq) {
		return app.Error400(c, "Request body is not a JSON.")
	}

	if err := cmreq.Validate(); err != nil {
		return app.Error(c, http.StatusNotAcceptable, err.Error())
	}

	userSess, err := app.ValidateSession(c, sessionReader)
	if err != nil || userSess == nil {
		return err
	}

	userExist, err := gUserDB.LoginExists(cmreq.NameReceiver)
	if err != nil {
		return app.Error500(c, err)
	}

	if userExist {
		return app.Error(c, http.StatusConflict, "")
	}

	var msg easyhttp.MessageResponse
	msg.Body.Message = "Send successful message."
	return c.JSON(http.StatusOK, msg.Body)
}