package endpoint

import (
	"net/http"

	"github.com/azevedoguigo/emailn/internal/contract"
	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewCampaign
	render.DecodeJSON(r.Body, &request)

	id, err := h.CampaignService.Create(request)

	return map[string]string{"id": id}, http.StatusCreated, err
}
