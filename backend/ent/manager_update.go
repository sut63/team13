// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team13/app/ent/manager"
	"github.com/team13/app/ent/orderproduct"
	"github.com/team13/app/ent/predicate"
)

// ManagerUpdate is the builder for updating Manager entities.
type ManagerUpdate struct {
	config
	hooks      []Hook
	mutation   *ManagerMutation
	predicates []predicate.Manager
}

// Where adds a new predicate for the builder.
func (mu *ManagerUpdate) Where(ps ...predicate.Manager) *ManagerUpdate {
	mu.predicates = append(mu.predicates, ps...)
	return mu
}

// SetName sets the name field.
func (mu *ManagerUpdate) SetName(s string) *ManagerUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetEmail sets the email field.
func (mu *ManagerUpdate) SetEmail(s string) *ManagerUpdate {
	mu.mutation.SetEmail(s)
	return mu
}

// SetPassword sets the password field.
func (mu *ManagerUpdate) SetPassword(s string) *ManagerUpdate {
	mu.mutation.SetPassword(s)
	return mu
}

// AddManagerIDs adds the managers edge to Orderproduct by ids.
func (mu *ManagerUpdate) AddManagerIDs(ids ...int) *ManagerUpdate {
	mu.mutation.AddManagerIDs(ids...)
	return mu
}

// AddManagers adds the managers edges to Orderproduct.
func (mu *ManagerUpdate) AddManagers(o ...*Orderproduct) *ManagerUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return mu.AddManagerIDs(ids...)
}

// Mutation returns the ManagerMutation object of the builder.
func (mu *ManagerUpdate) Mutation() *ManagerMutation {
	return mu.mutation
}

// RemoveManagerIDs removes the managers edge to Orderproduct by ids.
func (mu *ManagerUpdate) RemoveManagerIDs(ids ...int) *ManagerUpdate {
	mu.mutation.RemoveManagerIDs(ids...)
	return mu
}

// RemoveManagers removes managers edges to Orderproduct.
func (mu *ManagerUpdate) RemoveManagers(o ...*Orderproduct) *ManagerUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return mu.RemoveManagerIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (mu *ManagerUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := mu.mutation.Name(); ok {
		if err := manager.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := mu.mutation.Email(); ok {
		if err := manager.EmailValidator(v); err != nil {
			return 0, &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := mu.mutation.Password(); ok {
		if err := manager.PasswordValidator(v); err != nil {
			return 0, &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ManagerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *ManagerUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *ManagerUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *ManagerUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *ManagerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   manager.Table,
			Columns: manager.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: manager.FieldID,
			},
		},
	}
	if ps := mu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: manager.FieldName,
		})
	}
	if value, ok := mu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: manager.FieldEmail,
		})
	}
	if value, ok := mu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: manager.FieldPassword,
		})
	}
	if nodes := mu.mutation.RemovedManagersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manager.ManagersTable,
			Columns: []string{manager.ManagersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderproduct.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ManagersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manager.ManagersTable,
			Columns: []string{manager.ManagersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderproduct.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{manager.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ManagerUpdateOne is the builder for updating a single Manager entity.
type ManagerUpdateOne struct {
	config
	hooks    []Hook
	mutation *ManagerMutation
}

// SetName sets the name field.
func (muo *ManagerUpdateOne) SetName(s string) *ManagerUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetEmail sets the email field.
func (muo *ManagerUpdateOne) SetEmail(s string) *ManagerUpdateOne {
	muo.mutation.SetEmail(s)
	return muo
}

// SetPassword sets the password field.
func (muo *ManagerUpdateOne) SetPassword(s string) *ManagerUpdateOne {
	muo.mutation.SetPassword(s)
	return muo
}

// AddManagerIDs adds the managers edge to Orderproduct by ids.
func (muo *ManagerUpdateOne) AddManagerIDs(ids ...int) *ManagerUpdateOne {
	muo.mutation.AddManagerIDs(ids...)
	return muo
}

// AddManagers adds the managers edges to Orderproduct.
func (muo *ManagerUpdateOne) AddManagers(o ...*Orderproduct) *ManagerUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return muo.AddManagerIDs(ids...)
}

// Mutation returns the ManagerMutation object of the builder.
func (muo *ManagerUpdateOne) Mutation() *ManagerMutation {
	return muo.mutation
}

// RemoveManagerIDs removes the managers edge to Orderproduct by ids.
func (muo *ManagerUpdateOne) RemoveManagerIDs(ids ...int) *ManagerUpdateOne {
	muo.mutation.RemoveManagerIDs(ids...)
	return muo
}

// RemoveManagers removes managers edges to Orderproduct.
func (muo *ManagerUpdateOne) RemoveManagers(o ...*Orderproduct) *ManagerUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return muo.RemoveManagerIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (muo *ManagerUpdateOne) Save(ctx context.Context) (*Manager, error) {
	if v, ok := muo.mutation.Name(); ok {
		if err := manager.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := muo.mutation.Email(); ok {
		if err := manager.EmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if v, ok := muo.mutation.Password(); ok {
		if err := manager.PasswordValidator(v); err != nil {
			return nil, &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}

	var (
		err  error
		node *Manager
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ManagerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *ManagerUpdateOne) SaveX(ctx context.Context) *Manager {
	m, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return m
}

// Exec executes the query on the entity.
func (muo *ManagerUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *ManagerUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *ManagerUpdateOne) sqlSave(ctx context.Context) (m *Manager, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   manager.Table,
			Columns: manager.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: manager.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Manager.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := muo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: manager.FieldName,
		})
	}
	if value, ok := muo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: manager.FieldEmail,
		})
	}
	if value, ok := muo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: manager.FieldPassword,
		})
	}
	if nodes := muo.mutation.RemovedManagersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manager.ManagersTable,
			Columns: []string{manager.ManagersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderproduct.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ManagersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   manager.ManagersTable,
			Columns: []string{manager.ManagersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderproduct.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	m = &Manager{config: muo.config}
	_spec.Assign = m.assignValues
	_spec.ScanValues = m.scanValues()
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{manager.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return m, nil
}
