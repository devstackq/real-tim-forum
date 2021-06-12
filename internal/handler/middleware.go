package handler

import (
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) IsCookieValid(f http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//check expires cookie
		session, err := r.Cookie("session")
		if err != nil {
			log.Println(err, "expires timeout || cookie deleted")
			// utils.Logout(w, r, *session)
			return
		}
		uid, err := r.Cookie("user_id")
		//check client uuid - in Db Uuid if correct -> goToHandle
		fmt.Println(session, uid, "cookie")
		f.ServeHTTP(w, r)
		// best practice ?

		uuid, err := h.Services.User.GetDataInDb(uid.Value, "uuid")

		if uuid == session {
			log.Println("OK go to hanlde")
		}

		//|| db query here ?
		//cookie Browser -> send IsCookie(check if this user ->)
		// then call handler -> middleware
		// if isValidCookie, sessionF := utils.IsCookie(w, r, c.Value); isValidCookie {
		// 	err = DB.QueryRow("SELECT cookie_time FROM session WHERE user_id = ?", sessionF.UserID).Scan(&sessionF.Time)
		// 	if err != nil {
		// 		log.Println(err)
		// 	}
		// 	strToTime, _ := time.Parse(time.RFC3339, sessionF.Time)
		// 	diff := time.Now().Sub(strToTime)

		// 	if int(diff.Minutes()) > 290 && int(diff.Seconds()) < 298 {
		// 		uuid := utils.CreateUuid()
		// 		utils.SetCookie(w, uuid)
		// 		utils.ReSession(sessionF.UserID, session, "timeout", uuid)
		// 		fmt.Println("change cookie Browser and update sessiontime and uuid in Db")
		// 	}
		// 	*session = sessionF
		// 	f(w, r)
		// }
	}
}
