package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Orderonline holds the schema definition for the Orderonline entity.
type Orderonline struct {
	ent.Schema
}

// Fields of the Orderonline.
func (Orderonline) Fields() []ent.Field {
	return []ent.Field{
		field.Time("addedtime"),
		field.Int("stock"),
	}
}

// Edges of the Orderonline.
func (Orderonline) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("producton", Product.Type).
		Ref("formproductonline").
		Unique(),
		edge.From("paymentchannel", Paymentchannel.Type).
		Ref("formpaymentchannel").
		Unique(),
		edge.From("Typeproduct",Typeproduct.Type).
		Ref("fromTypeproductonline").
		Unique(),
		edge.From("customer",Customer.Type).
		Ref("formcustomer").
		Unique(),
	}
}