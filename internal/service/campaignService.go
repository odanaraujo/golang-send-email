package service

import (
	"errors"
	"github.com/odanaraujo/golang/send-emails/internal/contract"
	"github.com/odanaraujo/golang/send-emails/internal/domain/campaign"
)

type CampaignService struct {
	Repository campaign.Repository
}

func (s *CampaignService) Create(newCampaign *contract.NewCampaign) error {
	setCampaign := campaign.Campaign{
		Name:    newCampaign.Name,
		Body:    newCampaign.Body,
		Contact: newCampaign.Contacts,
	}

	err := s.Repository.Save(setCampaign.NewCampaign())

	if err != nil {
		return errors.New("unable to save campaign")
	}
	return nil
}
