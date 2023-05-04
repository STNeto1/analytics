package auth

import (
	"_schemas/ent"
	"context"

	"github.com/golang-jwt/jwt/v4"
)

type IAuthService interface {
	// Login logs in a user with the given email and password.
	// It returns a token on success, or an error on failure.
	Login(ctx context.Context, email, password string) (*ent.User, error)
	// Register registers a new user with the given name, email and password.
	// It returns a token on success, or an error on failure.
	Register(ctx context.Context, name, email, password string) (*ent.User, error)

	// GenerateToken generates a token for the given user
	// It returns a token on success, or an error on failure.
	GenerateToken(usr *ent.User) (string, error)

	// ValidateToken validates a token and returns the user ID if the token is valid.
	// It returns an error on failure.
	ValidateToken(token string) (*jwt.RegisteredClaims, error)
}
