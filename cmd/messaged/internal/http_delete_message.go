package internal

import (
	"github.com/labstack/echo"
	"net/http"
	"gitlab.com/projetAPI/easyhttp"
)

// swagger:route DELETE /v1/message/:messageID message messageID
//
// Handler to delete a message
//
// Security:
//    jwtToken: read
//
// Responses:
//    200: MessageResponse
//    400: ErrorResponse
//    500: ErrorResponse
func httpDeleteMessage(c echo.Context) error {
	userSess, err := app.ValidateSession(c, sessionReader)
	if err != nil || userSess == nil {
		return err
	}

	messageID := c.Param("messageID")
	if len(messageID) != 36 {
		return app.Error400(c, "The message identifier is incorrect")
	}

	permission, err := gUserDB.MessageDeletePermission(messageID, userSess.UserID)
	if err != nil {
		return app.Error500(c, err)
	}
	if !permission {
		return app.Error(c, http.StatusUnauthorized, "You are Unauthorized to delete this message")
	}


	// Check message
	messageExist, err := gUserDB.MessageExists(messageID)
	if err != nil {
		return app.Error500(c, err)
	}
	if !messageExist {
		return app.Error400(c, "The message you want to delete does not exist")
	}

	err = gUserDB.MessageDelete(messageID)
	if err != nil {
		return app.Error500(c, err)
	}

	var msg easyhttp.MessageResponse
	msg.Body.Message = "You message has just been deleted"
	return c.JSON(http.StatusOK, msg.Body)
}
