package auth_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestProfileWithInvalidId(t *testing.T) {
	s, client, l := CreateAuthService(t)
	defer client.Close()
	defer l.Sync()

	profile, err := s.GetUserFromId(context.Background(), uuid.New())
	assert.Nil(t, profile)
	assert.Error(t, err)
}

func TestProfileWithValidId(t *testing.T) {
	s, client, l := CreateAuthService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := generateAnyValidUser(client)
	assert.NotNil(t, usr)
	assert.NoError(t, err)

	profile, err := s.GetUserFromId(context.Background(), usr.ID)
	assert.NotNil(t, profile)
	assert.NoError(t, err)
}
