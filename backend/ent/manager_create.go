// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/manager"
	"github.com/tanapon395/playlist-video/ent/orderproduct"
)

// ManagerCreate is the builder for creating a Manager entity.
type ManagerCreate struct {
	config
	mutation *ManagerMutation
	hooks    []Hook
}

// SetName sets the name field.
func (mc *ManagerCreate) SetName(s string) *ManagerCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetEmail sets the email field.
func (mc *ManagerCreate) SetEmail(s string) *ManagerCreate {
	mc.mutation.SetEmail(s)
	return mc
}

// AddManagerIDs adds the managers edge to Orderproduct by ids.
func (mc *ManagerCreate) AddManagerIDs(ids ...int) *ManagerCreate {
	mc.mutation.AddManagerIDs(ids...)
	return mc
}

// AddManagers adds the managers edges to Orderproduct.
func (mc *ManagerCreate) AddManagers(o ...*Orderproduct) *ManagerCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return mc.AddManagerIDs(ids...)
}

// Mutation returns the ManagerMutation object of the builder.
func (mc *ManagerCreate) Mutation() *ManagerMutation {
	return mc.mutation
}

// Save creates the Manager in the database.
func (mc *ManagerCreate) Save(ctx context.Context) (*Manager, error) {
	if _, ok := mc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := mc.mutation.Name(); ok {
		if err := manager.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := mc.mutation.Email(); !ok {
		return nil, &ValidationError{Name: "email", err: errors.New("ent: missing required field \"email\"")}
	}
	if v, ok := mc.mutation.Email(); ok {
		if err := manager.EmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	var (
		err  error
		node *Manager
	)
	if len(mc.hooks) == 0 {
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ManagerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mc.mutation = mutation
			node, err = mc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *ManagerCreate) SaveX(ctx context.Context) *Manager {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mc *ManagerCreate) sqlSave(ctx context.Context) (*Manager, error) {
	m, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	m.ID = int(id)
	return m, nil
}

func (mc *ManagerCreate) createSpec() (*Manager, *sqlgraph.CreateSpec) {
	var (
		m     = &Manager{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: manager.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: manager.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: manager.FieldName,
		})
		m.Name = value
	}
	if value, ok := mc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: manager.FieldEmail,
		})
		m.Email = value
	}
	if nodes := mc.mutation.ManagersIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return m, _spec
}
