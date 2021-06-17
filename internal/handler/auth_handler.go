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
			JsonResponse(w, r, status, err.Error())
			return
		}
		//user.ID = id
		JsonResponse(w, r, status, id)
	default:
		JsonResponse(w, r, http.StatusBadRequest, "Bad Request")
	}
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("call signin handle Get")
	case "POST":
		fmt.Println("signin Post")
		user := &models.User{}
		resBody, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(resBody, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		status, session, err := h.Services.User.Signin(user)
		if err != nil {
			JsonResponse(w, r, status, "login or password incorrect")
			return
		}
		uid := strconv.Itoa(session.UserID)
		sessionCookie := &http.Cookie{Name: "session", Value: session.UUID, Path: "/", HttpOnly: false, Expires: time.Now().Add(20 * time.Minute)}
		uidCookie := &http.Cookie{Name: "user_id", Value: uid, Path: "/", HttpOnly: false, Expires: time.Now().Add(20 * time.Minute)}

		http.SetCookie(w, sessionCookie)
		http.SetCookie(w, uidCookie)
		//save global variabale session
		JsonResponse(w, r, status, session)
	}
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("call logout handle Get")
		s := models.Session{}
		resBody, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(resBody, &s)
		if err != nil {
			JsonResponse(w, r, http.StatusBadRequest, "")
			return
		}
		err = h.Services.User.Logout(&s)
		if err != nil {
			JsonResponse(w, r, http.StatusBadRequest, "")
			return
		}
		JsonResponse(w, r, http.StatusOK, "logout success")
	case "POST":
	}
}
