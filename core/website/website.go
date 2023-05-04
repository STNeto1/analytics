package website

import (
	"_schemas/ent"
	"context"

	"github.com/google/uuid"
)

type IWebsiteService interface {
	// CreateWebsite creates a new website with the given user, name and domain.
	// It returns a website on success, or an error on failure.
	CreateWebsite(ctx context.Context, user *ent.User, name, url string) (*ent.Website, error)

	// GetUserWebsites returns all websites for the given user.
	// It returns a list of websites on success, or an error on failure.
	GetUserWebsites(ctx context.Context, user *ent.User) ([]*ent.Website, error)

	// GetWebsiteByID returns the website with the given ID for the given user.
	// It returns a website on success, or an error on failure.
	GetWebsiteByID(ctx context.Context, user *ent.User, id uuid.UUID) (*ent.Website, error)

	// UpdateWebsite updates the website with the given ID for the given user.
	// It returns a website on success, or an error on failure.
	UpdateWebsite(ctx context.Context, user *ent.User, id uuid.UUID, name, domain string) (*ent.Website, error)

	// DeleteWebsite deletes the website with the given ID for the given user.
	// It returns an error on failure.
	DeleteWebsite(ctx context.Context, user *ent.User, id uuid.UUID) error
}
