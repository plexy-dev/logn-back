package schema

import "C"
import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Community holds the schema definition for the Community entity.
type Community struct {
	ent.Schema
}

// Fields of the Community.
func (Community) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MinLen(1).
			MaxLen(30).
			NotEmpty().
			Unique(),
		field.Text("about").
			MinLen(1).
			MaxLen(1000).
			NotEmpty().
			Unique(),
		field.Int("members").
			Min(1),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Community.
func (Community) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("socials", Social.Type).
			Ref("communities"),
		edge.From("companies", Company.Type).
			Ref("communities"),
		edge.From("areas", Area.Type).
			Ref("communities"),
		edge.To("users", User.Type),
	}
}
