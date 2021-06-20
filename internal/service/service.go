package service

import (
	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/devstackq/real-time-forum/internal/repository"
)

type User interface {
	Create(*models.User) (int, int, error)
	Signin(*models.User) (int, *models.Session, error)
	Logout(*models.Session) error
	GetDataInDb(string, string) (string, error)
	GetUserById(string) (*models.User, error)
	GetUserPosts(string) (*[]models.Post, error)
}
type Post interface {
	Create(*models.Post) (int, error)
	GetPostsByCategory(string) (*[]models.Post, error)
	GetPostById(string) (*models.Post, error)
}
type Vote interface {
	VotePost(*models.Vote) error
}
type Service struct {
	User
	Post
	Vote
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		User: NewUserService(r.User),
		Post: NewPostService(r.Post),
		Vote: NewVoteService(r.Vote)
	}
}
