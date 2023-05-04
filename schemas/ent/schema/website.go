package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Website holds the schema definition for the Website entity.
type Website struct {
	ent.Schema
}

// Fields of the Website.
func (Website) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name"),
		field.String("domain").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Website.
func (Website) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("websites").Unique(),
		edge.To("events", WebsiteEvent.Type).StorageKey(edge.Column("website_id")),
	}
}
