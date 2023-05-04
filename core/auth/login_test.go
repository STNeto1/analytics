package auth_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginWithNoValidEmail(t *testing.T) {
	s, client, l := CreateAuthService(t)
	defer client.Close()
	defer l.Sync()

	usr, err := s.Login(context.Background(), "some-mail@mail.com", "some-password")

	assert.Nil(t, usr)
	assert.NotNil(t, err)
}

func TestLoginWithNoValidPassword(t *testing.T) {
	s, client, l := CreateAuthService(t)
	defer client.Close()
	defer l.Sync()

	existingMail := "existing@mail.com"
	correctPassword := "some-password"

	pwdHash, err := bcrypt.GenerateFromPassword([]byte(correctPassword), bcrypt.DefaultCost)
	assert.Nil(t, err)

	_, err = client.User.Create().
		SetName("John Doe").
		SetEmail(existingMail).
		SetPassword(string(pwdHash)).
		Save(context.Background())
	assert.Nil(t, err)

	assert.Nil(t, err)
	usr, err := s.Login(context.Background(), existingMail, "wrong-password")

	assert.Nil(t, usr)
	assert.NotNil(t, err)
}

func TestLoginWithValidCredentials(t *testing.T) {
	s, client, l := CreateAuthService(t)
	defer client.Close()
	defer l.Sync()

	existingMail := "existing@mail.com"
	correctPassword := "some-password"

	pwdHash, err := bcrypt.GenerateFromPassword([]byte(correctPassword), bcrypt.DefaultCost)
	assert.Nil(t, err)

	_, err = client.User.
		Create().
		SetName("John Doe").
		SetEmail(existingMail).
		SetPassword(string(pwdHash)).
		Save(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, err)

	usr, err := s.Login(context.Background(), existingMail, correctPassword)

	assert.NotNil(t, usr)
	assert.Nil(t, err)
}
