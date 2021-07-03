package handler

import (
	"fmt"
	"net/http"
)

//route -> handler -> service -> repos -> dbFunc
func (h *Handler) LostsComment(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Println("get create comment")
	case "POST":
		// uid, err := r.Cookie("user_id")
		_, _, post, _, err := GetJsonData(w, r, "comment")
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
