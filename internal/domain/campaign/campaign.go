package campaign

import (
	"github.com/rs/xid"
	"time"
)

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=60"`
	CreatedOn time.Time `validate:"required"`
	Body      string    `validate:"min=5"`
	Contact   []Contact `validate:"min=1,dive"`
}

type Contact struct {
	Email string `validate:"email"`
}

func (campaign Campaign) NewCampaign() *Campaign {
	return &Campaign{
		ID:        xid.New().String(),
		Name:      campaign.Name,
		CreatedOn: time.Now(),
		Body:      campaign.Body,
		Contact:   campaign.Contact,
	}
}
