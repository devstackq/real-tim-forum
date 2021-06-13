package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) ProfileHandle(w http.ResponseWriter, r *http.Request) {
	//show statistics user: data, created post/comment, lost votes
	//after update, delete profile
	//get data from middleware -> query in Db, by userId
	switch r.Method {
	case "GET":
		fmt.Println("call proflie handle Get, need Auth")
		//posts/comment/votes/ user datas
		JsonResponse(w, r, http.StatusOK, "user data")
	case "POST":
		//update name, age, etc
		//delete user
		//update, delete request
	}
}
