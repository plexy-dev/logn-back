package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// TechnologyLevel holds the schema definition for the TechnologyLevel entity.
type TechnologyLevel struct {
	ent.Schema
}

// Fields of the TechnologyLevel.
func (TechnologyLevel) Fields() []ent.Field {
	return []ent.Field{
		field.Int("level").
			Min(1).
			Max(5).
			Default(1),
		field.Int("technology_id"),
		field.Int("vacancy_id"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the TechnologyLevel.
func (TechnologyLevel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("technology", Technology.Type).
			Required().
			Unique().
			Field("technology_id"),
		edge.To("vacancy", Vacancy.Type).
			Required().
			Unique().
			Field("vacancy_id"),
	}
}
