package internal

import (
	"flag"
	"gitlab.com/projetAPI/ProjetAPI/service"
	"os"
	"testing"
	"gitlab.com/projetAPI/ProjetAPI/db"
	_ "github.com/lib/pq"
	"gitlab.com/projetAPI/ProjetAPI/mock"
)

var (
	configFile = ""
)

func init() {
	flag.StringVar(&configFile, "config", "", "Configuration file")
}

// Init TU
func TestMain(m *testing.M) {
	flag.Parse()
	configLoaded := gconfig.load(configFile)

	if !configLoaded {
		gconfig.UsersDB.URL = "host=127.0.0.1 dbname=postgres user=postgres password=example sslmode=disable"
	}

	app = service.New(
		"glizou-user-TEST",
		&gconfig.HTTP,
		&gconfig.Log,
		AppVersion,
		AppBuildDate)

	verifUserDB := false
	gUserDB, verifUserDB = db.VerifyUserDB(gUserDB, app.Log, &gconfig.UsersDB)
	if !verifUserDB {
		app.Log.Fatalf("Critical server error. Can't connect to user database")
	}

	gUserDB.ClearTable()

	insertUser()

	code := m.Run()

	gUserDB.ClearTable()

	os.Exit(code)
}

func insertUser() {
	gUserDB.ExecSql(mock.InsertUserTest)
}
