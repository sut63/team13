// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team13/app/ent/stock"
	"github.com/team13/app/ent/zoneproduct"
)

// ZoneproductCreate is the builder for creating a Zoneproduct entity.
type ZoneproductCreate struct {
	config
	mutation *ZoneproductMutation
	hooks    []Hook
}

// SetZone sets the Zone field.
func (zc *ZoneproductCreate) SetZone(s string) *ZoneproductCreate {
	zc.mutation.SetZone(s)
	return zc
}

// AddZonestockIDs adds the zonestock edge to Stock by ids.
func (zc *ZoneproductCreate) AddZonestockIDs(ids ...int) *ZoneproductCreate {
	zc.mutation.AddZonestockIDs(ids...)
	return zc
}

// AddZonestock adds the zonestock edges to Stock.
func (zc *ZoneproductCreate) AddZonestock(s ...*Stock) *ZoneproductCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return zc.AddZonestockIDs(ids...)
}

// Mutation returns the ZoneproductMutation object of the builder.
func (zc *ZoneproductCreate) Mutation() *ZoneproductMutation {
	return zc.mutation
}

// Save creates the Zoneproduct in the database.
func (zc *ZoneproductCreate) Save(ctx context.Context) (*Zoneproduct, error) {
	if _, ok := zc.mutation.Zone(); !ok {
		return nil, &ValidationError{Name: "Zone", err: errors.New("ent: missing required field \"Zone\"")}
	}
	if v, ok := zc.mutation.Zone(); ok {
		if err := zoneproduct.ZoneValidator(v); err != nil {
			return nil, &ValidationError{Name: "Zone", err: fmt.Errorf("ent: validator failed for field \"Zone\": %w", err)}
		}
	}
	var (
		err  error
		node *Zoneproduct
	)
	if len(zc.hooks) == 0 {
		node, err = zc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ZoneproductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			zc.mutation = mutation
			node, err = zc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(zc.hooks) - 1; i >= 0; i-- {
			mut = zc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, zc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (zc *ZoneproductCreate) SaveX(ctx context.Context) *Zoneproduct {
	v, err := zc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (zc *ZoneproductCreate) sqlSave(ctx context.Context) (*Zoneproduct, error) {
	z, _spec := zc.createSpec()
	if err := sqlgraph.CreateNode(ctx, zc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	z.ID = int(id)
	return z, nil
}

func (zc *ZoneproductCreate) createSpec() (*Zoneproduct, *sqlgraph.CreateSpec) {
	var (
		z     = &Zoneproduct{config: zc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: zoneproduct.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: zoneproduct.FieldID,
			},
		}
	)
	if value, ok := zc.mutation.Zone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: zoneproduct.FieldZone,
		})
		z.Zone = value
	}
	if nodes := zc.mutation.ZonestockIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   zoneproduct.ZonestockTable,
			Columns: []string{zoneproduct.ZonestockColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: stock.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return z, _spec
}
