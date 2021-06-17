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
	CreatePost(*models.Post) (int64, error)
}

//struct -> User interface -> CreateUser
//User receive db, UserRepo Struct -> have method CreateUser()

type Repository struct {
	User
	Post
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
		Post: NewPostRepository(db),
	}
}

// Model -> CreateUser func, Repo - conn Db, -> service -> handler
