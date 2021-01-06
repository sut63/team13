package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Assessment holds the schema definition for the Assessment entity.
type Assessment struct {
	ent.Schema
}

// Fields of the Assessment.
func (Assessment) Fields() []ent.Field {
	return []ent.Field{
		field.String("assessment").NotEmpty(),
	}
}

// Edges of the Assessment.
func (Assessment) Edges() []ent.Edge {
	return nil
}
