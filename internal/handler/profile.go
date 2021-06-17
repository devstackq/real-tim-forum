package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) ProfileHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		uid, err := r.Cookie("user_id")
		if err != nil {
			JsonResponse(w, r, http.StatusUnauthorized, err)
		}
		//userExist() ?
		user, err := h.Services.User.GetUserById(uid.Value)

		if err != nil {
			JsonResponse(w, r, http.StatusNotFound, err.Error())
			return
		}
		posts, err := h.Services.User.GetUserPosts(uid.Value)

		if err != nil {
			JsonResponse(w, r, http.StatusNotFound, err.Error())
			return
		}
		//getCreatedComment() comment /array comment
		//getVotedPost() vote / array post
		fmt.Println(user, posts)
		JsonResponse(w, r, http.StatusOK, user)
	case "POST":
		//update name, age, etc, delete user, update, delete request
	}
}
