package service

import (
	"fmt"
	"github.com/labstack/echo"
	"gitlab.com/projetAPI/easyhttp"
	"net/http"
)

// ValidateSession validate session from http echo.Context
// return Session object and nil error on success
func (app *App) ValidateSession(c echo.Context, readerInterface ReaderInterface) (*Session, error) {
	jwtHeader, errResp, httpstatus := easyhttp.GetJWTAuthHeader(c.Request())
	if errResp != nil {
		return nil, c.JSON(httpstatus, errResp.Body)
	}

	if readerInterface == nil {
		var er easyhttp.ErrorResponse
		er.Body.Message = "Critical server error."
		return nil, easyhttp.WriteJSONError(c, app.Log, http.StatusInternalServerError, er.Body,
			"Fail to instantiate SessionReader")
	}

	userSess, err := readerInterface.LoadValidSessionFromJWT(jwtHeader.Authorization)
	if err != nil || userSess == nil {
		var errorMsg string
		if userSess != nil {
			errorMsg = fmt.Sprintf("Authorization failed for user %s with error %s", userSess.UserID, err)
		} else {
			errorMsg = fmt.Sprintf("Authorization failed with error %s", err)
		}
		var er easyhttp.ErrorResponse
		er.Body.Message = "Authorization failed."
		return nil, easyhttp.WriteJSONError(c, app.Log, http.StatusForbidden, er.Body,
			errorMsg)
	}

	return userSess, nil
}
