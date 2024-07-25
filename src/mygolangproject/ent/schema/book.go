package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Book struct {
	ent.Schema
}

// Fields of the User.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.Int("book_id").
		field.String("book_name"),
		field.String("author").
	}
}

// Edges of the User.
func (Book) Edges() []ent.Edge {
	return nil
}
