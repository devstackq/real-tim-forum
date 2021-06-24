package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var Authorized struct {
	UUID   string
	UserID int
}

//route -> handler -> service -> repos -> dbFunc
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("call signup handle Get")
	case "POST":
		_, _, user, err := GetJsonData(w, r, "user")
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
		_, _, user, err := GetJsonData(w, r, "user")
		if err != nil {
			JsonResponse(w, r, http.StatusBadRequest, err)
			return
		}
		status, session, err := h.Services.User.Signin(user)
		if err != nil {
			JsonResponse(w, r, status, "login or password incorrect")
			return
		}
		//set Authorized variable - when user signin
		Authorized.UUID = session.UUID
		Authorized.UserID = session.UserID
		//set in browser
		sessionCookie := &http.Cookie{Name: "session", Value: session.UUID, Path: "/", HttpOnly: false, Expires: time.Now().Add(24 * time.Hour)}
		uidCookie := &http.Cookie{Name: "user_id", Value: strconv.Itoa(session.UserID), Path: "/", HttpOnly: false, Expires: time.Now().Add(24 * time.Hour)}
		http.SetCookie(w, sessionCookie)
		http.SetCookie(w, uidCookie)
		JsonResponse(w, r, status, session)
	default:
		JsonResponse(w, r, http.StatusBadRequest, "Bad Request")
	}
}

//it func need ? -> each time update session
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		err := h.Services.User.Logout(Authorized.UUID)
		if err != nil {
			JsonResponse(w, r, http.StatusBadRequest, "")
			return
		}
		JsonResponse(w, r, http.StatusOK, "success")
	case "POST":
	default:
		JsonResponse(w, r, http.StatusBadRequest, "Bad Request")
	}
}

//handlers DRY ?
func (h *Handler) ProfileHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//userExist() ?
		user, err := h.Services.User.GetUserById(strconv.Itoa(Authorized.UserID))

		if err != nil {
			JsonResponse(w, r, http.StatusNotFound, err.Error())
			return
		}
		//posts, err := h.Services.User.GetUserPosts(uid.Value)
		//getCreatedComment() comment /array comment
		//getVotedPost() vote / array post
		JsonResponse(w, r, http.StatusOK, user)
	case "POST":
		//update name, age, etc, delete user, update, delete request
	default:
		JsonResponse(w, r, http.StatusBadRequest, "Bad Request")
	}
}
