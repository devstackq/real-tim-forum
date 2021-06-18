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
		post := &models.Post{}
		resBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(string(resBody))
		err = json.Unmarshal(resBody, post)

		// err := h.Services.Vote.VotePost(type)
	}

}
