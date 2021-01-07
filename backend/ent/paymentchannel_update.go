// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/orderonline"
	"github.com/tanapon395/playlist-video/ent/paymentchannel"
	"github.com/tanapon395/playlist-video/ent/predicate"
)

// PaymentchannelUpdate is the builder for updating Paymentchannel entities.
type PaymentchannelUpdate struct {
	config
	hooks      []Hook
	mutation   *PaymentchannelMutation
	predicates []predicate.Paymentchannel
}

// Where adds a new predicate for the builder.
func (pu *PaymentchannelUpdate) Where(ps ...predicate.Paymentchannel) *PaymentchannelUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetBank sets the Bank field.
func (pu *PaymentchannelUpdate) SetBank(s string) *PaymentchannelUpdate {
	pu.mutation.SetBank(s)
	return pu
}

// AddFormpaymentchannelIDs adds the formpaymentchannel edge to Orderonline by ids.
func (pu *PaymentchannelUpdate) AddFormpaymentchannelIDs(ids ...int) *PaymentchannelUpdate {
	pu.mutation.AddFormpaymentchannelIDs(ids...)
	return pu
}

// AddFormpaymentchannel adds the formpaymentchannel edges to Orderonline.
func (pu *PaymentchannelUpdate) AddFormpaymentchannel(o ...*Orderonline) *PaymentchannelUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pu.AddFormpaymentchannelIDs(ids...)
}

// Mutation returns the PaymentchannelMutation object of the builder.
func (pu *PaymentchannelUpdate) Mutation() *PaymentchannelMutation {
	return pu.mutation
}

// RemoveFormpaymentchannelIDs removes the formpaymentchannel edge to Orderonline by ids.
func (pu *PaymentchannelUpdate) RemoveFormpaymentchannelIDs(ids ...int) *PaymentchannelUpdate {
	pu.mutation.RemoveFormpaymentchannelIDs(ids...)
	return pu
}

// RemoveFormpaymentchannel removes formpaymentchannel edges to Orderonline.
func (pu *PaymentchannelUpdate) RemoveFormpaymentchannel(o ...*Orderonline) *PaymentchannelUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pu.RemoveFormpaymentchannelIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PaymentchannelUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.Bank(); ok {
		if err := paymentchannel.BankValidator(v); err != nil {
			return 0, &ValidationError{Name: "Bank", err: fmt.Errorf("ent: validator failed for field \"Bank\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentchannelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PaymentchannelUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PaymentchannelUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PaymentchannelUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PaymentchannelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   paymentchannel.Table,
			Columns: paymentchannel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: paymentchannel.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Bank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: paymentchannel.FieldBank,
		})
	}
	if nodes := pu.mutation.RemovedFormpaymentchannelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   paymentchannel.FormpaymentchannelTable,
			Columns: []string{paymentchannel.FormpaymentchannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderonline.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.FormpaymentchannelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   paymentchannel.FormpaymentchannelTable,
			Columns: []string{paymentchannel.FormpaymentchannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderonline.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{paymentchannel.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PaymentchannelUpdateOne is the builder for updating a single Paymentchannel entity.
type PaymentchannelUpdateOne struct {
	config
	hooks    []Hook
	mutation *PaymentchannelMutation
}

// SetBank sets the Bank field.
func (puo *PaymentchannelUpdateOne) SetBank(s string) *PaymentchannelUpdateOne {
	puo.mutation.SetBank(s)
	return puo
}

// AddFormpaymentchannelIDs adds the formpaymentchannel edge to Orderonline by ids.
func (puo *PaymentchannelUpdateOne) AddFormpaymentchannelIDs(ids ...int) *PaymentchannelUpdateOne {
	puo.mutation.AddFormpaymentchannelIDs(ids...)
	return puo
}

// AddFormpaymentchannel adds the formpaymentchannel edges to Orderonline.
func (puo *PaymentchannelUpdateOne) AddFormpaymentchannel(o ...*Orderonline) *PaymentchannelUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return puo.AddFormpaymentchannelIDs(ids...)
}

// Mutation returns the PaymentchannelMutation object of the builder.
func (puo *PaymentchannelUpdateOne) Mutation() *PaymentchannelMutation {
	return puo.mutation
}

// RemoveFormpaymentchannelIDs removes the formpaymentchannel edge to Orderonline by ids.
func (puo *PaymentchannelUpdateOne) RemoveFormpaymentchannelIDs(ids ...int) *PaymentchannelUpdateOne {
	puo.mutation.RemoveFormpaymentchannelIDs(ids...)
	return puo
}

// RemoveFormpaymentchannel removes formpaymentchannel edges to Orderonline.
func (puo *PaymentchannelUpdateOne) RemoveFormpaymentchannel(o ...*Orderonline) *PaymentchannelUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return puo.RemoveFormpaymentchannelIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *PaymentchannelUpdateOne) Save(ctx context.Context) (*Paymentchannel, error) {
	if v, ok := puo.mutation.Bank(); ok {
		if err := paymentchannel.BankValidator(v); err != nil {
			return nil, &ValidationError{Name: "Bank", err: fmt.Errorf("ent: validator failed for field \"Bank\": %w", err)}
		}
	}

	var (
		err  error
		node *Paymentchannel
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentchannelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PaymentchannelUpdateOne) SaveX(ctx context.Context) *Paymentchannel {
	pa, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pa
}

// Exec executes the query on the entity.
func (puo *PaymentchannelUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PaymentchannelUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PaymentchannelUpdateOne) sqlSave(ctx context.Context) (pa *Paymentchannel, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   paymentchannel.Table,
			Columns: paymentchannel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: paymentchannel.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Paymentchannel.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Bank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: paymentchannel.FieldBank,
		})
	}
	if nodes := puo.mutation.RemovedFormpaymentchannelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   paymentchannel.FormpaymentchannelTable,
			Columns: []string{paymentchannel.FormpaymentchannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderonline.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.FormpaymentchannelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   paymentchannel.FormpaymentchannelTable,
			Columns: []string{paymentchannel.FormpaymentchannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderonline.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pa = &Paymentchannel{config: puo.config}
	_spec.Assign = pa.assignValues
	_spec.ScanValues = pa.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{paymentchannel.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pa, nil
}