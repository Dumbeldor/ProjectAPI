package internal

import (
	"github.com/labstack/echo"
	"gitlab.com/projetAPI/easyhttp"
	"net/http"
)

// swagger:route PUT /v1/message/:messageID message modifyMessageRequest
//
// Handler to modify a message
//
// Security:
//    jwtToken: read
//
// Responses:
//    200: MessageResponse
//    400: ErrorResponse
//    401: ErrorResponse
//    409: ErrorResponse
//    500: ErrorResponse
func httpPutMessage(c echo.Context) error {
	userSess, err := app.ValidateSession(c, sessionReader)
	if err != nil || userSess == nil {
		return err
	}

	messageID := c.Param("messageID")
	if len(messageID) != 36 {
		return app.Error400(c, "The message identifier is incorrect")
	}

	var mreq modifyMessageRequest
	if !easyhttp.ReadJsonRequest(c.Request().Body, &mreq) {
		return app.Error400(c, "Request body is not a JSON.")
	}

	if err := mreq.Validate(); err != nil {
		return app.Error(c, http.StatusBadRequest, err.Error())
	}

	permission, err := gUserDB.MessagePutPermission(messageID, userSess.UserID)
	if err != nil {
		return app.Error500(c, err)
	}
	if !permission {
		return app.Error(c, http.StatusUnauthorized, "You are Unauthorized to modify this message")
	}

	// Check message
	messageExist, err := gUserDB.MessageExists(messageID)
	if err != nil {
		return app.Error500(c, err)
	}
	if !messageExist {
		return app.Error400(c, "The message you want to edit does not exist")
	}

	err = gUserDB.MessageUpdate(messageID, mreq.Message)
	if err != nil {
		return app.Error500(c, err)
	}

	var msg easyhttp.MessageResponse
	msg.Body.Message = "You message has just been modified"
	return c.JSON(http.StatusOK, msg.Body)
}