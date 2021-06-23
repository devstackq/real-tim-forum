package service

import (
	"log"
	"net/http"
	"time"

	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/devstackq/real-time-forum/internal/repository"
)

type CommentService struct {
	repository repository.Comment
}

func NewCommentService(repo repository.Comment) *CommentService {
	return &CommentService{repo}
}

func (cs *CommentService) Create(comment *models.Comment) (int, error) {
	status, err := cs.repository.CreateComment(comment)
	if err != nil {
		return http.StatusBadRequest, err
	}
	comment.CreatedTime = time.Now()

	log.Println(comment, status, "Created comment")
	return http.StatusOK, nil
}
