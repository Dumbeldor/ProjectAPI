package internal

import "github.com/labstack/echo"

// swagger:route GET /v1/message message getMessage
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

	mlr.Body.Messages, err := gUserDB.GetMessageForUser(userSess.UserID)
	
}
