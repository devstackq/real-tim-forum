package handler

import (
	"log"
	"net/http"

	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024, // read/write, count network call
	WriteBufferSize: 1024,
}

//1 create new obj chat - for each new user(conn)
var chat = &models.ChatStorage{
	ListUsers:   make(map[string]*models.Chat),
	NewMessage:  make(chan *models.Message), // 1 time - 10 user can write
	ListMessage: make(chan *models.Message),
	GetUsers:    make(chan *models.Chat),
	Join:        make(chan *models.Chat),
	Leave:       make(chan *models.Chat),
}

func (h *Handler) ChatHandler(w http.ResponseWriter, r *http.Request) {
	// write by channel new user ->   &chat.Join <- newUser
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	log.Println(conn.RemoteAddr(), conn.Subprotocol(), "connss")

	go h.Services.Chat.Run(chat)

	err = h.Services.Chat.ChatBerserker(conn, chat, Authorized.Name, Authorized.UUID)
	if err != nil {
		log.Println(err)
		return
	}
	//goroutine Run, action -> with channels(user state, leave, joikn, message) -> call concrete Method
	//listen event by channel -> select case : Join , Message, Leave
}

func (h *Handler) GetListUsers(w http.ResponseWriter, r *http.Request) {

}
