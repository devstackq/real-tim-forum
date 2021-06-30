package models

import (
	"time"

	"github.com/gorilla/websocket"
)

//User struct
type User struct {
	UserID      int             `json:"userid"`
	FullName    string          `json:"fullname"`
	Email       string          `json:"email"`
	Password    string          `json:"password"`
	Age         string          `json:"age"`
	Sex         string          `json:"gender"`
	CreatedTime time.Time       `json:"createdtime"`
	City        string          `json:"city"`
	Image       []byte          `json:"image"`
	ImageHTML   string          `json:"imagehtml"`
	SVG         bool            `json:"svg"`
	Username    string          `json:"username"`
	LastSeen    string          `json:"lastseen"`
	Session     *Session        `json:"session"`
	Conn        *websocket.Conn `json:"conn"`
	Global      *Chat
}

var Any struct {
	Posts Post
	Users User
}
