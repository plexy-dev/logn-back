package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Technology holds the schema definition for the Technology entity.
type Technology struct {
	ent.Schema
}

// Fields of the Technology.
func (Technology) Fields() []ent.Field {
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

// Edges of the Technology.
func (Technology) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("vacancies", Vacancy.Type).
			Through("technology_levels", TechnologyLevel.Type),
	}
}
