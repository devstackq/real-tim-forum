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

func (vs *VoteService) VoteItem(vote *models.Vote) error {
	//good practice
	if vote.VoteType == "like" {
		err := vs.repository.VoteLike(vote)
		if err != nil {
			return err
		}
		fmt.Println("like type post")
	}

	if vote.VoteType == "dislike" {
		//good practice, here -> getVoteState-> repos -> setVoteState(), updateVoteValue()
		err := vs.repository.VoteDislike(vote)
		if err != nil {
			return err
		}
		fmt.Println("dislike type post")
	}
	return nil
}
