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
	dbQueryGetLogin = `SELECT login FROM users WHERE user_id = $1`
	dbQueryUserLock = `UPDATE users SET locked = $2 WHERE user_id = $1`
	dbQueryUserDelete = `DELETE FROM users WHERE user_id = $1`
)

// Message
const (
	dbQueryCreateMessage = `INSERT INTO message(message_id, message, user_sender_id, user_receiver_id)
		VALUES(uuid_generate_v4(), $1, $2, (SELECT user_id FROM users WHERE login = $3))`
	dbQueryGetMessagesForUser = `SELECT message_id, message, message.creation_date, user_sender_id, users.login, user_receiver_id
		FROM message JOIN users ON message.user_sender_id = users.user_id WHERE user_receiver_id = $1`
	dbQueryMessageExists = `SELECT message_exists($1)`
	dbQueryMessageUpdate = `UPDATE message SET message=$1 WHERE message_id=$2`
	dbQueryMessageDelete = `DELETE FROM message WHERE message_id=$1`
	dbQueryMessagePutPermission = `SELECT 1 FROM message WHERE message_id=$1 AND user_sender_id=$2;`
	dbQueryMessageDeletePermission = `SELECT 1 FROM message WHERE message_id=$1 AND (user_sender_id=$2 OR user_receiver_id=$2);`
)
