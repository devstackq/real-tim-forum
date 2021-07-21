package models

import (
	"database/sql"
	"time"

	"github.com/gorilla/websocket"
)

//chatID99 -> from, who chatid JOIN
//query select * from messages where uiser_id1 => and user_id2=?
//59, 19 ->
//hi; how r u? uid1
//- hello, fine and u ? uid2
type Message struct {
	ID               int    `json:"id"`
	Content          string `json:"content"`
	ChatID           int    `json:"chatid"`
	Sender           string `json:"sender"`
	Receiver         string `json:"receiver"`
	Type             string `json:"type"`
	UserID           int    `json:"userid"`
	Conn             *websocket.Conn
	Room             string    `json:"room"`
	Name             string    `json:"aname"`
	SentTime         time.Time `json:"senttime"`
	LastIndexMessage int
	Indexs           []int
}

// 20 time,

// func NewMessage(body string, sender int) *Message {
// 	return &Message{
// 		Body:   body,
// 		Sender: sender,
// 	}
// }

type ChatStorage struct {
	OnlineUsers map[string]*Chat
	NewMessage  chan *Message
	Join        chan *Chat
	Leave       chan *Chat
	ListMessage chan *Message
	GetUsers    chan *Chat
}

type Chat struct {
	ID                    int             `json:"id"`
	UserName              string          `json:"fullname"`
	LastMessage           sql.NullString  `json:"lastmessage"`
	LastMessageSenderName sql.NullString  `json:"lastsender"`
	SentTime              sql.NullTime    `json:"senttime"`
	MesageID              sql.NullInt64   `json:"messageid"`
	Online                bool            `json:"online"`
	UUID                  string          `json:"uuid"`
	Conn                  *websocket.Conn `json:"conn"`
}

//user19 -> send msg -> user 59, from, who, chatid99
