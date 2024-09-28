package campaign

import (
	"github.com/azevedoguigo/emailn/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaing contract.NewCampaing) (string, error) {
	campaign, err := NewCampaing(newCampaing.Name, newCampaing.Content, newCampaing.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)

	if err != nil {
		return "", nil
	}

	return campaign.ID, nil
}
