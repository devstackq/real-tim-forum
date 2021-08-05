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
	ListMessage []models.Message        `json:"messages"`
	OnlineUsers map[string]*models.Chat `json:"online"`
	AllUsers    []*models.Chat          `json:"users"`
	Message     models.Message          `json:"message"`
	Type        string                  `json:"type"`
	Author      string                  `json:"author"`
	Receiver    string                  `json:"receiver"`
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

func (cs *ChatService) getMessages(m *models.Message, c *models.ChannelStorage) error {
	//find users room, if zero
	store := ChatStore{}
	store.Receiver = m.Receiver
	store.Author = m.Name

	room, err := cs.repository.IsExistRoom(m)
	if err != nil {
		log.Println(err, "empty room")
	}
	// log.Println(m.Name, "NAME")

	if room == "" {
		store.Type = "nomessages"
		// send []user - by concrete conn, lel
		// cs.addNewUser(u, c)
		m.Conn.WriteJSON(store)
		return nil
	}
	m.Room = room
	seq, err := cs.repository.GetMessages(m)
	if err != nil {
		log.Println(err, "get msg err")
	}

	store.ListMessage = seq
	store.Type = "listmessages"
	// log.Println(seq)
	err = m.Conn.WriteJSON(store)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ChatService) mergeUsers(dbUsers []*models.Chat, onlineUsers map[string]*models.Chat) []*models.Chat {
	// l := Tezt{1, "sd"}
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
	// }()
	return dbUsers
	// ChannelStorage.AllUsers = dbUsers
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
		store.Message.SentTime = time.Now().Format("2006-01-02 3:4:5 pm")
		// store.Message.SentTime = Format("2006-01-02 3:4:5 pm"))
		store.Message.Sender = m.Sender
		store.Message.Name = m.Name
		store.Type = "lastmessage"

		err = receiver.Conn.WriteJSON(store)
		if err != nil {
			log.Println(err, "err json add new msg")
		}
		//notify user, & send message in WriteMessage
	}
}

func (cs *ChatService) leaveUser(c *models.ChannelStorage, u *models.Chat) {
	leaveUser := NewUser{}
	leaveUser.Type = "leave"
	leaveUser.User = u

	delete(c.OnlineUsers, u.UUID)
	u.Conn.Close()
	//update users
	Online.Users = c.OnlineUsers
	for _, v := range Online.Users {
		v.Conn.WriteJSON(leaveUser)
	}
}

func (cs *ChatService) addNewUser(u *models.Chat, c *models.ChannelStorage) ChatStore {
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
	store.OnlineUsers = c.OnlineUsers

	sorted, err := cs.repository.GetSortedUsers(u.ID)
	if err != nil {
		log.Println(err)
	}
	store.AllUsers = cs.mergeUsers(sorted, store.OnlineUsers)
	// u.Uzers = sorted
	//another user send send list users & user which signin
	u.Online = true
	newUser := NewUser{}
	newUser.Type = "online"
	newUser.User = u

	//all connected user - observe - new user connect
	for _, v := range store.OnlineUsers {
		if v.UUID != newUser.User.UUID {
			v.Conn.WriteJSON(newUser)
		}
	}

	//reSortUsers -> if newuserSignup()
	// add new user handle, signup -> signin - other user observs
	// fix d- getProfileDate() - userid

	log.Println(store.OnlineUsers, newUser, "added user")
	//only get list user -> // store.AllUsers
	//user logged -> show own list user
	u.Conn.WriteJSON(store)
	return store
}

//sol1: adduser -> then use -> &user.ListUsers, by pointer in address
//sol2: again call newuser()
//sol3: send addUser() - listUser -> then update each change - client side
func (cs *ChatService) getUsers(user *models.Chat) {
	//no use global  var
	store := ChatStore{Type: "getusers", AllUsers: user.Uzers}
	println("memory address of => store", &user)
	user.Conn.WriteJSON(store)
}

//1 main -> Start() ->  createEmptyObjecetChat -> 2 ws Handler, newConn -> 3 go Run() // goruutine each newConn(user)
//handle if receive new user (from new conn(user) -> join), get new user -> by chan(goroutine)
func (cs *ChatService) Run(c *models.ChannelStorage) {
	//run every, wait data from chan
	for {
		//each conn - own goroutuine
		select {
		case newuser := <-c.Join:
			cs.addNewUser(newuser, c)
		case user := <-c.Leave:
			cs.leaveUser(c, user)
		case message := <-c.NewMessage:
			cs.sendMessage(c, message)
		// case listuser := <-c.GetUsers:
		// 	cs.addNewUser(listuser, c)
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
func (cs *ChatService) ChatBerserker(conn *websocket.Conn, c *models.ChannelStorage, name string, uuid string) error {

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

		if body.Type == "newuser" {
			log.Println("observe onlien users, update mergeusers")
			//get online users
			//updated listUsers
		}
		// if strings.TrimSpace(username) == "" {
		if body.Type == "online" || body.Type == "leave" {
			send uuid || id, leave || online user

			user := &models.Chat{
				UUID:     body.Sender,
				Conn:     conn, //set conn current user
				UserName: name,
			}

			if body.Type == "online" {
				id, err := cs.repository.GetUserID(body.Sender)
				if err != nil {
					log.Println(err, "erka")
				}
				user.ID = id
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
// if receiver  & message != nil -addNewUser
