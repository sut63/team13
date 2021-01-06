package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/edeg"
)

// Zoneproduct holds the schema definition for the Zoneproduct entity.
type Zoneproduct struct {
	ent.Schema
}

// Fields of the Zoneproduct.
func (Zoneproduct) Fields() []ent.Field {
	return []ent.Field{
		field.String("Zone").NotEmpty(),
}

// Edges of the Zoneproduct.
func (Zoneproduct) Edges() []ent.Edge {
	return []ent.Edge{
        edge.To("zonestock", Stock.Type),
    }
}
