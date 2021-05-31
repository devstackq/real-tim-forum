package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/devstackq/real-time-forum/internal/models"
)

// curl --header "Content-Type: application/json"   --request POST   --data '{"username":"xyz","password":"xyz", "email":"meow"}'   http://localhost:6969/signup

type Test struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//route -> handler -> service -> repos -> dbFunc
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl := template.Must(template.ParseFiles("../client/template/signup.html"))
		tmpl.Execute(w, nil)
	case "POST":
		user := &models.User{}
		resBody, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(resBody, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("after", user, user.Email)

		// ***
		status, id, err := h.Services.User.Create(user)
		if err != nil {
			w.WriteHeader(status)
			// tmpl := template.Must(template.ParseFiles("../client/template/error.html"))
			// tmpl.Execute(w, errorData{resp})
			//	writeResponse(w, status, err.Error())
			return
		}
		//user.ID = id
		fmt.Println(id, status, "id  status")
		http.Redirect(w, r, "/signin", http.StatusFound)
	default:
		//writeResponse(w, http.StatusBadRequest, "Bad Request")
	}
}
