package models

//chatID99 -> from, who chatid JOIN
//query select * from messages where uiser_id1 => and user_id2=?
//59, 19 ->
//hi; how r u? uid1
//- hello, fine and u ? uid2
type Message struct {
	ID       int    `json:"id"`
	Content  string `json:"message"`
	ChatID   int    `json:"chatid"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
}

// func NewMessage(body string, sender int) *Message {
// 	return &Message{
// 		Body:   body,
// 		Sender: sender,
// 	}
// }

type Chat struct {
	Users   map[string]*User
	Message chan *Message
	Join    chan *User
	Leave   chan *User
}

//user19 -> send msg -> user 59, from, who, chatid99
