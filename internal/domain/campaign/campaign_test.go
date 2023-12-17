package campaign

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCampaign_NewCampaign(t *testing.T) {
	name := "campaign X"
	body := "Body Content"
	contact := Contact{Email: "teste@gmail.com"}
	contacts := []Contact{contact}

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

	name := "campaign X"
	body := "Body Content"
	contact := Contact{Email: "teste@gmail.com"}
	contacts := []Contact{contact}

	campaign := Campaign{
		ID:      "1",
		Name:    name,
		Body:    body,
		Contact: contacts,
	}

	campaign = *campaign.NewCampaign()

	assert.NotNil(t, campaign.ID, "valor não pode ser vazio")
}

func TestCampaign_NewCampaign_DateIsNotNil(t *testing.T) {

	name := "campaign X"
	body := "Body Content"
	contact := Contact{Email: "teste@gmail.com"}
	contacts := []Contact{contact}

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
