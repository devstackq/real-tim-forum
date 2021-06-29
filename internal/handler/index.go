package handler

import (
	"html/template"
	"net/http"
)

func (h *Handler) IndexParse(w http.ResponseWriter, r *http.Request) {
	//1 time, execute template - like bridge back to front
	var count int
	if count == 0 {
		t, err := template.ParseFiles("../client/index.html")
		t.Execute(w, nil)
		if err != nil {
			JsonResponse(w, r, http.StatusInternalServerError, err)
			return
		}
		count++
	}
}
