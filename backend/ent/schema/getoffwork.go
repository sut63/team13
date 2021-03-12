package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// GetOffWork holds the schema definition for the GetOffWork entity.
type GetOffWork struct {
	ent.Schema
}

// Fields of the Shift.
func (GetOffWork) Fields() []ent.Field {
	return []ent.Field{
		field.Time("GetOffWork"),
	}
}

// Edges of the Shift.
func (GetOffWork) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("whenendwork", EmployeeWorkingHours.Type),
	}
}