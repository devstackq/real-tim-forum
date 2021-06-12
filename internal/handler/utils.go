package handler

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// fmt.Println(string(js), "send data")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
