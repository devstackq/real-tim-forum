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
	ToWho        int       `json:towho`
	FromWho      int       `json:fromwho`
	CreatedTime  time.Time `json:"createdtime"`
	UpdatedTime  time.Time `json:"updatedtime"`
	Image        []byte    `json:"image"`
	CountLike    int       `json:"countlike"`
	CountDislike int       `json:"countdislike"`
	Category     string    `json:"category"`
}
