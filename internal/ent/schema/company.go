package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

// Fields of the Company.
func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MinLen(1).
			MaxLen(50).
			NotEmpty().
			Unique(),
		field.Int("employ").
			Min(1),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("socials", Social.Type).
			Ref("companies"),
		edge.From("roles", Role.Type).
			Ref("companies"),
		edge.From("vacancies", Vacancy.Type).
			Ref("companies"),
		edge.From("areas", Area.Type).
			Ref("companies"),
		edge.To("users", User.Type),
		edge.To("communities", Community.Type),
	}
}
