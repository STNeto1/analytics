package website_test

import (
	"_core/website"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWebsiteEvents(t *testing.T) {
	s, client, l := CreateWebsiteService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := generateAnyValidUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	wbst, err := generateAnyValidWebsite(client, usr)
	assert.NotNil(t, wbst)
	assert.NoError(t, err)

	err = s.CreateWebsiteEvent(context.Background(), wbst.ID, &website.CreateWebsiteEventPayload{
		EventName: "test",
	})
	assert.NoError(t, err)

	events, err := s.GetWebsiteEvents(context.Background(), usr, wbst.ID)
	assert.NoError(t, err)
	assert.NotNil(t, events)
	assert.Len(t, events, 1)
}
