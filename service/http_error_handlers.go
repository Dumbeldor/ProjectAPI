package service

import (
	"github.com/labstack/echo"
	"gitlab.com/projetAPI/easyhttp"
	"net/http"
)

// Error500 Trigger a standard error 500 message in JSON format
func (app *App) Error500(c echo.Context, err error) error {
	if err == nil {
		app.Log.Errorf("%s - error %d: %s", c.Path(), http.StatusInternalServerError, "Critical server error.")
	} else {
		app.Log.Errorf("%s - error %d: %s", c.Path(), http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusInternalServerError, "Critical server error.")
}

// Error500 Trigger a standard error 500 with custom message in json format
func (app *App) Error500String(c echo.Context, err string) error {
	var er easyhttp.ErrorResponse
	er.Body.Message = err
	app.Log.Errorf("%s - error %d: %s", c.Path(), http.StatusInternalServerError, err)
	return c.JSON(http.StatusInternalServerError, er.Body)
}

// Error400 Trigger a standard error 400 with custom message in json format
func (app *App) Error400(c echo.Context, err string) error {
	var er easyhttp.ErrorResponse
	er.Body.Message = err
	return c.JSON(http.StatusBadRequest, er.Body)
}

// Error404 Trigger a standard error 404 with custom message in json format
func (app *App) Error404(c echo.Context, err string) error {
	var er easyhttp.ErrorResponse
	er.Body.Message = err
	return c.JSON(http.StatusNotFound, er.Body)
}

// Other code error a standard error with custom message in json format
func (app *App) Error(c echo.Context, code int, err string) error {
	var er easyhttp.ErrorResponse
	er.Body.Message = err
	return c.JSON(code, er.Body)
}
