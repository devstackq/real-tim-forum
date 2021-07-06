package handler

import (
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) IsCookieValid(f http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//check expires cookie
		session, err := r.Cookie("session")
		if err != nil {
			log.Println("session expires")
			JsonResponse(w, r, http.StatusUnauthorized, "cookie expires or not correct")
			return
		}
		userId, err := r.Cookie("user_id")
		if err != nil {
			log.Println("userid expires or incorrect")
			JsonResponse(w, r, http.StatusUnauthorized, "userId incorrect")
			return
		}
		//comapre db & browser cookie
		uuid, err := h.Services.User.GetDataInDb(userId.Value, "uuid")
		if err != nil {
			JsonResponse(w, r, http.StatusUnauthorized, "cant get cookie in db")
			return
		}
		name, err := h.Services.User.GetDataInDb(userId.Value, "name")
		log.Println(err)
		if err != nil {
			JsonResponse(w, r, http.StatusInternalServerError, "cant get name in db")
			return
		}
		if uuid == session.Value {
			//set Authorized variable - when user signin
			uid, _ := strconv.Atoi(userId.Value)
			Authorized.UUID = uuid
			Authorized.UserID = uid
			Authorized.Name = name
			f.ServeHTTP(w, r)
		} else {
			log.Println("session is not equal in Db")
			JsonResponse(w, r, http.StatusUnauthorized, "cookie incorrect")
		}
	}
}
