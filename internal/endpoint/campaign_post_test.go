package endpoint

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/azevedoguigo/emailn/internal/contract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	expectedCreatedBy = "creator@example.com"
	body              = contract.NewCampaign{
		Name:    "Test",
		Content: "Hi everyone",
		Emails:  []string{"test@example.com"},
	}
)

func Test_CampaignPost_Should_Save_New_Campaign(t *testing.T) {
	assert := assert.New(t)

	Setup()

	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		if request.Name == body.Name &&
			request.Content == body.Content &&
			request.CreatedBy == expectedCreatedBy &&
			request.Emails[0] == body.Emails[0] {
			return true
		} else {
			return false
		}
	})).Return("1xpTO", nil)

	req, rr := NewRequestAndRecorder(http.MethodPost, "/", body)
	req = AddContext(req, "email", expectedCreatedBy)

	_, status, err := handler.CampaignPost(rr, req)

	assert.Equal(http.StatusCreated, status)
	assert.Nil(err)
}

func Test_CampaignPost_Should_Inform_Error_When_Exist(t *testing.T) {
	assert := assert.New(t)

	Setup()

	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))

	req, rr := NewRequestAndRecorder(http.MethodPost, "/", body)
	req = AddContext(req, "email", expectedCreatedBy)

	_, _, err := handler.CampaignPost(rr, req)

	assert.NotNil(err)
}
