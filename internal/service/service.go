package service

import (
	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/devstackq/real-time-forum/internal/repository"
)

type User interface {
	Create(*models.User) (int, int, error)
	// Delete(id int) error
}
type Service struct {
	User
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		User: NewUserService(r.User),
	}
}
