package schema

import "github.com/facebookincubator/ent"

// Manager holds the schema definition for the Manager entity.
type Manager struct {
	ent.Schema
}

// Fields of the Manager.
func (Manager) Fields() []ent.Field {
	return nil
}

// Edges of the Manager.
func (Manager) Edges() []ent.Edge {
	return nil
}
