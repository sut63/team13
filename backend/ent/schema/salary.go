package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Salary holds the schema definition for the Salary entity.
type Salary struct {
	ent.Schema
}

// Fields of the Salary.
func (Salary) Fields() []ent.Field {
	return []ent.Field{
		field.Float("Salary").Positive(),
		field.Float("Bonus").Positive(),
		field.Time("SalaryDatetime"),
	}
}

// Edges of the Salary.
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