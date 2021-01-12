package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Orderproduct holds the schema definition for the Orderproduct entity.
type Orderproduct struct {
	ent.Schema
}

// Fields of the Orderproduct.
func (Orderproduct) Fields() []ent.Field {
	return []ent.Field{
		field.Time("addedtime"),
		field.Int("stock").Positive(),
	}
}

// Edges of the Orderproduct.
func (Orderproduct) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).Ref("products").Unique(),
		edge.From("company", Company.Type).Ref("companys").Unique(),
		edge.From("Typeproduct",Typeproduct.Type).Ref("Typeproducts").Unique(),
		edge.From("managers",Manager.Type).Ref("managers").Unique(),
	}
}
