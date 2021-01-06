package schema

import "github.com/facebookincubator/ent"

// Orderproduct holds the schema definition for the Orderproduct entity.
type Orderproduct struct {
	ent.Schema
}

// Fields of the Orderproduct.
func (Orderproduct) Fields() []ent.Field {
	return nil
}

// Edges of the Orderproduct.
func (Orderproduct) Edges() []ent.Edge {
	return nil
}
