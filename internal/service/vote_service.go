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

func (vs *VoteService) VoteTerminator(vote *models.Vote) error {
	//good practice?

	counts, err := vs.repository.GetCountVote(vote)

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

		if vote.VoteType == "like" {
			counts.CountLike += 1
		} else if vote.VoteType == "dislike" {
			counts.CountDislike += 1
		}

	} else if !state.LikeState && state.DislikeState {

		if vote.VoteType == "like" {
			vote.LikeState = true
			vote.DislikeState = false
			counts.CountLike += 1
			counts.CountDislike -= 1

		} else if vote.VoteType == "dislike" {
			vote.DislikeState = false
			counts.CountDislike -= 1
		}

	} else if state.LikeState && !state.DislikeState {

		if vote.VoteType == "like" {
			vote.LikeState = false
			counts.CountLike -= 1
		} else if vote.VoteType == "dislike" {
			vote.DislikeState = true
			vote.LikeState = false
			counts.CountLike -= 1
			counts.CountDislike += 1
		}

	} else if !state.LikeState && !state.DislikeState {
		if vote.VoteType == "like" {
			vote.LikeState = true
			counts.CountLike += 1
		} else if vote.VoteType == "dislike" {
			vote.DislikeState = false
			counts.CountDislike += 1
		}

	}

	err = vs.repository.UpdateVoteState(vote)
	if err != nil {
		return err
	}
	err = vs.repository.UpdateCountVote(vote)
	fmt.Println(vote, "vote value")

	if err != nil {
		return err
	}

	return nil
}
