package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/devstackq/real-time-forum/internal/models"
)

//getAllPost / PostByCategory() todo:

//route -> handler -> service -> repos -> dbFunc
func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Println("get create post")
	case "POST":
		post := &models.Post{}
		resBody, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(resBody, post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println("after", post, post.Content)

		// ***
		status, id, err := h.Services.Post.Create(post)
		if err != nil {
			w.WriteHeader(status)

			return
		}
		//user.ID = id
		log.Println(id, status, "id  status")

		http.Redirect(w, r, "/", http.StatusFound)

	default:
		// writeResponse(w, http.StatusBadRequest, "Bad Request")
	}
}

func isEmpty(text string) bool {
	for _, v := range text {
		if !(v <= 32) {
			return false
		}
	}
	return true
}
