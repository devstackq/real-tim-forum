package models

type Vote struct {
	ID        string `json:"id"`
	VoteType  string `json:"type"`
	CreatorID string `json:"creatorid"`
	VoteGroup string `json:"group"`
	CountLike int    `json:"countlike"`
	VoteState bool   `json:"likestate"`
}
