package campaign

import (
	"errors"
	"testing"

	"github.com/azevedoguigo/emailn/internal/contract"
	internalerros "github.com/azevedoguigo/emailn/internal/internal-erros"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Get() ([]Campaign, error) {
	//args := r.Called(campaign)
	return nil, nil
}

func (r *repositoryMock) GetByID(id string) (*Campaign, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Campaign), nil
}

func (m *repositoryMock) Delete(campaign *Campaign) error {
	args := m.Called(campaign)
	if args.Get(0) != nil {
		return args.Error(0)
	}

	return nil
}

var (
	newCampaign = contract.NewCampaing{
		Name:    "Test Campaign",
		Content: "Body content",
		Status:  StatusPending,
		Emails:  []string{"mail.test@example.com"},
	}
	service = ServiceImp{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaing{})

	assert.NotNil(err)
	assert.False(errors.Is(internalerros.ErrInternal, err))
}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service.Repository = repositoryMock
	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerros.ErrInternal, err))
}

func Test_GetByID_Return_Campaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaing(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("GetByID", mock.MatchedBy(func(id string) bool {
		return id == campaign.ID
	})).Return(campaign, nil)
	service.Repository = repositoryMock

	campaignReturned, err := service.GetByID(campaign.ID)

	assert.Nil(err)
	assert.Equal(campaign.Name, campaignReturned.Name)
	assert.Equal(campaign.Content, campaignReturned.Content)
	assert.Equal(campaign.Status, campaignReturned.Status)
}

func Test_GetByID_Return_Error(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaing(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("GetByID", mock.Anything).Return(nil, errors.New("error to get campaign"))
	service.Repository = repositoryMock

	_, err := service.GetByID(campaign.ID)

	assert.Equal(internalerros.ErrInternal, err)
}
