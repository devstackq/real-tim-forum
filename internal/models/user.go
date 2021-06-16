package models

import "time"

//User struct
type User struct {
	UserID      int       `json:"userid"`
	FullName    string    `json:"fullname"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	IsAdmin     bool      `json:"isadmin"`
	Age         string    `json:"age"`
	Sex         string    `json:"gender"`
	CreatedTime time.Time `json:"createdtime"`
	City        string    `json:"city"`
	Image       []byte    `json:"image"`
	ImageHTML   string    `json:"imagehtml"`
	Role        string    `json:"role"`
	SVG         bool      `json:"svg"`
	Type        string    `json:"type"`
	Temp        string    `json:"temp"`
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	Username    string    `json:"username"`
	LastTime    time.Time `json:"lasttime"`
	LastSeen    string    `json:"lastseen"`
}

// Session     *Session  `json:"session"`
