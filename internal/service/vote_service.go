package service

import (
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
		err := ps.repository.VotePostLike(vote)
		if err != nil {
			return err
		}
	}
	if vote.VoteType == "dislike" {
		err := ps.repository.VotePostDislike(vote)
		if err != nil {
			return err
		}
	}

	return nil
	//check if type == dislike -> VoteDislikePost()
}
