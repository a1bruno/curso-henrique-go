package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Header"
	contacts := []string{"e_mail1@me.com", "e_mail2@me.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}
