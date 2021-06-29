package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/devstackq/real-time-forum/internal/models"
)

var Authorized struct {
	UUID   string
	UserID int
}

var Profile struct {
	Posts      *[]models.Post
	User       *models.User
	VotedItems *[]models.Vote
	// Comment *models.Comment{}
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
		var err error
		Profile.User, err = h.Services.User.GetUserById(strconv.Itoa(Authorized.UserID))
		if err != nil {
			JsonResponse(w, r, http.StatusNotFound, err.Error())
			return
		}
		Profile.Posts, err = h.Services.User.GetCreatedUserPosts(Authorized.UserID)
		if err != nil {
			JsonResponse(w, r, http.StatusNotFound, err.Error())
			return
		}
		Profile.VotedItems, err = h.Services.User.GetUserVotedItems(Authorized.UserID)
		if err != nil {
			JsonResponse(w, r, http.StatusNotFound, err.Error())
			return
		}
		JsonResponse(w, r, http.StatusOK, Profile)

		//getCreatedComment() comment /array comment
	case "POST":
		//update name, age, etc, delete user, update, delete request
	default:
		JsonResponse(w, r, http.StatusBadRequest, "Bad Request")
	}
}
