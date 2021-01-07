package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// User holds the schema definition for the Salary entity.
type Salary struct {
	ent.Schema
}

// Fields of the User.
func (Salary) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("position").NotEmpty(),
		field.Int("Salary").Positive(),
	}
}

// Edges of the User.
func (Salary) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("assessment", Assessment.Type).
		Ref("formassessment").
		Unique(),
		edge.From("position", Position.Type).
		Ref("formposition").
		Unique(),
		edge.From("employee",Employee.Type).
		Ref("formemployee").
		Unique(),
	}
}
