package endpoint

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	internalmock "github.com/azevedoguigo/emailn/internal/test/internalmock"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignStart_Should_Return_Ok(t *testing.T) {
	assert := assert.New(t)

	service := new(internalmock.CampaignServiceMock)
	service.On("Start", mock.Anything).Return(nil)

	handler := Handler{
		CampaignService: service,
	}

	req, _ := http.NewRequest("PATCH", "/", nil)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignStart(rr, req)

	assert.Equal(http.StatusOK, status)
	assert.NoError(err)
}

func Test_CampaignStart_Should_Return_Err(t *testing.T) {
	assert := assert.New(t)

	campaignID := "1xpTO"
	expectedError := errors.New("something wrong")

	service := new(internalmock.CampaignServiceMock)
	service.On("Start", mock.MatchedBy(func(id string) bool {
		return id == campaignID
	})).Return(expectedError)

	handler := Handler{
		CampaignService: service,
	}

	req, _ := http.NewRequest("PATCH", "/", nil)
	rr := httptest.NewRecorder()

	chiContext := chi.NewRouteContext()
	chiContext.URLParams.Add("id", campaignID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))

	_, _, err := handler.CampaignStart(rr, req)

	assert.NotNil(err)
	assert.Equal(expectedError, err)
}
