package schema

import (
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// EmployeeWorkingHours holds the schema definition for the EmployeeWorkingHours entity.
type EmployeeWorkingHours struct {
	ent.Schema
}

// Fields of the EmployeeWorkingHours.
func (EmployeeWorkingHours) Fields() []ent.Field {
	return []ent.Field{
        field.String("CodeWork").Match(regexp.MustCompile("[ABC]\\d{7}")),
        field.String("IDNumber").MinLen(13).MaxLen(13).Match(regexp.MustCompile("[0-9]")),
        field.Float("Wages").Min(0).Positive(),
    }
}

// Edges of the EmployeeWorkingHours.
func (EmployeeWorkingHours) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("employee", Employee.Type).
			Ref("whose").
			Unique(),
		edge.From("day", Day.Type).
			Ref("whatday").
			Unique(),
		edge.From("startwork", StartWork.Type).
			Ref("whenwork").
			Unique(),
		edge.From("endwork", EndWork.Type).
			Ref("whenendwork").
			Unique(),
		edge.From("role", Role.Type).
			Ref("todo").
			Unique(),
	}
}
