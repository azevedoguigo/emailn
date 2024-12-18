package main

import (
	"net/http"

	"github.com/azevedoguigo/emailn/internal/domain/campaign"
	"github.com/azevedoguigo/emailn/internal/endpoint"
	"github.com/azevedoguigo/emailn/internal/infrastructure/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.Service{
		Repository: &database.CampaignRepository{},
	}
	handler := endpoint.Handler{
		CampaignService: campaignService,
	}

	r.Post("/campaign", endpoint.HandlerError(handler.CampaignPost))
	r.Get("/campaign", endpoint.HandlerError(handler.CampaignGet))

	http.ListenAndServe(":3000", r)
}
