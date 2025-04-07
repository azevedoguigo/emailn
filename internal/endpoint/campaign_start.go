package endpoint

import (
	"net/http"

	internalerros "github.com/azevedoguigo/emailn/internal/internal-erros"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignStart(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")

	err := h.CampaignService.Start(id)
	if err != nil {
		internalerros.ProcessErrorToReturn(err)
	}

	return nil, 200, err
}
