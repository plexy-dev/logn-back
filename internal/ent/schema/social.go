package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Social holds the schema definition for the Social entity.
type Social struct {
	ent.Schema
}

// Fields of the Social.
func (Social) Fields() []ent.Field {
	return []ent.Field{
		field.String("fb").
			MinLen(12).
			MaxLen(100).
			Unique().
			Optional(),
		field.String("twitter").
			MinLen(12).
			MaxLen(100).
			Unique().
			Optional(),
		field.String("discord").
			MinLen(12).
			MaxLen(100).
			Unique().
			Optional(),
		field.String("slack").
			MinLen(12).
			MaxLen(100).
			Unique().
			Optional(),
		field.String("other").
			MinLen(12).
			MaxLen(100).
			Unique().
			Optional(),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Social.
func (Social) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("companies", Company.Type),
		edge.To("communities", Community.Type),
	}
}
