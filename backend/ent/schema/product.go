package schema

import (

	"github.com/facebook/ent"
  "github.com/facebook/ent/schema/edge"
  "github.com/facebook/ent/schema/field"
	
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
    edge.To("stockproduct", Stock.Type).
			Unique(),
    edge.To("orderproducts", Orderproduct.Type),

	}
}
