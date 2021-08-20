package models

import (
	"time"
)

type Post struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Thread       string    `json:"thread"`
	Content      string    `json:"content"`
	Categories   []string  `json:"categories"`
	CreatorID    string    `json:"creatorid"`
	CreatedTime  time.Time `json:"time"`
	Time         string    `json:"createdtime"`
	UpdatedTime  time.Time `json:"updatedtime"`
	Image        []byte    `json:"image"`
	CountLike    int       `json:"countlike"`
	CountDislike int       `json:"countdislike"`
	Category     string    `json:"category"`
	Comments     []Comment `json:"comments"`
	Author       string    `json:"author"`
}
