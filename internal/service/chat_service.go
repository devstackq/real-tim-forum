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
	ListMessage []models.Message        `json:"messages"`
	ListUsers   map[string]*models.User `json:"users"`
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

//or global each time send client - use client side this var ?``

//open, then close conn //read/write - block //1 goproutine - read budffer, 2 goroutine write buffer, reusable buffer
func (cs *ChatService) addNewUser(u *models.User, c *models.Chat) {
	//fill user.Name -> in db, by uuid  u.Conn
	c.ListUsers[u.UUID] = u //add conn ws, & name by [uuid key]
	ChatStorage.Type = "listusers"
	ChatStorage.ListUsers = c.ListUsers // get gloabal var, map[uuid]name

	for _, v := range ChatStorage.ListUsers {
		v.Conn.WriteJSON(ChatStorage)
	}
	//add new user in map -> each connected user  send -> newuser
	// cs.getListUsers(u, c)
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
func (cs *ChatService) getUsers(u *models.User) {
	ChatStorage.Type = "listusers"
	// ChatStorage.ListUsers
	u.Conn.WriteJSON(ChatStorage)
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
			// default :
		}
	}
}

//out handler ?
// Handler -> each new client connect server, create handshake - save data - send service
//Service - check type from client -> if newmessage -> create obj Message, fill field -> run gorutine
//run func broadcast(), conn,WriteJson(data), find by uuid

func (cs *ChatService) ChatBerserker(conn *websocket.Conn, c *models.Chat, name string, uuid string) error {

	body := models.Message{}

	for {
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
				ChatStorage.Type = "leave"
				ChatStorage.Receiver = uuid
				delete(c.ListUsers, uuid)

				for _, v := range ChatStorage.ListUsers {
					v.Conn.WriteJSON(ChatStorage)
				}
				log.Println("user leave", ok)
				break
			}
		}

		// if strings.TrimSpace(username) == "" {
		if body.Type == "newuser" {
			user := &models.User{
				UUID:     body.Sender,
				Conn:     conn, //set conn current user
				FullName: name,
			}
			c.Join <- user
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

		if body.Type == "getusers" {
			users := &models.User{
				Conn: conn,
			}
			c.GetUsers <- users
		}
	}
	defer conn.Close()

	// cs.Reader(conn)
	return nil

}

// signin -> Authorization.UUID -> /api/chat -> ChatHandler->
// if receiver  & message != nil -> checkuserByUUID() if ok -> findUserConnByUUid() -> send message broadcast
