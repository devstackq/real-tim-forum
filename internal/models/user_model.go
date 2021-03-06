package models

import "time"

//User struct
type User struct {
	UserID      int       `json:"userid"`
	FullName    string    `json:"fullname"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Age         string    `json:"age"`
	Sex         string    `json:"gender"`
	CreatedTime time.Time `json:"time"`
	City        string    `json:"city"`
	Image       []byte    `json:"image"`
	Username    string    `json:"username"`
	LastSeen    string    `json:"lastseen"`
	Time        string    `json:"createdtime"`
}

type Tezt struct {
	V *Vote
	P *Post
	U *User
	M *Message
}

type Profile struct {
	User       *User
	Posts      []*Post //array pointer - each post
	VotedItems []*Post
	// Comment *models.Comment{}
}
type Error struct {
	Err   error  `json:"error"`
	Value string `json:"value"`
}
