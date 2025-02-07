package campaign_test

import (
	"errors"
	"testing"

	"github.com/azevedoguigo/emailn/internal/contract"
	"github.com/azevedoguigo/emailn/internal/domain/campaign"
	internalerros "github.com/azevedoguigo/emailn/internal/internal-erros"
	"github.com/azevedoguigo/emailn/internal/test/internalmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newCampaign = contract.NewCampaing{
		Name:    "Test Campaign",
		Content: "Body content",
		Status:  campaign.StatusPending,
		Emails:  []string{"mail.test@example.com"},
	}
	service = campaign.ServiceImp{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(internalmock.CampaignRepositoryMock)
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
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
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

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerros.ErrInternal, err))
}

func Test_GetByID_Return_Campaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := campaign.NewCampaing(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	repositoryMock := new(internalmock.CampaignRepositoryMock)
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

func Test_GetById_ReturnErrorWhenSomethingWrongExist(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := campaign.NewCampaing(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetByID", mock.Anything).Return(nil, errors.New("Something wrong'"))
	service.Repository = repositoryMock

	_, err := service.GetByID(campaign.ID)

	assert.Equal(internalerros.ErrInternal.Error(), err.Error())
}
