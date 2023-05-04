package website_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetUserWebsiteWithInvalidId(t *testing.T) {
	s, client, l := CreateWebsiteService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := generateAnyValidUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	website, err := s.GetWebsiteByID(context.Background(), usr, uuid.New())
	assert.Nil(t, website)
	assert.Error(t, err)
}

func TestGetUserWebsiteWithInvalidUser(t *testing.T) {
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

	website, err := s.GetWebsiteByID(context.Background(), usr2, wbst.ID)
	assert.Nil(t, website)
	assert.Error(t, err)
}

func TestGetUserWebsiteWithValidIdAndUser(t *testing.T) {
	s, client, l := CreateWebsiteService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := generateAnyValidUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	wbst, err := generateAnyValidWebsite(client, usr)
	assert.NotNil(t, wbst)
	assert.NoError(t, err)

	website, err := s.GetWebsiteByID(context.Background(), usr, wbst.ID)
	assert.NotNil(t, website)
	assert.NoError(t, err)
	assert.Equal(t, wbst.ID, website.ID)
}
