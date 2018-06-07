// Package internal authd app internals
// swagger:meta
package internal

import (
	"github.com/labstack/echo"
	"net/http"
	"gitlab.com/projetAPI/ProjetAPI/service"
)

// swagger:response sessionRemovalResponse
type sessRemovalResponse struct {
	// in: body
	Body struct {
		// Response status
		// required: true
		Status string `json:"status"`
	}
}

// swagger:route POST /v1/auth/logout user jwtToken
//
// Handler to logout
//
// Responses:
//    200: sessionRemovalResponse
//    403: ErrorResponse
//    500: ErrorResponse
func httpAuthLogout(c echo.Context) error {

	if sessionReader == nil {
		return app.Error500(c, &echo.HTTPError{Message: "Failed to instantiate sessionWriter"})
	}
	userSess, err := service.ValidateSession(c, sessionReader)
	if userSess == nil {
		return err
	}

	sw := newWriter(gconfig.Redis)
	if !sw.Destroy(userSess.UserID) {
		return app.Error500(c, &echo.HTTPError{Message: "Failed to remove session"})
	}

	srr := sessRemovalResponse{}
	srr.Body.Status = "OK"
	return c.JSON(http.StatusOK, srr.Body)
}
