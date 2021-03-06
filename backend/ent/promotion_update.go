// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team13/app/ent/discount"
	"github.com/team13/app/ent/giveaway"
	"github.com/team13/app/ent/predicate"
	"github.com/team13/app/ent/product"
	"github.com/team13/app/ent/promotion"
)

// PromotionUpdate is the builder for updating Promotion entities.
type PromotionUpdate struct {
	config
	hooks      []Hook
	mutation   *PromotionMutation
	predicates []predicate.Promotion
}

// Where adds a new predicate for the builder.
func (pu *PromotionUpdate) Where(ps ...predicate.Promotion) *PromotionUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetPromotionName sets the PromotionName field.
func (pu *PromotionUpdate) SetPromotionName(s string) *PromotionUpdate {
	pu.mutation.SetPromotionName(s)
	return pu
}

// SetStartPromotion sets the StartPromotion field.
func (pu *PromotionUpdate) SetStartPromotion(t time.Time) *PromotionUpdate {
	pu.mutation.SetStartPromotion(t)
	return pu
}

// SetEndPromotion sets the EndPromotion field.
func (pu *PromotionUpdate) SetEndPromotion(t time.Time) *PromotionUpdate {
	pu.mutation.SetEndPromotion(t)
	return pu
}

// SetPrice sets the Price field.
func (pu *PromotionUpdate) SetPrice(f float64) *PromotionUpdate {
	pu.mutation.ResetPrice()
	pu.mutation.SetPrice(f)
	return pu
}

// AddPrice adds f to Price.
func (pu *PromotionUpdate) AddPrice(f float64) *PromotionUpdate {
	pu.mutation.AddPrice(f)
	return pu
}

// SetSaleID sets the sale edge to Discount by id.
func (pu *PromotionUpdate) SetSaleID(id int) *PromotionUpdate {
	pu.mutation.SetSaleID(id)
	return pu
}

// SetNillableSaleID sets the sale edge to Discount by id if the given value is not nil.
func (pu *PromotionUpdate) SetNillableSaleID(id *int) *PromotionUpdate {
	if id != nil {
		pu = pu.SetSaleID(*id)
	}
	return pu
}

// SetSale sets the sale edge to Discount.
func (pu *PromotionUpdate) SetSale(d *Discount) *PromotionUpdate {
	return pu.SetSaleID(d.ID)
}

// SetGiveID sets the give edge to Giveaway by id.
func (pu *PromotionUpdate) SetGiveID(id int) *PromotionUpdate {
	pu.mutation.SetGiveID(id)
	return pu
}

// SetNillableGiveID sets the give edge to Giveaway by id if the given value is not nil.
func (pu *PromotionUpdate) SetNillableGiveID(id *int) *PromotionUpdate {
	if id != nil {
		pu = pu.SetGiveID(*id)
	}
	return pu
}

// SetGive sets the give edge to Giveaway.
func (pu *PromotionUpdate) SetGive(g *Giveaway) *PromotionUpdate {
	return pu.SetGiveID(g.ID)
}

// SetProductID sets the product edge to Product by id.
func (pu *PromotionUpdate) SetProductID(id int) *PromotionUpdate {
	pu.mutation.SetProductID(id)
	return pu
}

// SetNillableProductID sets the product edge to Product by id if the given value is not nil.
func (pu *PromotionUpdate) SetNillableProductID(id *int) *PromotionUpdate {
	if id != nil {
		pu = pu.SetProductID(*id)
	}
	return pu
}

// SetProduct sets the product edge to Product.
func (pu *PromotionUpdate) SetProduct(p *Product) *PromotionUpdate {
	return pu.SetProductID(p.ID)
}

// Mutation returns the PromotionMutation object of the builder.
func (pu *PromotionUpdate) Mutation() *PromotionMutation {
	return pu.mutation
}

// ClearSale clears the sale edge to Discount.
func (pu *PromotionUpdate) ClearSale() *PromotionUpdate {
	pu.mutation.ClearSale()
	return pu
}

// ClearGive clears the give edge to Giveaway.
func (pu *PromotionUpdate) ClearGive() *PromotionUpdate {
	pu.mutation.ClearGive()
	return pu
}

// ClearProduct clears the product edge to Product.
func (pu *PromotionUpdate) ClearProduct() *PromotionUpdate {
	pu.mutation.ClearProduct()
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PromotionUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.PromotionName(); ok {
		if err := promotion.PromotionNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "PromotionName", err: fmt.Errorf("ent: validator failed for field \"PromotionName\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Price(); ok {
		if err := promotion.PriceValidator(v); err != nil {
			return 0, &ValidationError{Name: "Price", err: fmt.Errorf("ent: validator failed for field \"Price\": %w", err)}
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
			mutation, ok := m.(*PromotionMutation)
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
func (pu *PromotionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PromotionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PromotionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PromotionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   promotion.Table,
			Columns: promotion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: promotion.FieldID,
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
	if value, ok := pu.mutation.PromotionName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: promotion.FieldPromotionName,
		})
	}
	if value, ok := pu.mutation.StartPromotion(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: promotion.FieldStartPromotion,
		})
	}
	if value, ok := pu.mutation.EndPromotion(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: promotion.FieldEndPromotion,
		})
	}
	if value, ok := pu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: promotion.FieldPrice,
		})
	}
	if value, ok := pu.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: promotion.FieldPrice,
		})
	}
	if pu.mutation.SaleCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.SaleIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.GiveCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.GiveIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{promotion.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PromotionUpdateOne is the builder for updating a single Promotion entity.
type PromotionUpdateOne struct {
	config
	hooks    []Hook
	mutation *PromotionMutation
}

// SetPromotionName sets the PromotionName field.
func (puo *PromotionUpdateOne) SetPromotionName(s string) *PromotionUpdateOne {
	puo.mutation.SetPromotionName(s)
	return puo
}

// SetStartPromotion sets the StartPromotion field.
func (puo *PromotionUpdateOne) SetStartPromotion(t time.Time) *PromotionUpdateOne {
	puo.mutation.SetStartPromotion(t)
	return puo
}

// SetEndPromotion sets the EndPromotion field.
func (puo *PromotionUpdateOne) SetEndPromotion(t time.Time) *PromotionUpdateOne {
	puo.mutation.SetEndPromotion(t)
	return puo
}

// SetPrice sets the Price field.
func (puo *PromotionUpdateOne) SetPrice(f float64) *PromotionUpdateOne {
	puo.mutation.ResetPrice()
	puo.mutation.SetPrice(f)
	return puo
}

// AddPrice adds f to Price.
func (puo *PromotionUpdateOne) AddPrice(f float64) *PromotionUpdateOne {
	puo.mutation.AddPrice(f)
	return puo
}

// SetSaleID sets the sale edge to Discount by id.
func (puo *PromotionUpdateOne) SetSaleID(id int) *PromotionUpdateOne {
	puo.mutation.SetSaleID(id)
	return puo
}

// SetNillableSaleID sets the sale edge to Discount by id if the given value is not nil.
func (puo *PromotionUpdateOne) SetNillableSaleID(id *int) *PromotionUpdateOne {
	if id != nil {
		puo = puo.SetSaleID(*id)
	}
	return puo
}

// SetSale sets the sale edge to Discount.
func (puo *PromotionUpdateOne) SetSale(d *Discount) *PromotionUpdateOne {
	return puo.SetSaleID(d.ID)
}

// SetGiveID sets the give edge to Giveaway by id.
func (puo *PromotionUpdateOne) SetGiveID(id int) *PromotionUpdateOne {
	puo.mutation.SetGiveID(id)
	return puo
}

// SetNillableGiveID sets the give edge to Giveaway by id if the given value is not nil.
func (puo *PromotionUpdateOne) SetNillableGiveID(id *int) *PromotionUpdateOne {
	if id != nil {
		puo = puo.SetGiveID(*id)
	}
	return puo
}

// SetGive sets the give edge to Giveaway.
func (puo *PromotionUpdateOne) SetGive(g *Giveaway) *PromotionUpdateOne {
	return puo.SetGiveID(g.ID)
}

// SetProductID sets the product edge to Product by id.
func (puo *PromotionUpdateOne) SetProductID(id int) *PromotionUpdateOne {
	puo.mutation.SetProductID(id)
	return puo
}

// SetNillableProductID sets the product edge to Product by id if the given value is not nil.
func (puo *PromotionUpdateOne) SetNillableProductID(id *int) *PromotionUpdateOne {
	if id != nil {
		puo = puo.SetProductID(*id)
	}
	return puo
}

// SetProduct sets the product edge to Product.
func (puo *PromotionUpdateOne) SetProduct(p *Product) *PromotionUpdateOne {
	return puo.SetProductID(p.ID)
}

// Mutation returns the PromotionMutation object of the builder.
func (puo *PromotionUpdateOne) Mutation() *PromotionMutation {
	return puo.mutation
}

// ClearSale clears the sale edge to Discount.
func (puo *PromotionUpdateOne) ClearSale() *PromotionUpdateOne {
	puo.mutation.ClearSale()
	return puo
}

// ClearGive clears the give edge to Giveaway.
func (puo *PromotionUpdateOne) ClearGive() *PromotionUpdateOne {
	puo.mutation.ClearGive()
	return puo
}

// ClearProduct clears the product edge to Product.
func (puo *PromotionUpdateOne) ClearProduct() *PromotionUpdateOne {
	puo.mutation.ClearProduct()
	return puo
}

// Save executes the query and returns the updated entity.
func (puo *PromotionUpdateOne) Save(ctx context.Context) (*Promotion, error) {
	if v, ok := puo.mutation.PromotionName(); ok {
		if err := promotion.PromotionNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "PromotionName", err: fmt.Errorf("ent: validator failed for field \"PromotionName\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Price(); ok {
		if err := promotion.PriceValidator(v); err != nil {
			return nil, &ValidationError{Name: "Price", err: fmt.Errorf("ent: validator failed for field \"Price\": %w", err)}
		}
	}

	var (
		err  error
		node *Promotion
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PromotionMutation)
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
func (puo *PromotionUpdateOne) SaveX(ctx context.Context) *Promotion {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *PromotionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PromotionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PromotionUpdateOne) sqlSave(ctx context.Context) (pr *Promotion, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   promotion.Table,
			Columns: promotion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: promotion.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Promotion.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.PromotionName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: promotion.FieldPromotionName,
		})
	}
	if value, ok := puo.mutation.StartPromotion(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: promotion.FieldStartPromotion,
		})
	}
	if value, ok := puo.mutation.EndPromotion(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: promotion.FieldEndPromotion,
		})
	}
	if value, ok := puo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: promotion.FieldPrice,
		})
	}
	if value, ok := puo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: promotion.FieldPrice,
		})
	}
	if puo.mutation.SaleCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.SaleIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.GiveCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.GiveIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pr = &Promotion{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{promotion.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}
