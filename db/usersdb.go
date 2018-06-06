package db

import (
	"database/sql"
	"github.com/op/go-logging"
)

type UsersDB struct {
	nativeDB *sql.DB
	log *logging.Logger
	config *UsersDBConfig
}

func NewUserDB(log *logging.Logger, config *UsersDBConfig) *UsersDB {
	db := &UsersDB{
		log: log,
	}

	if !db.init(config) {
		return nil
	}

	return db
}

func (db *UsersDB) init(config *UsersDBConfig) bool {
	db.log.Infof("Connecting to users DB at %s", config.URL)
	nativeDB, err := sql.Open("postgres", config.URL)
	if err != nil {
		db.log.Errorf("Failed to connect to content DB: %s", err)
		return false
	}

	db.nativeDB = nativeDB
	if !db.ValidationQuery() {
		db.nativeDB = nil
		return false
	}

	db.nativeDB.SetMaxIdleConns(config.MaxIdleConns)
	db.nativeDB.SetMaxOpenConns(config.MaxOpenConns)

	db.log.Infof("Connected to users DB.")
	return true

}

// ValidationQuery validates the connection pool
func (db *UsersDB) ValidationQuery() bool {
	rows, err := db.nativeDB.Query(ValidationQuery)
	if err != nil {
		db.log.Errorf("Failed to run UsersDB validation query: %s", err)
		return false
	}
	rows.Close()
	return true
}

// Begin starts a transaction
func (db *UsersDB) Begin() (*sql.Tx, error) {
	return db.nativeDB.Begin()
}
