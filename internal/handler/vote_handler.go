package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) VoteItemById(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Println("vote post by id")
	case "POST":
		vote, _, _, err := GetJsonData(w, r, "vote")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// if vote.VoteGroup == "post" {
		updatedVote, err := h.Services.Vote.VoteTerminator(vote)
		if err != nil {
			JsonResponse(w, r, http.StatusInternalServerError, err)
			return
		}
		JsonResponse(w, r, http.StatusOK, updatedVote)
	default:
		JsonResponse(w, r, http.StatusBadRequest, "Bad Request")
	}
}
