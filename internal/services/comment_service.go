package service

import (
	"net/http"

	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/devstackq/real-time-forum/internal/repository"
)

type CommentService struct {
	repository repository.Comment
}

func NewCommentService(repo repository.Comment) *CommentService {
	return &CommentService{repo}
}

func (cs *CommentService) LostComment(comment *models.Comment) (*models.Comment, int, error) {
	last, err := cs.repository.CreateComment(comment)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	return last, http.StatusOK, nil
}
