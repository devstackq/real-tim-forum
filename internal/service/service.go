package service

import (
	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/devstackq/real-time-forum/internal/repository"
)

type User interface {
	Create(*models.User) (int, int, error)
	// Delete(id int) error
}
type Post interface {
	Create(*models.Post) (int, int, error)
}
type Service struct {
	User
	Post
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		User: NewUserService(r.User),
		Post: NewPostService(r.Post),
	}
}
