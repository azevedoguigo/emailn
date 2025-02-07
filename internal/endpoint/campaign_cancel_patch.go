package endpoint

import (
	"net/http"

	internalerros "github.com/azevedoguigo/emailn/internal/internal-erros"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignCancelPatch(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")

	err := h.CampaignService.Cancel(id)
	if err != nil {
		internalerros.ProcessErrorToReturn(err)
	}

	return nil, http.StatusCreated, err
}
