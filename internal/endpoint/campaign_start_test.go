package endpoint

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignStart_Should_Return_Ok(t *testing.T) {
	Setup()
	assert := assert.New(t)

	service.On("Start", mock.Anything).Return(nil)

	req, rr := NewRequestAndRecorder(http.MethodPatch, "/", nil)

	campaignID := "1xpTO"
	req = AddParameter(req, "id", campaignID)

	_, status, err := handler.CampaignStart(rr, req)

	assert.Equal(http.StatusOK, status)
	assert.NoError(err)
}

func Test_CampaignStart_Should_Return_Err(t *testing.T) {
	Setup()
	assert := assert.New(t)

	campaignID := "1xpTO"
	expectedError := errors.New("something wrong")

	service.On("Start", mock.MatchedBy(func(id string) bool {
		return id == campaignID
	})).Return(expectedError)

	req, rr := NewRequestAndRecorder(http.MethodPatch, "/", nil)

	req = AddParameter(req, "id", campaignID)

	_, _, err := handler.CampaignStart(rr, req)

	assert.NotNil(err)
	assert.Equal(expectedError, err)
}
