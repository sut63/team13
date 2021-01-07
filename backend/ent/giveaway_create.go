// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/giveaway"
	"github.com/tanapon395/playlist-video/ent/promotion"
)

// GiveawayCreate is the builder for creating a Giveaway entity.
type GiveawayCreate struct {
	config
	mutation *GiveawayMutation
	hooks    []Hook
}

// SetGiveawayName sets the giveawayName field.
func (gc *GiveawayCreate) SetGiveawayName(i int) *GiveawayCreate {
	gc.mutation.SetGiveawayName(i)
	return gc
}

// AddForgiveawayIDs adds the forgiveaway edge to Promotion by ids.
func (gc *GiveawayCreate) AddForgiveawayIDs(ids ...int) *GiveawayCreate {
	gc.mutation.AddForgiveawayIDs(ids...)
	return gc
}

// AddForgiveaway adds the forgiveaway edges to Promotion.
func (gc *GiveawayCreate) AddForgiveaway(p ...*Promotion) *GiveawayCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gc.AddForgiveawayIDs(ids...)
}

// Mutation returns the GiveawayMutation object of the builder.
func (gc *GiveawayCreate) Mutation() *GiveawayMutation {
	return gc.mutation
}

// Save creates the Giveaway in the database.
func (gc *GiveawayCreate) Save(ctx context.Context) (*Giveaway, error) {
	if _, ok := gc.mutation.GiveawayName(); !ok {
		return nil, &ValidationError{Name: "giveawayName", err: errors.New("ent: missing required field \"giveawayName\"")}
	}
	var (
		err  error
		node *Giveaway
	)
	if len(gc.hooks) == 0 {
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GiveawayMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gc.mutation = mutation
			node, err = gc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GiveawayCreate) SaveX(ctx context.Context) *Giveaway {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gc *GiveawayCreate) sqlSave(ctx context.Context) (*Giveaway, error) {
	gi, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	gi.ID = int(id)
	return gi, nil
}

func (gc *GiveawayCreate) createSpec() (*Giveaway, *sqlgraph.CreateSpec) {
	var (
		gi    = &Giveaway{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: giveaway.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: giveaway.FieldID,
			},
		}
	)
	if value, ok := gc.mutation.GiveawayName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: giveaway.FieldGiveawayName,
		})
		gi.GiveawayName = value
	}
	if nodes := gc.mutation.ForgiveawayIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   giveaway.ForgiveawayTable,
			Columns: []string{giveaway.ForgiveawayColumn},
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
	return gi, _spec
}