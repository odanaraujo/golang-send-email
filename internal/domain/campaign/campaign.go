package campaign

import (
	"errors"
	"github.com/rs/xid"
	"time"
)

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
		ID:        xid.New().String(),
		Name:      campaign.Name,
		CreatedOn: time.Now(),
		Body:      campaign.Body,
		Contact:   campaign.Contact,
	}
}

func (campaign Campaign) Validate() error {
	if campaign.Name == "" {
		return errors.New("name is required")
	}

	if campaign.Body == "" {
		return errors.New("body is required")
	}

	if len(campaign.Contact) == 0 {
		return errors.New("contact is required")
	}
	return nil
}
