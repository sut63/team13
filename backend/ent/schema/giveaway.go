package schema

import (
	"github.com/facebook/ent/schema/field"
)
// Giveawayholds the schema definition for the Giveawayentity.
type Giveawaystruct {
}

// Fields of the Giveaway.
func (Giveaway) Fields() []ent.Field {
    return []ent.Field{
        field.Int("giveawayName"),
    }
}

// Edges of the Promotion.
func (Giveaway) Edges() []ent.Edge {
    return []ent.Edge{
		edge.To("forGiveaway", Promotion.Type),
	}
}