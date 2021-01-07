// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/assessment"
	"github.com/tanapon395/playlist-video/ent/salary"
)

// AssessmentCreate is the builder for creating a Assessment entity.
type AssessmentCreate struct {
	config
	mutation *AssessmentMutation
	hooks    []Hook
}

// SetAssessment sets the assessment field.
func (ac *AssessmentCreate) SetAssessment(s string) *AssessmentCreate {
	ac.mutation.SetAssessment(s)
	return ac
}

// AddFormassessmentIDs adds the formassessment edge to Salary by ids.
func (ac *AssessmentCreate) AddFormassessmentIDs(ids ...int) *AssessmentCreate {
	ac.mutation.AddFormassessmentIDs(ids...)
	return ac
}

// AddFormassessment adds the formassessment edges to Salary.
func (ac *AssessmentCreate) AddFormassessment(s ...*Salary) *AssessmentCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ac.AddFormassessmentIDs(ids...)
}

// Mutation returns the AssessmentMutation object of the builder.
func (ac *AssessmentCreate) Mutation() *AssessmentMutation {
	return ac.mutation
}

// Save creates the Assessment in the database.
func (ac *AssessmentCreate) Save(ctx context.Context) (*Assessment, error) {
	if _, ok := ac.mutation.Assessment(); !ok {
		return nil, &ValidationError{Name: "assessment", err: errors.New("ent: missing required field \"assessment\"")}
	}
	if v, ok := ac.mutation.Assessment(); ok {
		if err := assessment.AssessmentValidator(v); err != nil {
			return nil, &ValidationError{Name: "assessment", err: fmt.Errorf("ent: validator failed for field \"assessment\": %w", err)}
		}
	}
	var (
		err  error
		node *Assessment
	)
	if len(ac.hooks) == 0 {
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AssessmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ac.mutation = mutation
			node, err = ac.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AssessmentCreate) SaveX(ctx context.Context) *Assessment {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ac *AssessmentCreate) sqlSave(ctx context.Context) (*Assessment, error) {
	a, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	a.ID = int(id)
	return a, nil
}

func (ac *AssessmentCreate) createSpec() (*Assessment, *sqlgraph.CreateSpec) {
	var (
		a     = &Assessment{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: assessment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: assessment.FieldID,
			},
		}
	)
	if value, ok := ac.mutation.Assessment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: assessment.FieldAssessment,
		})
		a.Assessment = value
	}
	if nodes := ac.mutation.FormassessmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   assessment.FormassessmentTable,
			Columns: []string{assessment.FormassessmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: salary.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return a, _spec
}
