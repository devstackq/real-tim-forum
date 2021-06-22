package service

import (
	"fmt"

	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/devstackq/real-time-forum/internal/repository"
)

type VoteService struct {
	repository repository.Vote
}

func NewVoteService(repo repository.Vote) *VoteService {
	return &VoteService{repo}
}

func (vs *VoteService) VotePost(vote *models.Vote) error {
	if vote.VoteType == "like" {
		err := vs.repository.VotePostLike(vote)
		if err != nil {
			return err
		}

		fmt.Println("dislike type post")
	}
	if vote.VoteType == "dislike" {
		fmt.Println("dislike type post")
		err := vs.repository.VotePostDislike(vote)
		if err != nil {
			return err
		}
	}

	return nil
	//check if type == dislike -> VoteDislikePost()
}
