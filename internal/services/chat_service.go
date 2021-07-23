package service

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/devstackq/real-time-forum/internal/repository"
	"github.com/gorilla/websocket"
)

var ChatStorage struct {
	ListMessage []models.Message `json:"messages"`
	// OnlineUsers   map[string]*models.User `json:"users"`
	OnlineUsers map[string]*models.Chat `json:"online"`
	AllUsers    []models.Chat           `json:"users"`
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
	// store:= ChatStore{Messages}
	//return store.ListMessages

	ChatStorage.Receiver = m.Receiver
	if room == "" {
		ChatStorage.Type = "nomessages"
		// send []user - by concrete conn, lel
		// cs.addNewUser(u, c)
		m.Conn.WriteJSON(ChatStorage)
		return nil
	}
	m.Room = room
	// log.Println(m, err)
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

type Tezt struct {
	int
	string
}

func (cs *ChatService) mergeUsers(dbUsers []models.Chat, onlineUsers map[string]*models.Chat) *[]models.Chat {
	// l := Tezt{1, "sd"}
	// log.Println(l)
	for index, dbUser := range dbUsers { //sorted users from db
		for uuid, onlineUser := range onlineUsers { //server users
			if onlineUser.ID == dbUser.ID {
				dbUsers[index].UUID = uuid
				dbUsers[index].Online = true
			}
			// exlude -> ourselve
		}
	}
	return &dbUsers
	// ChatStorage.AllUsers = dbUsers
}

//if find userid -> update CS.OnlineUsers[id]=u, uuid, u.LastMessage, etc

//or global each time send client - use client side this var ?``
//open, then close conn //read/write - block //1 goproutine - read budffer, 2 goroutine write buffer, reusable buffer
type ChatStore struct {
	ListMessage []models.Message        `json:"messages"`
	OnlineUsers map[string]*models.Chat `json:"online"`
	AllUsers    *[]models.Chat          `json:"users"`
	Message     models.Message          `json:"message"`
	Type        string                  `json:"type"`
	Receiver    string                  `json:"receiver"`
}
type NewUser struct {
	User *models.Chat `json:"user"`
	Type string       `json:"type"`
}

func (cs *ChatService) addNewUser(u *models.Chat, c *models.ChatStorage) *[]models.Chat {
	//fill user.Name -> in db, by uuid  u.Conn
	store := ChatStore{}

	store.Type = "observeusers"
	//case - relogin, delete prev user in map, no duplicate
	if len(c.OnlineUsers) > 1 {
		for k := range c.OnlineUsers {
			_, err := cs.repository.GetUserID(k)
			if err != nil {
				delete(c.OnlineUsers, k)
			}
		}
	}
	c.OnlineUsers[u.UUID] = u
	store.OnlineUsers = c.OnlineUsers // get gloabal var, map[uuid]name
	//off -> on, findById, -> update data

	sorted, err := cs.repository.GetSortedUsers(u.ID)
	if err != nil {
		log.Println(err)
	}

	store.AllUsers = cs.mergeUsers(sorted, store.OnlineUsers)

	//ourselves -send list users
	u.Conn.WriteJSON(store)

	//another user send send list users & user which signin
	u.Online = true
	newUser := NewUser{}
	newUser.Type = "online"
	newUser.User = u

	for _, v := range store.OnlineUsers {
		if v.UUID != newUser.User.UUID {
			v.Conn.WriteJSON(newUser)
		}
	}
	//save store own conn || call each time - g
	return store.AllUsers
}

func (cs *ChatService) sendMessage(c *models.ChatStorage, m *models.Message) {
	//save db in message, caht table & message, add author, date message
	//send msg - to  conn - receiver if have in server
	room, err := cs.repository.IsExistRoom(m)
	if err != nil {
		log.Println(err, " room New msg3")
	}
	randomRoom := Randomaizer()

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
	// data := ChatStore{}
	//1delete from CS.Mapusers - global variable
	//2send leaved user, state = offline

	// data.Type = "leave"

	log.Println(len(ChatStorage.OnlineUsers), "prev")

	leaveUser := NewUser{}
	leaveUser.Type = "leave"
	leaveUser.User = u

	delete(c.OnlineUsers, u.UUID)
	u.Conn.Close()

	log.Println(len(ChatStorage.OnlineUsers), "after")
	//update users
	ChatStorage.OnlineUsers = c.OnlineUsers
	for _, v := range ChatStorage.OnlineUsers {
		v.Conn.WriteJSON(leaveUser)
	}
	// log.Println("user leave", ChatStorage.OnlineUsers)
}

//get adduser data - then use
func (cs *ChatService) getUsers(store *[]models.Chat, u *models.Chat) {
	//no use global  var
	//check if users have in session
	users := ChatStore{Type: "getusers", AllUsers: store}
	println("memory address of => store", &store)
	if store != nil {
		u.Conn.WriteJSON(users)
	}
}

//1 main -> Start() ->  createEmptyObjecetChat -> 2 ws Handler, newConn -> 3 go Run() // goruutine each newConn(user)
//handle if receive new user (from new conn(user) -> join), get new user -> by chan(goroutine)
func (cs *ChatService) Run(c *models.ChatStorage) {
	//run every, wait data from chan
	store := ChatStore{}
	for {
		select {
		case newuser := <-c.Join:
			store.AllUsers = cs.addNewUser(newuser, c)
		case user := <-c.GetUsers:
			cs.addNewUser(user, c)
			// cs.getUsers(store.AllUsers, user) //suk( wtf ??
		case user := <-c.Leave:
			cs.leaveUser(c, user)
		case message := <-c.NewMessage:
			cs.sendMessage(c, message)
		case list := <-c.ListMessage:
			cs.getMessages(list, c)
		default:
			// log.Println("nill chan", c.Leave)
		}
	}
}

// Handler -> each new client connect server, create handshake - save data - send service
//Service - check type from client -> if newmessage -> create obj Message, fill field -> run gorutine
//run func broadcast(), conn,WriteJson(data), find by uuid

//conn - hsndshake, client - server, server for loop: listen ws.send message -f if have -> another thread Run goroutine
//and slect case : wait data from channel
func (cs *ChatService) ChatBerserker(conn *websocket.Conn, c *models.ChatStorage, name string, uuid string) error {

	body := models.Message{}

	for {
		log.Println("accept more 1 ws.send from cleint")
		// err = conn.ReadJSON(&body)
		_, msg, errk := conn.ReadMessage()
		err := json.Unmarshal(msg, &body)
		if err != nil {
			return err
		}

		fmt.Println(body.Type, "ws type")

		if code, ok := errk.(*websocket.CloseError); ok {
			// if c.Code == 1000 {
			// 	// Never entering since c.Code == 1005
			// 	log.Println("ok status", k)
			// 	break
			// }
			//logout, close tab -> leave
			log.Println(code.Code, "codet")
			if code.Code == 1001 {
				log.Println(code.Code, "code 1001")
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
			//maybe set user id ?
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
		if body.Type == "newuser" || body.Type == "getusers" || body.Type == "leave" {
			id, err := cs.repository.GetUserID(body.Sender)
			if err != nil {
				log.Println(err, "erka")
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
			if body.Type == "leave" {
				user.ID = body.UserID
				c.Leave <- user
			}
		}
	}
	defer conn.Close()
	// cs.Reader(conn)
	return nil
}

// signin -> Authorization.UUID -> /api/chat -> ChatHandler->
// if receiver  & message != nil -> checkuserByUUID() if ok -> findUserConnByUUid() -> send message broadcast
