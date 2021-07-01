package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/devstackq/real-time-forum/internal/repository"
	"github.com/gorilla/websocket"
)

var UUID = ""

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

//open, then close conn //read/write - block //1 goproutine - read budffer, 2 goroutine write buffer, reusable buffer
func (cs *ChatService) addNewUser(u *models.User, c *models.Chat) {
	fmt.Println(u.UUID, "add user in map")
	c.Users[u.UUID] = u
}

//meow
func (cs *ChatService) broadcast(c *models.Chat, m *models.Message, sender string) error {

	if c.Users[sender] != nil {
		conn := c.Users[sender]
		conn.Conn.WriteJSON(m.Content)
		//notify user, & send message in WriteMessage
	}
	//else just save in db
	return nil
}

//1 main -> Start() ->  createEmptyObjecetChat -> 2 ws Handler, newConn -> 3 go Run() // goruutine each newConn(user)

//handle if receive new user (from new conn(user) -> join), get new user -> by chan(goroutine)
func (cs *ChatService) Run(c *models.Chat) {
	for {
		select {
		case user := <-c.Join:
			cs.addNewUser(user, c)
		case msg := <-c.Message:
			cs.broadcast(c, msg, UUID)
		}
		// default :
	}
}

//out handler ?
func (cs *ChatService) ChatBerserker(w http.ResponseWriter, r *http.Request, c *models.Chat, sender string) error {

	//senderUuid, connWS ->
	//if receiver online -> send notify else -> saveDb message, return  updated []messages -cleint Chat page
	// receiverUUID, connWS

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err, 1)
		return err
	}
	defer conn.Close()
	//create object - each new conn
	json := models.Message{}

	err = conn.ReadJSON(&json)
	if err != nil {
		return err
	}
	// err = conn.PingHandler()
	// if strings.TrimSpace(username) == "" {
	// getUserByConnUUID(json.Receiver)
	// senderId -> getMap[userID : conn] // when signin

	//add new User -> service and add in Map user like online

	///client side - /chat, get query (JoinUser get cookie UUID) ws.send(uuid, type join)
	//client side - click send msg  (sendMsg, )

	//set data - in cchat struct -> join user data
	user := &models.User{
		Conn:   conn, //set conn current user
		Global: c,    // set  User struct - chat object
	}

	c.Join <- user

	UUID = sender

	msg := &models.Message{
		Receiver: json.Receiver,
		Content:  json.Content,
	}

	c.Message <- msgAu
	//send message

	return nil
}

// signin -> Authorization.UUID -> /api/chat -> ChatHandler->
// if receiver  & message != nil -> checkuserByUUID() if ok -> findUserConnByUUid() -> send message broadcast
