package internal

import (
	"gitlab.com/projetAPI/ProjetAPI/service"
)

var app *service.App

// AppVersion application version (from git tag)
var AppVersion = "[unk]"

// AppBuildDate application build date
var AppBuildDate = "[unk]"

var sessionWriter service.WriterInterface
var sessionReader service.ReaderInterface

// StartApp initiate components
// Should be called from main function
func StartApp(configFile string) {
	gconfig.load(configFile)

	app = service.New(
		"uc-auth",
		&gconfig.HTTP,
		&gconfig.Log,
		AppVersion,
		AppBuildDate)

	app.Run(startCallback)
}

func startCallback() {
	verifyAuthDB()

	sessionWriter = newWriter(gconfig.Redis)
	sessionReader = service.NewReader(gconfig.Redis)

	app.Echo.POST("/v1/auth/login", httpAuthLogin)
	app.Echo.POST("/v1/auth/logout", httpAuthLogout)
}
