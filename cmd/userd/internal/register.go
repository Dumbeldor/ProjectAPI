package internal

import (
	"github.com/labstack/echo"
	"gitlab.com/projetAPI/auth"
	"gitlab.com/projetAPI/easyhttp"
	"math/rand"
	"net/http"
	"time"
)

const (
	pwSaltBytes = 256
)

// UserRegister swagger:route POST /v1/user/register user registerRequest
//
// Handler to register
//
// Responses:
//    200: MessageResponse
//    400: ErrorResponse
//    409: ErrorResponse
//    500: ErrorResponse
func httpRegister(c echo.Context) error {
	var rreq registerRequest
	if !easyhttp.ReadJsonRequest(c.Request().Body, &rreq) {
		return app.Error400(c, "Request body is not a JSON.")
	}

	if err := rreq.Validate(); err != nil {
		var er easyhttp.ErrorResponse
		er.Body.Message = err.Error()
		return easyhttp.WriteJSONError(c, app.Log, http.StatusNotAcceptable, er.Body, err.Error())
	}

	loginExist, err := gUserDB.LoginExists(rreq.Login)
	if err != nil {
		return app.Error500(c, err)
	}

	if loginExist {
		var er easyhttp.ErrorResponse
		er.Body.Message = "Login already taken."
		return easyhttp.WriteJSONError(c, app.Log, http.StatusConflict, er.Body, "")
	}

	emailExists, err := gUserDB.EmailExists(rreq.Email)
	if err != nil {
		return app.Error500(c, err)
	}

	if emailExists {
		var er easyhttp.ErrorResponse
		er.Body.Message = "Email already taken."
		return easyhttp.WriteJSONError(c, app.Log, http.StatusConflict, er.Body, "")
	}

	salt1, salt2 := generateSalt()

	encodedPassword := auth.EncodePassword(rreq.Login, rreq.Password, salt1, salt2)
	if encodedPassword == "" {
		return app.Error500(c, &echo.HTTPError{Message: "Unable to encode password."})
	}

	err = gUserDB.Register(rreq.Login, rreq.Email, encodedPassword, salt1, salt2)
	if err != nil {
		return app.Error500(c, err)
	}

	var msg easyhttp.MessageResponse
	msg.Body.Message = "Registration succeed."
	return c.JSON(http.StatusOK, msg.Body)
}

func generateSalt() (string, string) {
	rand.Seed(time.Now().UnixNano())
	salt1 := auth.GenerateSalt(pwSaltBytes)
	salt2 := auth.GenerateSalt(pwSaltBytes)
	return salt1, salt2
}
