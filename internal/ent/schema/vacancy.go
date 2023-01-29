package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Vacancy holds the schema definition for the Vacancy entity.
type Vacancy struct {
	ent.Schema
}

// Fields of the Vacancy.
func (Vacancy) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MinLen(1).
			MaxLen(60),
		field.Text("description").
			MinLen(100).
			NotEmpty().
			Unique(),
		field.Bool("is_negotiate"),
		field.Int("min_salary").
			Min(1),
		field.Int("max_salary"),
		field.Bool("is_remote"),
		field.Int("views"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Vacancy.
func (Vacancy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("technologies", Technology.Type).
			Ref("vacancies").
			Through("technology_levels", TechnologyLevel.Type),
		edge.From("locations", Location.Type).
			Ref("vacancies"),
		edge.From("areas", Area.Type).
			Ref("vacancies"),
	}
}
