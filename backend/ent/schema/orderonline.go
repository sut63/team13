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
		field.Int("stock").Min(0).Positive(),
		field.String("accountnumber").MinLen(10).MaxLen(13),
		field.String("cvv").MinLen(3).MaxLen(3),
	}
}

// Edges of the Orderonline.
func (Orderonline) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).
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