// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/employee"
	"github.com/tanapon395/playlist-video/ent/product"
	"github.com/tanapon395/playlist-video/ent/stock"
	"github.com/tanapon395/playlist-video/ent/typeproduct"
	"github.com/tanapon395/playlist-video/ent/zoneproduct"
)

// StockCreate is the builder for creating a Stock entity.
type StockCreate struct {
	config
	mutation *StockMutation
	hooks    []Hook
}

// SetPriceproduct sets the Priceproduct field.
func (sc *StockCreate) SetPriceproduct(s string) *StockCreate {
	sc.mutation.SetPriceproduct(s)
	return sc
}

// SetAmount sets the Amount field.
func (sc *StockCreate) SetAmount(s string) *StockCreate {
	sc.mutation.SetAmount(s)
	return sc
}

// SetTime sets the Time field.
func (sc *StockCreate) SetTime(t time.Time) *StockCreate {
	sc.mutation.SetTime(t)
	return sc
}

// SetProductID sets the product edge to Product by id.
func (sc *StockCreate) SetProductID(id int) *StockCreate {
	sc.mutation.SetProductID(id)
	return sc
}

// SetProduct sets the product edge to Product.
func (sc *StockCreate) SetProduct(p *Product) *StockCreate {
	return sc.SetProductID(p.ID)
}

// SetZoneproductID sets the zoneproduct edge to Zoneproduct by id.
func (sc *StockCreate) SetZoneproductID(id int) *StockCreate {
	sc.mutation.SetZoneproductID(id)
	return sc
}

// SetNillableZoneproductID sets the zoneproduct edge to Zoneproduct by id if the given value is not nil.
func (sc *StockCreate) SetNillableZoneproductID(id *int) *StockCreate {
	if id != nil {
		sc = sc.SetZoneproductID(*id)
	}
	return sc
}

// SetZoneproduct sets the zoneproduct edge to Zoneproduct.
func (sc *StockCreate) SetZoneproduct(z *Zoneproduct) *StockCreate {
	return sc.SetZoneproductID(z.ID)
}

// SetEmployeeID sets the employee edge to Employee by id.
func (sc *StockCreate) SetEmployeeID(id int) *StockCreate {
	sc.mutation.SetEmployeeID(id)
	return sc
}

// SetNillableEmployeeID sets the employee edge to Employee by id if the given value is not nil.
func (sc *StockCreate) SetNillableEmployeeID(id *int) *StockCreate {
	if id != nil {
		sc = sc.SetEmployeeID(*id)
	}
	return sc
}

// SetEmployee sets the employee edge to Employee.
func (sc *StockCreate) SetEmployee(e *Employee) *StockCreate {
	return sc.SetEmployeeID(e.ID)
}

// SetTypeproductID sets the Typeproduct edge to Typeproduct by id.
func (sc *StockCreate) SetTypeproductID(id int) *StockCreate {
	sc.mutation.SetTypeproductID(id)
	return sc
}

// SetNillableTypeproductID sets the Typeproduct edge to Typeproduct by id if the given value is not nil.
func (sc *StockCreate) SetNillableTypeproductID(id *int) *StockCreate {
	if id != nil {
		sc = sc.SetTypeproductID(*id)
	}
	return sc
}

// SetTypeproduct sets the Typeproduct edge to Typeproduct.
func (sc *StockCreate) SetTypeproduct(t *Typeproduct) *StockCreate {
	return sc.SetTypeproductID(t.ID)
}

// Mutation returns the StockMutation object of the builder.
func (sc *StockCreate) Mutation() *StockMutation {
	return sc.mutation
}

// Save creates the Stock in the database.
func (sc *StockCreate) Save(ctx context.Context) (*Stock, error) {
	if _, ok := sc.mutation.Priceproduct(); !ok {
		return nil, &ValidationError{Name: "Priceproduct", err: errors.New("ent: missing required field \"Priceproduct\"")}
	}
	if _, ok := sc.mutation.Amount(); !ok {
		return nil, &ValidationError{Name: "Amount", err: errors.New("ent: missing required field \"Amount\"")}
	}
	if _, ok := sc.mutation.Time(); !ok {
		return nil, &ValidationError{Name: "Time", err: errors.New("ent: missing required field \"Time\"")}
	}
	if _, ok := sc.mutation.ProductID(); !ok {
		return nil, &ValidationError{Name: "product", err: errors.New("ent: missing required edge \"product\"")}
	}
	var (
		err  error
		node *Stock
	)
	if len(sc.hooks) == 0 {
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StockCreate) SaveX(ctx context.Context) *Stock {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *StockCreate) sqlSave(ctx context.Context) (*Stock, error) {
	s, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}

func (sc *StockCreate) createSpec() (*Stock, *sqlgraph.CreateSpec) {
	var (
		s     = &Stock{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: stock.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: stock.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Priceproduct(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: stock.FieldPriceproduct,
		})
		s.Priceproduct = value
	}
	if value, ok := sc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: stock.FieldAmount,
		})
		s.Amount = value
	}
	if value, ok := sc.mutation.Time(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: stock.FieldTime,
		})
		s.Time = value
	}
	if nodes := sc.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   stock.ProductTable,
			Columns: []string{stock.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ZoneproductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   stock.ZoneproductTable,
			Columns: []string{stock.ZoneproductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: zoneproduct.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.EmployeeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   stock.EmployeeTable,
			Columns: []string{stock.EmployeeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: employee.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.TypeproductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   stock.TypeproductTable,
			Columns: []string{stock.TypeproductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: typeproduct.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return s, _spec
}