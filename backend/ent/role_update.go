// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/employeeworkinghours"
	"github.com/tanapon395/playlist-video/ent/predicate"
	"github.com/tanapon395/playlist-video/ent/role"
)

// RoleUpdate is the builder for updating Role entities.
type RoleUpdate struct {
	config
	hooks      []Hook
	mutation   *RoleMutation
	predicates []predicate.Role
}

// Where adds a new predicate for the builder.
func (ru *RoleUpdate) Where(ps ...predicate.Role) *RoleUpdate {
	ru.predicates = append(ru.predicates, ps...)
	return ru
}

// SetRole sets the Role field.
func (ru *RoleUpdate) SetRole(s string) *RoleUpdate {
	ru.mutation.SetRole(s)
	return ru
}

// AddTodoIDs adds the todo edge to Employeeworkinghours by ids.
func (ru *RoleUpdate) AddTodoIDs(ids ...int) *RoleUpdate {
	ru.mutation.AddTodoIDs(ids...)
	return ru
}

// AddTodo adds the todo edges to Employeeworkinghours.
func (ru *RoleUpdate) AddTodo(e ...*Employeeworkinghours) *RoleUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ru.AddTodoIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (ru *RoleUpdate) Mutation() *RoleMutation {
	return ru.mutation
}

// RemoveTodoIDs removes the todo edge to Employeeworkinghours by ids.
func (ru *RoleUpdate) RemoveTodoIDs(ids ...int) *RoleUpdate {
	ru.mutation.RemoveTodoIDs(ids...)
	return ru
}

// RemoveTodo removes todo edges to Employeeworkinghours.
func (ru *RoleUpdate) RemoveTodo(e ...*Employeeworkinghours) *RoleUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ru.RemoveTodoIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ru *RoleUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := ru.mutation.Role(); ok {
		if err := role.RoleValidator(v); err != nil {
			return 0, &ValidationError{Name: "Role", err: fmt.Errorf("ent: validator failed for field \"Role\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoleUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoleUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoleUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   role.Table,
			Columns: role.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: role.FieldID,
			},
		},
	}
	if ps := ru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Role(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldRole,
		})
	}
	if nodes := ru.mutation.RemovedTodoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.TodoTable,
			Columns: []string{role.TodoColumn},
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
	if nodes := ru.mutation.TodoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.TodoTable,
			Columns: []string{role.TodoColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RoleUpdateOne is the builder for updating a single Role entity.
type RoleUpdateOne struct {
	config
	hooks    []Hook
	mutation *RoleMutation
}

// SetRole sets the Role field.
func (ruo *RoleUpdateOne) SetRole(s string) *RoleUpdateOne {
	ruo.mutation.SetRole(s)
	return ruo
}

// AddTodoIDs adds the todo edge to Employeeworkinghours by ids.
func (ruo *RoleUpdateOne) AddTodoIDs(ids ...int) *RoleUpdateOne {
	ruo.mutation.AddTodoIDs(ids...)
	return ruo
}

// AddTodo adds the todo edges to Employeeworkinghours.
func (ruo *RoleUpdateOne) AddTodo(e ...*Employeeworkinghours) *RoleUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ruo.AddTodoIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (ruo *RoleUpdateOne) Mutation() *RoleMutation {
	return ruo.mutation
}

// RemoveTodoIDs removes the todo edge to Employeeworkinghours by ids.
func (ruo *RoleUpdateOne) RemoveTodoIDs(ids ...int) *RoleUpdateOne {
	ruo.mutation.RemoveTodoIDs(ids...)
	return ruo
}

// RemoveTodo removes todo edges to Employeeworkinghours.
func (ruo *RoleUpdateOne) RemoveTodo(e ...*Employeeworkinghours) *RoleUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ruo.RemoveTodoIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (ruo *RoleUpdateOne) Save(ctx context.Context) (*Role, error) {
	if v, ok := ruo.mutation.Role(); ok {
		if err := role.RoleValidator(v); err != nil {
			return nil, &ValidationError{Name: "Role", err: fmt.Errorf("ent: validator failed for field \"Role\": %w", err)}
		}
	}

	var (
		err  error
		node *Role
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoleUpdateOne) SaveX(ctx context.Context) *Role {
	r, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// Exec executes the query on the entity.
func (ruo *RoleUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoleUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RoleUpdateOne) sqlSave(ctx context.Context) (r *Role, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   role.Table,
			Columns: role.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: role.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Role.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ruo.mutation.Role(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: role.FieldRole,
		})
	}
	if nodes := ruo.mutation.RemovedTodoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.TodoTable,
			Columns: []string{role.TodoColumn},
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
	if nodes := ruo.mutation.TodoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.TodoTable,
			Columns: []string{role.TodoColumn},
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
	r = &Role{config: ruo.config}
	_spec.Assign = r.assignValues
	_spec.ScanValues = r.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return r, nil
}
