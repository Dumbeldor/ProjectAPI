package internal

import (
	"github.com/labstack/echo"
	"gitlab.com/projetAPI/easyhttp"
	"net/http"
)

// swagger:route PUT /v1/user/lock user jwtToken
//
// Handler to lock a user
//
// Security:
//    jwtToken: read
//
// Responses:
//    200: MessageResponse
//    403: ErrorResponse
//    500: ErrorResponse
func httpLockUser(c echo.Context) error {
	userSess, err := app.ValidateSession(c, sessionReader)
	if err != nil || userSess == nil {
		return err
	}

	err = gUserDB.UserLock(userSess.UserID, true)
	if err != nil {
		return app.Error500(c, err)
	}

	var msg easyhttp.MessageResponse
	msg.Body.Message = "You are locked out"
	return c.JSON(http.StatusOK, msg.Body)
}

// swagger:route PUT /v1/user/unlock user jwtToken
//
// Handler to unlock a user
//
// Security:
//    jwtToken: read
//
// Responses:
//    200: MessageResponse
//    403: ErrorResponse
//    500: ErrorResponse
func httpUnlockUser(c echo.Context) error {
	userSess, err := app.ValidateSession(c, sessionReader)
	if err != nil || userSess == nil {
		return err
	}

	err = gUserDB.UserLock(userSess.UserID, false)
	if err != nil {
		return app.Error500(c, err)
	}

	var msg easyhttp.MessageResponse
	msg.Body.Message = "You are unlocked"
	return c.JSON(http.StatusOK, msg.Body)
}