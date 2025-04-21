package main

import (
	"log"
	"time"

	"github.com/azevedoguigo/emailn/internal/domain/campaign"
	"github.com/azevedoguigo/emailn/internal/infrastructure/database"
	"github.com/azevedoguigo/emailn/internal/infrastructure/mail"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db := database.NewDB()
	repository := database.CampaignRepository{DB: db}
	campaignService := campaign.ServiceImp{
		Repository: &repository,
		SendMail:   mail.SendMail,
	}

	for {
		campaigns, err := repository.GetStartedCampaignsButNotExecuted()
		if err != nil {
			log.Fatalf("Error getting campaigns: %v", err)
		}

		for _, campaign := range campaigns {
			campaignService.SendEmail(&campaign)
		}

		time.Sleep(30 * time.Minute)
	}
}
