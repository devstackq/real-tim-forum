package models

import (
	"time"
)

type Comment struct {
	ID          int       `json:"id"`
	Content     string    `json:"content"`
	PostID      int       `json:"postid"`
	CreatorID   string    `json:"creatorid"`
	UpdatedTime time.Time `json:"updatedtime"`
	Image       []byte    `json:"image"`
	Author      string    `json:"author"`
	CreatedTime time.Time `json:"time"`
	Time        string    `json:"createdtime"`
	// CountLike    int       `json:"countlike"`
	// CountDislike int       `json:"countdislike"`
}
