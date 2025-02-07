package endpoint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/azevedoguigo/emailn/internal/contract"
	internalmock "github.com/azevedoguigo/emailn/internal/test/internalmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignPost_Should_Save_New_Campaign(t *testing.T) {
	assert := assert.New(t)

	body := contract.NewCampaing{
		Name:    "Test",
		Content: "Hi everyone",
		Emails:  []string{"test@example.com"},
	}

	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaing) bool {
		if request.Name == body.Name && request.Content == body.Content && request.Emails[0] == body.Emails[0] {
			return true
		} else {
			return false
		}
	})).Return("1xpTO", nil)

	handler := Handler{
		CampaignService: service,
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	req, _ := http.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignPost(rr, req)

	assert.Equal(http.StatusCreated, status)
	assert.Nil(err)
}

func Test_CampaignPost_Should_Inform_Error_When_Exist(t *testing.T) {
	assert := assert.New(t)

	body := contract.NewCampaing{
		Name:    "Test",
		Content: "Hi everyone",
		Emails:  []string{"test@example.com"},
	}

	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))

	handler := Handler{
		CampaignService: service,
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	req, _ := http.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()

	_, _, err := handler.CampaignPost(rr, req)

	assert.NotNil(err)
}
