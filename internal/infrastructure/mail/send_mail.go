package mail

import (
	"github.com/azevedoguigo/emailn/internal/domain/campaign"
	"gopkg.in/gomail.v2"
	"os"
)

func SendMail(campaign *campaign.Campaign) error {
	dialer := gomail.NewDialer(
		os.Getenv("EMAIL_SMTP"),
		587,
		os.Getenv("EMAIL_USER"),
		os.Getenv("EMAIL_PASSWORD"),
	)

	var emails []string
	for _, contact := range campaign.Contacts {
		emails = append(emails, contact.Email)
	}

	message := gomail.NewMessage()
	message.SetHeader("From", os.Getenv("MAIL_FROM"))
	message.SetHeader("To", emails...)
	message.SetHeader("Subject", campaign.Name)
	message.SetBody("text/html", campaign.Content)

	return dialer.DialAndSend(message)
}
