package auth_test

import (
	"_core/auth"
	"_schemas/ent"
	"_schemas/ent/enttest"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func CreateAuthService(t *testing.T) (*auth.AuthService, *ent.Client, *zap.Logger) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	logger := zap.NewNop()

	return auth.NewAuthService(client, "some-secret", logger), client, logger
}

func TestCreateAuthService(t *testing.T) {
	s, client, l := CreateAuthService(t)
	defer client.Close()
	defer l.Sync()

	assert.NotNil(t, s)
	assert.NotNil(t, client)
}

func TestImplementAuthServiceInterface(t *testing.T) {
	// s, _, _ := CreateAuthService(t)

	var _ auth.IAuthService = (*auth.AuthService)(nil)
}

func generateAnyValidUser(c *ent.Client) (*ent.User, error) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error while hashing password")
	}

	someEmail := fmt.Sprintf("some-%d@mail.com", rand.Int())

	return c.User.
		Create().
		SetName("some name").
		SetEmail(someEmail).
		SetPassword(string(pwdHash)).
		Save(context.Background())
}

// func TestFindUserWithInvalidId(t *testing.T) {
// 	s, client, l := CreateAuthService(t)
// 	defer client.Close()
// 	defer l.Sync()

// 	_, err := s.GetUserFromId(context.Background(), uuid.New())
// 	assert.Error(t, err)
// }

// func TestFindUserWithValidId(t *testing.T) {
// 	s, client, l := CreateAuthService(t)
// 	defer client.Close()
// 	defer l.Sync()

// 	usr, err := generateAnyValidUser(client)
// 	assert.Nil(t, err)

// 	foundUsr, err := s.GetUserFromId(context.Background(), usr.ID)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, foundUsr)
// 	assert.Equal(t, usr.ID.String(), foundUsr.ID.String())
// }
