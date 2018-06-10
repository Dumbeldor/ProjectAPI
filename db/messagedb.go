package db

func (db *UsersDB) CreateMessage(msg string, userSender string, userLoginReceiver string) error {
	_, err := db.nativeDB.Exec(dbQueryCreateMessage, msg, userSender, userLoginReceiver)
	return err
}
