package db

func (db *UsersDB) CreateMessage(msg string, userSender string, userReceiver string) error {
	_, err := db.nativeDB.Exec(dbQueryCreateMessage, msg, userSender, userReceiver)
	return err
}
