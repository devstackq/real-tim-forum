package handler

import (
	"log"
	"net/http"

	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/gorilla/websocket"
)

//add users in map
var chat = &models.Chat{
	// Users:   make(map[string]*models.User), // make for flexible size struct
	Users:      make(map[string]*websocket.Conn),
	ListsUsers: make(map[string]string),
	Message:    make(chan *models.Message),
	Join:       make(chan *models.User),
	Leave:      make(chan *models.User),
}

func (h *Handler) GetMessages(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":

	case "POST":
		message, _, _, _, err := GetJsonData(w, r, "message")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		seqMessages, err := h.Services.Chat.GetMessages(message)
		if err != nil {
			JsonResponse(w, r, http.StatusInternalServerError, err)
			return
		}
		JsonResponse(w, r, http.StatusOK, seqMessages)
	}
}

func (h *Handler) AddNewUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		//1 create new obj chat - for each new user(conn)
		// write by channel new user ->   &chat.Join <- newUser
		go h.Services.Chat.Run(chat)
		err := h.Services.Chat.ChatBerserker(w, r, chat, Authorized.UUID, Authorized.Name)
		if err != nil {
			log.Println(err)
			return
		}
		// JsonResponse(w, r, 200, list)
	//goroutine Run, action -> with channels(user state, leave, joikn, message) -> call concrete Method
	//listen event by channel -> select case : Join , Message, Leave
	// go h.Services.Chat.Run(chat)

	case "POST":
		log.Println("post QUERy")
		// 	// chat,_, _, _, err := GetJsonData(w, r, "chat")
		// 	if err != nil {
		// 		http.Error(w, err.Error(), http.StatusBadRequest)
		// 		return
		// 	}
		// 	if chat.Type == "listUser" {
		// 		//get online user & and chated users with current user
		// 		result, err := h.Services.GetListUser(chat.userId)
		// 	}
		// 	if chat.Type == "history" {
		// 		//if no have rromid, create new roomId -> sender, receiver, roomId save db, relation
		// 		//message table, message[id, text, from, who]
		// 		result, err := h.Services.GetHistoryUser(chat.roomId)
		// 	}
		// 	if chat.Type == "sendMessage" {
		// 		result, err := h.Services.sendMessage(chat.From, chat.Who)
		// 		//if no err -> h.Services.showNotify(chat.Who)
		// 	}
		// 	JsonResponse(w, r, http.StatusOK, result)
		// default:
		// 	JsonResponse(w, r, http.StatusBadRequest, "Bad Request")
	}
}

// ednpoint - api/caht/listuser; api/chat/historyuser/; api/chat/message

func (h *Handler) GetListUsers(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		//not show yourself uuid todo:
		// fmt.Println(chat.ListsUsers, 2)
		// users, err := h.Services.Chat.GetListUsers(w, r, chat)
		JsonResponse(w, r, http.StatusOK, chat.ListsUsers)
	case "POST":

	}
}
