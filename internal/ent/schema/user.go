package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MinLen(1).
			MaxLen(15).
			NotEmpty(),
		field.String("surname").
			MinLen(1).
			MaxLen(15).
			NotEmpty(),
		field.String("email").
			MinLen(4).
			MaxLen(50).
			Unique().
			NotEmpty(),
		field.Int("birth_day").
			Min(1).
			Max(31),
		field.Int("birth_month").
			Min(1).
			Max(12),
		field.Int("birth_year").
			Min(1930).
			Max(time.Now().Year()),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("passwords", Password.Type).
			Ref("users"),
		edge.From("roles", Role.Type).
			Ref("users"),
		edge.From("communities", Community.Type).
			Ref("users"),
		edge.From("companies", Company.Type).
			Ref("users"),
	}
}
