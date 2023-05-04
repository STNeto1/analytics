package website

import (
	"_schemas/ent"
	"_schemas/ent/user"
	"_schemas/ent/website"
	"context"
	"time"

	"github.com/google/uuid"
)

func (s *WebsiteService) DeleteWebsite(ctx context.Context, usr *ent.User, id uuid.UUID) error {
	// TODO should wrap this in a transaction

	_, err := s.GetWebsiteByID(ctx, usr, id)
	if err != nil {
		return err
	}

	return s.client.
		Website.
		Update().
		Where(website.ID(id)).
		Where(website.HasUserWith(user.ID(usr.ID))).
		SetDeletedAt(time.Now()).
		Exec(ctx)
}
