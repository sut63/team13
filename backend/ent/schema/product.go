package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("NameProduct").NotEmpty(),
		field.String("BarcodeProduct"),
		field.String("MFG"),
		field.String("EXP"),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
    	edge.To("products", Orderproduct.Type),
		edge.To("stockproduct", Stock.Type),
		edge.To("forproduct", Promotion.Type),
		edge.To("formproductonline", Orderonline.Type),
	}
}
