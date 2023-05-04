package website_test

import (
	"_core/website"
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

func CreateWebsiteService(t *testing.T) (*website.WebsiteService, *ent.Client, *zap.Logger) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	logger := zap.NewNop()

	return website.NewWebsiteService(client, logger), client, logger
}

func TestCreateWebsiteService(t *testing.T) {
	s, client, l := CreateWebsiteService(t)
	defer client.Close()
	defer l.Sync()

	assert.NotNil(t, s)
	assert.NotNil(t, client)
}

func TestImplementWebsiteServiceInterface(t *testing.T) {
	// s, _, _ := CreateAuthService(t)

	var _ website.IWebsiteService = (*website.WebsiteService)(nil)
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

func generateAnyValidWebsite(c *ent.Client, usr *ent.User) (*ent.Website, error) {
	return c.Website.
		Create().
		SetName("some name").
		SetDomain("some domain").
		SetUser(usr).
		Save(context.Background())
}
