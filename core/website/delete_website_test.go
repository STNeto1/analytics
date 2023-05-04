package website_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeleteWebsiteWithInvalidId(t *testing.T) {
	s, client, l := CreateWebsiteService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := generateAnyValidUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	err = s.DeleteWebsite(context.Background(), usr, uuid.New())
	assert.Error(t, err)
}

func TestDeleteWebsiteWithInvalidUser(t *testing.T) {
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

	err = s.DeleteWebsite(context.Background(), usr2, wbst.ID)
	assert.Error(t, err)
}

func TestDeleteWebsiteWithValidIdAndUser(t *testing.T) {
	s, client, l := CreateWebsiteService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := generateAnyValidUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	wbst, err := generateAnyValidWebsite(client, usr)
	assert.NotNil(t, wbst)
	assert.NoError(t, err)

	err = s.DeleteWebsite(context.Background(), usr, wbst.ID)
	assert.NoError(t, err)
}
