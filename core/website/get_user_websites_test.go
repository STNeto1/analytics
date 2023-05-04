package website_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReturnListOfUserWebsites(t *testing.T) {
	s, client, l := CreateWebsiteService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := generateAnyValidUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	wbst, err := generateAnyValidWebsite(client, usr)
	assert.NotNil(t, wbst)
	assert.NoError(t, err)

	websites, err := s.GetUserWebsites(context.Background(), usr)
	assert.NotNil(t, websites)
	assert.Len(t, websites, 1)
	assert.NoError(t, err)
}

func TestReturnListOfUserWebsitesWithNoDeletedRows(t *testing.T) {
	s, client, l := CreateWebsiteService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := generateAnyValidUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	wbst, err := generateAnyValidWebsite(client, usr)
	assert.NotNil(t, wbst)
	assert.NoError(t, err)

	err = client.Website.UpdateOne(wbst).SetDeletedAt(time.Now()).Exec(context.Background())
	assert.NoError(t, err)

	websites, err := s.GetUserWebsites(context.Background(), usr)
	assert.NotNil(t, websites)
	assert.Len(t, websites, 0)
	assert.NoError(t, err)
}
