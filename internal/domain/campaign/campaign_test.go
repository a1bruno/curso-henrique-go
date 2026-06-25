package campaign

import (
	"testing"
	"time"

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

func Test_NewCampaign_ID_Is_Not_Nil(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Header"
	contacts := []string{"e_mail1@me.com", "e_mail2@me.com"}

	campaign := NewCampaign(name, content, contacts)
	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOn_Greater_Than_Now(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Header"
	contacts := []string{"e_mail1@me.com", "e_mail2@me.com"}
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, contacts)
	assert.Greater(campaign.CreatedOn, now)
}
