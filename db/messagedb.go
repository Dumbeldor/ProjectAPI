package db

import (
	"gitlab.com/projetAPI/ProjetAPI/entity"
	"database/sql"
)

func (db *UsersDB) CreateMessage(msg string, userSender string, userLoginReceiver string) error {
	_, err := db.nativeDB.Exec(dbQueryCreateMessage, msg, userSender, userLoginReceiver)
	return err
}

func (db *UsersDB) GetMessageForUser(userID string) ([]entity.Message, error) {
	rows, err := db.nativeDB.Query(dbQueryGetMessagesForUser, userID)
	if rows != nil {
		defer rows.Close()
	}

	if err != nil {
		return nil, err
	}

	var results []entity.Message
	for rows.Next() {
		message := entity.Message{}
		if err := rows.Scan(&message.ID, &message.Message, &message.CreationDate, &message.UserSenderID, &message.UserSenderName,
			&message.UserReceiverID); err != nil {
				return nil, err
		}

		results = append(results, message)
	}

	return results, nil
}

func (db *UsersDB) MessageExists(messageID string) (bool, error) {
	var exists bool

	err := db.nativeDB.QueryRow(dbQueryMessageExists, messageID).Scan(&exists)

	if err != nil && err != sql.ErrNoRows {
		db.log.Errorf("%s", err)
		return false, err
	}

	return exists, nil
}

func (db *UsersDB) MessageUpdate(messageID string, message string) error {
	_, err := db.nativeDB.Exec(dbQueryMessageUpdate, message, messageID)
	if err != nil {
		db.log.Errorf("MessageUpdate error: %s", err)
	}

	return err
}

func (db *UsersDB) MessageDelete(messageID string) error {
	_, err := db.nativeDB.Exec(dbQueryMessageDelete, messageID)
	if err != nil {
		db.log.Errorf("MessageDelete error: %s", err)
	}

	return err
}

func (db *UsersDB) MessagePutPermission(messageID string, userID string) (bool, error) {
	var permission bool

	err := db.nativeDB.QueryRow(dbQueryMessagePutPermission, messageID, userID).Scan(&permission)

	if err != nil && err != sql.ErrNoRows {
		db.log.Errorf("%s", err)
		return false, err
	}

	return permission, nil
}

func (db *UsersDB) MessageDeletePermission(messageID string, userID string) (bool, error) {
	var permission bool

	err := db.nativeDB.QueryRow(dbQueryMessageDeletePermission, messageID, userID).Scan(&permission)

	if err != nil && err != sql.ErrNoRows {
		db.log.Errorf("%s", err)
		return false, err
	}

	return permission, nil
}
