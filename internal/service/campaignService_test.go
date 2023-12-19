package service

import (
	"errors"
	"github.com/odanaraujo/golang/send-emails/internal/contract"
	"github.com/odanaraujo/golang/send-emails/internal/domain/campaign"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	newCampaign = contract.NewCampaign{
		Name:     "new campaign",
		Body:     "body campaign",
		Contacts: []campaign.Contact{{Email: "test@gmail.com"}},
	}
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *campaign.Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func TestCampaignService_SaveCreate(t *testing.T) {

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

func TestCampaignService_ValidateCreate(t *testing.T) {

	newCampaign.Name = ""

	repositoryMock := new(repositoryMock)
	mockAndVaidateMock(repositoryMock, newCampaign)
	service := CampaignService{
		repositoryMock,
	}

	id, err := service.Create(&newCampaign)

	assert.Error(t, err, "validate campaign")
	assert.NotNil(t, id, "id required")
}

func TestCampaignService_CreateWithError(t *testing.T) {

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("unable to save campaignServic"))

	service := CampaignService{
		repositoryMock,
	}

	id, err := service.Create(&newCampaign)

	assert.Equal(t, err.Error(), "unable to save campaignServic")
	assert.Emptyf(t, id, "id is required")
}

func mockAndVaidateMock(repositoryMock *repositoryMock, newCampaign contract.NewCampaign) *mock.Call {
	return repositoryMock.On("Save", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Body != newCampaign.Body || len(campaign.Contact) != len(newCampaign.Contacts) {
			return false
		}

		return true
	})).Return(nil)
}
