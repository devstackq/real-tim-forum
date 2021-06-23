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

func (vs *VoteService) VoteTerminator(vote *models.Vote) error {
	//good practice?

	count, err := vs.repository.GetCountVote(vote)

	if err != nil {
		return err
	}

	state, err := vs.repository.GetVoteState(vote)
	//if no rows -> set likeState = true, count+=1
	if err != nil {
		//first row uid, pid
		err = vs.repository.SetVoteState(vote)
		if err != nil {
			return err
		}
		count += 1
	} else if !state.LikeState && state.DislikeState {

		if vote.VoteType == "like" {
			vote.LikeState = true
			vote.DislikeState = false
		} else if vote.VoteType == "dislike" {
			vote.DislikeState = false
		}
		err = vs.repository.UpdateVoteState(vote)
		if err != nil {
			return err
		}
		count += 1
		//like = true, dislkie= fasle, count+1
	} else if state.LikeState && !state.DislikeState {

		if vote.VoteType == "like" {
			vote.LikeState = false
		} else if vote.VoteType == "dislike" {
			vote.DislikeState = true
			vote.LikeState = false
		}

		err = vs.repository.UpdateVoteState(vote)
		if err != nil {
			return err
		}
		count -= 1

	}

	vote.Count = count

	err = vs.repository.UpdateCountVote(vote)

	if err != nil {
		return err
	}
	return nil

}
