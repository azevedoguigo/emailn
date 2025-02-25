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

	db := database.NewDB()

	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{DB: db},
	}
	handler := endpoint.Handler{
		CampaignService: &campaignService,
	}

	r.Route("/campaign", func(r chi.Router) {
		r.Use(endpoint.Auth)

		r.Post("/", endpoint.HandlerError(handler.CampaignPost))
		r.Get("/{id}", endpoint.HandlerError(handler.CampaignGetByID))
		r.Delete("/{id}", endpoint.HandlerError(handler.CampaignDelete))
	})

	http.ListenAndServe(":3000", r)
}
