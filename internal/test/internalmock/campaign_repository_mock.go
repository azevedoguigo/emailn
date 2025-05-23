package internalmock

import (
	"github.com/azevedoguigo/emailn/internal/domain/campaign"
	"github.com/stretchr/testify/mock"
)

type CampaignRepositoryMock struct {
	mock.Mock
}

func (r *CampaignRepositoryMock) Save(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *CampaignRepositoryMock) Get() ([]campaign.Campaign, error) {
	//args := r.Called(campaign)
	return nil, nil
}

func (r *CampaignRepositoryMock) GetByID(id string) (*campaign.Campaign, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*campaign.Campaign), nil
}

func (r *CampaignRepositoryMock) Update(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *CampaignRepositoryMock) Delete(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *CampaignRepositoryMock) GetStartedCampaignsButNotExecuted() ([]campaign.Campaign, error) {
	args := r.Called()
	return args.Get(0).([]campaign.Campaign), args.Error(1)
}
