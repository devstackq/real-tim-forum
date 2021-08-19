package models

import (
	"time"
)

type Comment struct {
	ID          int       `json:"id"`
	Content     string    `json:"content"`
	PostID      int       `json:"postid"`
	CreatorID   string    `json:"creatorid"`
	CreatedTime time.Time `json:"createdtime"`
	UpdatedTime time.Time `json:"updatedtime"`
	Image       []byte    `json:"image"`
	// CountLike    int       `json:"countlike"`
	// CountDislike int       `json:"countdislike"`
}
