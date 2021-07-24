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
func (vr *VoteRepository) GetVoteCount(vote *models.Vote) (*models.Vote, error) {
	query := `SELECT  count_like, count_dislike FROM ` + vote.VoteGroup + `s WHERE id=? `
	row := vr.db.QueryRow(query, vote.ID)
	err := row.Scan(&vote.CountLike, &vote.CountDislike)
	if err != nil {
		return nil, err
	}
	return vote, nil
}

//ref $1, vote.VoteGroup+"s"
func (vr *VoteRepository) UpdateVoteCount(vote *models.Vote) (*models.Vote, error) {

	query, err := vr.db.Prepare(`UPDATE ` + vote.VoteGroup + `s SET count_like=?, count_dislike=? WHERE id=?`)
	if err != nil {
		return nil, err
	}
	_, err = query.Exec(vote.CountLike, vote.CountDislike, vote.ID)

	if err != nil {
		return nil, err
	}
	return vote, nil
}

func (vr *VoteRepository) GetVoteState(vote *models.Vote) (*models.Vote, error) {
	query := `SELECT like_state, dislike_state FROM votes WHERE ` + vote.VoteGroup + `_id=? AND user_id=?`
	row := vr.db.QueryRow(query, vote.ID, vote.CreatorID)
	err := row.Scan(&vote.LikeState, &vote.DislikeState)
	if err != nil {
		return nil, err
	}
	return vote, nil
}

func (vr *VoteRepository) SetVoteState(vote *models.Vote) error {

	query, err := vr.db.Prepare(`INSERT INTO votes(` + vote.VoteType + `_state, post_id, user_id) VALUES(?,?,?)`)
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
	query, err := vr.db.Prepare(`UPDATE votes SET like_state=?, dislike_state=? WHERE post_id=? AND user_id=?`)
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
