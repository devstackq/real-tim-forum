package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/devstackq/real-time-forum/internal/models"
)

//connect index html
func (h *Handler) IndexParse(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		t, err := template.ParseFiles("../client/index.html")
		t.Execute(w, nil)
		category := r.URL.Path[1:]

		if err != nil {
			log.Println(err)
		}

		if r.URL.Path == "/" {
			JsonResponse(w, r, http.StatusOK, "all")
		}
		if category == "love" {
			//service.GetPostsByCategory( category )
			JsonResponse(w, r, http.StatusOK, "love")
		}
		if category == "nature" {
			JsonResponse(w, r, http.StatusOK, "nature")
		}
		if category == "science" {
			JsonResponse(w, r, http.StatusOK, "science")
		}

	case "POST":

		category := models.Category{}

		resBody, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(resBody, &category.Name)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if category.Name == "/" {
			//all post
			fmt.Println(category.Name, "all")
		} else if category.Name == "/love" {
			fmt.Println(category.Name)
		} else if category.Name == "/nature" {
			fmt.Println(category.Name)
		} else if category.Name == "/science" {
			fmt.Println(category.Name)
		}
		// status, id, err := h.Services.User.Create(user)

	default:
		JsonResponse(w, r, http.StatusBadRequest, "Bad Request")
	}
}
