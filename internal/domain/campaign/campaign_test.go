package campaign

import (
	"testing"
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

	if campaign.ID != "1" {
		t.Error("expected 1", campaign.ID)
	}

	if campaign.Name != name {
		t.Error("expected ", name)
	}
}
