package models

import (
	"time"
)

type Comment struct {
	ID           string    `json:"id"`
	ParentID     int       `json:"parentid"`
	Content      string    `json:"content"`
	PostID       int       `json:"postid"`
	CreatorID    string    `json:"creatorid"`
	To           int       `json:towho`   //integer??
	From         int       `json:fromwho` //integer??
	CreatedTime  time.Time `json:"createdtime"`
	UpdatedTime  time.Time `json:"updatedtime"`
	Image        []byte    `json:"image"`
	CountLike    int       `json:"countlike"`
	CountDislike int       `json:"countdislike"`
	Category     string    `json:"category"`
}
