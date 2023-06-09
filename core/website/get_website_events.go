package website

import (
	"_schemas/ent"
	"_schemas/ent/websiteevent"
	"context"

	"github.com/google/uuid"
)

func (s *WebsiteService) GetWebsiteEvents(ctx context.Context, user *ent.User, id uuid.UUID) ([]*ent.WebsiteEvent, error) {
	website, err := s.GetWebsiteByID(ctx, user, id)
	if err != nil {
		return nil, err
	}

	return website.
		QueryEvents().
		Order(websiteevent.ByCreatedAt()).
		All(ctx)
}
