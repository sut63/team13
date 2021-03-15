package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// BeginWork holds the schema definition for the BeginWork entity.
type StartWork struct {
	ent.Schema
}

// Fields of the Shift.
func (StartWork) Fields() []ent.Field {
	return []ent.Field{
		field.Time("StartWork"),
	}
}

// Edges of the Shift.
func (StartWork) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("whenwork", EmployeeWorkingHours.Type),
	}
}