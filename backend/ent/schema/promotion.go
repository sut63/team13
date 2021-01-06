package schema

import "github.com/facebook/ent/schema/field"

// Promotion holds the schema definition for the Promotion entity.
type Promotion struct {
}

// Fields of the Promotion.
func (Promotion) Fields() []ent.Field {
    return []ent.Field{
        field.String("PromotionName").Unique(),
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