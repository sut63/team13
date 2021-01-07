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
<<<<<<< HEAD
		field.Float("Salary").Positive(),
		field.Time("SalaryDatetime"),
=======
		field.Int("Salary").Positive(),
>>>>>>> f34210ab6b6442c2024f1f2cc6eb75a8ccfbe5ef
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