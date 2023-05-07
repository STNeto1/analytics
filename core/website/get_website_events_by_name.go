package website

import (
	"_schemas/ent"
	"_schemas/ent/websiteevent"
	"context"

	"github.com/google/uuid"
)

func (s *WebsiteService) GetWebsiteEventsByName(ctx context.Context, user *ent.User, id uuid.UUID, eventName string) ([]*ent.WebsiteEvent, error) {
	// Get the website.
	website, err := s.GetWebsiteByID(ctx, user, id)
	if err != nil {
		return nil, err
	}

	// Get the events.
	events, err := website.QueryEvents().Where(websiteevent.EventName(eventName)).All(ctx)
	if err != nil {
		return nil, err
	}

	return events, nil

}
