package auth_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestCreatingUser(t *testing.T) {
	s, client, l := CreateAuthService(t)
	defer client.Close()
	defer l.Sync()

	_, err := s.Register(context.Background(),
		"John Doe",
		"some-mail@mail.com",
		"some-password",
	)

	assert.Nil(t, err)
}

func TestCreatingUserWithExistingEmail(t *testing.T) {
	s, client, l := CreateAuthService(t)
	defer client.Close()
	defer l.Sync()

	existingMail := "existing@mail.com"

	_, err := client.User.
		Create().
		SetName("John Doe").
		SetEmail(existingMail).
		SetPassword("some-password").
		Save(context.Background())
	assert.Nil(t, err)

	_, err = s.Register(context.Background(),
		"John Doe",
		existingMail,
		"some-password",
	)

	assert.NotNil(t, err)
}
