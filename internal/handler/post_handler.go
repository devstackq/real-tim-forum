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
		uid, _ := r.Cookie("user_id")
		post.CreatorID = uid.Value

		status, err := h.Services.Post.Create(post)

		if err != nil {
			fmt.Println(err, "error")
			JsonResponse(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		log.Println(status, "id  status")
		//or redierct created cpost ?
		// http.Redirect(w, r, "/", http.StatusFound)
	default:
		// writeResponse(w, http.StatusBadRequest, "Bad Request")
	}
}
func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		c, err := r.Cookie("category")
		if err != nil {
			JsonResponse(w, r, http.StatusBadRequest, "cookie incorrect")
		}
		posts, err := h.Services.GetPostsByCategory(c.Value[1:])
		// fmt.Println(err, 22)
		if err != nil {
			JsonResponse(w, r, http.StatusBadRequest, err)
			return
		}
		JsonResponse(w, r, http.StatusOK, posts)
	}
}

func (h *Handler) GetPostById(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		fmt.Println("get post by id")
	case "POST":
		post := &models.Post{}
		resBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(resBody, post)
		if err != nil {
			JsonResponse(w, r, http.StatusBadRequest, err)
			return
		}
		result, err := h.Services.GetPostById(post.ID)
		if err != nil {
			fmt.Println(err)
			JsonResponse(w, r, http.StatusBadRequest, err)
			return
		}
		JsonResponse(w, r, http.StatusOK, result)
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
