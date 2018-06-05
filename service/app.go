package service

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/op/go-logging"
	"os"
	"os/signal"
	"syscall"
)

// App standard webapp
type App struct {
	Name           string
	version        string
	buildDate      string
	Log            *logging.Logger
	Echo           *echo.Echo
	httpConfig     *HTTPConfig
	LogConfig      *LogConfig
	sigHUPHandlers map[string]func()
	sigHUPChannel  chan os.Signal
}

// New Creates a new application
func New(name string, config *HTTPConfig, logConfig *LogConfig, version string, buildDate string) *App {
	app := &App{
		Name:           name,
		httpConfig:     config,
		LogConfig:      logConfig,
		version:        version,
		buildDate:      buildDate,
		Echo:           echo.New(),
		sigHUPChannel:  make(chan os.Signal, 1),
		sigHUPHandlers: make(map[string]func()),
	}

	signal.Notify(app.sigHUPChannel, syscall.SIGHUP)

	app.initLogger()
	app.initSignalHandlers()

	return app
}

// Run run the webapp
func (app *App) Run(startCallback func()) {
	app.Log.Infof("Starting %s version %s.", app.Name, app.version)
	app.Log.Infof("Build date: %s.", app.buildDate)

	app.Echo.GET("/v1/version", app.httpGetVersion)

	if app.httpConfig.EnableCORS {
		app.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS, echo.HEAD},
			AllowHeaders: []string{"Authorization"},
		}))
	}

	if startCallback != nil {
		startCallback()
	}

	httpListeningAdress := fmt.Sprintf(":%d", app.httpConfig.Port)

	app.Log.Error("%", app.Echo.Start(httpListeningAdress))
	app.Log.Infof("Exiting %s", app.Name)
}

func (app *App) initSignalHandlers() {
	go func() {
		for sig := range app.sigHUPChannel {
			for name, handler := range app.sigHUPHandlers {
				app.Log.Infof("SIGHUP(%s) received, running '%s' handler", sig, name)
				handler()
			}
		}
	}()
}

// OnSigHUP register sighup callback with a name
func (app *App) OnSigHUP(name string, sigCallback func()) {
	if sigCallback == nil {
		app.Log.Fatalf("sigCallback is nil, it's not permitted!")
	}

	app.sigHUPHandlers[name] = sigCallback
}
