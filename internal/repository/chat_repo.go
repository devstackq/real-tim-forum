package repository

import (
	"database/sql"

	"github.com/devstackq/real-time-forum/internal/models"
)

type ChatRepository struct {
	db *sql.DB
}

func NewChatRepository(db *sql.DB) *ChatRepository {
	return &ChatRepository{db}
}

//add message in db
func (cr *ChatRepository) AddMessage(m *models.Message) error {
	return nil
}

func (cr *ChatRepository) GetUserName(uid int) (string, error) {

	var name string
	query := `SELECT full_name FROM users WHERE id=?`
	row := cr.db.QueryRow(query, uid)
	err := row.Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func (cr *ChatRepository) GetMessages(m *models.Message) ([]models.Message, error) {

	var messages = []models.Message{}
	var room string
	msg := models.Message{}

	senderUserID, err := cr.GetUserID(m.Sender)
	if err != nil {
		return nil, err
	}
	receiverUserID, err := cr.GetUserID(m.Receiver)
	if err != nil {
		return nil, err
	}
	var row *sql.Row

	row = cr.db.QueryRow("SELECT room FROM chats WHERE user_id1=? AND user_id2=?", senderUserID, receiverUserID)
	err = row.Scan(&room)
	//2,1 -> swap query ?
	if err != nil {
		row = cr.db.QueryRow("SELECT room FROM chats WHERE user_id1=? AND user_id2=?", receiverUserID, senderUserID)
		err = row.Scan(&room)
		if err != nil {
			return nil, err
		}
	}

	queryStmt, err := cr.db.Query("SELECT content, user_id FROM messages  WHERE room=?", room)
	// queryStmt, err := cr.db.Query("SELECT content, user_id FROM messages  WHERE room=? ORDER BY  message_sent_time DESC", room)
	if err != nil {
		return nil, err
	}

	for queryStmt.Next() {
		//or when create msg - set name ?
		nameSender, err := cr.GetUserName(senderUserID)
		if err != nil {
			return nil, err
		}
		nameReceiver, err := cr.GetUserName(receiverUserID)
		if err != nil {
			return nil, err
		}

		msg.Sender = nameSender
		msg.Receiver = nameReceiver

		if err := queryStmt.Scan(&msg.Content, &msg.UserID); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	// fmt.Println(room, messages)
	return messages, nil
}

//good practice ? or service switch type -> 2 differnet repo ?
func (cr *ChatRepository) ChatManager(m *models.Chat) error {
	//save db message
	return nil
}

//utils repo ? - anotehr service use ? todo:
func (cr *ChatRepository) GetUserID(uuid string) (int, error) {

	var userid int
	query := `SELECT user_id FROM sessions WHERE uuid=?`
	row := cr.db.QueryRow(query, uuid)
	err := row.Scan(&userid)

	if err != nil {
		return 0, err
	}

	return userid, nil
}
