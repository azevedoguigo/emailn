package database

import (
	"github.com/azevedoguigo/emailn/internal/domain/campaign"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	DB *gorm.DB
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	tx := c.DB.Create(campaign)

	return tx.Error
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	tx := c.DB.Find(&campaigns)

	return campaigns, tx.Error
}

func (c *CampaignRepository) GetByID(id string) (*campaign.Campaign, error) {
	var campaign campaign.Campaign
	tx := c.DB.Preload("Contacts").First(&campaign, "id = ?", id)

	return &campaign, tx.Error
}

func (c *CampaignRepository) Update(campaign *campaign.Campaign) error {
	tx := c.DB.Save(campaign)

	return tx.Error
}

func (c *CampaignRepository) Delete(campaign *campaign.Campaign) error {
	tx := c.DB.Select("Contacts").Delete(campaign)

	return tx.Error
}

func (c *CampaignRepository) GetStartedCampaignsButNotExecuted() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	tx := c.DB.Preload("Contacts").
		Find(
			&campaigns,
			"status = ? and date_part('minute', now()::timestamp - updated_at::timestamp) > ?",
			campaign.StatusStarted,
			1,
		)

	return campaigns, tx.Error
}
