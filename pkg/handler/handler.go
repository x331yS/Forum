package handler

import (
	"net/http"

	"github.com/anatolethien/forum/pkg/service"
)

type Handler struct {
	services *service.Service
}

type route struct {
	Path       string
	Handler    http.HandlerFunc
	NeedAuth   bool
	OnlyUnauth bool
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) InitRouter() *http.ServeMux {
	routes := h.createRoutes()
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	images := http.FileServer(http.Dir("./assets/images"))
	mux.Handle("/images/", http.StripPrefix("/images/", images))

	for _, route := range routes {

		if route.NeedAuth {
			route.Handler = h.NeedAuthMiddleware(route.Handler)
		}

		if route.OnlyUnauth {
			route.Handler = h.OnlyUnauthMiddleware(route.Handler)
		}

		route.Handler = h.CookiesCheckMiddleware(route.Handler)

		mux.HandleFunc(route.Path, route.Handler)
	}

	return mux
}

func (h *Handler) createRoutes() []route {
	return []route{
		{
			Path:       "/",
			Handler:    h.Index(),
			NeedAuth:   false,
			OnlyUnauth: false,
		},
		{
			Path:       "/home",
			Handler:    h.Home(),
			NeedAuth:   false,
			OnlyUnauth: false,
		},
		{
			Path:       "/register",
			Handler:    h.Register,
			NeedAuth:   false,
			OnlyUnauth: true,
		},
		{
			Path:       "/login",
			Handler:    h.Login,
			NeedAuth:   false,
			OnlyUnauth: true,
		},
		{
			Path:       "/logout",
			Handler:    h.LogOut,
			NeedAuth:   true,
			OnlyUnauth: false,
		},

		{
			Path:       "/post/create",
			Handler:    h.CreatePost(),
			NeedAuth:   true,
			OnlyUnauth: false,
		},
		{
			Path:       "/post/rate",
			Handler:    h.RatePost,
			NeedAuth:   true,
			OnlyUnauth: false,
		},
		{
			Path:       "/post/",
			Handler:    h.GetPost(),
			NeedAuth:   false,
			OnlyUnauth: false,
		},
		{
			Path:       "/filter/",
			Handler:    h.Filter(),
			NeedAuth:   false,
			OnlyUnauth: false,
		},

		{
			Path:       "/comment/create",
			Handler:    h.CreateComment,
			NeedAuth:   true,
			OnlyUnauth: false,
		},
		{
			Path:       "/comment/rate",
			Handler:    h.RateComment,
			NeedAuth:   true,
			OnlyUnauth: false,
		},
	}
}
