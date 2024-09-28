package campaign

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	Content   string
	Contacts  []Contact
	CreatedAt time.Time
}

func NewCampaing(name string, content string, emails []string) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	if content == "" {
		return nil, errors.New("content is required")
	}

	if len(emails) == 0 {
		return nil, errors.New("contact is required")
	}

	contacts := make([]Contact, len(emails))
	for index, value := range emails {
		contacts[index].Email = value
	}

	return &Campaign{
		ID:        uuid.NewString(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedAt: time.Now(),
	}, nil
}
