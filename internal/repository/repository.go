package repository

import (
	"database/sql"

	"github.com/devstackq/real-time-forum/internal/models"
)

type User interface {
	CreateUser(*models.User) (int64, error)
	SigninUser(*models.User) (int, string, error)
	Logout(string) error
	UpdateSession(*models.Session) error
	GetUserUuid(string) (string, error)
	GetUserName(string) (string, error)
	GetUserById(string) (*models.User, error)
	GetCreatedUserPosts(int) (*[]models.Post, error)
	GetUserVotedItems(int) (*[]models.Vote, error)
}

type Post interface {
	CreatePost(*models.Post) (string, int, error)
	GetPostsByCategory(string) (*[]models.Post, error)
	GetPostById(string) (*models.Post, error)
	JoinCategoryPost(string, string) error
}
type Vote interface {
	GetVoteCount(*models.Vote) (*models.Vote, error)
	UpdateVoteCount(*models.Vote) (*models.Vote, error)
	GetVoteState(*models.Vote) (*models.Vote, error)
	SetVoteState(*models.Vote) error
	UpdateVoteState(*models.Vote) error
}
type Comment interface {
	CreateComment(*models.Comment) (int, error)
}
type Chat interface {
	ChatManager(*models.Chat) error
	GetMessages(m *models.Message) ([]models.Message, error)
}
type Repository struct {
	User
	Post
	Vote
	Comment
	Chat
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
		Post: NewPostRepository(db),
		Vote: NewVoteRepository(db),
		Chat: NewChatRepository(db),
	}
}
