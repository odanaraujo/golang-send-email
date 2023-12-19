package service

import (
	"github.com/odanaraujo/golang/send-emails/internal/contract"
	"github.com/odanaraujo/golang/send-emails/internal/domain/campaign"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	name     = "new campaign"
	body     = "body campaign"
	contact  = campaign.Contact{Email: "teste@gmail.com"}
	contacts = []campaign.Contact{contact}
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func TestCampaignService_SaveCreate(t *testing.T) {

	newCampaign := contract.NewCampaign{
		Name:     name,
		Body:     body,
		Contacts: contacts,
	}

	repositoryMock := new(repositoryMock)
	mockAndVaidateMock(repositoryMock, newCampaign)
	service := CampaignService{
		repositoryMock,
	}

	id, err := service.Create(&newCampaign)

	assert.NotNil(t, id, "ID is required")
	assert.Nil(t, err, "Unable to create campaign in service")
	repositoryMock.AssertExpectations(t)
}

func mockAndVaidateMock(repositoryMock *repositoryMock, newCampaign contract.NewCampaign) *mock.Call {
	return repositoryMock.On("Save", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Body != newCampaign.Body || campaign.Contact[0] != newCampaign.Contacts[0] {
			return false
		}

		return true
	})).Return(nil)
}
