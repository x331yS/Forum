package handler

import (
	"net/http"

	"github.com/anatolethien/forum/internal/app/models"
)

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		content := r.FormValue("content")
		postID := r.FormValue("post_id")

		c, _ := r.Cookie("forum")
		userID, err := h.services.User.GetUserIDByToken(c.Value)
		if err != nil {
			writeResponse(w, http.StatusForbidden, "Invalid Token")
			return
		}

		comment := &models.Comment{
			UserID:  userID,
			Content: content,
		}

		code, _, err := h.services.Comment.Create(comment, postID)
		if err != nil {
			writeResponse(w, code, err.Error())
		} else {
			http.Redirect(w, r, "/post/"+postID, http.StatusFound)
		}
	default:
		writeResponse(w, http.StatusBadRequest, "Bad Method")
	}
}

func (h *Handler) RateComment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		commentID := r.FormValue("comment_id")
		postID := r.FormValue("post_id")
		types := r.FormValue("type")

		c, _ := r.Cookie("forum")
		userID, err := h.services.User.GetUserIDByToken(c.Value)
		if err != nil {
			writeResponse(w, http.StatusForbidden, "Invalid Token")
			return
		}

		if err := h.services.Comment.EstimateComment(commentID, userID, types); err != nil {
			writeResponse(w, http.StatusBadRequest, err.Error())
		} else {
			http.Redirect(w, r, "/post/"+postID, http.StatusFound)
		}
	default:
		writeResponse(w, http.StatusBadRequest, "Bad Method")
	}
}
