package campaign

import (
	"errors"

	"github.com/azevedoguigo/emailn/internal/contract"
	internalerros "github.com/azevedoguigo/emailn/internal/internal-erros"
)

type Service interface {
	Create(newCampaign contract.NewCampaing) (string, error)
	GetByID(id string) (*contract.GetCampaign, error)
	Delete(string) error
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

func (s *ServiceImp) GetByID(id string) (*contract.GetCampaign, error) {
	campaign, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, internalerros.ProcessErrorToReturn(err)
	}

	return &contract.GetCampaign{
		ID:                   campaign.ID,
		Name:                 campaign.Name,
		Content:              campaign.Content,
		Status:               campaign.Status,
		AmountOfEmailsToSend: len(campaign.Contacts),
	}, nil
}

func (s *ServiceImp) Delete(id string) error {
	campaign, err := s.Repository.GetByID(id)

	if err != nil {
		return internalerros.ProcessErrorToReturn(err)
	}

	if campaign.Status != StatusPending {
		return errors.New("Campaign status invalid")
	}

	campaign.Delete()
	err = s.Repository.Delete(campaign)
	if err != nil {
		return internalerros.ErrInternal
	}

	return nil
}
