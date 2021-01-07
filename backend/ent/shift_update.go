// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/employeeworkinghours"
	"github.com/tanapon395/playlist-video/ent/predicate"
	"github.com/tanapon395/playlist-video/ent/shift"
)

// ShiftUpdate is the builder for updating Shift entities.
type ShiftUpdate struct {
	config
	hooks      []Hook
	mutation   *ShiftMutation
	predicates []predicate.Shift
}

// Where adds a new predicate for the builder.
func (su *ShiftUpdate) Where(ps ...predicate.Shift) *ShiftUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetTimeStart sets the TimeStart field.
func (su *ShiftUpdate) SetTimeStart(t time.Time) *ShiftUpdate {
	su.mutation.SetTimeStart(t)
	return su
}

// SetTimeEnd sets the TimeEnd field.
func (su *ShiftUpdate) SetTimeEnd(t time.Time) *ShiftUpdate {
	su.mutation.SetTimeEnd(t)
	return su
}

// AddWhenIDs adds the when edge to Employeeworkinghours by ids.
func (su *ShiftUpdate) AddWhenIDs(ids ...int) *ShiftUpdate {
	su.mutation.AddWhenIDs(ids...)
	return su
}

// AddWhen adds the when edges to Employeeworkinghours.
func (su *ShiftUpdate) AddWhen(e ...*Employeeworkinghours) *ShiftUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return su.AddWhenIDs(ids...)
}

// Mutation returns the ShiftMutation object of the builder.
func (su *ShiftUpdate) Mutation() *ShiftMutation {
	return su.mutation
}

// RemoveWhenIDs removes the when edge to Employeeworkinghours by ids.
func (su *ShiftUpdate) RemoveWhenIDs(ids ...int) *ShiftUpdate {
	su.mutation.RemoveWhenIDs(ids...)
	return su
}

// RemoveWhen removes when edges to Employeeworkinghours.
func (su *ShiftUpdate) RemoveWhen(e ...*Employeeworkinghours) *ShiftUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return su.RemoveWhenIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *ShiftUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShiftMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *ShiftUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ShiftUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ShiftUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *ShiftUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shift.Table,
			Columns: shift.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shift.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.TimeStart(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shift.FieldTimeStart,
		})
	}
	if value, ok := su.mutation.TimeEnd(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shift.FieldTimeEnd,
		})
	}
	if nodes := su.mutation.RemovedWhenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   shift.WhenTable,
			Columns: []string{shift.WhenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employeeworkinghours.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.WhenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   shift.WhenTable,
			Columns: []string{shift.WhenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employeeworkinghours.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shift.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ShiftUpdateOne is the builder for updating a single Shift entity.
type ShiftUpdateOne struct {
	config
	hooks    []Hook
	mutation *ShiftMutation
}

// SetTimeStart sets the TimeStart field.
func (suo *ShiftUpdateOne) SetTimeStart(t time.Time) *ShiftUpdateOne {
	suo.mutation.SetTimeStart(t)
	return suo
}

// SetTimeEnd sets the TimeEnd field.
func (suo *ShiftUpdateOne) SetTimeEnd(t time.Time) *ShiftUpdateOne {
	suo.mutation.SetTimeEnd(t)
	return suo
}

// AddWhenIDs adds the when edge to Employeeworkinghours by ids.
func (suo *ShiftUpdateOne) AddWhenIDs(ids ...int) *ShiftUpdateOne {
	suo.mutation.AddWhenIDs(ids...)
	return suo
}

// AddWhen adds the when edges to Employeeworkinghours.
func (suo *ShiftUpdateOne) AddWhen(e ...*Employeeworkinghours) *ShiftUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return suo.AddWhenIDs(ids...)
}

// Mutation returns the ShiftMutation object of the builder.
func (suo *ShiftUpdateOne) Mutation() *ShiftMutation {
	return suo.mutation
}

// RemoveWhenIDs removes the when edge to Employeeworkinghours by ids.
func (suo *ShiftUpdateOne) RemoveWhenIDs(ids ...int) *ShiftUpdateOne {
	suo.mutation.RemoveWhenIDs(ids...)
	return suo
}

// RemoveWhen removes when edges to Employeeworkinghours.
func (suo *ShiftUpdateOne) RemoveWhen(e ...*Employeeworkinghours) *ShiftUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return suo.RemoveWhenIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (suo *ShiftUpdateOne) Save(ctx context.Context) (*Shift, error) {

	var (
		err  error
		node *Shift
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShiftMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ShiftUpdateOne) SaveX(ctx context.Context) *Shift {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *ShiftUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ShiftUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *ShiftUpdateOne) sqlSave(ctx context.Context) (s *Shift, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shift.Table,
			Columns: shift.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shift.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Shift.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.TimeStart(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shift.FieldTimeStart,
		})
	}
	if value, ok := suo.mutation.TimeEnd(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shift.FieldTimeEnd,
		})
	}
	if nodes := suo.mutation.RemovedWhenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   shift.WhenTable,
			Columns: []string{shift.WhenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employeeworkinghours.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.WhenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   shift.WhenTable,
			Columns: []string{shift.WhenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employeeworkinghours.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	s = &Shift{config: suo.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shift.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}
