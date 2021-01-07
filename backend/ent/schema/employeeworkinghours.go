package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

// Employeeworkinghours holds the schema definition for the Employeeworkinghours entity.
type Employeeworkinghours struct {
	ent.Schema
}

// Fields of the Employeeworkinghours.
func (Employeeworkinghours) Fields() []ent.Field {
	return nil
}

// Edges of the Employeeworkinghours.
func (Employeeworkinghours) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("workinghour", Employee.Type).
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