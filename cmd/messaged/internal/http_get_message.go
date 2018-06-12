package internal

import (
	"github.com/labstack/echo"
	"database/sql"
	"net/http"
)

// swagger:route GET /v1/message message messageListResponse
//
// Returns messages received by the logged-in user
//
// Security:
//    jwtToken: read
//
// Handler to register
//
// Responses:
//    200: messageListResponse
//    400: ErrorResponse
//    409: ErrorResponse
//    500: ErrorResponse
func httpGetMessage(c echo.Context) error {
	userSess, err := app.ValidateSession(c, sessionReader)
	if err != nil || userSess == nil {
		return err
	}

	mlr := messageListResponse{}

	mlr.Body.Messages, err = gUserDB.GetMessageForUser(userSess.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return app.Error(c, http.StatusNoContent, "You have no message !")
		}
		return app.Error500(c, err)
	}

	if mlr.Body.Messages == nil {
		return app.Error(c, http.StatusNoContent, "You have no message !")
	}

	return c.JSON(http.StatusOK, mlr.Body)
}
