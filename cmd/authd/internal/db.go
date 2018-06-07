package internal

import (
	"database/sql"
	_ "github.com/lib/pq" // pq requires blank import
)

type authdb struct {
	nativeDB *sql.DB
}

var gAuthDB *authdb

func newAuthDB() *authdb {
	db := &authdb{}

	if !db.init() {
		return nil
	}

	return db
}

func verifyAuthDB() bool {
	if gAuthDB == nil {
		gAuthDB = newAuthDB()
		// init failed, abort event reading
		if gAuthDB == nil {
			return false
		}
	}

	return gAuthDB.validationQuery()
}

func (db *authdb) init() bool {
	app.Log.Infof("Connecting to auth DB at %s", gconfig.Postgresql.Connstr)
	nativeDB, err := sql.Open("postgres", gconfig.Postgresql.Connstr)
	if err != nil {
		app.Log.Errorf("Failed to connect to auth DB: %s", err)
		return false
	}

	db.nativeDB = nativeDB
	if !db.validationQuery() {
		db.nativeDB = nil
		return false
	}

	app.Log.Infof("Connected to auth DB.")
	return true
}

func (db *authdb) validationQuery() bool {
	rows, err := db.nativeDB.Query(dbQueryValidation)
	if err != nil {
		app.Log.Errorf("Failed to run authdb validation query: %s", err)
		return false
	}
	rows.Close()
	return true
}

func (db *authdb) getLoginInfos(login string) (*loginInfos, error) {
	linfo := &loginInfos{Login: login}
	err := db.nativeDB.QueryRow(dbQueryGetUserInfosByLogin, login).
		Scan(&linfo.UserID, &linfo.Password, &linfo.Locked, &linfo.Salt1, &linfo.Salt2)
	if err != nil {
		app.Log.Errorf("%s", err)
		return nil, err
	}

	return linfo, nil
}
