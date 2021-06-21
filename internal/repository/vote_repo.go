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

func (vr *VoteRepository) VotePost(vote *models.Vote) error {
	return nil
}
