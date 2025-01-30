package endpoint

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignGetByID(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	campaigns, err := h.CampaignService.GetByID(id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return campaigns, http.StatusOK, err
}
