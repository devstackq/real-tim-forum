package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devstackq/real-time-forum/internal/models"
)

//add username : c.Users[Users.username]

func (h *Handler) ChatHandler(w http.ResponseWriter, r *http.Request) {

	// switch r.Method {
	// case "GET":

	//1 create new obj chat - for each new user
	chat := &models.Chat{
		Users:   make(map[string]*models.User), // make for flexible size struct
		Message: make(chan *models.Message),
		Join:    make(chan *models.User),
		Leave:   make(chan *models.User),
	}
	fmt.Println(r.Method, "methot type", chat)
	//hanler each new conn
	//get username, conn, user.global : *chat
	// write by channel new user ->   &chat.Join <- newUser
	go h.Services.Chat.Run(chat)

	err := h.Services.Chat.ChatBerserker(w, r, chat, Authorized.UUID)
	if err != nil {
		log.Fatal(err)
	}

	//goroutine Run, action -> with channels(user state, leave, joikn, message) -> call concrete Method
	//listen event by channel -> select case : Join , Message, Leave
	// go h.Services.Chat.Run(chat)

	// case "POST":
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

// ednpoint - api/caht/listuser; api/chat/historyuser/; api/chat/message
