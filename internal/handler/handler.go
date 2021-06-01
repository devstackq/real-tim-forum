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
	fileServer := http.FileServer(http.Dir("../client/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	for _, route := range routes {
		if route.NeedAuth {
			// route.Handler = h.needAuthMiddleware(route.Handler), route /signin
		}
		if route.UnAuth {
			// route.Handler =   h.UnAuthMiddleware(route.Handler), route /
		}
		//default
		// route.Handler = h.CookieIsValid(route.Handler) //
		mux.HandleFunc(route.Path, route.Handler)
	}
	return mux
}

// /h.Signup undefined (type *Handler has no field or method Signup, but does have SignUp)
// cannot use h.IndexParse (type func(http.ResponseWriter)) as type http.HandlerFunc in field valu
func (h *Handler) createRoutes() []Route {

	return []Route{
		{
			Path:     "/",
			Handler:  h.IndexParse,
			NeedAuth: false,
			UnAuth:   false,
		},
		{
			Path:     "/signup",
			Handler:  h.SignUp,
			NeedAuth: false,
			UnAuth:   true,
		},
		// {
		// 	Path: "/signin",
		// 	Handler:    h.SignIn,
		// 	NeedAuth: false,
		// 	UnAuth:   true,
		// },
		{
			Path:     "/createpost",
			Handler:  h.CreatePost,
			NeedAuth: true,
			UnAuth:   false,
		},
	}
}
