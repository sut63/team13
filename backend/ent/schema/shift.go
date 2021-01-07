package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Shift holds the schema definition for the Shift entity.
type Shift struct {
	ent.Schema
}

// Fields of the Shift.
func (Shift) Fields() []ent.Field {
	return []ent.Field{
		field.Time("TimeStart"),
		field.Time("TimeEnd"),
	}
}

// Edges of the Shift.
func (Shift) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("when", Employeeworkinghours.Type),
	}
}