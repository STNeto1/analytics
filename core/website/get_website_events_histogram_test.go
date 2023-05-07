package website_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWebsiteEventHistogram(t *testing.T) {
	s, client, l := CreateWebsiteService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := generateAnyValidUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	wbst, err := generateAnyValidWebsite(client, usr)
	assert.NotNil(t, wbst)
	assert.NoError(t, err)

	event1, err := client.WebsiteEvent.Create().SetEventName("event1").SetWebsite(wbst).Save(context.Background())
	assert.NotNil(t, event1)
	assert.NoError(t, err)

	histogram, err := s.GetWebsiteEventHistogram(context.Background(), usr, wbst.ID, "days")
	assert.NotNil(t, histogram)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(histogram))

}
