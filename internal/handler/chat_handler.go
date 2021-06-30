package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024, // read/write, count network call
	WriteBufferSize: 1024,
}

//open, then close conn
//read/write - block
//1 goproutine - read budffer, 2 goroutine write buffer
//reusable buffer

func (c *models.Chat) Run() {
	for {
		select {
		case user := <-c.Join:
			c.add(user)
		}
	}
}

func (chat *models.Chat) ChatService(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	log.Fatal(err)

	keys := r.URL.Query()

	username := keys.Get("username")
	if strings.TrimSpace(username) == "" {
		username = "Usheq" // random name or send error -> empty field
	}

	chat.Join <- &models.User{
		Username: username,
		Conn:     conn,
		Global:   chat,
	}
}

// func (c *models.Chat) add(user *models.User) {
// 	log.Println(c)
// }

//add username : c.Users[Users.username]

func (h *Handler) ChatHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		// ChatService(w,r)
		chat := &models.Chat{
			Users:    make(map[string]*models.User),
			Messages: make(chan *models.Message),
			Join:     make(chan *models.User),
			Leave:    make(chan *models.User),
		}

		// http.HandleFunc("api/chat", chat.ChatService(w, r))
		chat.ChatService(w, r)
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
}

// ednpoint - api/caht/listuser; api/chat/historyuser/; api/chat/message
