package service

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/devstackq/real-time-forum/internal/repository"
	"github.com/gorilla/websocket"
)

var ChatStorage struct {
	ListMessage []models.Message `json:"messages"`
	// OnlineUsers   map[string]*models.User `json:"users"`
	OnlineUsers map[string]*models.Chat `json:"online"`
	AllUsers    map[string]*models.Chat `json:"users"`
	Message     models.Message          `json:"message"`
	Type        string                  `json:"type"`
	Receiver    string                  `json:"receiver"`
}

type ChatService struct {
	repository repository.Chat
}

func NewChatService(repo repository.Chat) *ChatService {
	return &ChatService{repo}
}

func (cs *ChatService) getMessages(m *models.Message, c *models.ChatStorage) error {
	//find users room, if zero
	room, err := cs.repository.IsExistRoom(m)
	if err != nil {
		log.Println(err, "empty room getMsg()")
	}
	ChatStorage.Receiver = m.Receiver
	if room == "" {
		ChatStorage.Type = "nomessages"
		m.Conn.WriteJSON(ChatStorage)
		return nil
	}
	m.Room = room
	log.Println(m, err)
	seq, err := cs.repository.GetMessages(m)
	if err != nil {
		log.Println(err, "get msg err")
	}
	ChatStorage.Type = "listmessages"
	ChatStorage.ListMessage = seq
	err = m.Conn.WriteJSON(ChatStorage)
	if err != nil {
		return err
	}
	return nil
}

//map[string]*models.Chat
func (cs *ChatService) mergeUsers(dbUsers []models.Chat, onlineUsers map[string]*models.Chat) {
	log.Println(ChatStorage.OnlineUsers, 0)
	//1 init - 1 user then add in CS.OnlineUsers
	//2 userX sigin -> update OnlineUsers or /chat js -> getusers -> already preapared data work
	//case offline -> online, [14]user,
	//[qwerty]=user,
	//11 in 4,
	//add new user, append -> in dbUsers, updatefield - send client

	//1 get sorted user, 2 getOnlineUser -> add field -> and online = true -> send client obj
	//sortusers Cs.ChatUsers
	//join or leave -> change state -> ChatUsers

	sort, condition, append in dbUsers, then - send AllUsers -> send client
	for uuid, onlineUser := range onlineUsers { //server users
		for _, dbUser := range dbUsers { //sorted users from db
			log.Println(ChatStorage.OnlineUsers[strconv.Itoa(dbUser.ID)], "key map", strconv.Itoa(dbUser.ID))
			//2 strcut - merge 1 ?
			if _, ok := ChatStorage.OnlineUsers[uuid]; !ok && len(uuid) == 36 {
				if dbUser.ID == onlineUser.ID {
					onlineUser.UserName = dbUser.UserName
					onlineUser.LastMessage = dbUser.LastMessage
					onlineUser.Online = true
					ChatStorage.OnlineUsers[uuid] = onlineUser
				}
			}
			if _, ok := ChatStorage.OnlineUsers[strconv.Itoa(dbUser.ID)]; !ok && dbUser.ID != onlineUser.ID {
				// if ChatStorage.OnlineUsers[strconv.Itoa(user.ID)] == nil {
				// log.Println(user.ID, "add offline users", user.UserName)
				c := models.Chat{}
				c.ID = dbUser.ID
				c.UserName = dbUser.UserName
				c.LastMessage = dbUser.LastMessage
				c.Online = false
				ChatStorage.OnlineUsers[strconv.Itoa(c.ID)] = &c
			}
		}
	}
}

//if find userid -> update CS.OnlineUsers[id]=u, uuid, u.LastMessage, etc
func (cs *ChatService) updateStateUser(user *models.Chat) {

}

//or global each time send client - use client side this var ?``
//open, then close conn //read/write - block //1 goproutine - read budffer, 2 goroutine write buffer, reusable buffer
func (cs *ChatService) addNewUser(u *models.Chat, c *models.ChatStorage) {
	//fill user.Name -> in db, by uuid  u.Conn
	ChatStorage.Type = "observeusers"
	// log.Println("add user prev", u.UUID, u.ID)

	//case - relogin, delete prev user in map, no duplicate
	if len(c.OnlineUsers) > 1 {
		for k := range c.OnlineUsers {
			_, err := cs.repository.GetUserID(k)
			if err != nil {
				delete(c.OnlineUsers, k)
			}
		}
	}
	//not find by userid, exit
	//
	// if updated, user := cs.updateStateUser(u); !updated {

	c.OnlineUsers[u.UUID] = u
	ChatStorage.OnlineUsers = c.OnlineUsers // get gloabal var, map[uuid]name

	//off -> on, findById, -> update data

	sorted, err := cs.repository.GetSortedUsers(u.ID)
	if err != nil {
		log.Println(err)
	}
	go cs.mergeUsers(sorted, c.OnlineUsers)

	log.Println(len(ChatStorage.OnlineUsers), "res", ChatStorage.OnlineUsers)

	// for _, v := range ChatStorage.OnlineUsers {
	// 	v.Conn.WriteJSON(ChatStorage)
	// }
	// }
}

func (cs *ChatService) getUsers(u *models.Chat) {
	//check if users have in session
	ChatStorage.Type = "getusers"
	// cs.onlineUsers(sorted)
	u.Conn.WriteJSON(ChatStorage)
}

func (cs *ChatService) sendMessage(c *models.ChatStorage, m *models.Message) {
	// if c.Users[m.Receiver] == uuidReceiverInDB?
	//save db in message, caht table & message, add author, date message
	//send msg - to  conn - receiver if have in server
	room, err := cs.repository.IsExistRoom(m)
	if err != nil {
		log.Println(err, " room New msg3")
	}
	randomRoom := Randomaizer()
	// fmt.Println(room, "reciece", m.Receiver, "sen", m.Sender, "pre add message")

	if room == "" {
		m.Room = randomRoom
		log.Println(m.Room, "send new room in db func")
		err = cs.repository.AddNewRoom(m)
		if err != nil {
			log.Println(err, "add room err")
		}
	} else {
		m.Room = room
	}
	// log.Println(m.Sender, m.Receiver)
	err = cs.repository.AddNewMessage(m)
	if err != nil {
		log.Println(err, "err add new msg")
	}
	if c.OnlineUsers[m.Receiver] != nil {
		receiver := c.OnlineUsers[m.Receiver]
		ChatStorage.Message.Content = m.Content
		ChatStorage.Message.SentTime = time.Now()
		ChatStorage.Message.Name = m.Name
		ChatStorage.Type = "lastmessage"
		ChatStorage.Message.Receiver = m.Sender

		err = receiver.Conn.WriteJSON(ChatStorage)
		if err != nil {
			log.Println(err, "err json add new msg")
		}
		//notify user, & send message in WriteMessage
		// defer rec.Close()
	}
}

func (cs *ChatService) leaveUser(c *models.ChatStorage, u *models.Chat) {
	ChatStorage.Type = "leave"
	u.Conn.Close()
	delete(c.OnlineUsers, u.UUID)
	ChatStorage.OnlineUsers = c.OnlineUsers
	for _, v := range ChatStorage.OnlineUsers {
		v.Conn.WriteJSON(ChatStorage)
	}
	// log.Println("user leave", ChatStorage.OnlineUsers)
}

//1 main -> Start() ->  createEmptyObjecetChat -> 2 ws Handler, newConn -> 3 go Run() // goruutine each newConn(user)
//handle if receive new user (from new conn(user) -> join), get new user -> by chan(goroutine)
func (cs *ChatService) Run(c *models.ChatStorage) {
	for {
		select {
		case newuser := <-c.Join:
			cs.addNewUser(newuser, c)
		case message := <-c.NewMessage:
			cs.sendMessage(c, message)
		case list := <-c.ListMessage:
			cs.getMessages(list, c)
		case users := <-c.GetUsers:
			cs.getUsers(users)
		case user := <-c.Leave:
			cs.leaveUser(c, user)
		}
		// default :

	}
}

//out handler ?
// Handler -> each new client connect server, create handshake - save data - send service
//Service - check type from client -> if newmessage -> create obj Message, fill field -> run gorutine
//run func broadcast(), conn,WriteJson(data), find by uuid

func (cs *ChatService) ChatBerserker(conn *websocket.Conn, c *models.ChatStorage, name string, uuid string) error {

	body := models.Message{}

	for {
		// log.Println(c.OnlineUsers, "chatBers")
		// err = conn.ReadJSON(&body)
		_, msg, errk := conn.ReadMessage()
		err := json.Unmarshal(msg, &body)
		if err != nil {
			return err
		}

		fmt.Println(body.Type, "jzon")

		if code, ok := errk.(*websocket.CloseError); ok {
			// if c.Code == 1000 {
			// 	// Never entering since c.Code == 1005
			// 	log.Println("ok status", k)
			// 	break
			// }
			//logout, close tab -> leave
			log.Println(code.Code, "codet")
			if code.Code == 1001 {
				cs.leaveUser(c, &models.Chat{UUID: body.Sender})
				break
			}
		}

		if body.Type == "getmessages" {
			messages := &models.Message{
				Conn:     conn,
				Sender:   body.Sender,
				Receiver: body.Receiver,
				Name:     name,
			}
			c.ListMessage <- messages
		}

		if body.Type == "newmessage" {

			message := &models.Message{
				Conn:     conn, //set conn current user
				Receiver: body.Receiver,
				Sender:   body.Sender,
				Content:  body.Content,
				UserID:   body.UserID,
				Name:     name,
			}
			c.NewMessage <- message
		}
		// if strings.TrimSpace(username) == "" {
		if body.Type == "newuser" || body.Type == "getusers" {
			id, err := cs.repository.GetUserID(body.Sender)
			if err != nil {
				log.Println(err)
			}
			user := &models.Chat{
				ID:       id,
				UUID:     body.Sender,
				Conn:     conn, //set conn current user
				UserName: name,
			}
			if body.Type == "getusers" {
				c.GetUsers <- user
			}
			if body.Type == "newuser" {
				c.Join <- user
			}
		}
		if body.Type == "leave" {
			user := &models.Chat{
				Conn: conn,
				UUID: body.Sender,
			}
			c.Leave <- user
		}
	}
	defer conn.Close()
	// cs.Reader(conn)
	return nil

}

// signin -> Authorization.UUID -> /api/chat -> ChatHandler->
// if receiver  & message != nil -> checkuserByUUID() if ok -> findUserConnByUUid() -> send message broadcast
