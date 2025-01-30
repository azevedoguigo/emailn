package campaign

import (
	"time"

	internalerros "github.com/azevedoguigo/emailn/internal/internal-erros"
	"github.com/google/uuid"
)

type Contact struct {
	ID         string `validate:"required"`
	Email      string `validate:"email"`
	CampaignID string `validate:"required"`
}

const (
	StatusPending string = "pending"
)

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	Content   string    `validate:"min=5,max=1024"`
	Status    string    `validate:"required"`
	Contacts  []Contact `validate:"min=1,dive"`
	CreatedAt time.Time `validate:"required"`
}

func NewCampaing(name string, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))
	for index, value := range emails {
		contacts[index].ID = uuid.NewString()
		contacts[index].Email = value
	}

	campaign := &Campaign{
		ID:        uuid.NewString(),
		Name:      name,
		Content:   content,
		Status:    StatusPending,
		Contacts:  contacts,
		CreatedAt: time.Now(),
	}
	err := internalerros.ValidateStruct(campaign)

	if err == nil {
		return campaign, nil
	}

	return nil, err
}
