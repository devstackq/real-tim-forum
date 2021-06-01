package handler

import (
	"html/template"
	"log"
	"net/http"
)

//connect index html
func (h *Handler) IndexParse(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../client/index.html")
	t.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
