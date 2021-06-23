package handler

import (
	"fmt"
	"net/http"
)

//route -> handler -> service -> repos -> dbFunc
func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		fmt.Println("get create post")
	case "POST":
		// uid, err := r.Cookie("user_id")
		post, _, err := GetJsonData(w, r, "post")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		status, err := h.Services.Post.Create(post)

		if err != nil {
			JsonResponse(w, r, http.StatusInternalServerError, err)
			return
		}
		JsonResponse(w, r, status, "success")
	default:
		JsonResponse(w, r, http.StatusBadRequest, "Bad Request")
	}
}

func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		category, err := r.Cookie("category")
		if err != nil {
			JsonResponse(w, r, http.StatusBadRequest, "cookie incorrect")
			return
		}
		posts, err := h.Services.GetPostsByCategory(category.Value[1:])
		if err != nil {
			JsonResponse(w, r, http.StatusInternalServerError, err)
			return
		}
		JsonResponse(w, r, http.StatusOK, posts)
	default:
		JsonResponse(w, r, http.StatusBadRequest, "Bad Request")
	}
}

func (h *Handler) GetPostById(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		fmt.Println("get post by id")
	case "POST":
		post, _, err := GetJsonData(w, r, "post")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//postHaveById(id) valid ?, if Post.ID > 0 && post.ID < lastInsertedPost or getCountPosts
		result, err := h.Services.GetPostById(post.ID)
		if err != nil {
			JsonResponse(w, r, http.StatusInternalServerError, err)
			return
		}
		JsonResponse(w, r, http.StatusOK, result)
	default:
		JsonResponse(w, r, http.StatusBadRequest, "Bad Request")
	}
}
