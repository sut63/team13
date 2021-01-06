package schema

import (
	"github.com/facebook/ent/schema/field"
)
// Discountholds the schema definition for the Discountentity.
type Discountstruct {
}

// Fields of the Discount.
func (Discount) Fields() []ent.Field {
    return []ent.Field{
        field.Int("sale"),
    }
}

// Edges of the Promotion.
func (Discount) Edges() []ent.Edge {
    return []ent.Edge{
		edge.To("fordiscount", Promotion.Type),
	}
}