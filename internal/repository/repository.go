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
	CreatePost(*models.Post) (string, int, error)
	GetPostsByCategory(string) (*[]models.Post, error)
	GetPostById(string) (*models.Post, error)
	JoinCategoryPost(string, string) error
}
type Vote interface {
	GetCountVote(*models.Vote) (*models.Vote, error)
	UpdateCountVote(*models.Vote) error
	GetVoteState(*models.Vote) (*models.Vote, error)
	SetVoteState(*models.Vote) error
	UpdateVoteState(*models.Vote) error
}
type Comment interface {
	CreateComment(*models.Comment) (int, error)
}
type Repository struct {
	User
	Post
	Vote
	Comment
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
		Post: NewPostRepository(db),
		Vote: NewVoteRepository(db),
	}
}
