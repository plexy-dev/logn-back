package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Area holds the schema definition for the Area entity.
type Area struct {
	ent.Schema
}

// Fields of the Area.
func (Area) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MinLen(1).
			MaxLen(15).
			NotEmpty().
			Unique(),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Area.
func (Area) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("vacancies", Vacancy.Type),
		edge.To("companies", Company.Type),
		edge.To("communities", Community.Type),
	}
}
