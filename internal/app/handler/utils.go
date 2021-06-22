package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type errorData struct {
	Data interface{}
}

func writeResponse(w http.ResponseWriter, code int, resp interface{}) {
	w.WriteHeader(code)
	tmpl := template.Must(template.ParseFiles("./web/template/error.html"))
	tmpl.Execute(w, errorData{resp})
}

func getPostIDFromURL(url string) int {
	idStr := strings.TrimPrefix(url, "/post/")

	if idStr == "" {
		return -1
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return -1
	}

	return id
}

func getFiltersFieldFromURL(url string) string {
	return strings.Title(strings.TrimPrefix(url, "/filter/"))
}

func IsLoggedUser(r *http.Request) bool {
	c, _ := r.Cookie("forum")
	if c != nil {
		return true
	}
	return false
}
