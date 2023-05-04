package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WebsiteEvent holds the schema definition for the WebsiteEvent entity.
type WebsiteEvent struct {
	ent.Schema
}

// Fields of the WebsiteEvent.
func (WebsiteEvent) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("event_name"),
		field.String("url_path").Optional(),
		field.String("url_query").Optional(),
		field.String("referrer_path").Optional(),
		field.String("referrer_query").Optional(),
		field.String("referrer_domain").Optional(),
		field.String("page_title").Optional(),
		field.String("page_data").Optional(), // JSON string
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

// Edges of the WebsiteEvent.
func (WebsiteEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("website", Website.Type).Ref("events").Unique(),
	}
}
