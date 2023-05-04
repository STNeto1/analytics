package website_test

import (
	"_core/website"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateWebsiteEventWithInvalidWebsiteId(t *testing.T) {
	s, client, l := CreateWebsiteService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := generateAnyValidUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	err = s.CreateWebsiteEvent(context.Background(), usr.ID, &website.CreateWebsiteEventPayload{
		EventName: "test",
	})
	assert.Error(t, err)
}

func TestCreateWebsiteEventWithValidWebsiteId(t *testing.T) {
	s, client, l := CreateWebsiteService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := generateAnyValidUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	site, err := generateAnyValidWebsite(client, usr)
	assert.NotNil(t, site)
	assert.NoError(t, err)

	err = s.CreateWebsiteEvent(context.Background(), site.ID, &website.CreateWebsiteEventPayload{
		EventName: "test",
	})
	assert.NoError(t, err)
}

func TestCreateWebsiteEventWithCompletePayload(t *testing.T) {
	s, client, l := CreateWebsiteService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := generateAnyValidUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	site, err := generateAnyValidWebsite(client, usr)
	assert.NotNil(t, site)
	assert.NoError(t, err)

	var testString = "test"
	err = s.CreateWebsiteEvent(context.Background(), site.ID, &website.CreateWebsiteEventPayload{
		EventName:      "test",
		UrlPath:        &testString,
		UrlQuery:       &testString,
		ReferrerPath:   &testString,
		ReferrerQuery:  &testString,
		ReferrerDomain: &testString,
		PageTitle:      &testString,
		PageData:       &testString,
	})
	assert.NoError(t, err)
}
