package repository

import (
	"database/sql"

	"github.com/devstackq/real-time-forum/internal/models"
)

type User interface {
	CreateUser(*models.User) (int64, error)
	SigninUser(*models.User) (int64, error)
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
// route -> handle ->service-> repo -> dbFunc
