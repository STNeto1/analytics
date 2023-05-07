package website_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEventsByName(t *testing.T) {

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

	event2, err := client.WebsiteEvent.Create().SetEventName("event2").SetWebsite(wbst).Save(context.Background())
	assert.NotNil(t, event2)
	assert.NoError(t, err)

	events, err := s.GetWebsiteEventsByName(context.Background(), usr, wbst.ID, "event1")
	assert.NotNil(t, events)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(events))
}
