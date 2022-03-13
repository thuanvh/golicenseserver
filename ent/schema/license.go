package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// License holds the schema definition for the License entity.
type License struct {
	ent.Schema
}

// Fields of the License.
func (License) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").NotEmpty(),
		field.Int("duration").NonNegative(),
		field.Int("cpu").Default(0),
		field.Int("storage").Default(0),
		field.Int("number_of_nodes").Default(0),
	}
}

// Edges of the License.
func (License) Edges() []ent.Edge {
	return nil
}
