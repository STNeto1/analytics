package auth

import (
	"_schemas/ent"
	"_schemas/ent/user"
	"context"
	"errors"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Register(ctx context.Context, name, email, password string) (*ent.User, error) {
	_, err := s.client.User.Query().Where(user.Email(email)).Only(ctx)
	if err == nil {
		s.logger.Error("user already exists", zap.Error(err))
		return nil, errors.New("user already exists")
	}

	pwdHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error("error while hashing password", zap.Error(err))
		return nil, errors.New("error while hashing password")
	}

	usr, err := s.client.User.
		Create().
		SetName(name).
		SetEmail(email).
		SetPassword(string(pwdHash)).
		Save(ctx)

	if err != nil {
		s.logger.Error("error while creating user", zap.Error(err))
		return nil, errors.New("error while creating user")
	}

	return usr, nil
}
