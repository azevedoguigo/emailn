package endpoint

import "github.com/azevedoguigo/emailn/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
