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
		"glizou-user",
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
	verifyUserDB()

	sessionReader = auth.NewReader(gconfig.Redis)

	app.Echo.POST("/v1/user/register", httpRegister)
	//app.Echo.GET("/v1/user/user/:uuid", httpGetUser)
	//app.Echo.PUT("/v1/user/user/:uuid", httpModifyUser)
}
