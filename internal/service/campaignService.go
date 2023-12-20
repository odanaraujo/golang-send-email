package service

import (
	"fmt"
	"github.com/odanaraujo/golang/send-emails/configurations/exception"
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
		fmt.Println(err.Error())
		return "", exception.NewBadRequestErr(err.Error())
	}

	if err := s.Repository.Save(campaignService.NewCampaign()); err != nil {
		return "", exception.NewInternalServerError(err.Error())
	}

	return campaignService.ID, nil
}
