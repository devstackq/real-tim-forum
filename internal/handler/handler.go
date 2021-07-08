package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devstackq/real-time-forum/internal/service"
)

type Handler struct {
	Services *service.Service
}
type Route struct {
	Path    string
	Handler http.HandlerFunc
	IsAuth  bool
}

func NewHandler(s *service.Service) *Handler {
	fmt.Println("created handler ")
	return &Handler{s}
}

func (h *Handler) IndexParse(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/index.html")
	fmt.Println("serve file")
}

//valid params  handler or service ?
func (h *Handler) InitRouter() *http.ServeMux {

	routes := h.createRoutes()
	log.Println("created routers")
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("../client/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "../client/index.html")
	// })

	//add middleware each auth route
	for _, route := range routes {
		if route.IsAuth {
			route.Handler = h.IsCookieValid(route.Handler)
		}
		mux.HandleFunc(route.Path, route.Handler)
	}
	return mux
}

func (h *Handler) createRoutes() []Route {

	return []Route{
		{
			Path:    "/",
			Handler: h.IndexParse,
			IsAuth:  false,
		},
		{
			Path:    "/api/signup",
			Handler: h.SignUp,
			IsAuth:  false,
		},
		{
			Path:    "/api/signin",
			Handler: h.SignIn,
			IsAuth:  false,
		},
		{
			Path:    "/api/logout",
			Handler: h.Logout,
			IsAuth:  true,
		},
		{
			Path:    "/api/profile",
			Handler: h.ProfileHandle,
			IsAuth:  true,
		},
		{
			Path:    "/api/post/create",
			Handler: h.CreatePost,
			IsAuth:  true,
		},
		{
			Path:    "/api/post",
			Handler: h.GetPosts,
			IsAuth:  false,
		},
		{
			Path:    "/api/post/id",
			Handler: h.GetPostById,
			IsAuth:  false,
		},
		{
			Path:    "/api/vote",
			Handler: h.VoteItemById,
			IsAuth:  true,
		},
		{
			Path:    "/api/chat",
			Handler: h.ChatHandler,
			IsAuth:  true,
		},
	}
}
