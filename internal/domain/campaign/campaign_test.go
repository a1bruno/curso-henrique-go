package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "Campaign X"
	content := "Header"
	contacts := []string{"e_mail1@me.com", "e_mail2@me.com", "e_mail3@me.com"}

	campaign := NewCampaign(name, content, contacts)

	if campaign.ID != "1" {
		t.Errorf("found %s expected 1", campaign.ID)
	} else if campaign.Name != name {
		t.Errorf("found %s expected Body", campaign.Name)
	} else if len(campaign.Contacts) != len(contacts) {
		t.Errorf("found %d, expected %d", len(campaign.Contacts), len(contacts))
	}
}
