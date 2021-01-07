// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/discount"
	"github.com/tanapon395/playlist-video/ent/promotion"
)

// DiscountCreate is the builder for creating a Discount entity.
type DiscountCreate struct {
	config
	mutation *DiscountMutation
	hooks    []Hook
}

// SetSale sets the sale field.
func (dc *DiscountCreate) SetSale(i int) *DiscountCreate {
	dc.mutation.SetSale(i)
	return dc
}

// AddFordiscountIDs adds the fordiscount edge to Promotion by ids.
func (dc *DiscountCreate) AddFordiscountIDs(ids ...int) *DiscountCreate {
	dc.mutation.AddFordiscountIDs(ids...)
	return dc
}

// AddFordiscount adds the fordiscount edges to Promotion.
func (dc *DiscountCreate) AddFordiscount(p ...*Promotion) *DiscountCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return dc.AddFordiscountIDs(ids...)
}

// Mutation returns the DiscountMutation object of the builder.
func (dc *DiscountCreate) Mutation() *DiscountMutation {
	return dc.mutation
}

// Save creates the Discount in the database.
func (dc *DiscountCreate) Save(ctx context.Context) (*Discount, error) {
	if _, ok := dc.mutation.Sale(); !ok {
		return nil, &ValidationError{Name: "sale", err: errors.New("ent: missing required field \"sale\"")}
	}
	var (
		err  error
		node *Discount
	)
	if len(dc.hooks) == 0 {
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiscountMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DiscountCreate) SaveX(ctx context.Context) *Discount {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DiscountCreate) sqlSave(ctx context.Context) (*Discount, error) {
	d, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	d.ID = int(id)
	return d, nil
}

func (dc *DiscountCreate) createSpec() (*Discount, *sqlgraph.CreateSpec) {
	var (
		d     = &Discount{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: discount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: discount.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.Sale(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: discount.FieldSale,
		})
		d.Sale = value
	}
	if nodes := dc.mutation.FordiscountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   discount.FordiscountTable,
			Columns: []string{discount.FordiscountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: promotion.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return d, _spec
}
