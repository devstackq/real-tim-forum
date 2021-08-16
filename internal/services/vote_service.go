package service

import (
	"fmt"
	"log"

	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/devstackq/real-time-forum/internal/repository"
)

type VoteService struct {
	repository repository.Vote
}

func NewVoteService(repo repository.Vote) *VoteService {
	return &VoteService{repo}
}

func (vs *VoteService) VoteTerminator(vote *models.Vote) (*models.Vote, error) {
	//good practice?
	counts, err := vs.repository.GetVoteCount(vote)
	if err != nil {
		log.Println(err)
	}
	for {
		state, err := vs.repository.GetVoteState(vote)
		//if no rows -> set likeState = true, count+=1

		if err != nil {
			//first row uid, pid
			err = vs.repository.SetVoteState(vote)
			log.Println(vote, 1, state, err)

			if err != nil {
				return nil, err
			}
			if vote.VoteType == "like" {
				counts.CountLike += 1
				vote.LikeState = true
			} else if vote.VoteType == "dislike" {
				counts.CountDislike += 1
				vote.DislikeState = true
			}
			fmt.Println("first vote")
			break
		}
		if !state.LikeState && state.DislikeState {
			if vote.VoteType == "like" {
				vote.LikeState = true
				vote.DislikeState = false
				counts.CountLike += 1
				counts.CountDislike -= 1
			} else if vote.VoteType == "dislike" {
				vote.DislikeState = false
				counts.CountDislike -= 1
			}
			break
		}
		if state.LikeState && !state.DislikeState {
			if vote.VoteType == "like" {
				vote.LikeState = false
				counts.CountLike -= 1
			} else if vote.VoteType == "dislike" {
				vote.DislikeState = true
				vote.LikeState = false
				counts.CountLike -= 1
				counts.CountDislike += 1
			}
			break
		}
		if !state.LikeState && !state.DislikeState {
			if vote.VoteType == "like" {
				counts.CountLike += 1
				vote.LikeState = true
			} else if vote.VoteType == "dislike" {
				vote.DislikeState = true
				counts.CountDislike += 1
			}
			break
		}
	}
	log.Println(vote)

	err = vs.repository.UpdateVoteState(vote)
	if err != nil {
		return nil, err
	}
	vote, err = vs.repository.UpdateVoteCount(counts)
	if err != nil {
		return nil, err
	}
	return vote, nil
}
