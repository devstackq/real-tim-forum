package repository

import (
	"database/sql"
	"fmt"

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
	query := `SELECT count_like FROM ` + vote.VoteGroup + `s WHERE id=?`
	row := vr.db.QueryRow(query, vote.ID)
	err := row.Scan(&vote.CountLike)
	fmt.Println(err, "query getCountVote")
	if err != nil {
		return -1, err
	}
	return vote.CountLike, nil
}

func (vr *VoteRepository) getVoteState(vote *models.Vote) (*models.Vote, error) {

	// where := vote.VoteGroup[:len(vote.VoteGroup)-1]
	// fmt.Println(where)
	query := `SELECT like_state, dislike_state FROM voteState WHERE ` + vote.VoteGroup + `_id=?`
	row := vr.db.QueryRow(query, vote.ID)
	err := row.Scan(&vote.LikeState, &vote.DislikeState)
	if err != nil {
		return nil, err
	}
	//if voteState, like, dislike - 0,1,2 -> if type == like, & voteState == 1, -> voteState -> 0, countLike-1
	// voteState == 0 && type==like -> countLike+=1, voteState -> 1
	// voteState == 1 && type==dislike, countLike-=1, countDislike +=1, likeState = 0, dislikeState = 1
	return vote, nil
}
func (vr *VoteRepository) setVoteState(vote *models.Vote) error {

	query, err := vr.db.Prepare(`INSERT INTO voteState(likeState, post_id, user_id) VALUES(?,?,?)`)
	if err != nil {
		return err
	}
	_, err = query.Exec(true, vote.ID, vote.CreatorID)
	if err != nil {
		return err
	}
	return nil
}
func (vr *VoteRepository) VoteLike(vote *models.Vote) error {

	//postId, userId
	//get Bypostid ->   count like
	//update post -> countLike value

	//isVote, err := vr.getVoteState(vote)
	//setVoteState
	//setCountLike
	//isVote { update by postId - voteState} else insert -> like, dislike

	//if voteState, 1 post - 1 user, vs -
	// getVote state -> table by uid, & postid

	//good practice ?
	// getState, write logic, change countLike, chageVoteState
	//3 state, !Like & !Dis -> Like, count+=1, Like & !Dis-> !Like, like-=1, !Like && Dis -> Like, !Dis, count+1

	states, err := vr.getVoteState(vote)
	//if no rows -> set likeState = true, count+=1
	if err != nil {
		//first row uid, pid
		err = vr.setVoteState(vote)
		if err != nil {
			return err
		}
	}

	fmt.Println(err, "err", "state", states)
	//if count == nil
	count, err := vr.getCountLike(vote)
	if err != nil {
		return err
	}
	fmt.Println(count, "count")
	//update queryb
	return nil
}

func (vr *VoteRepository) VoteDislike(vote *models.Vote) error {
	return nil
}
