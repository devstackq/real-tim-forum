package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/devstackq/real-time-forum/internal/models"
)

func (h *Handler) VotePostById(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		fmt.Println("get post by id")
	case "POST":
		vote := &models.Vote{}
		resBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(resBody, vote)
		fmt.Println(string(resBody), "vote post")
		if err != nil {
			JsonResponse(w, r, http.StatusBadRequest, err)
			return
		}
		fmt.Println(vote, "vote data handler")
		err = h.Services.Vote.VotePost(vote)
		if err != nil {
			JsonResponse(w, r, http.StatusInternalServerError, err)
			return
		}
	}
}
