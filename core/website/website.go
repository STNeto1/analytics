package website

import (
	"_schemas/ent"
	"context"
)

type IWebsiteService interface {
	// CreateWebsite creates a new website with the given user, name and domain.
	// It returns a website on success, or an error on failure.
	CreateWebsite(ctx context.Context, user *ent.User, name, url string) (*ent.Website, error)
}
