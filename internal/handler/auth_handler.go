package handler

import (
	"net/http"
	"html/template"
	"github.com/devstackq/real-time-forum/internal/models"
	"fmt"
)

//route -> handler -> service -> repos -> dbFunc
// todo: json api
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl := template.Must(template.ParseFiles("./client/template/signup.html"))
		tmpl.Execute(w, nil)
	case "POST":
		user := &models.User{
			Email: r.FormValue("email"),
			//etc
		}
		// ***
		status, id, err := h.Services.User.Create(user)
		if err != nil {
		//	writeResponse(w, status, err.Error())
			return
		}
		//user.ID = id
		fmt.Println(id, status)
		http.Redirect(w, r, "/signin", http.StatusFound)
	default:
		//writeResponse(w, http.StatusBadRequest, "Bad Request")
	}
}
