package contract

type GetCampaign struct {
	ID                   string
	Name                 string
	Content              string
	Status               string
	AmountOfEmailsToSend int
}
