package contract

type NewCampaign struct {
	Name      string
	Content   string
	Status    string
	CreatedBy string
	Emails    []string
}
