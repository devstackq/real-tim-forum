package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//connect index html
func (h *Handler) IndexParse(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		t, err := template.ParseFiles("../client/index.html")
		t.Execute(w, nil)
		category := r.URL.Path[1:]

		if err != nil {
			log.Println(err, 1)
		}

		if r.URL.Path == "/" {
			JsonResponse(w, r, http.StatusOK, "all")
		}
		if category == "love" {
			posts, err := h.Services.GetPostsByCategory(category)
			fmt.Println(posts, err, "hande")
			JsonResponse(w, r, http.StatusOK, posts)
		}
		if category == "nature" {
			JsonResponse(w, r, http.StatusOK, "nature")
		}
		if category == "science" {
			JsonResponse(w, r, http.StatusOK, "science")
		}

	case "POST":
	//handle postById()

	default:
		JsonResponse(w, r, http.StatusBadRequest, "Bad Request")
	}
}
