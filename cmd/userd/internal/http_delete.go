package internal

import (
	"github.com/labstack/echo"
	"net/http"
	"gitlab.com/projetAPI/easyhttp"
)

// swagger:route DELETE /v1/user user jwtToken
//
// Handler to delete a user
//
// Security:
//    jwtToken: read
//
// Responses:
//    200: MessageResponse
//    403: ErrorResponse
//    500: ErrorResponse
func httpDeleteUser(c echo.Context) error {
	userSess, err := app.ValidateSession(c, sessionReader)
	if err != nil || userSess == nil {
		return err
	}

	err = gUserDB.UserDelete(userSess.UserID)
	if err != nil {
		return app.Error500(c, err)
	}

	var msg easyhttp.MessageResponse
	msg.Body.Message = "You are deleted"
	return c.JSON(http.StatusOK, msg.Body)
}
