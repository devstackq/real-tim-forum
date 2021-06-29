package models

type Chat struct {
	ID     int
	ChatID string
	Type   string
}

//chatID99 -> from, who chatid JOIN
//query select * from messages where uiser_id1 => and user_id2=?
//59, 19 ->
//hi; how r u? uid1
//- hello, fine and u ? uid2
type Message struct {
	ID       int
	Content  string
	ChatID   int //99
	Sender   int
	Receiver int
}

//user19 -> send msg -> user 59, from, who, chatid99
