package contract

import "github.com/odanaraujo/golang/send-emails/internal/domain/campaign"

type NewCampaign struct {
	Name     string
	Body     string
	Contacts []campaign.Contact
}
