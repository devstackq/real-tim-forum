package handler

import (
	"encoding/json"
	"fmt"
	// "html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	// "log"

	"github.com/devstackq/real-time-forum/internal/models"
)

// curl --header "Content-Type: application/json"   --request POST   --data '{"username":"xyz","password":"xyz", "email":"meow"}'   http://localhost:6969/signup


//route -> handler -> service -> repos -> dbFunc
func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var data *models.PageData
	switch r.Method {
	case "GET":
		// tmpl := template.Must(template.ParseFiles("../client/template/index.html"))
		// tmpl.Execute(w, nil)
		// every index html, change route
	case "POST":
		
		regex := regexp.MustCompile(`^.*\.(jpg|JPG|jpeg|JPEG|gif|GIF|png|PNG|svg|SVG)$`)

		mf, fh, _ := r.FormFile("image")
		if fh != nil {
			defer mf.Close()
		}


		categories := r.Form["categories"]
		categoryexist := make(map[string]bool)
		thread := r.FormValue("thread")
		content := r.FormValue("content")

		for _, category := range data.Categories {
			categoryexist[category] = true
		}

		for _, category := range categories {
			if !categoryexist[category] {
				data.Data = "Invalid category " + category
				w.WriteHeader(http.StatusUnprocessableEntity)
				// InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
				return
			}
		}

		if len(categories) == 0 {
			data.Data = "Categories must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			// InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
			return
		} else if isEmpty(thread) {
			data.Data = "Title must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			// InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
			return
		} else if isEmpty(content) {
			data.Data = "Content must not be empty"
			w.WriteHeader(http.StatusUnprocessableEntity)
			// InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
			return
		} else if fh != nil && fh.Size > 20000000 {
			data.Data = "File too large, please limit size to 20MB"
			w.WriteHeader(http.StatusUnprocessableEntity)
			// InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
			return
		} else if fh != nil && !regex.MatchString(fh.Filename) {
			data.Data = "Invalid file type, please upload jpg, jpeg, png, gif, svg"
			w.WriteHeader(http.StatusUnprocessableEntity)
			// InternalError(w, r, tpl.ExecuteTemplate(w, "createpost.html", data))
			return
		}
		
		
		
		post := &models.Post{}
		resBody, err := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(resBody, post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("after", post, post.Thread)

		// err = repository.CreatePost(&post)
		// if InternalError(w, r, err) {
		// 	return
		// }

		// if fh != nil {
		// 	f, err := os.Create(fmt.Sprintf("./static/images/%v", newPost.PostID))
		// 	defer f.Close()
		// 	_, err = io.Copy(f, mf)
		// 	if InternalError(w, r, err) {
		// 		return
		// 	}
		// }

		// http.Redirect(w, r, fmt.Sprintf("/posts/id/%d", newPost.PostID), http.StatusSeeOther)
		// ***
		// status, id, err := h.Services
		// if err != nil {
		// 	w.WriteHeader(status)
		// 	// tmpl := template.Must(template.ParseFiles("../client/template/error.html"))
		// 	// tmpl.Execute(w, errorData{resp})
		// 	//	writeResponse(w, status, err.Error())
		// 	return
		// }
		//user.ID = id
		// fmt.Println(id, status, "id  status")

		// http.Redirect(w, r, "/signin", http.StatusFound)
	default:
		//writeResponse(w, http.StatusBadRequest, "Bad Request")
	}
}

func isEmpty(text string) bool {
	for _, v := range text {
		if !(v <= 32) {
			return false
		}
	}
	return true
}