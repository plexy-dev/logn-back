package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Password holds the schema definition for the Password entity.
type Password struct {
	ent.Schema
}

// Fields of the Password.
func (Password) Fields() []ent.Field {
	return []ent.Field{
		field.String("hash").
			Sensitive(),
		field.String("salt").
			Sensitive(),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Password.
func (Password) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
