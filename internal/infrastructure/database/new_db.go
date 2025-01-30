package database

import (
	"log"

	"github.com/azevedoguigo/emailn/internal/domain/campaign"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=emailn port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection error: %v", err.Error())
		panic(err)
	}

	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	return db
}
