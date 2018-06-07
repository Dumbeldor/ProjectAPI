package internal

import (
	"flag"
	"testing"
	"gitlab.com/projetAPI/ProjetAPI/service"
	"gitlab.com/projetAPI/ProjetAPI/db"
	_ "github.com/lib/pq"
	"os"
	"gitlab.com/projetAPI/ProjetAPI/mock"
)

var (
	configFile = ""
	sessionWriter service.WriterInterface
)

func init() {
	flag.StringVar(&configFile, "config", "", "Configuration file")
}

func TestMain(m *testing.M) {
	flag.Parse()
	configLoaded := gconfig.load(configFile)

	if !configLoaded {
		gconfig.UsersDB.URL = "host=127.0.0.1 dbname=postgres user=postgres password=example sslmode=disable"
	}

	app = service.New(
		"glizou-message-TEST",
		&gconfig.HTTP,
		&gconfig.Log,
		AppVersion,
		AppBuildDate)

	verifUserDB := false
	gUserDB, verifUserDB = db.VerifyUserDB(gUserDB, app.Log, &gconfig.UsersDB)
	if !verifUserDB {
		app.Log.Fatalf("Critical server error. Can't connect to user database")
	}

	createSessionWriter()

	createSessionReader()

	mock.InsertSession(sessionWriter, app.Log)

	code := m.Run()

	os.Exit(code)
}

func createSessionWriter() {
	sessionWriter = mock.NewSessionMock()
}

func createSessionReader() {
	sessionReader = mock.NewSessionMock()
}
