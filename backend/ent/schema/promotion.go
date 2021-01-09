package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Promotion holds the schema definition for the Promotion entity.
type Promotion struct {
    ent.Schema
}

// Fields of the Promotion.
func (Promotion) Fields() []ent.Field {
    return []ent.Field{
        field.String("PromotionName").Unique(),
        field.Float("Price"),
    }
}

// Edges of the Promotion.
func (Promotion) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("sale",Discount.Type).
            Ref("fordiscount").
            Unique(),
        edge.From("give",Giveaway.Type).
            Ref("forgiveaway").
            Unique(),
        edge.From("product",Product.Type).
            Ref("forproduct").
            Unique(),
    }
}