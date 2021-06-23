package handler

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"

	"github.com/anatolethien/forum/pkg/models"
)

func (h *Handler) CreatePost() http.HandlerFunc {
	type viewData struct {
		Categories []string
	}
	const maxUploadImage = 20 * 1024 * 1024

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			tmpl := template.Must(template.ParseFiles("./web/template/create_post.html"))
			if categories, err := h.services.Post.GetValidCategories(); err != nil {
				writeResponse(w, http.StatusInternalServerError, err.Error())
			} else {
				tmpl.Execute(w, viewData{categories})
			}

		case "POST":
			c, _ := r.Cookie("forum")
			userID, err := h.services.User.GetUserIDByToken(c.Value)
			if err != nil {
				writeResponse(w, http.StatusForbidden, "Invalid Token")
				return
			}

			r.Body = http.MaxBytesReader(w, r.Body, maxUploadImage)
			if err := r.ParseMultipartForm(maxUploadImage); err != nil {
				writeResponse(w, http.StatusInternalServerError, "Images size over 20Mb")
				return
			}

			r.ParseForm()

			post := &models.Post{
				UserID:     userID,
				Title:      r.FormValue("title"),
				Content:    r.FormValue("content"),
				Categories: r.Form["categories"],
			}

			formdata := r.MultipartForm
			files := formdata.File["files"]

			filesPaths, err := h.services.Post.GenerateImagesFromFiles(files)
			if err != nil {
				writeResponse(w, http.StatusInternalServerError, err)
				return
			}

			code, id, err := h.services.Post.Create(post)
			if err != nil {
				writeResponse(w, code, err.Error())
				return
			}
			post.ID = id

			for _, path := range filesPaths {
				if err := h.services.Post.SetImage(post.ID, path); err != nil {
					writeResponse(w, http.StatusInternalServerError, err.Error())
					return
				}
			}

			http.Redirect(w, r, fmt.Sprintf("/post/%d", post.ID), http.StatusFound)
		default:
			writeResponse(w, http.StatusBadRequest, "Bad Method")
		}
	}
}

func (h *Handler) GetPost() http.HandlerFunc {
	type viewData struct {
		Aleatoire int
		Post      *models.Post
		PostID    int
		LoggedIn  bool
	}
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			id := getPostIDFromURL(r.URL.Path)
			if post, err := h.services.Post.Get(id); err != nil {
				writeResponse(w, http.StatusBadRequest, err.Error())
			} else {
				tmpl := template.Must(template.ParseFiles("./web/template/view_post.html"))
				ok := IsLoggedUser(r)
				max := 8
				min := 1
				aleatoire := rand.Intn(max-min) + min
				tmpl.Execute(w, viewData{aleatoire, post, post.ID, ok})
			}
		default:
			writeResponse(w, http.StatusBadRequest, "Bad Method")
		}
	}
}

func (h *Handler) RatePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		postID := r.FormValue("post_id")
		types := r.FormValue("type")

		c, _ := r.Cookie("forum")
		userID, err := h.services.User.GetUserIDByToken(c.Value)
		if err != nil {
			writeResponse(w, http.StatusForbidden, "Invalid Token")
			return
		}

		if err := h.services.Post.EstimatePost(postID, userID, types); err != nil {
			writeResponse(w, http.StatusBadRequest, err.Error())
		} else {
			http.Redirect(w, r, "/post/"+postID, http.StatusFound)
		}
	default:
		writeResponse(w, http.StatusBadRequest, "Bad Method")
	}
}

func (h *Handler) Filter() http.HandlerFunc {
	type templateData struct {
		Posts           []*models.Post
		LoggedIn        bool
		ValidCategories []string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			tmpl := template.Must(template.ParseFiles("./web/template/home.html"))
			field := getFiltersFieldFromURL(r.URL.Path)

			userID := 0
			var err error

			c, _ := r.Cookie("forum")
			if c != nil {
				userID, err = h.services.User.GetUserIDByToken(c.Value)
				if err != nil {
					writeResponse(w, http.StatusForbidden, "Invalid Token")
					return
				}
			}

			posts, err := h.services.Post.Filter(field, userID)
			if err != nil {
				if err.Error() == "Unauthorized" {
					http.Redirect(w, r, "/login", http.StatusFound)
				} else {
					writeResponse(w, http.StatusInternalServerError, err.Error())
				}
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
