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
	//HAVING COUNT(is_read) IS NOT m.user_id=$1

	queryStmt, err := cr.db.Query(`SELECT u.id, u.full_name, m.content, 
	 m.name as lastMessage,  m.sent_time, 
	 SUM(CASE WHEN m.user_id!=u.id AND is_read = 0
		THEN +1
		ELSE 0
		END) as countunread,
		MAX(m.id)
	FROM users u
	left join chats c ON c.user_id1 = $1 AND c.user_id2 = u.id OR c.user_id2 = $1 AND c.user_id1 = u.id
	left join messages m ON m.room = c.room 
	GROUP BY u.id
	ORDER by m.sent_time DESC, u.full_name ASC`, userid)

	if err != nil {
		return nil, err
	}

	chatUsers := []*models.Chat{}
	for queryStmt.Next() {
		c := models.Chat{}
		if err := queryStmt.Scan(&c.ID, &c.UserName, &c.LastMessage, &c.LastMessageSenderName, &c.SentTime, &c.CountUnreadMessage, &c.MesageID); err != nil {
			return nil, err
		}
		c.Time = c.SentTime.Time.Format(time.Stamp)
		chatUsers = append(chatUsers, &c)
	}
	return chatUsers, nil
}

func (cr *ChatRepository) AddNewRoom(m *models.Message) error {
	//save in db message in chat, messages table
	receiverUserID, senderUserID, err := cr.getUsersID(m)
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

func (cr *ChatRepository) GetMessages(m *models.Message) ([]models.Message, int, error) {

	var lastId int
	var messages = []models.Message{}

	msg := models.Message{}
	var id string
	var err error

	if len(m.Receiver) == 1 {
		id = m.Receiver
	} else {
		uid, err := cr.GetUserID(m.Receiver)
		if err != nil {
			log.Println(err, -1, m.Receiver)
			return nil, 0, err
		}
		id = strconv.Itoa(uid)
	}
update -> is-read = state
	log.Println(err, 1)
	//update read states
	queryChat, err := cr.db.Prepare(`UPDATE  messages SET is_read=1 WHERE room=? AND user_id = ?`)
	if err != nil {
		log.Println(err, 2)
		return nil, 0, err
	}

	_, err = queryChat.Exec(m.Room, id)
	log.Println(err, 0)

	queryStmt, err := cr.db.Query("SELECT id, content, user_id, name, sent_time, is_read  FROM messages  WHERE room=? ORDER BY sent_time DESC LIMIT 10 OFFSET?", m.Room, m.Offset)
	if err != nil {
		log.Println(err, 3)
		return nil, 0, err
	}
	for queryStmt.Next() {
		//or when create msg - set name ?
		if err := queryStmt.Scan(&msg.ID, &msg.Content, &msg.UserID, &msg.Name, &msg.Time, &msg.IsRead); err != nil {
			log.Println(err, 4)
			return nil, 0, err
		}

		msg.SentTime = msg.Time.Format(time.Stamp)
		msg.Receiver = m.Receiver
		msg.Sender = m.Sender
		msg.IsRead = m.IsRead
		messages = append(messages, msg)
		lastId = msg.ID
	}
	log.Println(lastId, "last in db")

	return messages, lastId, nil
}

func (cr *ChatRepository) getUsersID(m *models.Message) (int, int, error) {

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
	var room string

	receiverUserID, senderUserID, err := cr.getUsersID(m)
	if err != nil {
		return "", err
	}
	// log.Println(receiverUserID, senderUserID, "after")
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
