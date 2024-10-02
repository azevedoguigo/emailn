package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Test Campaign"
	content  = "Body content"
	contatcs = []string{"testmail.one@example.com", "testmail.two@example.com"}
	fake     = faker.New()
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaing(name, content, contatcs)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contatcs))
}

func Test_NewCampaign_IdIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaing(name, content, contatcs)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedAtMustBeNow(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaing(name, content, contatcs)

	actualTime := time.Now().Add(-time.Minute)

	assert.Greater(campaign.CreatedAt, actualTime)
}

func Test_NewCampaign_MustValidateNameMinLength(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing("", content, contatcs)

	assert.Equal("Name is required with min: 5", err.Error())
}

func Test_NewCampaign_MustValidateNameMaxLength(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(fake.Lorem().Text(25), content, contatcs)

	assert.Equal("Name is required with max: 24", err.Error())
}

func Test_NewCampaign_MustValidateContentMinLength(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, "", contatcs)

	assert.Equal("Content is required with min: 5", err.Error())
}

func Test_NewCampaign_MustValidateContentMaxLength(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, fake.Lorem().Text(1100), contatcs)

	assert.Equal("Content is required with max: 1024", err.Error())
}

func Test_NewCampaign_MustValidateValidEmail(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, content, []string{"invalidmail.com"})

	assert.Equal("Email is invalid.", err.Error())
}
