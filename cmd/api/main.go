package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"

	"github.com/azevedoguigo/emailn/internal/domain/campaign"
	"github.com/azevedoguigo/emailn/internal/endpoint"
	"github.com/azevedoguigo/emailn/internal/infrastructure/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

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

	log.Default().Println("Starting server on :3000")
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal("Error starting server", err)
	}
}
