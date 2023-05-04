package website

import (
	"_schemas/ent"
	"_schemas/ent/user"
	"_schemas/ent/website"
	"context"

	"github.com/google/uuid"
)

func (s *WebsiteService) UpdateWebsite(ctx context.Context, usr *ent.User, id uuid.UUID, name, domain string) (*ent.Website, error) {
	return s.client.
		Website.
		UpdateOneID(id).
		Where(website.HasUserWith(user.ID(usr.ID))).
		SetName(name).
		SetDomain(domain).
		Save(ctx)
}
