// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team13/app/ent/assessment"
	"github.com/team13/app/ent/predicate"
)

// AssessmentDelete is the builder for deleting a Assessment entity.
type AssessmentDelete struct {
	config
	hooks      []Hook
	mutation   *AssessmentMutation
	predicates []predicate.Assessment
}

// Where adds a new predicate to the delete builder.
func (ad *AssessmentDelete) Where(ps ...predicate.Assessment) *AssessmentDelete {
	ad.predicates = append(ad.predicates, ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *AssessmentDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ad.hooks) == 0 {
		affected, err = ad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AssessmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ad.mutation = mutation
			affected, err = ad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ad.hooks) - 1; i >= 0; i-- {
			mut = ad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *AssessmentDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *AssessmentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: assessment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: assessment.FieldID,
			},
		},
	}
	if ps := ad.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
}

// AssessmentDeleteOne is the builder for deleting a single Assessment entity.
type AssessmentDeleteOne struct {
	ad *AssessmentDelete
}

// Exec executes the deletion query.
func (ado *AssessmentDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{assessment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *AssessmentDeleteOne) ExecX(ctx context.Context) {
	ado.ad.ExecX(ctx)
}
