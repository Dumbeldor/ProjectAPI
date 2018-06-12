package db

import (
	"database/sql"
	"github.com/op/go-logging"
	_ "github.com/lib/pq"
)

type UsersDB struct {
	nativeDB *sql.DB
	log      *logging.Logger
	config   *UsersDBConfig
}

func VerifyUserDB(db *UsersDB, log *logging.Logger, config *UsersDBConfig) (*UsersDB, bool) {
	if db == nil {
		db = NewUserDB(log, config)
		if db == nil {
			return db, false
		}
	}

	return db, db.ValidationQuery()
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

func (db *UsersDB) ClearTable() {
	db.nativeDB.Exec("TRUNCATE users RESTART IDENTITY CASCADE")
}

func (db *UsersDB) ExecSql(sql string) error {
	if _, err := db.nativeDB.Exec(sql); err != nil {
		db.log.Fatal(err)
		return err
	}
	return nil
}

func (db *UsersDB) LoginExists(login string) (bool, error) {
	var exists bool

	err := db.nativeDB.QueryRow(dbQueryLoginExists, login).Scan(&exists)

	if err != nil && err != sql.ErrNoRows {
		db.log.Errorf("%s", err)
		return false, err
	}

	return exists, nil
}

func (db *UsersDB) EmailExists(email string) (bool, error) {
	var exists bool

	err := db.nativeDB.QueryRow(dbQueryEmailExists, email).Scan(&exists)

	if err != nil && err != sql.ErrNoRows {
		db.log.Errorf("%s", err)
		return false, err
	}

	return exists, nil
}

func (db *UsersDB) GetLogin(userID string) (string, error) {
	var login string

	err := db.nativeDB.QueryRow(dbQueryGetLogin, userID).Scan(&login)

	if err != nil && err != sql.ErrNoRows {
		db.log.Errorf("%s", err)
		return "", err
	}

	return login, nil
}

func (db *UsersDB) Register(login string, email string, encodedPassword string, salt1 string, salt2 string) error {
	_, err := db.nativeDB.Exec(dbQueryRegister, login, email, encodedPassword, salt1, salt2)
	return err
}

func (db *UsersDB) UserLock(userID string, lock bool) error {
	_, err := db.nativeDB.Exec(dbQueryUserLock, userID, lock)
	return err
}

func (db *UsersDB) UserDelete(userID string) error {
	_, err := db.nativeDB.Exec(dbQueryUserDelete, userID)
	return err
}
