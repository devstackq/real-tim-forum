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

	//check if type == dislike -> VoteDislikePost()

}
