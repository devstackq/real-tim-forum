package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024, // read/write, count network call
	WriteBufferSize: 1024,
}

//add users in map
var chat = &models.Chat{
	// Users:   make(map[string]*models.User), // make for flexible size struct
	Users:       make(map[string]*websocket.Conn),
	ListsUsers:  make(map[string]string),
	NewMessage:  make(chan *models.Message),
	Join:        make(chan *models.User),
	Leave:       make(chan *models.User),
	ListMessage: make(chan *models.Message),
}

func (h *Handler) ChatHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		//1 create new obj chat - for each new user(conn)
		// write by channel new user ->   &chat.Join <- newUser
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(conn.RemoteAddr(), conn.Subprotocol(), "conn")
		go h.Services.Chat.Run(chat)
		err = h.Services.Chat.ChatBerserker(conn, chat, Authorized.Name)
		if err != nil {
			log.Println(err)
			return
		}

		// h.GetListUsers(w, r)

	//goroutine Run, action -> with channels(user state, leave, joikn, message) -> call concrete Method
	//listen event by channel -> select case : Join , Message, Leave
	// go h.Services.Chat.Run(chat)
	case "POST":
		fmt.Println("post quer")
		// message, _, _, _, err := GetJsonData(w, r, "message")
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }
		// seqMessages, err := h.Services.Chat.GetMessages(message)

		// if err != nil {
		// 	log.Println(err)
		// 	JsonResponse(w, r, http.StatusInternalServerError, err)
		// 	return
		// }
		// log.Println("post QUERy")
	}
	//check type, if newuser -> call NewUser service, etc
}

func (h *Handler) GetMessages(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":

	case "POST":

		// message, _, _, _, err := GetJsonData(w, r, "message")
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }
		// seqMessages, err := h.Services.Chat.GetMessages(message)

		// if err != nil {
		// 	log.Println(err)
		// 	JsonResponse(w, r, http.StatusInternalServerError, err)
		// 	return
		// }
		// JsonResponse(w, r, http.StatusOK, seqMessages)
	}
}

//add newuser in map[string]string
func (h *Handler) AddNewUser(w http.ResponseWriter, r *http.Request) {

	//call getListusers ->
	switch r.Method {
	case "GET":
		//1 create new obj chat - for each new user(conn)
		// write by channel new user ->   &chat.Join <- newUser
		// conn, err := upgrader.Upgrade(w, r, nil)
		// if err != nil {
		// 	log.Println(err)
		// 	return
		// }
		//call ds
		go h.Services.Chat.Run(chat)
		// err := h.Services.Chat.ChatBerserker(w, r, chat, Authorized.UUID, Authorized.Name)
		// if err != nil {
		// 	log.Println(err)
		// 	return
		// }
		h.GetListUsers(w, r)
	//goroutine Run, action -> with channels(user state, leave, joikn, message) -> call concrete Method
	//listen event by channel -> select case : Join , Message, Leave
	// go h.Services.Chat.Run(chat)
	case "POST":
		log.Println("post QUERy")
	}
}

// ednpoint - api/caht/listuser; api/chat/historyuser/; api/chat/message

//return arr obj name:uuid
func (h *Handler) GetListUsers(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		//not show yourself uuid todo:
		// users, err := h.Services.Chat.GetListUsers(w, r, chat) ?
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		err = conn.WriteJSON(chat.ListsUsers)
		if err != nil {
			log.Println(err)
			return
		}
	case "POST":
	}
}
