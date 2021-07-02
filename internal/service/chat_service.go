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
	fmt.Println("uid", u.UUID, "add user in map")
	//fill user.Name -> in db, by uuid
	c.Users[u.UUID] = u
}

//meow
func (cs *ChatService) broadcast(c *models.Chat, m *models.Message) error {

	// if c.Users[m.Receiver] == uuidReceiverInDB?

	if c.Users[m.Receiver] != nil {
		// fmt.Println(c.Users[m.Receiver], "brodcast2")
		rec := c.Users[m.Receiver]
		//save message & sender, receiver id in DB
		err := rec.Conn.WriteJSON(m)
		fmt.Println(err)
		//notify user, & send message in WriteMessage
		defer rec.Conn.Close()

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
			cs.broadcast(c, msg)
		}
		// default :
	}
}

//out handler ?
func (cs *ChatService) ChatBerserker(w http.ResponseWriter, r *http.Request, c *models.Chat, sender string) (map[string]*models.User, error) {

	//senderUuid, connWS ->
	//if receiver online -> send notify else -> saveDb message, return  updated []messages -cleint Chat page
	// receiverUUID, connWS

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err, 1)
		return nil, err
	}

	//create object - each new conn
	json := models.Message{}

	err = conn.ReadJSON(&json)
	if err != nil {
		return nil, err
	}
	// err = conn.PingHandler()
	// if strings.TrimSpace(username) == "" {
	// getUserByConnUUID(json.Receiver)
	// senderId -> getMap[userID : conn] // when signin

	//add new User -> service and add in Map user like online

	///client side - /chat, get query (JoinUser get cookie UUID) ws.send(uuid, type join)
	//client side - click send msg  (sendMsg, )

	// UUID = sender
	// sender, err := r.Cookie("session")
	fmt.Println(json.Type, "type")

	if json.Type == "listusers" {
		//call 1 service
		//return conn:nameUser
		// _, m, _ := conn.WriteMessage(c.Users)
		conn.WriteJSON(c.Users)
		send json data -> then show client side
		return c.Users, nil
	}
	if json.Type == "message" {
		//list users -> click -> msg -> send -> get receiverId, msg -> ws.send
		// getUserById(id) -> uuid -> json.Receiver
		// recevier search by uuid in Db -> broadecast user else save Db

		msg := &models.Message{}
		if json.Receiver != "" {
			msg.Receiver = json.Receiver
			msg.Content = json.Content
			msg.Sender = sender

			c.Message <- msg
		}
	}
	if json.Type == "newuser" {
		//set data - in cchat struct -> join user data
		user := &models.User{
			UUID:   sender,
			Conn:   conn, //set conn current user
			Global: c,    // set  User struct - chat object
		}

		c.Join <- user
		//add user -> /api/chat
		//broadcase -> when send messgae click

	}

	// fmt.Println(msg, "come mbsg")
	//send message
	return nil, nil

}

// signin -> Authorization.UUID -> /api/chat -> ChatHandler->
// if receiver  & message != nil -> checkuserByUUID() if ok -> findUserConnByUUid() -> send message broadcast
