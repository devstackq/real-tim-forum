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
var chat = &models.Chat{
	ListUsers:   make(map[string]*models.User, 100),
	NewMessage:  make(chan *models.Message), // 1 time - 10 user can write
	GetUsers:    make(chan *models.User),
	Join:        make(chan *models.User),
	Leave:       make(chan *models.User),
	ListMessage: make(chan *models.Message),
}

func (h *Handler) ChatHandler(w http.ResponseWriter, r *http.Request) {

	// write by channel new user ->   &chat.Join <- newUser
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	log.Println(conn.RemoteAddr(), conn.Subprotocol(), "conn")

	go h.Services.Chat.Run(chat)

	err = h.Services.Chat.ChatBerserker(conn, chat, Authorized.Name)
	if err != nil {
		log.Println(err)
		return
	}
	//goroutine Run, action -> with channels(user state, leave, joikn, message) -> call concrete Method
	//listen event by channel -> select case : Join , Message, Leave
}

func (h *Handler) GetMessages(w http.ResponseWriter, r *http.Request) {
}
