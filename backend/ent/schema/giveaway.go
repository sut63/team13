package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Giveaway holds the schema definition for the Giveaway entity.
type Giveaway struct {
    ent.Schema
}

// Fields of the Giveaway.
func (Giveaway) Fields() []ent.Field {
    return []ent.Field{
        field.String("giveawayName"),
    }
}

// Edges of the Promotion.
func (Giveaway) Edges() []ent.Edge {
    return []ent.Edge{
		edge.To("forgiveaway", Promotion.Type),
	}
}