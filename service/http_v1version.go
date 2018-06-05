package service

import (
	"github.com/labstack/echo"
	"net/http"
)

// swagger:response versionResponse
type versionResponse struct {
	// in: body
	Body struct {
		// Service version
		// required: true
		Version string `json:"version"`

		// Service build date
		BuildDate string `json:"build_date"`
	}
}

// swagger:route POST /v1/version version
//
// Handler to get version
//
// Responses:
//    200: versionResponse
func (app *App) httpGetVersion(c echo.Context) error {
	r := versionResponse{}
	r.Body.Version = app.version
	r.Body.BuildDate = app.buildDate
	return c.JSON(http.StatusOK, r.Body)
}
