package website

import (
	"_schemas/ent"

	"go.uber.org/zap"
)

type WebsiteService struct {
	client *ent.Client
	logger *zap.Logger
}

func NewWebsiteService(client *ent.Client, logger *zap.Logger) *WebsiteService {
	return &WebsiteService{client: client, logger: logger}
}
