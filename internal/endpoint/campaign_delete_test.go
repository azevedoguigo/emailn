package endpoint

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsDelete_Success(t *testing.T) {
	Setup()

	campaignId := "xpto"
	service.On("Delete", mock.MatchedBy(func(id string) bool {
		return id == campaignId
	})).Return(nil)
	req, rr := NewRequestAndRecorder("PATCH", "/", nil)
	req = AddParameter(req, "id", campaignId)

	_, status, err := handler.CampaignDelete(rr, req)

	assert.Equal(t, http.StatusOK, status)
	assert.Nil(t, err)
}

func Test_CampaignsDelete_Err(t *testing.T) {
	Setup()

	errExpected := errors.New("something wrong")
	service.On("Delete", mock.Anything).Return(errExpected)
	req, rr := NewRequestAndRecorder("PATCH", "/", nil)

	_, _, err := handler.CampaignDelete(rr, req)

	assert.Equal(t, errExpected, err)
}
