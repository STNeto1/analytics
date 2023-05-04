package auth

import (
	"_schemas/ent"
	"_schemas/ent/user"
	"context"
	"errors"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Login(ctx context.Context, email, password string) (*ent.User, error) {
	usr, err := s.client.User.Query().Where(user.Email(email)).Only(ctx)
	if err != nil {
		s.logger.Error("user not found", zap.Error(err))
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
	if err != nil {
		s.logger.Error("invalid password", zap.Error(err))
		return nil, errors.New("invalid password")
	}

	return usr, nil
}
