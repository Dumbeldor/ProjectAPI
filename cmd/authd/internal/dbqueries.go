package internal

const (
	dbQueryValidation = `SELECT 1`
)

const (
	dbQueryGetUserInfosByLogin = `SELECT user_id, password, locked, salt1, salt2 FROM users WHERE login = $1`
)
