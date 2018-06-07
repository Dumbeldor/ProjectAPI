package internal

import (
	"gitlab.com/projetAPI/ProjetAPI/service"
	"gitlab.com/projetAPI/ProjetAPI/db"
)

var app *service.App

var gUserDB *db.UsersDB

// AppVersion application version (from git tag)
var AppVersion = "[unk]"

// AppBuildDate application build date
var AppBuildDate = "[unk]"

var sessionReader service.ReaderInterface

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
	verifUserDB := false
	gUserDB, verifUserDB = db.VerifyUserDB(gUserDB, app.Log, &gconfig.UsersDB)
	if !verifUserDB {
		app.Log.Fatalf("Critical server error. Can't connect to user database")
	}

	sessionReader = service.NewReader(gconfig.Redis)

	app.Echo.POST("/v1/message/create", httpCreateMessage)
	//app.Echo.GET("/v1/user/user/:uuid", httpGetUser)
	//app.Echo.PUT("/v1/user/user/:uuid", httpModifyUser)
}

