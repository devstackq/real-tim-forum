package repository

import (
	"database/sql"
	"log"

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
func (vr *VoteRepository) GetCountVote(vote *models.Vote) (int, error) {
	query := `SELECT  count_` + vote.VoteType + ` FROM ` + vote.VoteGroup + `s WHERE id=?`
	row := vr.db.QueryRow(query, vote.ID)
	err := row.Scan(&vote.Count)
	if err != nil {
		return -1, err
	}
	return vote.Count, nil
}

func (vr *VoteRepository) UpdateCountVote(vote *models.Vote) error {

	query, err := vr.db.Prepare(`UPDATE ` + vote.VoteGroup + `s SET , ` + vote.VoteType + `s=? WHERE id=?`)
	if err != nil {
		return err
	}
	_, err = query.Exec(vote.Count, vote.ID)

	if err != nil {
		return err
	}
	return nil
}

func (vr *VoteRepository) GetVoteState(vote *models.Vote) (*models.Vote, error) {
	query := `SELECT like_state, dislike_state FROM votes WHERE ` + vote.VoteGroup + `_id=?`
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

func (vr *VoteRepository) SetVoteState(vote *models.Vote) error {

	query, err := vr.db.Prepare(`INSERT INTO votes(` + vote.VoteType + `State, post_id, user_id) VALUES(?,?,?)`)
	if err != nil {
		return err
	}
	_, err = query.Exec(true, vote.ID, vote.CreatorID)
	if err != nil {
		return err
	}
	return nil
}

func (vr *VoteRepository) UpdateVoteState(vote *models.Vote) error {
	// case  !L, D, L, !D,
	query, err := vr.db.Prepare(`UPDATE votes SET like_state, dislike_state? WHERE post_id=? AND user_id=?`)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = query.Exec(vote.LikeState, vote.DislikeState, vote.ID, vote.CreatorID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
