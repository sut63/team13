// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team13/app/ent/orderonline"
	"github.com/team13/app/ent/orderproduct"
	"github.com/team13/app/ent/predicate"
	"github.com/team13/app/ent/product"
	"github.com/team13/app/ent/promotion"
	"github.com/team13/app/ent/stock"
)

// ProductUpdate is the builder for updating Product entities.
type ProductUpdate struct {
	config
	hooks      []Hook
	mutation   *ProductMutation
	predicates []predicate.Product
}

// Where adds a new predicate for the builder.
func (pu *ProductUpdate) Where(ps ...predicate.Product) *ProductUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetNameProduct sets the NameProduct field.
func (pu *ProductUpdate) SetNameProduct(s string) *ProductUpdate {
	pu.mutation.SetNameProduct(s)
	return pu
}

// SetBarcodeProduct sets the BarcodeProduct field.
func (pu *ProductUpdate) SetBarcodeProduct(s string) *ProductUpdate {
	pu.mutation.SetBarcodeProduct(s)
	return pu
}

// SetMFG sets the MFG field.
func (pu *ProductUpdate) SetMFG(s string) *ProductUpdate {
	pu.mutation.SetMFG(s)
	return pu
}

// SetEXP sets the EXP field.
func (pu *ProductUpdate) SetEXP(s string) *ProductUpdate {
	pu.mutation.SetEXP(s)
	return pu
}

// AddProductIDs adds the products edge to Orderproduct by ids.
func (pu *ProductUpdate) AddProductIDs(ids ...int) *ProductUpdate {
	pu.mutation.AddProductIDs(ids...)
	return pu
}

// AddProducts adds the products edges to Orderproduct.
func (pu *ProductUpdate) AddProducts(o ...*Orderproduct) *ProductUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pu.AddProductIDs(ids...)
}

// AddStockproductIDs adds the stockproduct edge to Stock by ids.
func (pu *ProductUpdate) AddStockproductIDs(ids ...int) *ProductUpdate {
	pu.mutation.AddStockproductIDs(ids...)
	return pu
}

// AddStockproduct adds the stockproduct edges to Stock.
func (pu *ProductUpdate) AddStockproduct(s ...*Stock) *ProductUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.AddStockproductIDs(ids...)
}

// SetForproductID sets the forproduct edge to Promotion by id.
func (pu *ProductUpdate) SetForproductID(id int) *ProductUpdate {
	pu.mutation.SetForproductID(id)
	return pu
}

// SetNillableForproductID sets the forproduct edge to Promotion by id if the given value is not nil.
func (pu *ProductUpdate) SetNillableForproductID(id *int) *ProductUpdate {
	if id != nil {
		pu = pu.SetForproductID(*id)
	}
	return pu
}

// SetForproduct sets the forproduct edge to Promotion.
func (pu *ProductUpdate) SetForproduct(p *Promotion) *ProductUpdate {
	return pu.SetForproductID(p.ID)
}

// AddFormproductonlineIDs adds the formproductonline edge to Orderonline by ids.
func (pu *ProductUpdate) AddFormproductonlineIDs(ids ...int) *ProductUpdate {
	pu.mutation.AddFormproductonlineIDs(ids...)
	return pu
}

// AddFormproductonline adds the formproductonline edges to Orderonline.
func (pu *ProductUpdate) AddFormproductonline(o ...*Orderonline) *ProductUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pu.AddFormproductonlineIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (pu *ProductUpdate) Mutation() *ProductMutation {
	return pu.mutation
}

// RemoveProductIDs removes the products edge to Orderproduct by ids.
func (pu *ProductUpdate) RemoveProductIDs(ids ...int) *ProductUpdate {
	pu.mutation.RemoveProductIDs(ids...)
	return pu
}

// RemoveProducts removes products edges to Orderproduct.
func (pu *ProductUpdate) RemoveProducts(o ...*Orderproduct) *ProductUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pu.RemoveProductIDs(ids...)
}

// RemoveStockproductIDs removes the stockproduct edge to Stock by ids.
func (pu *ProductUpdate) RemoveStockproductIDs(ids ...int) *ProductUpdate {
	pu.mutation.RemoveStockproductIDs(ids...)
	return pu
}

// RemoveStockproduct removes stockproduct edges to Stock.
func (pu *ProductUpdate) RemoveStockproduct(s ...*Stock) *ProductUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.RemoveStockproductIDs(ids...)
}

// ClearForproduct clears the forproduct edge to Promotion.
func (pu *ProductUpdate) ClearForproduct() *ProductUpdate {
	pu.mutation.ClearForproduct()
	return pu
}

// RemoveFormproductonlineIDs removes the formproductonline edge to Orderonline by ids.
func (pu *ProductUpdate) RemoveFormproductonlineIDs(ids ...int) *ProductUpdate {
	pu.mutation.RemoveFormproductonlineIDs(ids...)
	return pu
}

// RemoveFormproductonline removes formproductonline edges to Orderonline.
func (pu *ProductUpdate) RemoveFormproductonline(o ...*Orderonline) *ProductUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pu.RemoveFormproductonlineIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *ProductUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.NameProduct(); ok {
		if err := product.NameProductValidator(v); err != nil {
			return 0, &ValidationError{Name: "NameProduct", err: fmt.Errorf("ent: validator failed for field \"NameProduct\": %w", err)}
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
			mutation, ok := m.(*ProductMutation)
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
func (pu *ProductUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProductUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProductUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProductUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: product.FieldID,
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
	if value, ok := pu.mutation.NameProduct(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldNameProduct,
		})
	}
	if value, ok := pu.mutation.BarcodeProduct(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldBarcodeProduct,
		})
	}
	if value, ok := pu.mutation.MFG(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldMFG,
		})
	}
	if value, ok := pu.mutation.EXP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldEXP,
		})
	}
	if nodes := pu.mutation.RemovedProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductsTable,
			Columns: []string{product.ProductsColumn},
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
	if nodes := pu.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductsTable,
			Columns: []string{product.ProductsColumn},
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
	if nodes := pu.mutation.RemovedStockproductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.StockproductTable,
			Columns: []string{product.StockproductColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.StockproductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.StockproductTable,
			Columns: []string{product.StockproductColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ForproductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   product.ForproductTable,
			Columns: []string{product.ForproductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: promotion.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ForproductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   product.ForproductTable,
			Columns: []string{product.ForproductColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := pu.mutation.RemovedFormproductonlineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.FormproductonlineTable,
			Columns: []string{product.FormproductonlineColumn},
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
	if nodes := pu.mutation.FormproductonlineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.FormproductonlineTable,
			Columns: []string{product.FormproductonlineColumn},
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
			err = &NotFoundError{product.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ProductUpdateOne is the builder for updating a single Product entity.
type ProductUpdateOne struct {
	config
	hooks    []Hook
	mutation *ProductMutation
}

// SetNameProduct sets the NameProduct field.
func (puo *ProductUpdateOne) SetNameProduct(s string) *ProductUpdateOne {
	puo.mutation.SetNameProduct(s)
	return puo
}

// SetBarcodeProduct sets the BarcodeProduct field.
func (puo *ProductUpdateOne) SetBarcodeProduct(s string) *ProductUpdateOne {
	puo.mutation.SetBarcodeProduct(s)
	return puo
}

// SetMFG sets the MFG field.
func (puo *ProductUpdateOne) SetMFG(s string) *ProductUpdateOne {
	puo.mutation.SetMFG(s)
	return puo
}

// SetEXP sets the EXP field.
func (puo *ProductUpdateOne) SetEXP(s string) *ProductUpdateOne {
	puo.mutation.SetEXP(s)
	return puo
}

// AddProductIDs adds the products edge to Orderproduct by ids.
func (puo *ProductUpdateOne) AddProductIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.AddProductIDs(ids...)
	return puo
}

// AddProducts adds the products edges to Orderproduct.
func (puo *ProductUpdateOne) AddProducts(o ...*Orderproduct) *ProductUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return puo.AddProductIDs(ids...)
}

// AddStockproductIDs adds the stockproduct edge to Stock by ids.
func (puo *ProductUpdateOne) AddStockproductIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.AddStockproductIDs(ids...)
	return puo
}

// AddStockproduct adds the stockproduct edges to Stock.
func (puo *ProductUpdateOne) AddStockproduct(s ...*Stock) *ProductUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.AddStockproductIDs(ids...)
}

// SetForproductID sets the forproduct edge to Promotion by id.
func (puo *ProductUpdateOne) SetForproductID(id int) *ProductUpdateOne {
	puo.mutation.SetForproductID(id)
	return puo
}

// SetNillableForproductID sets the forproduct edge to Promotion by id if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableForproductID(id *int) *ProductUpdateOne {
	if id != nil {
		puo = puo.SetForproductID(*id)
	}
	return puo
}

// SetForproduct sets the forproduct edge to Promotion.
func (puo *ProductUpdateOne) SetForproduct(p *Promotion) *ProductUpdateOne {
	return puo.SetForproductID(p.ID)
}

// AddFormproductonlineIDs adds the formproductonline edge to Orderonline by ids.
func (puo *ProductUpdateOne) AddFormproductonlineIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.AddFormproductonlineIDs(ids...)
	return puo
}

// AddFormproductonline adds the formproductonline edges to Orderonline.
func (puo *ProductUpdateOne) AddFormproductonline(o ...*Orderonline) *ProductUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return puo.AddFormproductonlineIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (puo *ProductUpdateOne) Mutation() *ProductMutation {
	return puo.mutation
}

// RemoveProductIDs removes the products edge to Orderproduct by ids.
func (puo *ProductUpdateOne) RemoveProductIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.RemoveProductIDs(ids...)
	return puo
}

// RemoveProducts removes products edges to Orderproduct.
func (puo *ProductUpdateOne) RemoveProducts(o ...*Orderproduct) *ProductUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return puo.RemoveProductIDs(ids...)
}

// RemoveStockproductIDs removes the stockproduct edge to Stock by ids.
func (puo *ProductUpdateOne) RemoveStockproductIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.RemoveStockproductIDs(ids...)
	return puo
}

// RemoveStockproduct removes stockproduct edges to Stock.
func (puo *ProductUpdateOne) RemoveStockproduct(s ...*Stock) *ProductUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.RemoveStockproductIDs(ids...)
}

// ClearForproduct clears the forproduct edge to Promotion.
func (puo *ProductUpdateOne) ClearForproduct() *ProductUpdateOne {
	puo.mutation.ClearForproduct()
	return puo
}

// RemoveFormproductonlineIDs removes the formproductonline edge to Orderonline by ids.
func (puo *ProductUpdateOne) RemoveFormproductonlineIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.RemoveFormproductonlineIDs(ids...)
	return puo
}

// RemoveFormproductonline removes formproductonline edges to Orderonline.
func (puo *ProductUpdateOne) RemoveFormproductonline(o ...*Orderonline) *ProductUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return puo.RemoveFormproductonlineIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *ProductUpdateOne) Save(ctx context.Context) (*Product, error) {
	if v, ok := puo.mutation.NameProduct(); ok {
		if err := product.NameProductValidator(v); err != nil {
			return nil, &ValidationError{Name: "NameProduct", err: fmt.Errorf("ent: validator failed for field \"NameProduct\": %w", err)}
		}
	}

	var (
		err  error
		node *Product
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
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
func (puo *ProductUpdateOne) SaveX(ctx context.Context) *Product {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *ProductUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProductUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProductUpdateOne) sqlSave(ctx context.Context) (pr *Product, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: product.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Product.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.NameProduct(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldNameProduct,
		})
	}
	if value, ok := puo.mutation.BarcodeProduct(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldBarcodeProduct,
		})
	}
	if value, ok := puo.mutation.MFG(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldMFG,
		})
	}
	if value, ok := puo.mutation.EXP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldEXP,
		})
	}
	if nodes := puo.mutation.RemovedProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductsTable,
			Columns: []string{product.ProductsColumn},
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
	if nodes := puo.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductsTable,
			Columns: []string{product.ProductsColumn},
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
	if nodes := puo.mutation.RemovedStockproductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.StockproductTable,
			Columns: []string{product.StockproductColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.StockproductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.StockproductTable,
			Columns: []string{product.StockproductColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ForproductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   product.ForproductTable,
			Columns: []string{product.ForproductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: promotion.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ForproductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   product.ForproductTable,
			Columns: []string{product.ForproductColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := puo.mutation.RemovedFormproductonlineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.FormproductonlineTable,
			Columns: []string{product.FormproductonlineColumn},
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
	if nodes := puo.mutation.FormproductonlineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.FormproductonlineTable,
			Columns: []string{product.FormproductonlineColumn},
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
	pr = &Product{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}
