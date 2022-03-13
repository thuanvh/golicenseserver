package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CustomerLicense holds the schema definition for the CustomerLicense entity.
type CustomerLicense struct {
	ent.Schema
}

// Fields of the CustomerLicense.
func (CustomerLicense) Fields() []ent.Field {
	return []ent.Field{
		field.Text("license_code").NotEmpty(),
		field.Bool("active").Default(true),
		field.Time("start_date").Default(time.Now),
		field.Time("end_date").Default(time.Now),
		field.Int("cpu").Default(0),
		field.Int("storage").Default(0),
		field.Int("number_of_nodes").Default(0),
	}
}

// Edges of the CustomerLicense.
func (CustomerLicense) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("customer", Customer.Type).Unique(),
		edge.To("license", License.Type),
	}
}
