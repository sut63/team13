package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Manager holds the schema definition for the Manager entity.
type Manager struct {
	ent.Schema
}

// Fields of the Manager.
func (Manager) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.String("email").NotEmpty().Unique(),
	}
}

// Edges of the Manager.
func (Manager) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("managers", Orderproduct.Type),
	}
}