package service

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/devstackq/real-time-forum/internal/repository"
)

type PostService struct {
	repository repository.Post
}

//repoPost wrapper struct - PostService
func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo}
}
func (ps *PostService) Create(post *models.Post) (int, error) {

	// if !ps.isValid(post) {

	status, err := ps.repository.CreatePost(post)
	if err != nil {
		return http.StatusBadRequest, err
	}
	post.CreatedTime = time.Now()

	log.Println(post, status, "Created post")
	return http.StatusOK, nil

	// } else {
	// 	return http.StatusBadRequest, errors.New("content is empty")
	// }
}

// func (ps *PostService) isImageValid(post *models.Post) bool {
// 	regex := regexp.MustCompile(`^.*\.(jpg|JPG|jpeg|JPEG|gif|GIF|png|PNG|svg|SVG)$`)
// 		// 	mf, fh, _ := r.FormFile("image")
// 		// if fh != nil{
// 		// 	defer mf.Close()
// 		// }
// 	// return regex.MatchString(post.Image)
// 	return true
// }

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
		if !(v <= 32) {
			return false
		}
	}
	return true
}

// regex := regexp.MustCompile(`^.*\.(jpg|JPG|jpeg|JPEG|gif|GIF|png|PNG|svg|SVG)$`)

// mf, fh, _ := r.FormFile("image")
// if fh != nil {
// 	defer mf.Close()
// }

// categories := r.Form["categories"]
// categoryexist := make(map[string]bool)
// thread := r.FormValue("thread")
// content := r.FormValue("content")

// for _, category := range data.Categories {
// 	categoryexist[category] = true
// }

// for _, category := range categories {
// 	if !categoryexist[category] {
// 		data.Data = "Invalid category " + category
// 		w.WriteHeader(http.StatusUnprocessableEntity)
// 		// InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
// 		return
// 	}
// }

// if len(categories) == 0 {
// 	data.Data = "Categories must not be empty"
// 	w.WriteHeader(http.StatusUnprocessableEntity)
// 	// InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
// 	return
// } else if isEmpty(thread) {
// 	data.Data = "Title must not be empty"
// 	w.WriteHeader(http.StatusUnprocessableEntity)
// 	// InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
// 	return
// } else if isEmpty(content) {
// 	data.Data = "Content must not be empty"
// 	w.WriteHeader(http.StatusUnprocessableEntity)
// 	// InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
// 	return
// } else if fh != nil && fh.Size > 20000000 {
// 	data.Data = "File too large, please limit size to 20MB"
// 	w.WriteHeader(http.StatusUnprocessableEntity)
// 	// InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
// 	return
// 	data.Data = "Invalid file type, please upload jpg, jpeg, png, gif, svg"
// 	w.WriteHeader(http.StatusUnprocessableEntity)
// 	// InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
// 	return
// }

func (ps *PostService) GetPostsByCategory(category string) (*[]models.Post, error) {

	posts, err := ps.repository.GetPostsByCategory(category)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (ps *PostService) GetPostById(postId int) (*models.Post, error) {

	post, err := ps.repository.GetPostById(postId)
	if err != nil {
		return nil, err
	}
	return post, nil
}
