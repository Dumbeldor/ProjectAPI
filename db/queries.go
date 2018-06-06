package db

const (
	ValidationQuery = `SELECT 1`
)

// Users DB
const (
	dbQueryLoginExists = `SELECT login_exists($1)`
	dbQueryEmailExists = `SELECT email_exists($1)`
	dbQueryRegister    = `INSERT INTO users(user_id, login, email, password, salt1, salt2)
		VALUES(uuid_generate_v4(), $1, $2, $3, $4, $5)`
)
