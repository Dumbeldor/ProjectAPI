package service

import (
	"fmt"
	"github.com/op/go-logging"
	"os"
)

var format = logging.MustStringFormatter(
	`%{color}%{time:02/01/2006 15:04:05.000} %{longfunc} - %{level:.5s} %{message}`,
)

func (app *App) logLevelFromConfig(name string, backend logging.LeveledBackend) logging.Level {
	switch app.LogConfig.LogLevel {
	case "debug":
		fmt.Printf("Logging level set to %s for backend %s.\n", app.LogConfig.LogLevel, name)
		backend.SetLevel(logging.DEBUG, "")
		return logging.DEBUG
	case "info":
		fmt.Printf("Logging level set to %s for backend %s.\n", app.LogConfig.LogLevel, name)
		backend.SetLevel(logging.INFO, "")
		return logging.INFO
	case "notice":
		fmt.Printf("Logging level set to %s for backend %s.\n", app.LogConfig.LogLevel, name)
		backend.SetLevel(logging.NOTICE, "")
		return logging.NOTICE
	case "warning":
		fmt.Printf("Logging level set to %s for backend %s.\n", app.LogConfig.LogLevel, name)
		backend.SetLevel(logging.WARNING, "")
		return logging.WARNING
	case "error":
		fmt.Printf("Logging level set to %s for backend %s.\n", app.LogConfig.LogLevel, name)
		backend.SetLevel(logging.ERROR, "")
		return logging.ERROR
	default:
		fmt.Printf("Logging level set to info for backend %s.\n", name)
		backend.SetLevel(logging.INFO, "")
		return logging.INFO
	}
}

func (app *App) initLogger() bool {
	app.Log = logging.MustGetLogger(app.Name)

	stderrLeveled := logging.AddModuleLevel(logging.NewLogBackend(os.Stderr, "", 0))
	app.logLevelFromConfig("stderr", stderrLeveled)

	if app.LogConfig.EnableSyslog {
		syslogBackend, err := logging.NewSyslogBackend(app.Name)
		if err != nil {
			app.Log.Error("Failed to setup syslog backend.")
			return false
		}

		syslogLeveled := logging.AddModuleLevel(syslogBackend)
		app.logLevelFromConfig("syslog", syslogLeveled)

		logging.SetBackend(logging.NewBackendFormatter(stderrLeveled, format), syslogLeveled)
	} else {
		logging.SetBackend(logging.NewBackendFormatter(stderrLeveled, format))
	}

	return true
}
