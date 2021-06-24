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

func (vs *VoteService) VoteTerminator(vote *models.Vote) (*models.Vote, error) {
	//good practice?
	fmt.Println(vote, "json")
	// counts, err := vs.repository.GetVoteCount(vote)
	for {
		state, err := vs.repository.GetVoteState(vote)
		//if no rows -> set likeState = true, count+=1
		fmt.Println(state, "get votestate")
		if err != nil {
			//first row uid, pid
			err = vs.repository.SetVoteState(vote)
			if err != nil {
				return nil, err
			}
			if vote.VoteType == "like" {
				vote.CountLike += 1
				vote.LikeState = true
			} else if vote.VoteType == "dislike" {
				vote.CountDislike += 1
				vote.DislikeState = true
			}
			fmt.Println("first vote")
			break
		}
		if !state.LikeState && state.DislikeState {
			if vote.VoteType == "like" {
				vote.LikeState = true
				vote.DislikeState = false
				vote.CountLike += 1
				vote.CountDislike -= 1
			} else if vote.VoteType == "dislike" {
				vote.DislikeState = false
				vote.CountDislike -= 1
			}
			break
		}
		if state.LikeState && !state.DislikeState {
			if vote.VoteType == "like" {
				vote.LikeState = false
				vote.CountLike -= 1
			} else if vote.VoteType == "dislike" {
				vote.DislikeState = true
				vote.LikeState = false
				vote.CountLike -= 1
				vote.CountDislike += 1
			}
			break
		}
		if !state.LikeState && !state.DislikeState {
			if vote.VoteType == "like" {
				vote.CountLike += 1
				vote.LikeState = true
			} else if vote.VoteType == "dislike" {
				vote.DislikeState = true
				vote.CountDislike += 1
			}
			break
		}
	}
	err := vs.repository.UpdateVoteState(vote)
	if err != nil {
		return nil, err
	}
	vote, err = vs.repository.UpdateVoteCount(vote)
	if err != nil {
		return nil, err
	}
	fmt.Println(vote, "after")
	return vote, nil
}
check uniq userId, postId - sql
debug like/dislike post -> refactor logic