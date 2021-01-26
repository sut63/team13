package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Typeproduct holds the schema definition for the Typeproduct entity.
type Typeproduct struct {
	ent.Schema
}

// Fields of the Typeproduct.
func (Typeproduct) Fields() []ent.Field {
	return []ent.Field{
		field.String("Typeproduct").NotEmpty().Unique(),
	}
}

// Edges of the Typeproduct.
func (Typeproduct) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("typestock", Stock.Type),
		edge.To("Typeproducts", Orderproduct.Type),
		edge.To("fromTypeproductonline", Orderonline.Type),
	}
}
