package internalmock

import (
	"github.com/azevedoguigo/emailn/internal/contract"
	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (m *CampaignServiceMock) Create(newCampaign contract.NewCampaing) (string, error) {
	args := m.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func (m *CampaignServiceMock) GetByID(id string) (*contract.GetCampaign, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*contract.GetCampaign), args.Error(1)
}

func (r *CampaignServiceMock) Delete(id string) error {
	return nil
}
