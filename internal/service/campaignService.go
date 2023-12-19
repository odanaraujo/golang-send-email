package service

import (
	"errors"
	"github.com/odanaraujo/golang/send-emails/internal/contract"
	"github.com/odanaraujo/golang/send-emails/internal/domain/campaign"
)

type CampaignService struct {
	Repository campaign.Repository
}

func (s *CampaignService) Create(newCampaign *contract.NewCampaign) (string, error) {
	campaignService := campaign.Campaign{
		Name:    newCampaign.Name,
		Body:    newCampaign.Body,
		Contact: newCampaign.Contacts,
	}

	if err := campaignService.Validate(); err != nil {
		return "", errors.New("validate error")
	}

	if err := s.Repository.Save(campaignService.NewCampaign()); err != nil {
		return "", errors.New("unable to save campaignService")
	}

	return campaignService.ID, nil
}
