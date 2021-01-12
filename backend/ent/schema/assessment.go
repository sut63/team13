package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Assessment holds the schema definition for the Assessment entity.
type Assessment struct {
	ent.Schema
}

// Fields of the Assessment.
func (Assessment) Fields() []ent.Field {
	return []ent.Field{
		field.String("assessment").NotEmpty().Unique(),
	}
}

// Edges of the Assessment.
func (Assessment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("formassessment", Salary.Type),
	}
}