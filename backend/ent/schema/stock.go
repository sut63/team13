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
		field.Float("Priceproduct").Min(0).Positive(),
		field.Int("Amount").Min(0).Positive(),
        field.Time("Time"),
	}
}

func (Stock) Edges() []ent.Edge {
	return []ent.Edge{
        edge.From("product", Product.Type).
			Ref("stockproduct").
			Unique(),
		edge.From("zoneproduct", Zoneproduct.Type).
            Ref("zonestock").
			Unique(),
		edge.From("employee", Employee.Type).
            Ref("employeestock").
			Unique(),
		edge.From("Typeproduct", Typeproduct.Type).
            Ref("typestock").
            Unique(),
	}
}

