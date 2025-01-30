package endpoint

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/azevedoguigo/emailn/internal/contract"
	internalmock "github.com/azevedoguigo/emailn/internal/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignGetByID_Should_Return_Campaign(t *testing.T) {
	assert := assert.New(t)

	campaign := contract.GetCampaign{
		ID:      "1xPTO",
		Name:    "Test",
		Content: "Hi everyone",
		Status:  "pending",
	}

	service := new(internalmock.CampaignServiceMock)
	service.On("GetByID", mock.Anything).Return(&campaign, nil)

	handler := Handler{
		CampaignService: service,
	}

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	response, status, _ := handler.CampaignGetByID(rr, req)

	assert.Equal(http.StatusOK, status)
	assert.Equal(campaign.ID, response.(*contract.GetCampaign).ID)
	assert.Equal(campaign.Name, response.(*contract.GetCampaign).Name)
	assert.Equal(campaign.Status, response.(*contract.GetCampaign).Status)
}

func Test_CampaignGetByID_Should_Return_Error(t *testing.T) {
	assert := assert.New(t)

	service := new(internalmock.CampaignServiceMock)

	expectedError := errors.New("something wrong")
	service.On("GetByID", mock.Anything).Return(nil, expectedError)

	handler := Handler{
		CampaignService: service,
	}

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	_, _, responseError := handler.CampaignGetByID(rr, req)

	assert.NotNil(responseError)
	assert.Equal(responseError, expectedError)
}
