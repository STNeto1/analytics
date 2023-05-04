package website

import (
	"_schemas/ent"
	"context"
)

func (s *WebsiteService) CreateWebsite(ctx context.Context, user *ent.User, name, domain string) (*ent.Website, error) {
	qb := s.client.Website.
		Create().
		SetName(name).
		SetDomain(domain).
		SetUser(user)

	if domain != "" {
		qb.SetDomain(domain)
	}

	return qb.Save(ctx)
}
