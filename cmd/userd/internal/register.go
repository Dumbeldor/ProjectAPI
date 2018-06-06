package internal

import (
	"github.com/labstack/echo"
	"gitlab.com/projetAPI/easyhttp"
	"net/http"
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

	var msg easyhttp.MessageResponse
	msg.Body.Message = "Registration succeed."
	return c.JSON(http.StatusOK, msg.Body)
}
