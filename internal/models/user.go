package models

import "time"

//User struct
type User struct {
	UserID      int       `json:"userid"`
	FullName    string    `json:"fullName"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	IsAdmin     bool      `json:"isAdmin"`
	Age         int       `json:"age"`
	Sex         string    `json:"sex"`
	CreatedTime time.Time `json:"createdTime"`
	City        string    `json:"city"`
	Image       []byte    `json:"image"`
	ImageHTML   string    `json:"imageHtml"`
	Role        string    `json:"role"`
	SVG         bool      `json:"svg"`
	Type        string    `json:"type"`
	Temp        string    `json:"temp"`
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	Username    string    `json:"username"`
	// Session *general.Session
	LastTime time.Time `json:"lastTime"`
	LastSeen string    `json:"lastSeen"`
}
