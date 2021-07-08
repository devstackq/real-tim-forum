package models

import (
	"time"

	"github.com/gorilla/websocket"
)

//chatID99 -> from, who chatid JOIN
//query select * from messages where uiser_id1 => and user_id2=?
//59, 19 ->
//hi; how r u? uid1
//- hello, fine and u ? uid2
type Message struct {
	ID       int    `json:"id"`
	Content  string `json:"content"`
	ChatID   int    `json:"chatid"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Type     string `json:"type"`
	UserID   int    `json:"userid"`
	Conn     *websocket.Conn
	Room     string    `json:"room"`
	Name     string    `json:"name"`
	SentTime time.Time `json:"senttime"`
}

// func NewMessage(body string, sender int) *Message {
// 	return &Message{
// 		Body:   body,
// 		Sender: sender,
// 	}
// }

type Chat struct {
	// Users      map[string]*websocket.Conn `json:"users"`
	// ListsUsers map[string]string
	ListUsers   map[string]*User `json:"uzers"` // key =
	NewMessage  chan *Message    `json:"newmessage"`
	Join        chan *User
	Leave       chan *User
	ListMessage chan *Message `json:"listmessages"`
	GetUsers    chan *User
}

//user19 -> send msg -> user 59, from, who, chatid99
