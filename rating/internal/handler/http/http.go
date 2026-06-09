package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	rating "movieexample.com/rating/internal/controller"
	model "movieexample.com/rating/pkg"
)

type Handler struct {
	ctrl *rating.Controller
}

func New(ctrl *rating.Controller) *Handler {
	return &Handler{}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	recordID := model.RecordID(r.FormValue("id"))
	if recordID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	recordType := model.RecordType(r.FormValue("type"))
	if recordType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		v, err := h.ctrl.GetAggregatedRating(r.Context(), recordID, recordType)
		if err != nil && err.Error() == "Not Found" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("Response encode error: %v\n", err)
		}

	case http.MethodPut:
		userID := model.UserID(r.FormValue("userId"))
		v, err := strconv.ParseFloat(r.FormValue("value"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := h.ctrl.PutRating(r.Context(), recordID, recordType, model.Rating{
			UserID: userID,
			Value:  model.RatingValue(v),
		}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusBadGateway)
	}
}
