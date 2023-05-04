package auth_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	s, client, l := CreateAuthService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := generateAnyValidUser(client)
	assert.NotNil(t, usr)
	assert.Nil(t, err)

	token, err := s.GenerateToken(usr)
	assert.NotEmpty(t, token)
	assert.Nil(t, err)
}

func TestValidateToken(t *testing.T) {
	s, client, l := CreateAuthService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := generateAnyValidUser(client)
	assert.NotNil(t, usr)
	assert.Nil(t, err)

	token, err := s.GenerateToken(usr)
	assert.NotEmpty(t, token)
	assert.Nil(t, err)

	claims, err := s.ValidateToken(token)
	assert.NotNil(t, claims)
	assert.Nil(t, err)
}
