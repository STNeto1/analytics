package website_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateWebsiteWithNoDomain(t *testing.T) {
	s, client, l := CreateWebsiteService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := generateAnyValidUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	website, err := s.CreateWebsite(context.Background(), usr, "some name", "")
	assert.NotNil(t, website)
	assert.NoError(t, err)
}

func TestCreateWebsiteWithDomain(t *testing.T) {
	s, client, l := CreateWebsiteService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := generateAnyValidUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	website, err := s.CreateWebsite(context.Background(), usr, "some name", "some domain")
	assert.NotNil(t, website)
	assert.NoError(t, err)
}
