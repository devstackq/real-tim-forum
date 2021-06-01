package handler

import (
	"net/http"

	"github.com/devstackq/real-time-forum/internal/service"
)

type Handler struct {
	Services *service.Service
}

type Route struct {
	Path     string
	Handler  http.HandlerFunc
	NeedAuth bool
	UnAuth   bool
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) InitRouter() *http.ServeMux {

	routes := h.createRoutes()
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./client/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	for _, route := range routes {
		if route.NeedAuth {
			//add middleware redirect - signin
			// route.Handler = h.needAuthMiddleware(route.Handler), route /signup
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
		// {
		// 	Path: "/",
		// 	//	Handler: h.Index(), // called
		// 	NeedAuth: false,
		// 	UnAuth:   false,
		// },
		{
			Path:     "/signup",
			Handler:  h.SignUp,
			NeedAuth: false,
			UnAuth:   true,
		},
		// {
		// 	Path: "/signin",
		// 	//Handler:    h.SignIn,
		// 	NeedAuth: false,
		// 	UnAuth:   true,
		// },
		{
			Path: "/createpost",
			Handler: h.CreatePost,
			NeedAuth: true,
			UnAuth: false,
		},
	}
}
