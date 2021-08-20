package service

import (
	"errors"
	"log"
	"net/http"

	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/devstackq/real-time-forum/internal/repository"
)

type PostService struct {
	repository repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo}
}
func (ps *PostService) Create(post *models.Post) (int, error) {

	if IsEmpty(post.Thread) {
		return 200, errors.New("empty thread field")
	} else if IsEmpty(post.Content) {
		return 200, errors.New("empty content field")
	} else if len(post.Categories) == 0 {
		return 200, errors.New("empty category field")
	} else {

		lastID, _, err := ps.repository.CreatePost(post)
		if err != nil {
			return http.StatusBadRequest, err
		}
		// send userid not uuid
		for _, v := range post.Categories {
			if v == "love" {
				err = ps.repository.JoinCategoryPost(lastID, "2")
				if err != nil {
					return -1, err
				}
			} else if v == "nature" {
				err = ps.repository.JoinCategoryPost(lastID, "3")
				if err != nil {
					return -1, err
				}
			} else if v == "science" {
				err = ps.repository.JoinCategoryPost(lastID, "1")
				if err != nil {
					return -1, err
				}
			}
		}
	}
	log.Println(post, "succes create post")
	return http.StatusOK, nil
}

func (ps *PostService) GetPostsByCategory(category string) (*[]models.Post, error) {

	posts, err := ps.repository.GetPostsByCategory(category)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (ps *PostService) GetPostById(postId string) (*models.Post, error) {
	post, err := ps.repository.GetPostById(postId)
	if err != nil {
		log.Println(post, "pid", err)

		return nil, err
	}

	return post, nil
}
