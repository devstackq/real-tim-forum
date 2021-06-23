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
		sessionCookie := &http.Cookie{Name: "session", Value: session.UUID, Path: "/", HttpOnly: false, Expires: time.Now().Add(24 * time.Hour)}
		uidCookie := &http.Cookie{Name: "user_id", Value: uid, Path: "/", HttpOnly: false, Expires: time.Now().Add(24 * time.Hour)}

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
		// posts, err := h.Services.User.GetUserPosts(uid.Value)

		if err != nil {
			JsonResponse(w, r, http.StatusNotFound, err.Error())
			return
		}
		//getCreatedComment() comment /array comment
		//getVotedPost() vote / array post
		// fmt.Println(user, "posts created", posts)
		JsonResponse(w, r, http.StatusOK, user)
	case "POST":
		//update name, age, etc, delete user, update, delete request
	}
}
