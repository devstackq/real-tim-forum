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

var Online struct {
	Users map[string]*models.Chat `json:"online"`
}

//or global each time send client - use client side this var ?``
//open, then close conn //read/write - block //1 goproutine - read budffer, 2 goroutine write buffer, reusable buffer
type ChatStore struct {
	ListMessage     []models.Message        `json:"messages"`
	OnlineUsers     map[string]*models.Chat `json:"online"`
	AllUsers        []*models.Chat          `json:"users"`
	Message         models.Message          `json:"message"`
	Type            string                  `json:"type"`
	Author          string                  `json:"author"`
	Receiver        string                  `json:"receiver"`
	Sender          string                  `json:"sender"`
	Offset          int                     `json:"offset"`
	CountNewMessage int                     `json:"countnewmessage"`
}

type NewUser struct {
	User *models.Chat `json:"user"`
	Type string       `json:"type"`
}

type ChatService struct {
	repository repository.Chat
}

func NewChatService(repo repository.Chat) *ChatService {
	return &ChatService{repo}
}

//receiver, sender uuid
func (cs *ChatService) getMessages(m *models.Message, c *models.ChannelStorage) {
	//send client = lastID Msg & RoomName - offset
	store := ChatStore{}
	store.Receiver = m.Receiver
	store.Sender = m.Sender
	store.Author = m.Name
	//set and send last result offset
	store.Offset = m.Offset

	room, err := cs.repository.IsExistRoom(m)
	if err != nil {
		log.Println(err, "empty room")
	}
	if room == "" {
		store.Type = "nomessages"
		m.Conn.WriteJSON(store)
		return
	}
	//offset -> sdata change -> scroll
	m.Room = room
	seq, _, err := cs.repository.GetMessages(m)
	if err != nil {
		log.Println(err, "get msg err")
	}
	log.Println(2)
	store.ListMessage = seq
	store.Type = "chathistory"
	err = m.Conn.WriteJSON(store)
	if err != nil {
		log.Println(err)
	}
	// log.Println(seq, 2, store)
}

func (cs *ChatService) mergeUsers(dbUsers []*models.Chat, onlineUsers map[string]*models.Chat) []*models.Chat {
	// go func() {
	for index, dbUser := range dbUsers { //sorted users from db
		for uuid, onlineUser := range onlineUsers { //server users
			if onlineUser.ID == dbUser.ID {
				dbUsers[index].UUID = uuid
				dbUsers[index].Online = true
			}
			// exlude -> ourselve
		}
	}
	return dbUsers
}

//if find userid -> update CS.OnlineUsers[id]=u, uuid, u.LastMessage, etc
func (cs *ChatService) sendMessage(c *models.ChannelStorage, m *models.Message) {
	//save db in message, caht table & message, add author, date message
	//send msg - to  conn - receiver if have in server
	room, err := cs.repository.IsExistRoom(m)
	if err != nil {
		log.Println(err, " Room New msg3")
	}
	randomRoom := Randomaizer()
	store := ChatStore{}
	if room == "" {
		m.Room = randomRoom
		log.Println(m.Room, "Send new room in db func")
		err = cs.repository.AddNewRoom(m)
		if err != nil {
			log.Println(err, "Add room err")
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
		store.Message.Content = m.Content
		store.Message.SentTime = time.Now().Format(time.Stamp)
		store.Message.Sender = m.Sender
		store.Message.Name = m.Name                    //sender msg name
		store.Message.ReceiverName = receiver.UserName // receiver name
		store.Type = "lastmessage"

		err = receiver.Conn.WriteJSON(store)
		if err != nil {
			log.Println(err, "err json add new msg")
		}
		//notify user, & send message in WriteMessage
	}
}

func (cs *ChatService) leaveUser(c *models.ChannelStorage, u *models.Chat) {
	user := NewUser{}
	user.Type = "leaveuser"
	user.User = u
	// user.UserID =
	delete(c.OnlineUsers, u.UUID)
	u.Conn.Close()
	//update users
	Online.Users = c.OnlineUsers
	for _, v := range Online.Users {
		v.Conn.WriteJSON(user)
	}
}

func (cs *ChatService) addGetUpdateUser(u *models.Chat, c *models.ChannelStorage, wsType string) {
	//fill user.Name -> in db, by uuid  u.Conn
	store := ChatStore{}
	store.Type = "chatusers"
	//case - relogin, delete prev user in map, no duplicate
	if len(c.OnlineUsers) > 1 {
		for k := range c.OnlineUsers {
			_, err := cs.repository.GetUserID(k)
			if err != nil {
				delete(c.OnlineUsers, k)
			}
		}
	}
	//add onlien user
	c.OnlineUsers[u.UUID] = u
	store.OnlineUsers = c.OnlineUsers

	sorted, err := cs.repository.GetSortedUsers(u.ID)
	if err != nil {
		log.Println(err)
	}
	//add online state user, for own
	store.AllUsers = cs.mergeUsers(sorted, store.OnlineUsers)
	u.Conn.WriteJSON(store)
	//for another user -> list users & user which signin
	u.Online = true
	newUser := NewUser{}
	newUser.Type = wsType
	newUser.User = u
	// newUser.UUID = u.Sende

	//all connected user - observe - new user connect
	for _, v := range store.OnlineUsers {
		if v.UUID != newUser.User.UUID {
			v.Conn.WriteJSON(newUser)
		}
	}
}

//1 main -> Start() ->  createEmptyObjecetChat -> 2 ws Handler, newConn -> 3 go Run() // goruutine each newConn(user)
//handle if receive new user (from new conn(user) -> join), get new user -> by chan(goroutine)
func (cs *ChatService) Run(c *models.ChannelStorage) {
	//run every, wait data from chan
	for {
		//each conn - own goroutuine
		select {
		case newuser := <-c.NewUser:
			cs.addGetUpdateUser(newuser, c, "newuser")
		case onlineUser := <-c.Join:
			cs.addGetUpdateUser(onlineUser, c, "online")
		case user := <-c.Leave:
			cs.leaveUser(c, user)
		case message := <-c.NewMessage:
			cs.sendMessage(c, message)
		// case listuser := <-c.GetUsers:
		// 	cs.getUsers(listuser, c)
		case last := <-c.LastMessages:
			cs.getMessages(last, c)
		}
	}
}

// Handler -> each new client connect server, create handshake - save data - send service
//Service - check type from client -> if newmessage -> create obj Message, fill field -> run gorutine
//run func broadcast(), conn,WriteJson(data), find by uuid

//conn - hsndshake, client - server, server for loop: listen ws.send message -f if have -> another thread Run goroutine
//and slect case : wait data from channel
func (cs *ChatService) ChatBerserker(conn *websocket.Conn, c *models.ChannelStorage, name string, uuid string) error {

	body := models.Message{}

	for {
		log.Println("accept more 1 ws.send from cleint")
		_, msg, errk := conn.ReadMessage()
		err := json.Unmarshal(msg, &body)
		if err != nil {
			return err
		}

		fmt.Println(body.Type, "ws type")

		if code, ok := errk.(*websocket.CloseError); ok {
			//logout, close tab -> leave, break loop
			if code.Code == 1001 || code.Code == 1006 {
				log.Println(code.Code, "err conn")
				cs.leaveUser(c, &models.Chat{UUID: body.Sender})
				break
			}
			log.Println(code.Code, "err conn")
		}

		if body.Type == "last10msg" {
			messages := &models.Message{
				Conn:     conn,
				Sender:   body.Sender,
				SenderID: body.SenderID,
				Receiver: body.Receiver,
				Name:     name,
				Offset:   body.Offset,
			}
			c.LastMessages <- messages
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
		if body.Type == "getusers" || body.Type == "signin" || body.Type == "leave" || body.Type == "newuser" {
			// send uuid || id, leave || online user
			log.Println(body, "body")
			user := &models.Chat{
				UUID:     body.Sender,
				Conn:     conn, //set conn current user
				UserName: name,
			}
			id, err := cs.repository.GetUserID(body.Sender)
			if err != nil {
				log.Println(err, "erka")
			}
			user.ID = id

			if body.Type == "newuser" {
				c.NewUser <- user
			}

			if body.Type == "getusers" || body.Type == "signin" {
				c.Join <- user
			}

			if body.Type == "leave" {
				user.ID = body.UserID
				//user.UUID = body.Sender
				c.Leave <- user
			}
		}
	}
	defer conn.Close()
	return nil
}
