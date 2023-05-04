package website_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdateWebsiteWithInvalidId(t *testing.T) {
	s, client, l := CreateWebsiteService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := generateAnyValidUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	wbst, err := generateAnyValidWebsite(client, usr)
	assert.NotNil(t, wbst)
	assert.NoError(t, err)

	website, err := s.UpdateWebsite(context.Background(), usr, uuid.New(), "name", "domain")
	assert.Nil(t, website)
	assert.Error(t, err)
}

func TestUpdateWebsiteWithInvalidUser(t *testing.T) {
	s, client, l := CreateWebsiteService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := generateAnyValidUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	usr2, err := generateAnyValidUser(client)
	assert.NotNil(t, usr2)
	assert.NoError(t, err)

	wbst, err := generateAnyValidWebsite(client, usr)
	assert.NotNil(t, wbst)
	assert.NoError(t, err)

	website, err := s.UpdateWebsite(context.Background(), usr2, wbst.ID, "name", "domain")
	assert.Nil(t, website)
	assert.Error(t, err)
}

func TestUpdateWebsiteWithValidIdAndUser(t *testing.T) {
	s, client, l := CreateWebsiteService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := generateAnyValidUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	wbst, err := generateAnyValidWebsite(client, usr)
	assert.NotNil(t, wbst)
	assert.NoError(t, err)

	website, err := s.UpdateWebsite(context.Background(), usr, wbst.ID, "name", "domain")
	assert.NotNil(t, website)
	assert.NoError(t, err)
}
