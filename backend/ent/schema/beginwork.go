package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// BeginWork holds the schema definition for the BeginWork entity.
type BeginWork struct {
	ent.Schema
}

// Fields of the Shift.
func (BeginWork) Fields() []ent.Field {
	return []ent.Field{
		field.Time("BeginWork"),
	}
}

// Edges of the Shift.
func (BeginWork) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("whenwork", EmployeeWorkingHours.Type),
	}
}