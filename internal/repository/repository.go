package repository

import (
	"database/sql"

	"github.com/devstackq/real-time-forum/internal/models"
)

type User interface {
	CreateUser(*models.User) (int64, error)
	SigninUser(*models.User) (int, string, error)
	UpdateSession(*models.Session) error
	Logout(*models.Session) error
	GetUuidInDb(string) (string, error)
	GetUserById(string) (*models.User, error)
	GetUserPosts(string) (*[]models.Post, error)
}

type Post interface {
	CreatePost(*models.Post) (int, error)
	GetPostsByCategory(string) (*[]models.Post, error)
	GetPostById(string) (*models.Post, error)
}
type Vote interface {
	VotePost(*models.Vote) error
	VotePostLike(*models.Vote) error
	VotePostDislike(*models.Vote) error
}
type Repository struct {
	User
	Post
	Vote
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
		Post: NewPostRepository(db),
		Vote: NewVoteRepository(db),
	}
}
