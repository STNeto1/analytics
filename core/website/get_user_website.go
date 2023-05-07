package website

import (
	"_schemas/ent"
	"_schemas/ent/user"
	"_schemas/ent/website"
	"context"

	"github.com/google/uuid"
)

func (s *WebsiteService) GetWebsiteByID(ctx context.Context, usr *ent.User, id uuid.UUID) (*ent.Website, error) {
	return s.client.
		Website.
		Query().
		Where(website.ID(id)).
		Where(website.HasUserWith(user.ID(usr.ID))).
		Where(website.DeletedAtIsNil()).
		First(ctx)
}
