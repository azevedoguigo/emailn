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
	"gorm.io/gorm"
)

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test Campaign",
		Content: "Body content",
		Status:  campaign.StatusPending,
		Emails:  []string{"mail.test@example.com"},
	}
	repositoryMock *internalmock.CampaignRepositoryMock
	service        = campaign.ServiceImp{}
)

func setup() {
	repositoryMock = new(internalmock.CampaignRepositoryMock)
	service.Repository = repositoryMock
}

func Test_Create_Campaign(t *testing.T) {
	setup()

	assert := assert.New(t)

	repositoryMock.On("Save", mock.Anything).Return(nil)

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	setup()

	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaign{})

	assert.NotNil(err)
	assert.False(errors.Is(internalerros.ErrInternal, err))
}

func Test_Create_SaveCampaign(t *testing.T) {
	setup()

	repositoryMock.On("Save", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Content != newCampaign.Content ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	setup()

	assert := assert.New(t)

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerros.ErrInternal, err))
}

func Test_GetByID_Return_Campaign(t *testing.T) {
	setup()

	assert := assert.New(t)

	savedCampaign, _ := campaign.NewCampaing(
		newCampaign.Name,
		newCampaign.Content,
		newCampaign.CreatedBy,
		newCampaign.Emails,
	)

	repositoryMock.On("GetByID", mock.MatchedBy(func(id string) bool {
		return id == savedCampaign.ID
	})).Return(savedCampaign, nil)

	campaignReturned, err := service.GetByID(savedCampaign.ID)

	assert.Nil(err)
	assert.Equal(savedCampaign.Name, campaignReturned.Name)
	assert.Equal(savedCampaign.Content, campaignReturned.Content)
	assert.Equal(savedCampaign.Status, campaignReturned.Status)
}

func Test_GetById_ReturnErrorWhenSomethingWrongExist(t *testing.T) {
	setup()

	assert := assert.New(t)

	savedCampaign, _ := campaign.NewCampaing(
		newCampaign.Name,
		newCampaign.Content,
		newCampaign.CreatedBy,
		newCampaign.Emails,
	)
	repositoryMock.On("GetByID", mock.Anything).Return(nil, errors.New("something wrong"))

	_, err := service.GetByID(savedCampaign.ID)

	assert.Equal(internalerros.ErrInternal.Error(), err.Error())
}

func Test_Delete_ReturnNilWhenDeleteHasSuccess(t *testing.T) {
	setup()

	assert := assert.New(t)

	campaignData, _ := campaign.NewCampaing(
		"Campaign Name",
		"Campaign content",
		newCampaign.CreatedBy,
		[]string{"test@example.com"},
	)

	repositoryMock.On("GetByID", mock.Anything).Return(campaignData, nil)
	repositoryMock.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaignData == campaign
	})).Return(nil)

	err := service.Delete(campaignData.ID)

	assert.Nil(err)
}

func Test_Delete_ReturnsErrorWhenCampaignDoesNoExists(t *testing.T) {
	setup()

	assert := assert.New(t)

	repositoryMock.On("GetByID", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	err := service.Delete("invalid_id")

	assert.NotNil(err)
	assert.Equal(err, gorm.ErrRecordNotFound)
}

func Test_Delete_ReturnStatusInvalidWhenCampaignHasStatusNotEqualsPending(t *testing.T) {
	setup()

	assert := assert.New(t)

	savedCampaign := campaign.Campaign{
		ID:      "631884b4-d7e0-4ee7-b8d8-eb8b393R4Sf1",
		Name:    "Test Campaign",
		Content: "This is test campaign",
		Status:  campaign.StatusDone,
	}

	repositoryMock.On("GetByID", savedCampaign.ID).Return(&savedCampaign, nil)

	err := service.Delete(savedCampaign.ID)

	assert.NotNil(err)
	assert.Equal(err.Error(), "campaign status invalid")
}

func Test_Delete_ReturnInternalErrorWhenDeleteWasProblem(t *testing.T) {
	setup()

	assert := assert.New(t)

	campaignData, _ := campaign.NewCampaing(
		"Campaign Name",
		"Campaign content",
		newCampaign.CreatedBy,
		[]string{"test@example.com"},
	)

	repositoryMock.On("GetByID", mock.Anything).Return(campaignData, nil)
	repositoryMock.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaignData == campaign
	})).Return(errors.New("error to delete campaign"))

	err := service.Delete(campaignData.ID)

	assert.Equal(internalerros.ErrInternal.Error(), err.Error())
}

func Test_Start_ReturnsErrorWhenCampaignDoesNoExists(t *testing.T) {
	setup()

	assert := assert.New(t)

	repositoryMock.On("GetByID", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	err := service.Start("invalid_id")

	assert.NotNil(err)
	assert.Equal(err, gorm.ErrRecordNotFound)
}

func Test_Start_ReturnStatusInvalidWhenCampaignHasStatusNotEqualsPending(t *testing.T) {
	setup()

	assert := assert.New(t)

	savedCampaign := campaign.Campaign{
		ID:      "631884b4-d7e0-4ee7-b8d8-eb8b393R4Sf1",
		Name:    "Test Campaign",
		Content: "This is test campaign",
		Status:  campaign.StatusDone,
	}

	repositoryMock.On("GetByID", savedCampaign.ID).Return(&savedCampaign, nil)

	err := service.Start(savedCampaign.ID)

	assert.NotNil(err)
	assert.Equal(err.Error(), "campaign status invalid")
}

func Test_Start_ShouldSendMail(t *testing.T) {
	setup()

	assert := assert.New(t)

	savedCampaign := campaign.Campaign{
		ID:      "631884b4-d7e0-4ee7-b8d8-eb8b393R4Sf1",
		Name:    "Test Campaign",
		Content: "This is test campaign",
		Status:  campaign.StatusPending,
	}

	repositoryMock.On("GetByID", savedCampaign.ID).Return(&savedCampaign, nil)

	sentMail := false
	sendMail := func(campaign *campaign.Campaign) error {
		if campaign.ID == savedCampaign.ID {
			sentMail = true
		}

		return nil
	}
	service.SendMail = sendMail

	_ = service.Start(savedCampaign.ID)

	assert.True(sentMail)
}

func Test_Start_ReturnErrorWhenFuncSendMailError(t *testing.T) {
	setup()

	assert := assert.New(t)

	savedCampaign := campaign.Campaign{
		ID:      "631884b4-d7e0-4ee7-b8d8-eb8b393R4Sf1",
		Name:    "Test Campaign",
		Content: "This is test campaign",
		Status:  campaign.StatusPending,
	}

	repositoryMock.On("GetByID", savedCampaign.ID).Return(&savedCampaign, nil)

	sendMail := func(campaign *campaign.Campaign) error {
		return errors.New("error to send mail")
	}
	service.SendMail = sendMail

	err := service.Start(savedCampaign.ID)

	assert.NotNil(err)
	assert.Equal(internalerros.ErrInternal.Error(), err.Error())
}

func Test_Start_ReturnNilWhenUpdatedToDone(t *testing.T) {
	setup()

	assert := assert.New(t)

	savedCampaign := campaign.Campaign{
		ID:      "631884b4-d7e0-4ee7-b8d8-eb8b393R4Sf1",
		Name:    "Test Campaign",
		Content: "This is test campaign",
		Status:  campaign.StatusPending,
	}

	repositoryMock.On("GetByID", savedCampaign.ID).Return(&savedCampaign, nil)
	repositoryMock.On("Update", mock.MatchedBy(func(campaignToUpdate *campaign.Campaign) bool {
		return savedCampaign.ID == campaignToUpdate.ID && savedCampaign.Status == campaign.StatusDone
	})).Return(nil)

	sentMail := false
	sendMail := func(campaign *campaign.Campaign) error {
		if campaign.ID == savedCampaign.ID {
			sentMail = true
		}

		return nil
	}
	service.SendMail = sendMail

	_ = service.Start(savedCampaign.ID)

	assert.True(sentMail)
	assert.Equal(savedCampaign.Status, campaign.StatusDone)
}
