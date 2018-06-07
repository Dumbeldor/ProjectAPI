// Package internal authd app internals
// swagger:meta
package internal

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	"gitlab.com/projetAPI/easyhttp"
	"net/http"
	"gitlab.com/projetAPI/ProjetAPI/service"
)

// swagger:response authResponse
type authResponse struct {
	// in: body
	Body struct {
		// The JWT response token
		// required: true
		Token string `json:"token"`

		// User ID
		// required: true
		UserID string `json:"user_id"`
	}
}

type loginInfos struct {
	UserID   string
	Login    string
	Password string
	Locked   bool
	Salt1    string
	Salt2    string
}

// AuthLogin swagger:route POST /v1/auth/login user login
//
// Handler to login
//
// Responses:
//    200: authResponse
//    400: ErrorResponse
//    403: ErrorResponse
//    500: ErrorResponse
func httpAuthLogin(c echo.Context) error {
	var areq authRequest
	if !easyhttp.ReadJsonRequest(c.Request().Body, &areq.Body) {
		return app.Error400(c, "Request body is not a JSON.")
	}

	if err := areq.Validate(); err != nil {
		return app.Error400(c, "Unable to validate auth request.")
	}

	if !verifyAuthDB() {
		return app.Error500(c, &echo.HTTPError{Message: "Failed to verify authentication database"})
	}

	linfos, err := gAuthDB.getLoginInfos(areq.Body.Login)
	if err != nil {
		if err == sql.ErrNoRows {
			var er easyhttp.ErrorResponse
			er.Body.Message = "Invalid user/password."
			return easyhttp.WriteJSONError(c, app.Log, http.StatusForbidden, er.Body,
				fmt.Sprintf("Invalid user %s", areq.Body.Login))
		}
		return app.Error500(c, err)
	}

	if linfos.Locked {
		var er easyhttp.ErrorResponse
		er.Body.Message = "Account locked."
		return easyhttp.WriteJSONError(c, app.Log, http.StatusForbidden, er.Body,
			fmt.Sprintf("User %s tried to auth but account is locked.", areq.Body.Login))
	}

	encodedPassword := service.EncodePassword(areq.Body.Login, areq.Body.Password, linfos.Salt1, linfos.Salt2)
	if encodedPassword != linfos.Password {
		var er easyhttp.ErrorResponse
		er.Body.Message = "Invalid user/password."
		return easyhttp.WriteJSONError(c, app.Log, http.StatusForbidden, er.Body,
			fmt.Sprintf("Invalid password for user %s", areq.Body.Login))
	}

	var arep authResponse
	var tokenSecret string

	arep.Body.Token, tokenSecret, err = createJWT(areq, linfos)
	if err != nil {
		return app.Error500(c, err)
	}

	if sessionWriter == nil {
		return app.Error500(c, &echo.HTTPError{Message: "Failed to instantiate sessionWriter"})
	}

	if !sessionWriter.Write(linfos.UserID, tokenSecret) {
		return app.Error500(c, &echo.HTTPError{
			Message: fmt.Sprintf("Unable to create session for user %s.", areq.Body.Login),
		})
	}

	arep.Body.UserID = linfos.UserID

	err = c.JSON(http.StatusOK, arep.Body)
	app.Log.Infof("Authentication succeed for user %s.", areq.Body.Login)

	return err
}
