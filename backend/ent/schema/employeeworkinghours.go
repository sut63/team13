package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

// EmployeeWorkingHours holds the schema definition for the EmployeeWorkingHours entity.
type EmployeeWorkingHours struct {
	ent.Schema
}

// Fields of the EmployeeWorkingHours.
func (EmployeeWorkingHours) Fields() []ent.Field {
	return nil
}

// Edges of the EmployeeWorkingHours.
func (EmployeeWorkingHours) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("EmployeeWorkingHours", Employee.Type).
			Ref("whose").
			Unique(),
		edge.From("day", Day.Type).
			Ref("whatday").
			Unique(),
		edge.From("shift", Shift.Type).
			Ref("when").
			Unique(),
		edge.From("role", Role.Type).
			Ref("todo").
			Unique(),
	}
}
