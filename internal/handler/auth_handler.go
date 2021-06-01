package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/devstackq/real-time-forum/internal/models"
)

//route -> handler -> service -> repos -> dbFunc
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// JsonResponse(w, http.StatusOK, "signup page")
	case "POST":
		user := &models.User{}
		resBody, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(resBody, user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		status, id, err := h.Services.User.Create(user)
		if err != nil {
			w.WriteHeader(status)
			JsonResponse(w, status, err.Error())
			return
		}
		//user.ID = id
		JsonResponse(w, status, id)
		// http.Redirect(w, r, "/signin", http.StatusFound)
	default:
		//JsonResponse(w, http.StatusBadRequest, "Bad Request")
	}
}
