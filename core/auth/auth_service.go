package auth

import (
	"_schemas/ent"

	"go.uber.org/zap"
)

type AuthService struct {
	client *ent.Client
	logger *zap.Logger

	secret string
}

func NewAuthService(client *ent.Client, secret string, logger *zap.Logger) *AuthService {
	return &AuthService{client: client, secret: secret, logger: logger}
}
