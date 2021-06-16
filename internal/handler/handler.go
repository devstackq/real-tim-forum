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
			//route /create/post -> if have sesiion & session correct -> createPost -> else signin page
			route.Handler = h.IsCookieValid(route.Handler)
		}
		// if route.UnAuth {
		// route.Handler =   h.UnAuthMiddleware(route.Handler), route /
		// }
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
			Path:    "/api/create/post",
			Handler: h.CreatePost,
			IsAuth:  false,
		},
		{
			Path:    "/api/profile",
			Handler: h.ProfileHandle,
			IsAuth:  true,
		},
	}
}
