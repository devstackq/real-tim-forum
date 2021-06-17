package handler

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	r.Header.Add("Accept-Charset", "UTF-8;q=1, ISO-8859-1;q=0")
	w.WriteHeader(status)
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, string(js), status)
		return
	}
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	// fmt.Println(string(js), "send data client")
	w.Write(js)
}
