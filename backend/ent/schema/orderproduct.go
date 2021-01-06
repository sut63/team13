package schema

import (
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
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
		edge.From("product", Product.Type).Ref("products").Unique(),
		edge.From("company", Company.Type).Ref("companys").Unique(),
		edge.From("typeproduct",TypeProduct.Type).Ref("typeproducts").Unique(),
	}
}
