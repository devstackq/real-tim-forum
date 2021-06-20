package models

type Vote struct {
	PostID   string `json:"postid"`
	UserID   string `json:"userid"`
	VoteType string `json:"type"`
}
