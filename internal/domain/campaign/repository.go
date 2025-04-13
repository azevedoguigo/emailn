package campaign

type Repository interface {
	Save(campaign *Campaign) error
	Get() ([]Campaign, error)
	GetByID(id string) (*Campaign, error)
	Update(campaign *Campaign) error
	Delete(campaign *Campaign) error
	GetStartedCampaignsButNotExecuted() ([]Campaign, error)
}
