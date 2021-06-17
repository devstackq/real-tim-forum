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
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(string(resBody), "json")

		err = json.Unmarshal(resBody, post)
		if err != nil {
			fmt.Println(err, "error")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println("after", post, post.Content)

		status, err := h.Services.Post.Create(post)

		if err != nil {
			fmt.Println(err, "error")
			JsonResponse(w, r, http.StatusInternalServerError, err.Error())
			return
		}
		log.Println(status, "id  status")
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
