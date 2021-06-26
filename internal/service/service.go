package service

import (
	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/devstackq/real-time-forum/internal/repository"
)

type User interface {
	Create(*models.User) (int, int, error)
	Signin(*models.User) (int, *models.Session, error)
	Logout(string) error
	GetDataInDb(string, string) (string, error)
	GetUserById(string) (*models.User, error)
	GetCreatedUserPosts(int) (*[]models.Post, error)
	GetUserVotedItems(int) (*[]models.Vote, error)
}
type Post interface {
	Create(*models.Post) (int, error)
	GetPostsByCategory(string) (*[]models.Post, error)
	GetPostById(string) (*models.Post, error)
}
type Comment interface {
	Create(*models.Post) (int, error)
}
type Vote interface {
	VoteTerminator(*models.Vote) (*models.Vote, error)
}

type Service struct {
	User
	Post
	Vote
	Comment
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		User: NewUserService(r.User),
		Post: NewPostService(r.Post),
		Vote: NewVoteService(r.Vote),
		// Comment: NewCommentService(r.Comment),
	}
}
