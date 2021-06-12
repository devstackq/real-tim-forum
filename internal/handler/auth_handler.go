package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/devstackq/real-time-forum/internal/models"
)

//route -> handler -> service -> repos -> dbFunc
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("call signup handle Get")
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
		fmt.Println("send id", err, id)

		if err != nil {
			JsonResponse(w, status, err.Error())
			return
		}
		//user.ID = id
		JsonResponse(w, status, id)
	default:
		JsonResponse(w, http.StatusBadRequest, "Bad Request")
	}
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("call signin handle Get")
		// JsonResponse(w, http.StatusOK, "signin page")
	case "POST":
		user := &models.User{}
		resBody, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(resBody, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		status, session, err := h.Services.User.Signin(user)
		if err != nil {
			JsonResponse(w, status, "login or password incorrect")
			return
		}
		uid := strconv.Itoa(session.UserID)
		sessionCookie := &http.Cookie{Name: "session", Value: session.UUID, HttpOnly: false, Expires: time.Now().Add(20 * time.Minute)}
		uidCookie := &http.Cookie{Name: "user_id", Value: uid, HttpOnly: false, Expires: time.Now().Add(20 * time.Minute)}

		http.SetCookie(w, sessionCookie)
		http.SetCookie(w, uidCookie)
		//save global variabale session
		JsonResponse(w, status, session)
	}
}
