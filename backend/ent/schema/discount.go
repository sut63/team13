package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Discount holds the schema definition for the Discount entity.
type Discount struct {
	ent.Schema
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