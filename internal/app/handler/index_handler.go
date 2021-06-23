package handler

import (
	"html/template"
	"math/rand"
	"net/http"

	"github.com/anatolethien/forum/internal/app/models"
)

func (h *Handler) Home() http.HandlerFunc {
	type templateData struct {
		Posts           []*models.Post
		LoggedIn        bool
		ValidCategories []string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			if r.URL.Path != "/home" {
				writeResponse(w, http.StatusNotFound, "Page Not Found")
				return
			}
			tmpl := template.Must(template.ParseFiles("./web/template/home.html"))
			posts, err := h.services.Post.GetAll()
			if err != nil {
				writeResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			validcategories, err := h.services.Post.GetValidCategories()
			if err != nil {
				writeResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			ok := IsLoggedUser(r)
			tmpl.Execute(w, templateData{posts, ok, validcategories})
		default:
			writeResponse(w, http.StatusBadRequest, "Bad Method")
		}
	}
}
func (h *Handler) Index() http.HandlerFunc {
	type templateData struct {
		Aleatoire      int
		Posts           []*models.Post
		LoggedIn        bool
		ValidCategories []string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			if r.URL.Path != "/" {
				writeResponse(w, http.StatusNotFound, "Page Not Found")
				return
			}
			tmpl := template.Must(template.ParseFiles("./web/template/index.html"))
			posts, err := h.services.Post.GetAll()
			if err != nil {
				writeResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			validcategories, err := h.services.Post.GetValidCategories()
			if err != nil {
				writeResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			max := 8
			min := 1
			aleatoire := rand.Intn(max - min) + min

			ok := IsLoggedUser(r)

			tmpl.Execute(w, templateData{aleatoire, posts, ok, validcategories})
		default:
			writeResponse(w, http.StatusBadRequest, "Bad Method")
		}
	}
}