package models

type Vote struct {
	ID           string `json:"id"`
	VoteType     string `json:"type"`
	CreatorID    string `json:"creatorid"`
	VoteGroup    string `json:"group"`
	CountLike    int    `json:"countlike"`
	CountDislike int    `json:"countdislike"`
	LikeState    bool   `json:"likestate"`
	DislikeState bool   `json:"dislikestate"`
}
