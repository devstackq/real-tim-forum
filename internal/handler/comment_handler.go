package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) LostComment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("get create comment")
	case "POST":
		comment, _, _, _, _, err := GetJsonData(w, r, "comment")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		lastComment, status, err := h.Services.Comment.LostComment(comment)
		if err != nil {
			JsonResponse(w, r, http.StatusInternalServerError, err)
			return
		}
		JsonResponse(w, r, status, lastComment)
	default:
		JsonResponse(w, r, http.StatusBadRequest, "Bad Request")
	}
}
