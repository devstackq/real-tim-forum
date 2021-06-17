package models

import (
	"time"
)

type Post struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Thread       string    `json:"thread"`
	Content      string    `json:"content"`
	Categories   []string  `json:"categories"`
	CreatorID    int       `json:"creatorID"`
	CreatedTime  time.Time `json:"createdTime"`
	UpdatedTime  time.Time `json:"updatedTime"`
	Image        []byte    `json:"image"`
	CountLike    int       `json:"countLike"`
	CountDislike int       `json:"countDislike"`
}

type Category struct {
	Name string `json:"name"`
}
