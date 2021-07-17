package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/devstackq/real-time-forum/internal/models"
)

type ChatRepository struct {
	db *sql.DB
}

func NewChatRepository(db *sql.DB) *ChatRepository {
	return &ChatRepository{db}
}

func (cr *ChatRepository) AddNewRoom(m *models.Message) error {
	//save in db message in chat, messages table
	senderUserID, err := cr.GetUserID(m.Sender)
	if err != nil {
		return err
	}
	receiverUserID, err := cr.GetUserID(m.Receiver)
	if err != nil {
		return err
	}

	queryChat, err := cr.db.Prepare(`INSERT INTO chats(user_id1, user_id2, room) VALUES(?,?,?)`)
	if err != nil {
		return err
	}
	_, err = queryChat.Exec(senderUserID, receiverUserID, m.Room)
	if err != nil {
		return err
	}

	defer queryChat.Close()
	return nil
}

//addnewMessage - save name, no need fix
//add message in db
func (cr *ChatRepository) AddNewMessage(m *models.Message) error {
	log.Println(m.Name, "name, room", m.Room)
	//save in db message in chat, messages table
	queryMsg, err := cr.db.Prepare(`INSERT INTO messages(content, room, user_id, name, sent_time) VALUES(?,?,?, ?, ?)`)
	if err != nil {
		return err
	}
	_, err = queryMsg.Exec(m.Content, m.Room, m.UserID, m.Name, time.Now())
	if err != nil {
		return err
	}
	defer queryMsg.Close()
	return nil
}

func (ur *ChatRepository) GetUserName(id int) (string, error) {

	var name string
	query := `SELECT full_name FROM users WHERE id=?`
	row := ur.db.QueryRow(query, id)
	err := row.Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func (cr *ChatRepository) GetMessages(m *models.Message) ([]models.Message, error) {

	var messages = []models.Message{}
	msg := models.Message{}

	queryStmt, err := cr.db.Query("SELECT content, user_id, name, sent_time FROM messages  WHERE room=?", m.Room)
	// queryStmt, err := cr.db.Query("SELECT content, user_id FROM messages  WHERE room=? ORDER BY  message_sent_time DESC", room)
	if err != nil {
		return nil, err
	}

	for queryStmt.Next() {
		//or when create msg - set name ?
		if err := queryStmt.Scan(&msg.Content, &msg.UserID, &msg.Name, &msg.SentTime); err != nil {
			return nil, err
		}

		msg.Receiver = m.Receiver
		msg.Sender = m.Sender
		messages = append(messages, msg)
	}
	return messages, nil
}

func (cr *ChatRepository) GetLastMessageIndex(room string, userid int) (lastindex int, err error) {

	row, err := cr.db.Query("SELECT id FROM messages  WHERE room=? AND  user_id=?", room, userid)
	if err != nil {
		return 0, err
	}
	for row.Next() {
		if err := row.Scan(&lastindex); err != nil {
			return 0, err
		}
	}

	return lastindex, nil
}

func (cr *ChatRepository) IsExistRoom(m *models.Message) (string, error) {

	senderUserID, err := cr.GetUserID(m.Sender)
	if err != nil {
		return "", err
	}
	receiverUserID, err := cr.GetUserID(m.Receiver)
	if err != nil {
		return "", err
	}
	var room string

	row := cr.db.QueryRow("SELECT room FROM chats WHERE user_id1=? AND user_id2=?", senderUserID, receiverUserID)
	err = row.Scan(&room)
	//2,1 -> swap query ?
	if err != nil {
		row = cr.db.QueryRow("SELECT room FROM chats WHERE user_id1=? AND user_id2=?", receiverUserID, senderUserID)
		err = row.Scan(&room)
		if err != nil {
			return "", err
		}
	}
	return room, nil
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
