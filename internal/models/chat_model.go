package models

import (
	"database/sql"
	"time"

	"github.com/gorilla/websocket"
)

type Message struct {
	ID           int    `json:"id"`
	Content      string `json:"content"`
	ChatID       int    `json:"chatid"`
	Sender       string `json:"sender"`
	Receiver     string `json:"receiver"`
	Type         string `json:"type"`
	UserID       int    `json:"userid"`
	Conn         *websocket.Conn
	Room         string    `json:"room"`
	Name         string    `json:"sendername"`
	ReceiverName string    `json:"receivername"`
	SentTime     string    `json:"senttime"`
	Time         time.Time `json:"time"`
}

// func NewMessage(body string, sender int) *Message {
// 	return &Message{
// 		Body:   body,
// 		Sender: sender,
// 	}
// }

type ChannelStorage struct {
	OnlineUsers map[string]*Chat
	NewMessage  chan *Message
	Join        chan *Chat
	Leave       chan *Chat
	ListMessage chan *Message
	GetUsers    chan *Chat
	NewUser     chan *Chat
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
	Uzers                 []*Chat         `json:"uzers"`
}
