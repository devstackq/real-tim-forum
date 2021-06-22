package repository

import (
	"database/sql"

	"github.com/devstackq/real-time-forum/internal/models"
)

//have db - connect
type VoteRepository struct {
	db *sql.DB
}

func NewVoteRepository(db *sql.DB) *VoteRepository {
	return &VoteRepository{db}
}

//good practice ? or service switch type -> 2 differnet repo ?

func (vr *VoteRepository) getCountLike(vote *models.Vote) (int, error) {

	query := `SELECT count_like FROM ` + vote.VoteGroup + ` WHERE id=?`
	row := vr.db.QueryRow(query, vote.ID)

	err := row.Scan(&vote.CountLike)
	if err != nil {
		return -1, err
	}
	return vote.CountLike, nil
}

func (vr *VoteRepository) getVoteState(vote *models.Vote) (bool, error) {

	query := `SELECT ` + vote.VoteType + `_state FROM ` + vote.VoteGroup + ` WHERE id=?`
	row := vr.db.QueryRow(query, vote.ID)
	err := row.Scan(&vote.VoteState)
	if err != nil {
		return false, err
	}
	//if voteState, like, dislike - 0,1,2 -> if type == like, & voteState == 1, -> voteState -> 0, countLike-1
	// voteState == 0 && type==like -> countLike+=1, voteState -> 1
	// voteState == 1 && type==dislike, countLike-=1, countDislike +=1, likeState = 0, dislikeState = 1
	//etc
	return vote.VoteState, nil

}

func (vr *VoteRepository) VotePostLike(vote *models.Vote) error {
	//postId, userId
	//get Bypostid ->   count like
	//update post -> countLike value

	// isVote, err := vr.getVoteState(vote)
	//setVoteState
	//setCountLike
	//isVote { update by postId - voteState} else insert -> like, dislike

	//if voteState, 1 post - 1 user, vs -
	// getVote state -> table by uid, & postid
	count, _ := vr.getCountLike(vote)
	if count != -1 {

	}
	//update queryb
	return nil
}

func (vr *VoteRepository) VotePostDislike(vote *models.Vote) error {
	return nil
}

func (vr *VoteRepository) VotePost(vote *models.Vote) error {
	return nil
}
