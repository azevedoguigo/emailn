package campaign

import (
	"errors"

	"github.com/azevedoguigo/emailn/internal/contract"
	internalerros "github.com/azevedoguigo/emailn/internal/internal-erros"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	GetByID(id string) (*contract.GetCampaign, error)
	Delete(id string) error
	Start(id string) error
}

type ServiceImp struct {
	Repository Repository
	SendMail   func(campaign *Campaign) error
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := NewCampaing(newCampaign.Name, newCampaign.Content, newCampaign.CreatedBy, newCampaign.Emails)
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
		return errors.New("campaign status invalid")
	}

	campaign.Delete()
	err = s.Repository.Delete(campaign)
	if err != nil {
		return internalerros.ErrInternal
	}

	return nil
}

func (s *ServiceImp) Start(id string) error {
	campaign, err := s.Repository.GetByID(id)
	if err != nil {
		return internalerros.ProcessErrorToReturn(err)
	}

	if campaign.Status != StatusPending {
		return errors.New("campaign status invalid")
	}

	go func() {
		err = s.SendMail(campaign)
		if err != nil {
			campaign.Failed()
		} else {
			campaign.Done()
		}

		s.Repository.Update(campaign)
	}()

	campaign.Started()
	err = s.Repository.Update(campaign)
	if err != nil {
		return internalerros.ProcessErrorToReturn(err)
	}

	return nil
}
