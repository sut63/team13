// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/discount"
	"github.com/tanapon395/playlist-video/ent/giveaway"
	"github.com/tanapon395/playlist-video/ent/product"
	"github.com/tanapon395/playlist-video/ent/promotion"
)

// PromotionCreate is the builder for creating a Promotion entity.
type PromotionCreate struct {
	config
	mutation *PromotionMutation
	hooks    []Hook
}

// SetPromotionName sets the PromotionName field.
func (pc *PromotionCreate) SetPromotionName(s string) *PromotionCreate {
	pc.mutation.SetPromotionName(s)
	return pc
}

// SetSaleID sets the sale edge to Discount by id.
func (pc *PromotionCreate) SetSaleID(id int) *PromotionCreate {
	pc.mutation.SetSaleID(id)
	return pc
}

// SetNillableSaleID sets the sale edge to Discount by id if the given value is not nil.
func (pc *PromotionCreate) SetNillableSaleID(id *int) *PromotionCreate {
	if id != nil {
		pc = pc.SetSaleID(*id)
	}
	return pc
}

// SetSale sets the sale edge to Discount.
func (pc *PromotionCreate) SetSale(d *Discount) *PromotionCreate {
	return pc.SetSaleID(d.ID)
}

// SetGiveID sets the give edge to Giveaway by id.
func (pc *PromotionCreate) SetGiveID(id int) *PromotionCreate {
	pc.mutation.SetGiveID(id)
	return pc
}

// SetNillableGiveID sets the give edge to Giveaway by id if the given value is not nil.
func (pc *PromotionCreate) SetNillableGiveID(id *int) *PromotionCreate {
	if id != nil {
		pc = pc.SetGiveID(*id)
	}
	return pc
}

// SetGive sets the give edge to Giveaway.
func (pc *PromotionCreate) SetGive(g *Giveaway) *PromotionCreate {
	return pc.SetGiveID(g.ID)
}

// SetProductID sets the product edge to Product by id.
func (pc *PromotionCreate) SetProductID(id int) *PromotionCreate {
	pc.mutation.SetProductID(id)
	return pc
}

// SetNillableProductID sets the product edge to Product by id if the given value is not nil.
func (pc *PromotionCreate) SetNillableProductID(id *int) *PromotionCreate {
	if id != nil {
		pc = pc.SetProductID(*id)
	}
	return pc
}

// SetProduct sets the product edge to Product.
func (pc *PromotionCreate) SetProduct(p *Product) *PromotionCreate {
	return pc.SetProductID(p.ID)
}

// Mutation returns the PromotionMutation object of the builder.
func (pc *PromotionCreate) Mutation() *PromotionMutation {
	return pc.mutation
}

// Save creates the Promotion in the database.
func (pc *PromotionCreate) Save(ctx context.Context) (*Promotion, error) {
	if _, ok := pc.mutation.PromotionName(); !ok {
		return nil, &ValidationError{Name: "PromotionName", err: errors.New("ent: missing required field \"PromotionName\"")}
	}
	var (
		err  error
		node *Promotion
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PromotionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PromotionCreate) SaveX(ctx context.Context) *Promotion {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PromotionCreate) sqlSave(ctx context.Context) (*Promotion, error) {
	pr, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pr.ID = int(id)
	return pr, nil
}

func (pc *PromotionCreate) createSpec() (*Promotion, *sqlgraph.CreateSpec) {
	var (
		pr    = &Promotion{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: promotion.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: promotion.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.PromotionName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: promotion.FieldPromotionName,
		})
		pr.PromotionName = value
	}
	if nodes := pc.mutation.SaleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   promotion.SaleTable,
			Columns: []string{promotion.SaleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: discount.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.GiveIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   promotion.GiveTable,
			Columns: []string{promotion.GiveColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: giveaway.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   promotion.ProductTable,
			Columns: []string{promotion.ProductColumn},
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
	return pr, _spec
}