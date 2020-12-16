package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/edge"

)
// User holds the schema definition for the Salary entity.
type salary struct {
	ent.Schema
}

// Fields of the User.
func (salary) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("position").NotEmpty(),
		field.String("salary").NotEmpty(),
	}
}

// Edges of the User.
func (salary) Edges() []ent.Edge {
	return nil
}
