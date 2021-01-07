package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Employees holds the chema definition for the Employee entity.
type Employees struct {
	ent.Schema
}

// Fields of the Employee.
func (Employees) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("email").Unique(),
	}
}

// Edges of the Employee.
func (Employees) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("whose", EmployeeWorkingHours.Type),
	}
}
