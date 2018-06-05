package internal

import (
	"gitlab.com/projetAPI/auth"
	"gitlab.com/projetAPI/ProjetAPI/service"
)

var app *service.App

// AppVersion application version (from git tag)
var AppVersion = "[unk]"

// AppBuildDate application build date
var AppBuildDate = "[unk]"

var sessionReader auth.ReaderInterface

// StartApp initiate components
// Should be called from main function
func StartApp(configFile string) {
	gconfig.load(configFile)

	app = service.New(
		"uc-profile",
		&gconfig.HTTP,
		&gconfig.Log,
		AppVersion,
		AppBuildDate)

	app.OnSigHUP("config-reload", func() {
		gconfig.load(configFile)
	})

	app.Run(startCallback)
}

func startCallback() {
	verifyProfileDB()

	sessionReader = auth.NewReader(gconfig.Redis)

	app.Echo.POST("/v1/profile/register", httpRegister)
	app.Echo.GET("/v1/profile/user/:uuid", httpGetUser)
	app.Echo.PUT("/v1/profile/user/:uuid", httpModifyUser)
}
