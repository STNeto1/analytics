package website

import (
	"_schemas/ent"
	"context"
)

type IWebsiteService interface {
	// CreateWebsite creates a new website with the given user, name and domain.
	// It returns a website on success, or an error on failure.
	CreateWebsite(ctx context.Context, user *ent.User, name, url string) (*ent.Website, error)

	// GetUserWebsites returns all websites for the given user.
	// It returns a list of websites on success, or an error on failure.
	GetUserWebsites(ctx context.Context, user *ent.User) ([]*ent.Website, error)
}
