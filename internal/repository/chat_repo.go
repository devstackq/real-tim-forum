package repository

import (
	"database/sql"

	"github.com/devstackq/real-time-forum/internal/models"
)

type ChatRepository struct {
	db *sql.DB
}

func NewChatRepository(db *sql.DB) *ChatRepository {
	return &ChatRepository{db}
}

//good practice ? or service switch type -> 2 differnet repo ?
func (cr *ChatRepository) ChatManager(vote *models.Chat) error {
	//save db message
	return nil
}
