package schema

import "github.com/facebookincubator/ent"

// TypeProduct holds the schema definition for the TypeProduct entity.
type TypeProduct struct {
	ent.Schema
}

// Fields of the TypeProduct.
func (TypeProduct) Fields() []ent.Field {
	return []ent.Field{
		field.String("TypeProduct").NotEmpty(),
}

// Edges of the TypeProduct.
func (TypeProduct) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("typestock", Stock.Type),
		edge.To("typeproducts", Orderproduct.Type),
	}
}
