package repository

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/devstackq/real-time-forum/internal/models"
)

type ChatRepository struct {
	db *sql.DB
}

func NewChatRepository(db *sql.DB) *ChatRepository {
	return &ChatRepository{db}
}

func (cr *ChatRepository) GetSortedUsers(userid int) ([]*models.Chat, error) {
	//save in db message in chat, c.room,  m.user_id, m.id,, WHERE u.id NOT IN($1)
	queryStmt, err := cr.db.Query(`SELECT u.id, u.full_name, m.content,  m.name as lastMessageSenderName,  m.sent_time, MAX(m.id) FROM users u  
	left join chats c ON c.user_id1 = $1  AND c.user_id2 = u.id  OR  c.user_id2 = $1 AND c.user_id1 = u.id
	left join messages m ON m.room = c.room   GROUP BY u.id ORDER by m.sent_time DESC, u.full_name ASC`, userid)
	if err != nil {
		return nil, err
	}

	chatUsers := []*models.Chat{}
	for queryStmt.Next() {
		c := models.Chat{}
		if err := queryStmt.Scan(&c.ID, &c.UserName, &c.LastMessage, &c.LastMessageSenderName, &c.SentTime, &c.MesageID); err != nil {
			return nil, err
		}
		chatUsers = append(chatUsers, &c)
	}
	return chatUsers, nil
}

func (cr *ChatRepository) AddNewRoom(m *models.Message) error {
	//save in db message in chat, messages table

	receiverUserID, senderUserID, err := cr.getUserIDs(m)
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
	//save in db message in chat, messages table
	queryMsg, err := cr.db.Prepare(`INSERT INTO messages(content, room, user_id, name, sent_time) VALUES(?,?,?, ?, ?)`)
	if err != nil {
		log.Println(err, 1)
		return err
	}
	_, err = queryMsg.Exec(m.Content, m.Room, m.UserID, m.Name, time.Now())
	if err != nil {
		log.Println(err, 2)
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
		if err := queryStmt.Scan(&msg.Content, &msg.UserID, &msg.Name, &msg.Time); err != nil {
			return nil, err
		}
		msg.SentTime = msg.Time.Format(time.RFC3339)
		msg.Receiver = m.Receiver
		msg.Sender = m.Sender
		messages = append(messages, msg)
	}
	return messages, nil
}

func (cr *ChatRepository) GetAllUsers() ([]models.User, error) {

	users := []models.User{}
	row, err := cr.db.Query("SELECT id, full_name FROM users")
	if err != nil {
		return nil, err
	}
	for row.Next() {
		user := models.User{}
		if err := row.Scan(&user.UserID, &user.FullName); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (cr *ChatRepository) GetLastMessageIndex(room string, userid int) (lastindex int, err error) {

	//	row, err := cr.db.Query("SELECT messages.id FROM messages LEFT JOIN chats ON chats.room=messages.room WHERE user_id=? AND messages.room=? ORDER BY sent_time ASC", userid, room)
	row, err := cr.db.Query("SELECT id FROM messages WHERE  room=?", room)
	if err != nil {
		return 0, err
	}
	for row.Next() {
		if err := row.Scan(&lastindex); err != nil {
			return 0, err
		}
	}
	// log.Println("last indx", lastindex)
	return lastindex, nil
}

func (cr *ChatRepository) getUserIDs(m *models.Message) (int, int, error) {

	senderUserID := m.ID
	receiverUserID := m.UserID

	var err error

	if len(m.Receiver) == 1 {
		receiverUserID, err = strconv.Atoi(m.Receiver)
		if err != nil {
			return 0, 0, err
		}
		senderUserID, err = cr.GetUserID(m.Sender)
		if err != nil {
			return 0, 0, err
		}
	} else if (senderUserID == 0 || receiverUserID == 0) && len(m.Receiver) == 36 {

		receiverUserID, err = cr.GetUserID(m.Receiver)
		if err != nil {
			return 0, 0, err
		}
		senderUserID, err = cr.GetUserID(m.Sender)
		if err != nil {
			return 0, 0, err
		}
	}

	return receiverUserID, senderUserID, nil

}

func (cr *ChatRepository) IsExistRoom(m *models.Message) (string, error) {

	receiverUserID, senderUserID, err := cr.getUserIDs(m)
	if err != nil {
		return "", err
	}

	log.Println(receiverUserID, senderUserID, "after")

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
