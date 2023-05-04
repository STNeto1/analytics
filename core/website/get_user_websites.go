package website

import (
	"_schemas/ent"
	"_schemas/ent/user"
	"_schemas/ent/website"
	"context"
)

func (s *WebsiteService) GetUserWebsites(ctx context.Context, usr *ent.User) ([]*ent.Website, error) {
	qb := s.client.
		Website.
		Query().
		Where(website.HasUserWith(user.ID(usr.ID))).
		Where(website.DeletedAtIsNil())

	return qb.All(ctx)
}
