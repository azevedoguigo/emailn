package campaign

import (
	"time"

	internalerros "github.com/azevedoguigo/emailn/internal/internal-erros"
	"github.com/google/uuid"
)

type Contact struct {
	ID         string `gorm:"size:50" validate:"required"`
	Email      string `gorm:"size:50" validate:"email"`
	CampaignID string `gorm:"size:50"`
}

const (
	StatusPending string = "pending"
)

type Campaign struct {
	ID        string    `gorm:"size:50" validate:"required"`
	Name      string    `gorm:"size:50" validate:"min=5,max=24"`
	Content   string    `gorm:"size:1024" validate:"min=5,max=1024"`
	Status    string    `gorm:"size:12" validate:"required"`
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
