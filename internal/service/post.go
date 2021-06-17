package service

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/devstackq/real-time-forum/internal/models"
	"github.com/devstackq/real-time-forum/internal/repository"
	sqlite "github.com/mattn/go-sqlite3"
)

type PostService struct {
	repository repository.Post
}

//repoPost wrapper struct - PostService
func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo}
}
func (ps *PostService) Create(post *models.Post) (int, int, error) {

	if !ps.isValid(post) {

		lastID, err := ps.repository.CreatePost(post)

		post.CreatedTime = time.Now()
		if err != nil {
			if sqliteErr, ok := err.(sqlite.Error); ok {
				if sqliteErr.ExtendedCode == sqlite.ErrConstraintUnique {
					return http.StatusBadRequest, -1, errors.New("Post already created")
				}
			}
			return http.StatusInternalServerError, -1, err
		}

		log.Println(post, "Create service")
		return http.StatusOK, int(lastID), nil
	} else {
		return http.StatusBadRequest, 0, errors.New("content is empty")
	}
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
