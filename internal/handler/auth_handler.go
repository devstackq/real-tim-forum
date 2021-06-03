package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/devstackq/real-time-forum/internal/models"
)

// add in files -> changes Yelnur, push -> pull Yelanur??

//route -> handler -> service -> repos -> dbFunc
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("call signup handle Get")
		// JsonResponse(w, http.StatusOK, "signup page")
	case "POST":
		fmt.Println("call signup handle Post")
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
