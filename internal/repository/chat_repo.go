package repository

import (
	"database/sql"
	"fmt"

	"github.com/devstackq/real-time-forum/internal/models"
)

type ChatRepository struct {
	db *sql.DB
}

func NewChatRepository(db *sql.DB) *ChatRepository {
	return &ChatRepository{db}
}

func (cr *ChatRepository) AddMessage(m *models.Message) error {
	return nil
}

func (cr *ChatRepository) GetMessages(m *models.Message) ([]models.Message, error) {

	var messages = []models.Message{}
	var room string
	msg := models.Message{}
	//getUserIdByUUID(m.Sender, m.Receiver)
	client send uuidSender, and receiver -> get byUUid - user_id m.Sender

	row := cr.db.QueryRow("SELECT room FROM chats WHERE user_id1=? AND user_id2=?", m.Sender, m.Receiver)
	err := row.Scan(&room)
	if err != nil {
		return nil, err
	}
	queryStmt, err := cr.db.Query("SELECT content, user_id FROM messages  WHERE room=?  ", room)
	if err != nil {
		return nil, err
	}
	for queryStmt.Next() {
		//query getUsername byPostId
		if err := queryStmt.Scan(&msg.Content, &msg.UserID); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	fmt.Println(room, messages)
	return messages, nil
}

//good practice ? or service switch type -> 2 differnet repo ?
func (cr *ChatRepository) ChatManager(m *models.Chat) error {
	//save db message
	return nil
}
