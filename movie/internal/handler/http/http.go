package http

import (
	"encoding/json"
	"net/http"

	"movieexample.com/movie/internal/controller/movie"
)

type Handler struct {
	ctrl *movie.Controller
}

func New(ctrl *movie.Controller) *Handler {
	return &Handler{ctrl}
}

func (h *Handler) GetMovieDetails(w http.ResponseWriter, r *http.Request) {
	// Implementation of HTTP handling logic goes here
	id := r.FormValue("id")
	details, err := h.ctrl.GetMovieDetails(r.Context(), id)
	if err != nil {
		if err == movie.ErrNotFound {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(details); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
