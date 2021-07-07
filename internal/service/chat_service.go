package service

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/devstackq/real-time-forum/internal/repository"
	"github.com/gorilla/websocket"
)

var ChatJSON struct {
	ListMessage []models.Message  `json:"messages"`
	Type        string            `json:"type"`
	ListUsers   map[string]string `json:"users"`
	Receiver    string            `json:"receiver"`
	Message     string            `json:"message"`
}

type ChatService struct {
	repository repository.Chat
}

func NewChatService(repo repository.Chat) *ChatService {
	return &ChatService{repo}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024, // read/write, count network call
	WriteBufferSize: 1024,
}

func (cs *ChatService) getMessages(m *models.Message, c *models.Chat) error {
	//find users room, if zero
	room, err := cs.repository.IsExistRoom(m)
	if err != nil {
		log.Println(err, "empty room get msg")
	}
	// m.Room = room
	if room == "" {
		ChatJSON.Type = "nomessages"
		ChatJSON.Receiver = m.Receiver
		m.Conn.WriteJSON(ChatJSON)
		return nil
	}
	m.Room = room
	seq, err := cs.repository.GetMessages(m)
	if err != nil {
		log.Println(err, "get msg err")
	}

	ChatJSON.Type = "listmessages"
	ChatJSON.ListMessage = seq
	err = m.Conn.WriteJSON(ChatJSON)
	if err != nil {
		return err
	}
	return nil
}

//open, then close conn //read/write - block //1 goproutine - read budffer, 2 goroutine write buffer, reusable buffer
func (cs *ChatService) addNewUser(u *models.User, c *models.Chat) {

	//fill user.Name -> in db, by uuid
	c.Users[u.UUID] = u.Conn
	c.ListsUsers[u.UUID] = u.FullName // if leave remove

	//updatelist users
	ChatJSON.Type = "listusers"
	ChatJSON.ListUsers = c.ListsUsers

	err := u.Conn.WriteJSON(ChatJSON)
	if err != nil {
		log.Println(err)
		return
	}
}

//utils ?
func randomaizer() string {
	rand.Seed(time.Now().Unix())
	charSet := "abcdedfghijklmnopqrst"
	var output strings.Builder
	length := 10
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}

func (cs *ChatService) broadcast(c *models.Chat, m *models.Message) {

	//comapere ListUsers {
	// get Users[m.Receiver]
	// if c.Users[m.Receiver] == uuidReceiverInDB?

	//save db in message, caht table & message, add author, date message
	//send msg - to  conn - receiver if have server

	//logic todo here if exist
	room, err := cs.repository.IsExistRoom(m)
	if err != nil {
		log.Println(err, " room New msg3")
	}
	randomRoom := randomaizer()
	fmt.Println(room, "rec", m.Receiver, "sen", m.Sender, "pre add message")

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
	// fix each time new room create new message
	m.Name, err = cs.repository.GetUserName(m.UserID)
	if err != nil {
		log.Println(err, "get username err")
	}

	err = cs.repository.AddNewMessage(m)
	if err != nil {
		log.Println(err, "err add new msg")
	}
	if c.Users[m.Receiver] != nil {
		log.Println(m.Content, "send another conn")
		rec := c.Users[m.Receiver]
		// ChatJSON.Type = "lastmessage"
		ChatJSON.Type = "listmessages"
		ChatJSON.Message = m.Content
		err = rec.WriteJSON(ChatJSON)
		if err != nil {
			log.Println(err, "err json add new msg")
		}
		//notify user, & send message in WriteMessage
		// defer rec.Close()
	}
}

//1 main -> Start() ->  createEmptyObjecetChat -> 2 ws Handler, newConn -> 3 go Run() // goruutine each newConn(user)
//handle if receive new user (from new conn(user) -> join), get new user -> by chan(goroutine)
func (cs *ChatService) Run(c *models.Chat) {
	for {
		select {
		case user := <-c.Join:
			cs.addNewUser(user, c)
		case message := <-c.NewMessage:
			cs.broadcast(c, message)
		case list := <-c.ListMessage:
			cs.getMessages(list, c)
			// default :
		}
	}
}

//out handler ?
// Handler -> each new client connect server, create handshake - save data - send service
//Service - check type from client -> if newmessage -> create obj Message, fill field -> run gorutine
//run func broadcast(), conn,WriteJson(data), find by uuid

func (cs *ChatService) ChatBerserker(conn *websocket.Conn, c *models.Chat, name string) error {

	json := models.Message{}
	for {

		err := conn.ReadJSON(&json)
		if err != nil {
			return err
		}
		fmt.Println(json.Type, json.UserID, "jzon")
		// err = conn.PingHandler()
		// if strings.TrimSpace(username) == "" {
		//add new User -> service and add in Map user like online

		//set data - in cchat struct -> join user data
		if json.Type == "newuser" {
			user := &models.User{
				UUID:     json.Sender,
				Conn:     conn, //set conn current user
				Global:   c,    // set  User struct - chat object
				FullName: name,
			}
			c.Join <- user
		}

		if json.Type == "getmessages" {
			getMsg := &models.Message{
				Conn:     conn,
				Sender:   json.Sender,
				Receiver: json.Receiver,
				Name:     name,
			}
			c.ListMessage <- getMsg
		}

		if json.Type == "newmessage" {
			message := &models.Message{
				Conn:     conn, //set conn current user
				Receiver: json.Receiver,
				Sender:   json.Sender,
				Content:  json.Content,
				UserID:   json.UserID,
			}
			c.NewMessage <- message
		}
	}
	// cs.Reader(conn)
	return nil

}

// signin -> Authorization.UUID -> /api/chat -> ChatHandler->
// if receiver  & message != nil -> checkuserByUUID() if ok -> findUserConnByUUid() -> send message broadcast
