package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Employee holds the schema definition for the Employee entity.
type Employee struct {
	ent.Schema
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("email").NotEmpty().Unique(),
		field.Int("age").Positive(),
	}
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return nil
}