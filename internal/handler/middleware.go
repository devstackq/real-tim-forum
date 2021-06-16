package handler

import (
	"log"
	"net/http"
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
		// best practice ?
		if err != nil {
			log.Println("userid expires or incorrect")
			JsonResponse(w, r, http.StatusUnauthorized, "userId incorrect")
			return
		}
		uuid, err := h.Services.User.GetDataInDb(userId.Value, "uuid")

		if uuid == session.Value {
			log.Println("call hanlde, ok")
			JsonResponse(w, r, http.StatusOK, "all right")
			f.ServeHTTP(w, r)
		} else {
			log.Println("session not equal in Db")
			JsonResponse(w, r, http.StatusUnauthorized, "cookie incorrect")
		}
		// 	*session = sessionF
	}
}
