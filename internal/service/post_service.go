package service

import (
	"fmt"
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

	if ps.isValid(post) {
		lastID, status, err := ps.repository.CreatePost(post)
		if err != nil {
			return http.StatusBadRequest, err
		}

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
		log.Println(post, status, "Created post")
	}
	return http.StatusOK, nil
}

func (ps *PostService) isValid(post *models.Post) bool {
	var (
		isEmptyContent = false
		isEmptyThread  = false
		// largeFile = false
		// invalidFile = false
	)
	// regex := regexp.MustCompile(`^.*\.(jpg|JPG|jpeg|JPEG|gif|GIF|png|PNG|svg|SVG)$`)
	if ps.isEmpty(post.Content) {
		isEmptyContent = true
	} else if ps.isEmpty(post.Thread) {
		isEmptyThread = true
		// return
	}
	return !isEmptyContent && !isEmptyThread
}

func (ps *PostService) isEmpty(text string) bool {
	fmt.Println(text)
	for _, v := range text {
		fmt.Println("ture", v)

		if !(v <= 32) {
			return false
		}
	}
	return true
}

//todo:
// regex := regexp.MustCompile(`^.*\.(jpg|JPG|jpeg|JPEG|gif|GIF|png|PNG|svg|SVG)$`)
// if len(categories) == 0 {
// 	data.Data = "Categories must not be empty"
// } else if isEmpty(thread) {
// 	data.Data = "Title must not be empty"
// } else if isEmpty(content) {
// 	data.Data = "Content must not be empty"
// } else if fh != nil && fh.Size > 20000000 {
// 	data.Data = "File too large, please limit size to 20MB"
// 	w.WriteHeader(http.StatusUnprocessableEntity)
// 	data.Data = "Invalid file type, please upload jpg, jpeg, png, gif, svg"
// 	w.WriteHeader(http.StatusUnprocessableEntity)
// }

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
		return nil, err
	}
	return post, nil
}
