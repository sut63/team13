package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/edeg"
)

// Stock holds the schema definition for the Stock entity.
type Stock struct {
	ent.Schema
}

// Fields of the Stock.
func (Stock) Fields() []ent.Field {
	return []ent.Field{
        field.String("Priceproduct"),
        field.Time("Time").
            Default(time.Now),
}

// Edges of the Stock.
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
