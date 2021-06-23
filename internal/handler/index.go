package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func (h *Handler) IndexParse(w http.ResponseWriter, r *http.Request) {
	//1 time, execute template - like bridge back to front
	var count int
	if count == 0 {
		fmt.Println(count)
		t, err := template.ParseFiles("../client/index.html")
		t.Execute(w, nil)
		if err != nil {
			JsonResponse(w, r, http.StatusInternalServerError, err)
		}
		count++
	}
}
