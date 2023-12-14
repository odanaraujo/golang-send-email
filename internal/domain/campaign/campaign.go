package campaign

import "time"

type Campaign struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Body      string
	Contact   []Contact
}

type Contact struct {
	Email string
}

func (campaign Campaign) NewCampaign() *Campaign {
	return &Campaign{
		ID:        campaign.ID,
		Name:      campaign.Name,
		CreatedOn: time.Now(),
		Body:      campaign.Body,
		Contact:   campaign.Contact,
	}
}
