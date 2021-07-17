package service

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/devstackq/real-time-forum/internal/repository"
	"github.com/gorilla/websocket"
)

var ChatStorage struct {
	ListMessage []models.Message        `json:"messages"`
	ListUsers   map[string]*models.User `json:"users"`
	NewUser     *models.User            `json:"newuser"`
	ByAlpha     map[string]*models.User `json:"alphausers"`
	ByTime      map[string]*models.User `json:"timeusers"`
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

func (cs *ChatService) getMessages(m *models.Message, c *models.Chat) error {
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

func sortUsers(seq []*models.Message, sortType string, indexs []int) error {
	sortedUserByTime := make(map[string]*models.User)
	sortedUserByAlpha := make(map[string]*models.User)
	if len(indexs) > 1 {
		sort.Ints(indexs)
	}

	if sortType == "time" {
		for _, message := range seq {
			for _, i := range indexs {
				if message.ID == i {
					sortedUser := models.User{UUID: message.Receiver, FullName: message.Name}
					sortedUserByTime[message.Receiver] = &sortedUser
					// ChatStorage.ByTime = append(ChatStorage.ByTime, *models.User{UUID: message.Receiver, FullName: message.Name})
				}
			}
		}
	} else if sortType == "alpha" {
		// for _, m := range seq {
			sort string alpha
		sort.Slice(seq, func(i1, i2 int) bool {
			return len(seq[i1].Name) < len(seq[i2].Name)
		})
		// sort.Sort(sort.StringSlice(m.Name))
	}

	ChatStorage.ByAlpha = sortedUserByAlpha
	ChatStorage.ByTime = sortedUserByTime

	log.Println(sortedUserByAlpha, "by alpha empty user")
	return nil
}

func (cs *ChatService) prepareListUsers(sender string) (err error) {

	seqLastMessageTime := []*models.Message{}
	seqAlpha := []*models.Message{}
	indexs := []int{}

	for k, v := range ChatStorage.ListUsers {
		if k != sender {
			temp := models.Message{}
			// log.Println(k, receiver, v.UUID, 1234)
			m := models.Message{Sender: sender, Receiver: k}
			room, err := cs.repository.IsExistRoom(&m)
			if err != nil {
				log.Println(err, " room err")
				log.Println(v.FullName, "no has message")
				temp.Name = v.FullName
				seqAlpha = append(seqAlpha, &temp)
			} else {
				temp.Sender = sender
				temp.Receiver = k

				index, err := cs.repository.GetLastMessageIndex(room, v.UserID)
				if err != nil {
					log.Println(err)
					return err
				}
				temp.ID = index
				temp.Name = v.FullName
				seqLastMessageTime = append(seqLastMessageTime, &temp)
				indexs = append(indexs, index)
			}
		}
	}
	// indexs = sort.Ints(indexs)

	err = sortUsers(seqLastMessageTime, "time", indexs)
	if err != nil {
		return err
	}
	err = sortUsers(seqAlpha, "alpha", nil)
	if err != nil {
		return err
	}
	return nil
}

//or global each time send client - use client side this var ?``
//open, then close conn //read/write - block //1 goproutine - read budffer, 2 goroutine write buffer, reusable buffer
func (cs *ChatService) addNewUser(u *models.User, c *models.Chat) {
	//fill user.Name -> in db, by uuid  u.Conn
	ChatStorage.Type = "observeusers"
	//no duplicate add user

	log.Println("add user", u.UUID, u.UserID)
	if len(c.ListUsers) > 1 {

		for k := range c.ListUsers {
			id, err := cs.repository.GetUserID(k)
			if err != nil {
				log.Println("del user", k)
				delete(c.ListUsers, k)
			} else {
				log.Println("add user", k, "idK", id)
				//delete prev user -> resession case
				//if userid > 0 {
				u.UserID = id
				c.ListUsers[u.UUID] = u
			}
		}
	} else {
		//add first user
		id, err := cs.repository.GetUserID(u.UUID)
		if err != nil {
			log.Println(err)
		}
		u.UserID = id
		c.ListUsers[u.UUID] = u
	}
	ChatStorage.ListUsers = c.ListUsers // get gloabal var, map[uuid]name
	//added new user -> sort -> send update ByTime, ByHistory Users
	if len(c.ListUsers) > 1 {
		cs.prepareListUsers(u.UUID)
		// prepareUsers(u.UUID) here ?
	}
	for _, v := range ChatStorage.ListUsers {
		v.Conn.WriteJSON(ChatStorage)
	}
	// ChatStorage.NewUser = nil
	// ChatStorage.NewUser = u
}

func (cs *ChatService) getUsers(u *models.User) {
	//check if users have in session
	ChatStorage.Type = "getusers"
	cs.prepareListUsers(u.UUID)
	u.Conn.WriteJSON(ChatStorage)
}

func (cs *ChatService) sendMessage(c *models.Chat, m *models.Message) {
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
		log.Println(m.Room, "send new room in db fucn")
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
	if c.ListUsers[m.Receiver] != nil {
		receiver := c.ListUsers[m.Receiver]
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

func (cs *ChatService) leaveUser(c *models.Chat, u *models.User) {
	ChatStorage.Type = "leave"
	u.Conn.Close()
	delete(c.ListUsers, u.UUID)
	ChatStorage.ListUsers = c.ListUsers
	for _, v := range ChatStorage.ListUsers {
		v.Conn.WriteJSON(ChatStorage)
	}
	// log.Println("user leave", ChatStorage.ListUsers)
}

//1 main -> Start() ->  createEmptyObjecetChat -> 2 ws Handler, newConn -> 3 go Run() // goruutine each newConn(user)
//handle if receive new user (from new conn(user) -> join), get new user -> by chan(goroutine)
func (cs *ChatService) Run(c *models.Chat) {
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

func (cs *ChatService) ChatBerserker(conn *websocket.Conn, c *models.Chat, name string, uuid string) error {

	body := models.Message{}

	for {
		// log.Println(c.ListUsers, "chatBers")
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
				cs.leaveUser(c, &models.User{UUID: body.Sender})
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
		if body.Type == "newuser" {
			// log.Println(body, "newuser")
			user := &models.User{
				UUID:     body.Sender,
				Conn:     conn, //set conn current user
				FullName: name,
			}
			c.Join <- user
		}
		if body.Type == "getusers" {
			users := &models.User{
				Conn: conn,
			}
			c.GetUsers <- users
		}
		if body.Type == "leave" {
			user := &models.User{
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
