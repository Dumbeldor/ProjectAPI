package internal

import (
	_ "github.com/lib/pq" // pq requires blank import
	"gitlab.com/projetAPI/ProjetAPI/db"
)

var gUserDB *db.UsersDB

func verifyUserDB() bool {
	if gUserDB == nil {
		gUserDB = db.NewUserDB(app.Log, &gconfig.UsersDB)
		if gUserDB == nil {
			return false
		}
	}

	return gUserDB.ValidationQuery()
}
