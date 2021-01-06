package schema

import (
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebookincubator/ent"
)

// Orderproduct holds the schema definition for the Orderproduct entity.
type Orderproduct struct {
	ent.Schema
}

// Fields of the Orderproduct.
func (Orderproduct) Fields() []ent.Field {
	return []ent.Field{
		field.Time("addedtime"),
		field.Int("stock"),
	}
}

// Edges of the Orderproduct.
func (Orderproduct) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).Ref("orderproducts").Unique(),
	}
}
