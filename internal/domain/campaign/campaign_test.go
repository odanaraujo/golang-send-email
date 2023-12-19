package campaign

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	name     = "campaign X"
	body     = "Body Content"
	contact  = Contact{Email: "teste@gmail.com"}
	contacts = []Contact{contact}
)

func TestCampaign_NewCampaign(t *testing.T) {

	campaign := Campaign{
		ID:      "1",
		Name:    name,
		Body:    body,
		Contact: contacts,
	}

	campaign = *campaign.NewCampaign()

	assert.Equal(t, campaign.Name, "campaign X")
}

func TestCampaign_NewCampaign_IDIsNotNil(t *testing.T) {

	campaign := Campaign{
		ID:      "1",
		Name:    name,
		Body:    body,
		Contact: contacts,
	}

	campaign = *campaign.NewCampaign()

	assert.NotNil(t, campaign.ID, "valor não pode ser vazio")
}

func TestCampaign_NewCampaign_CreatedOnMustBeNow(t *testing.T) {

	campaign := Campaign{
		ID:      "1",
		Name:    name,
		Body:    body,
		Contact: contacts,
	}

	now := time.Now().Add(-time.Minute)

	campaign = *campaign.NewCampaign()

	assert.Greater(t, campaign.CreatedOn, now)
}

func TestCampaign_Validate(t *testing.T) {
	campaign := Campaign{
		ID:      "1",
		Name:    name,
		Body:    body,
		Contact: contacts,
	}

	err := campaign.Validate()

	assert.Nil(t, err, "validate ok")
}

func TestCampaign_ValidateName(t *testing.T) {
	campaign := Campaign{
		ID:      "1",
		Name:    "",
		Body:    body,
		Contact: contacts,
	}

	err := campaign.Validate()

	assert.Error(t, err, "name is required")
}

func TestCampaign_ValidateBody(t *testing.T) {
	campaign := Campaign{
		ID:      "1",
		Name:    name,
		Body:    "",
		Contact: contacts,
	}

	err := campaign.Validate()

	assert.Error(t, err, "body is required")
}

func TestCampaign_ValidateContact(t *testing.T) {
	campaign := Campaign{
		ID:      "1",
		Name:    name,
		Body:    body,
		Contact: []Contact{},
	}

	err := campaign.Validate()

	assert.Error(t, err, "contact is required")
}
