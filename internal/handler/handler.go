package handler

import (
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
	UnAuth  bool
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) InitRouter() *http.ServeMux {

	routes := h.createRoutes()
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("../client/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	for _, route := range routes {
		if route.IsAuth {
			//add middleware redirect - signin
			//route /create/post -> if have sesiion & session correct -> createPost -> else signin page
			route.Handler = h.IsCookieValid(route.Handler)
		}
		if route.UnAuth {
			//addMiddleware redirect =
			// route.Handler =   h.UnAuthMiddleware(route.Handler), route /
		}
		//default
		// route.Handler = h.CookieIsValid(route.Handler) //
		//add  mux, handler
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
			UnAuth:  false,
		},
		{
			Path:    "/api/signup",
			Handler: h.SignUp,
			IsAuth:  false,
			UnAuth:  true,
		},
		{
			Path:    "/api/signin",
			Handler: h.SignIn,
			IsAuth:  false,
			UnAuth:  true,
		},
		{
			Path:    "/api/create/post",
			Handler: h.CreatePost,
			IsAuth:  false,
			UnAuth:  true,
		},
		{
			Path:    "/api/profile",
			Handler: h.ProfileHandle,
			IsAuth:  true,
			UnAuth:  false,
		},
	}
}
