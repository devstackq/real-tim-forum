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
			log.Println("session expires or incorrect")
			// utils.Logout(w, r, *session)
			JsonResponse(w, r, http.StatusUnauthorized, "cookie expires or not correct")
			return
		}
		// fmt.Println(r, "qwe")
		userId, err := r.Cookie("user_id")
		// best practice ?
		if err != nil {
			log.Println("userid expires or incorrect")
			JsonResponse(w, r, http.StatusUnauthorized, "userId incorrect")
			return
		}
		uuid, err := h.Services.User.GetDataInDb(userId.Value, "uuid")

		if uuid == session.Value {
			log.Println("OK go to hanlde")
			JsonResponse(w, r, http.StatusOK, "all right")
			f.ServeHTTP(w, r)
		} else {
			log.Println("session not equal in Db")

			JsonResponse(w, r, http.StatusUnauthorized, "cookie incorrect")
		}
		// 	*session = sessionF
	}
}
