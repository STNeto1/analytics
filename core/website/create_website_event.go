package website

import (
	"_schemas/ent/website"
	"context"

	"github.com/google/uuid"
)

type CreateWebsiteEventPayload struct {
	EventName      string
	UrlPath        *string
	UrlQuery       *string
	ReferrerPath   *string
	ReferrerQuery  *string
	ReferrerDomain *string
	PageTitle      *string
	PageData       *string
}

func (s *WebsiteService) CreateWebsiteEvent(ctx context.Context, id uuid.UUID, payload *CreateWebsiteEventPayload) error {
	// TODO wrap in transaction

	site, err := s.client.Website.Query().
		Where(website.ID(id)).
		Only(ctx)
	if err != nil {
		return err
	}

	qb := s.client.WebsiteEvent.
		Create().
		SetWebsite(site).
		SetEventName(payload.EventName)

	if payload.UrlPath != nil {
		qb.SetURLPath(*payload.UrlPath)
	}

	if payload.UrlQuery != nil {
		qb.SetURLQuery(*payload.UrlQuery)
	}

	if payload.ReferrerPath != nil {
		qb.SetReferrerPath(*payload.ReferrerPath)
	}

	if payload.ReferrerQuery != nil {
		qb.SetReferrerQuery(*payload.ReferrerQuery)
	}

	if payload.ReferrerDomain != nil {
		qb.SetReferrerDomain(*payload.ReferrerDomain)
	}

	if payload.PageTitle != nil {
		qb.SetPageTitle(*payload.PageTitle)
	}

	if payload.PageData != nil {
		qb.SetPageData(*payload.PageData)
	}

	_, err = qb.Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
