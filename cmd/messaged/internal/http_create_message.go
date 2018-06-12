package internal

import (
	"github.com/labstack/echo"
	"gitlab.com/projetAPI/easyhttp"
	"net/http"
)

// swagger:route POST /v1/message message messageRequest
//
// Sent a message to a user
//
// Security:
//    jwtToken: read
//
// Responses:
//    200: MessageResponse
//    400: ErrorResponse
//    409: ErrorResponse
//    500: ErrorResponse
func httpCreateMessage(c echo.Context) error {
	var cmreq messageRequest
	if !easyhttp.ReadJsonRequest(c.Request().Body, &cmreq) {
		return app.Error400(c, "Request body is not a JSON.")
	}

	if err := cmreq.Validate(); err != nil {
		return app.Error(c, http.StatusBadRequest, err.Error())
	}

	userSess, err := app.ValidateSession(c, sessionReader)
	if err != nil || userSess == nil {
		return err
	}

	userExist, err := gUserDB.LoginExists(cmreq.NameReceiver)
	if err != nil {
		return app.Error500(c, err)
	}

	if !userExist {
		return app.Error(c, http.StatusBadRequest, "The user does not exist.")
	}

	err = gUserDB.CreateMessage(cmreq.Message, userSess.UserID, cmreq.NameReceiver)
	if err != nil {
		return app.Error500(c, err)
	}

	var msg easyhttp.MessageResponse
	msg.Body.Message = "The sending of the message: "+ cmreq.Message + " to the " + cmreq.NameReceiver + " user has gone well."
	return c.JSON(http.StatusOK, msg.Body)
}