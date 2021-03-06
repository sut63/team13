// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team13/app/ent/endwork"
	"github.com/team13/app/ent/predicate"
)

// EndWorkDelete is the builder for deleting a EndWork entity.
type EndWorkDelete struct {
	config
	hooks      []Hook
	mutation   *EndWorkMutation
	predicates []predicate.EndWork
}

// Where adds a new predicate to the delete builder.
func (ewd *EndWorkDelete) Where(ps ...predicate.EndWork) *EndWorkDelete {
	ewd.predicates = append(ewd.predicates, ps...)
	return ewd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ewd *EndWorkDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ewd.hooks) == 0 {
		affected, err = ewd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EndWorkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ewd.mutation = mutation
			affected, err = ewd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ewd.hooks) - 1; i >= 0; i-- {
			mut = ewd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ewd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ewd *EndWorkDelete) ExecX(ctx context.Context) int {
	n, err := ewd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ewd *EndWorkDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: endwork.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: endwork.FieldID,
			},
		},
	}
	if ps := ewd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ewd.driver, _spec)
}

// EndWorkDeleteOne is the builder for deleting a single EndWork entity.
type EndWorkDeleteOne struct {
	ewd *EndWorkDelete
}

// Exec executes the deletion query.
func (ewdo *EndWorkDeleteOne) Exec(ctx context.Context) error {
	n, err := ewdo.ewd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{endwork.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ewdo *EndWorkDeleteOne) ExecX(ctx context.Context) {
	ewdo.ewd.ExecX(ctx)
}
