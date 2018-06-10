package db

import "gitlab.com/projetAPI/ProjetAPI/entity"

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
		if err := rows.Scan(&message.ID, &message.Message, &message.CreationDate, &message.UserSenderID,
			&message.UserReceiverID); err != nil {
				return nil, err
		}

		results = append(results, message)
	}

	return results, nil
}
