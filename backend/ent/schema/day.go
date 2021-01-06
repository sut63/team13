package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Day holds the schema definition for the Day entity.
type Day struct {
	ent.Schema
}

// Fields of the Day.
func (Day) Fields() []ent.Field {
	return []ent.Field{
		field.String("Day"),
	}
}

// Edges of the Day.
func (Day) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("whatday", EmployeeWorkingHours.Type),
	}
}