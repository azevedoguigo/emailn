package campaign

import (
	"github.com/azevedoguigo/emailn/internal/contract"
	internalerros "github.com/azevedoguigo/emailn/internal/internal-erros"
)

type Service interface {
	Create(newCampaign contract.NewCampaing) (string, error)
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaing) (string, error) {

	campaign, err := NewCampaing(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internalerros.ErrInternal
	}

	return campaign.ID, nil
}
