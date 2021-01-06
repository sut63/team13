package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Stock holds the schema definition for the Stock entity.
type Stock struct {
	ent.Schema
}

// Fields of the Stock.
func (Stock) Fields() []ent.Field {
	return []ent.Field{
        field.String("Priceproduct"),
        field.Time("Time"),
	}
}

func (Stock) Edges() []ent.Edge {
	return []ent.Edge{
        edge.From("product", Product.Type).
			Ref("stockproduct").
			Unique().
			Required(),
		edge.From("zoneproduct", Zoneproduct.Type).
            Ref("zonestock").
			Unique(),
		edge.From("employee", Employee.Type).
            Ref("employeestock").
			Unique(),
		edge.From("typeproduct", TypeProduct.Type).
            Ref("typestock").
            Unique(),
	}
}
