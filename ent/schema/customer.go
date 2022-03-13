package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").NotEmpty(),
		field.Text("address").NotEmpty(),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return nil
}
