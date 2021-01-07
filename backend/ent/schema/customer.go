package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Paymentchannel holds the schema definition for the Paymentchannel entity.
type Paymentchannel struct {
	ent.Schema
}

// Fields of the Paymentchannel.
func (Paymentchannel) Fields() []ent.Field {
	return []ent.Field{
		field.String("Bank").NotEmpty(),
	}
}

// Edges of the Paymentchannel.
func (Paymentchannel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("formpaymentchannel", Orderonline.Type),
	}
}