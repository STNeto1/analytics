package website

import (
	"_schemas/ent"
	"context"

	"github.com/google/uuid"
)

type Histogram = map[string][]*ent.WebsiteEvent

func (s *WebsiteService) GetWebsiteEventHistogram(ctx context.Context, user *ent.User, id uuid.UUID, level string) (Histogram, error) {
	// Get the website.
	website, err := s.GetWebsiteByID(ctx, user, id)
	if err != nil {
		return nil, err
	}

	// Get the events.
	events, err := website.QueryEvents().All(ctx)
	if err != nil {
		return nil, err
	}

	// Create the histogram.
	histogram := make(Histogram)

	// hours or days levels
	if level == "days" {

		// Iterate over the events.
		for _, event := range events {
			// Get the date.
			date := event.CreatedAt.Format("2006-01-02")

			// Add the event to the histogram.
			histogram[date] = append(histogram[date], event)
		}
	} else {
		// Iterate over the events.
		for _, event := range events {
			// Get the date.
			date := event.CreatedAt.Format("2006-01-02 15")

			// Add the event to the histogram.
			histogram[date] = append(histogram[date], event)
		}
	}

	return histogram, nil

}
