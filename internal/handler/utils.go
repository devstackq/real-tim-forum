package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/devstackq/real-time-forum/internal/models"
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

func GetJsonData(w http.ResponseWriter, r *http.Request, signature string) (*models.Comment, *models.Message, *models.Vote, *models.Post, *models.User, error) {

	var v models.Vote
	var p models.Post
	var u models.User
	var m models.Message
	var c models.Comment

	// err = json.Unmarshal(resBody, &j.M)
	resBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	if signature == "message" {
		err = json.Unmarshal(resBody, &m)
	}
	if signature == "vote" {
		err = json.Unmarshal(resBody, &v)
	}
	if signature == "post" {
		err = json.Unmarshal(resBody, &p)
	}
	if signature == "user" {
		err = json.Unmarshal(resBody, &u)
	}
	if signature == "comment" {
		err = json.Unmarshal(resBody, &c)
	}

	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	return &c, &m, &v, &p, &u, nil
}
