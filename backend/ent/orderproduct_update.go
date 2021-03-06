// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team13/app/ent/company"
	"github.com/team13/app/ent/manager"
	"github.com/team13/app/ent/orderproduct"
	"github.com/team13/app/ent/predicate"
	"github.com/team13/app/ent/product"
	"github.com/team13/app/ent/typeproduct"
)

// OrderproductUpdate is the builder for updating Orderproduct entities.
type OrderproductUpdate struct {
	config
	hooks      []Hook
	mutation   *OrderproductMutation
	predicates []predicate.Orderproduct
}

// Where adds a new predicate for the builder.
func (ou *OrderproductUpdate) Where(ps ...predicate.Orderproduct) *OrderproductUpdate {
	ou.predicates = append(ou.predicates, ps...)
	return ou
}

// SetAddedtime sets the addedtime field.
func (ou *OrderproductUpdate) SetAddedtime(t time.Time) *OrderproductUpdate {
	ou.mutation.SetAddedtime(t)
	return ou
}

// SetStock sets the stock field.
func (ou *OrderproductUpdate) SetStock(i int) *OrderproductUpdate {
	ou.mutation.ResetStock()
	ou.mutation.SetStock(i)
	return ou
}

// AddStock adds i to stock.
func (ou *OrderproductUpdate) AddStock(i int) *OrderproductUpdate {
	ou.mutation.AddStock(i)
	return ou
}

// SetShipment sets the shipment field.
func (ou *OrderproductUpdate) SetShipment(s string) *OrderproductUpdate {
	ou.mutation.SetShipment(s)
	return ou
}

// SetDetail sets the detail field.
func (ou *OrderproductUpdate) SetDetail(s string) *OrderproductUpdate {
	ou.mutation.SetDetail(s)
	return ou
}

// SetProductID sets the product edge to Product by id.
func (ou *OrderproductUpdate) SetProductID(id int) *OrderproductUpdate {
	ou.mutation.SetProductID(id)
	return ou
}

// SetNillableProductID sets the product edge to Product by id if the given value is not nil.
func (ou *OrderproductUpdate) SetNillableProductID(id *int) *OrderproductUpdate {
	if id != nil {
		ou = ou.SetProductID(*id)
	}
	return ou
}

// SetProduct sets the product edge to Product.
func (ou *OrderproductUpdate) SetProduct(p *Product) *OrderproductUpdate {
	return ou.SetProductID(p.ID)
}

// SetCompanyID sets the company edge to Company by id.
func (ou *OrderproductUpdate) SetCompanyID(id int) *OrderproductUpdate {
	ou.mutation.SetCompanyID(id)
	return ou
}

// SetNillableCompanyID sets the company edge to Company by id if the given value is not nil.
func (ou *OrderproductUpdate) SetNillableCompanyID(id *int) *OrderproductUpdate {
	if id != nil {
		ou = ou.SetCompanyID(*id)
	}
	return ou
}

// SetCompany sets the company edge to Company.
func (ou *OrderproductUpdate) SetCompany(c *Company) *OrderproductUpdate {
	return ou.SetCompanyID(c.ID)
}

// SetTypeproductID sets the Typeproduct edge to Typeproduct by id.
func (ou *OrderproductUpdate) SetTypeproductID(id int) *OrderproductUpdate {
	ou.mutation.SetTypeproductID(id)
	return ou
}

// SetNillableTypeproductID sets the Typeproduct edge to Typeproduct by id if the given value is not nil.
func (ou *OrderproductUpdate) SetNillableTypeproductID(id *int) *OrderproductUpdate {
	if id != nil {
		ou = ou.SetTypeproductID(*id)
	}
	return ou
}

// SetTypeproduct sets the Typeproduct edge to Typeproduct.
func (ou *OrderproductUpdate) SetTypeproduct(t *Typeproduct) *OrderproductUpdate {
	return ou.SetTypeproductID(t.ID)
}

// SetManagersID sets the managers edge to Manager by id.
func (ou *OrderproductUpdate) SetManagersID(id int) *OrderproductUpdate {
	ou.mutation.SetManagersID(id)
	return ou
}

// SetNillableManagersID sets the managers edge to Manager by id if the given value is not nil.
func (ou *OrderproductUpdate) SetNillableManagersID(id *int) *OrderproductUpdate {
	if id != nil {
		ou = ou.SetManagersID(*id)
	}
	return ou
}

// SetManagers sets the managers edge to Manager.
func (ou *OrderproductUpdate) SetManagers(m *Manager) *OrderproductUpdate {
	return ou.SetManagersID(m.ID)
}

// Mutation returns the OrderproductMutation object of the builder.
func (ou *OrderproductUpdate) Mutation() *OrderproductMutation {
	return ou.mutation
}

// ClearProduct clears the product edge to Product.
func (ou *OrderproductUpdate) ClearProduct() *OrderproductUpdate {
	ou.mutation.ClearProduct()
	return ou
}

// ClearCompany clears the company edge to Company.
func (ou *OrderproductUpdate) ClearCompany() *OrderproductUpdate {
	ou.mutation.ClearCompany()
	return ou
}

// ClearTypeproduct clears the Typeproduct edge to Typeproduct.
func (ou *OrderproductUpdate) ClearTypeproduct() *OrderproductUpdate {
	ou.mutation.ClearTypeproduct()
	return ou
}

// ClearManagers clears the managers edge to Manager.
func (ou *OrderproductUpdate) ClearManagers() *OrderproductUpdate {
	ou.mutation.ClearManagers()
	return ou
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ou *OrderproductUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := ou.mutation.Stock(); ok {
		if err := orderproduct.StockValidator(v); err != nil {
			return 0, &ValidationError{Name: "stock", err: fmt.Errorf("ent: validator failed for field \"stock\": %w", err)}
		}
	}
	if v, ok := ou.mutation.Shipment(); ok {
		if err := orderproduct.ShipmentValidator(v); err != nil {
			return 0, &ValidationError{Name: "shipment", err: fmt.Errorf("ent: validator failed for field \"shipment\": %w", err)}
		}
	}
	if v, ok := ou.mutation.Detail(); ok {
		if err := orderproduct.DetailValidator(v); err != nil {
			return 0, &ValidationError{Name: "detail", err: fmt.Errorf("ent: validator failed for field \"detail\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(ou.hooks) == 0 {
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderproductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderproductUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderproductUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderproductUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ou *OrderproductUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderproduct.Table,
			Columns: orderproduct.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderproduct.FieldID,
			},
		},
	}
	if ps := ou.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.Addedtime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderproduct.FieldAddedtime,
		})
	}
	if value, ok := ou.mutation.Stock(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderproduct.FieldStock,
		})
	}
	if value, ok := ou.mutation.AddedStock(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderproduct.FieldStock,
		})
	}
	if value, ok := ou.mutation.Shipment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderproduct.FieldShipment,
		})
	}
	if value, ok := ou.mutation.Detail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderproduct.FieldDetail,
		})
	}
	if ou.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderproduct.ProductTable,
			Columns: []string{orderproduct.ProductColumn},
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
	if nodes := ou.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderproduct.ProductTable,
			Columns: []string{orderproduct.ProductColumn},
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
	if ou.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderproduct.CompanyTable,
			Columns: []string{orderproduct.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: company.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderproduct.CompanyTable,
			Columns: []string{orderproduct.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: company.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.TypeproductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderproduct.TypeproductTable,
			Columns: []string{orderproduct.TypeproductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: typeproduct.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.TypeproductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderproduct.TypeproductTable,
			Columns: []string{orderproduct.TypeproductColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ou.mutation.ManagersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderproduct.ManagersTable,
			Columns: []string{orderproduct.ManagersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: manager.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.ManagersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderproduct.ManagersTable,
			Columns: []string{orderproduct.ManagersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: manager.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderproduct.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// OrderproductUpdateOne is the builder for updating a single Orderproduct entity.
type OrderproductUpdateOne struct {
	config
	hooks    []Hook
	mutation *OrderproductMutation
}

// SetAddedtime sets the addedtime field.
func (ouo *OrderproductUpdateOne) SetAddedtime(t time.Time) *OrderproductUpdateOne {
	ouo.mutation.SetAddedtime(t)
	return ouo
}

// SetStock sets the stock field.
func (ouo *OrderproductUpdateOne) SetStock(i int) *OrderproductUpdateOne {
	ouo.mutation.ResetStock()
	ouo.mutation.SetStock(i)
	return ouo
}

// AddStock adds i to stock.
func (ouo *OrderproductUpdateOne) AddStock(i int) *OrderproductUpdateOne {
	ouo.mutation.AddStock(i)
	return ouo
}

// SetShipment sets the shipment field.
func (ouo *OrderproductUpdateOne) SetShipment(s string) *OrderproductUpdateOne {
	ouo.mutation.SetShipment(s)
	return ouo
}

// SetDetail sets the detail field.
func (ouo *OrderproductUpdateOne) SetDetail(s string) *OrderproductUpdateOne {
	ouo.mutation.SetDetail(s)
	return ouo
}

// SetProductID sets the product edge to Product by id.
func (ouo *OrderproductUpdateOne) SetProductID(id int) *OrderproductUpdateOne {
	ouo.mutation.SetProductID(id)
	return ouo
}

// SetNillableProductID sets the product edge to Product by id if the given value is not nil.
func (ouo *OrderproductUpdateOne) SetNillableProductID(id *int) *OrderproductUpdateOne {
	if id != nil {
		ouo = ouo.SetProductID(*id)
	}
	return ouo
}

// SetProduct sets the product edge to Product.
func (ouo *OrderproductUpdateOne) SetProduct(p *Product) *OrderproductUpdateOne {
	return ouo.SetProductID(p.ID)
}

// SetCompanyID sets the company edge to Company by id.
func (ouo *OrderproductUpdateOne) SetCompanyID(id int) *OrderproductUpdateOne {
	ouo.mutation.SetCompanyID(id)
	return ouo
}

// SetNillableCompanyID sets the company edge to Company by id if the given value is not nil.
func (ouo *OrderproductUpdateOne) SetNillableCompanyID(id *int) *OrderproductUpdateOne {
	if id != nil {
		ouo = ouo.SetCompanyID(*id)
	}
	return ouo
}

// SetCompany sets the company edge to Company.
func (ouo *OrderproductUpdateOne) SetCompany(c *Company) *OrderproductUpdateOne {
	return ouo.SetCompanyID(c.ID)
}

// SetTypeproductID sets the Typeproduct edge to Typeproduct by id.
func (ouo *OrderproductUpdateOne) SetTypeproductID(id int) *OrderproductUpdateOne {
	ouo.mutation.SetTypeproductID(id)
	return ouo
}

// SetNillableTypeproductID sets the Typeproduct edge to Typeproduct by id if the given value is not nil.
func (ouo *OrderproductUpdateOne) SetNillableTypeproductID(id *int) *OrderproductUpdateOne {
	if id != nil {
		ouo = ouo.SetTypeproductID(*id)
	}
	return ouo
}

// SetTypeproduct sets the Typeproduct edge to Typeproduct.
func (ouo *OrderproductUpdateOne) SetTypeproduct(t *Typeproduct) *OrderproductUpdateOne {
	return ouo.SetTypeproductID(t.ID)
}

// SetManagersID sets the managers edge to Manager by id.
func (ouo *OrderproductUpdateOne) SetManagersID(id int) *OrderproductUpdateOne {
	ouo.mutation.SetManagersID(id)
	return ouo
}

// SetNillableManagersID sets the managers edge to Manager by id if the given value is not nil.
func (ouo *OrderproductUpdateOne) SetNillableManagersID(id *int) *OrderproductUpdateOne {
	if id != nil {
		ouo = ouo.SetManagersID(*id)
	}
	return ouo
}

// SetManagers sets the managers edge to Manager.
func (ouo *OrderproductUpdateOne) SetManagers(m *Manager) *OrderproductUpdateOne {
	return ouo.SetManagersID(m.ID)
}

// Mutation returns the OrderproductMutation object of the builder.
func (ouo *OrderproductUpdateOne) Mutation() *OrderproductMutation {
	return ouo.mutation
}

// ClearProduct clears the product edge to Product.
func (ouo *OrderproductUpdateOne) ClearProduct() *OrderproductUpdateOne {
	ouo.mutation.ClearProduct()
	return ouo
}

// ClearCompany clears the company edge to Company.
func (ouo *OrderproductUpdateOne) ClearCompany() *OrderproductUpdateOne {
	ouo.mutation.ClearCompany()
	return ouo
}

// ClearTypeproduct clears the Typeproduct edge to Typeproduct.
func (ouo *OrderproductUpdateOne) ClearTypeproduct() *OrderproductUpdateOne {
	ouo.mutation.ClearTypeproduct()
	return ouo
}

// ClearManagers clears the managers edge to Manager.
func (ouo *OrderproductUpdateOne) ClearManagers() *OrderproductUpdateOne {
	ouo.mutation.ClearManagers()
	return ouo
}

// Save executes the query and returns the updated entity.
func (ouo *OrderproductUpdateOne) Save(ctx context.Context) (*Orderproduct, error) {
	if v, ok := ouo.mutation.Stock(); ok {
		if err := orderproduct.StockValidator(v); err != nil {
			return nil, &ValidationError{Name: "stock", err: fmt.Errorf("ent: validator failed for field \"stock\": %w", err)}
		}
	}
	if v, ok := ouo.mutation.Shipment(); ok {
		if err := orderproduct.ShipmentValidator(v); err != nil {
			return nil, &ValidationError{Name: "shipment", err: fmt.Errorf("ent: validator failed for field \"shipment\": %w", err)}
		}
	}
	if v, ok := ouo.mutation.Detail(); ok {
		if err := orderproduct.DetailValidator(v); err != nil {
			return nil, &ValidationError{Name: "detail", err: fmt.Errorf("ent: validator failed for field \"detail\": %w", err)}
		}
	}

	var (
		err  error
		node *Orderproduct
	)
	if len(ouo.hooks) == 0 {
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderproductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			mut = ouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderproductUpdateOne) SaveX(ctx context.Context) *Orderproduct {
	o, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return o
}

// Exec executes the query on the entity.
func (ouo *OrderproductUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderproductUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouo *OrderproductUpdateOne) sqlSave(ctx context.Context) (o *Orderproduct, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderproduct.Table,
			Columns: orderproduct.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderproduct.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Orderproduct.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ouo.mutation.Addedtime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderproduct.FieldAddedtime,
		})
	}
	if value, ok := ouo.mutation.Stock(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderproduct.FieldStock,
		})
	}
	if value, ok := ouo.mutation.AddedStock(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderproduct.FieldStock,
		})
	}
	if value, ok := ouo.mutation.Shipment(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderproduct.FieldShipment,
		})
	}
	if value, ok := ouo.mutation.Detail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderproduct.FieldDetail,
		})
	}
	if ouo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderproduct.ProductTable,
			Columns: []string{orderproduct.ProductColumn},
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
	if nodes := ouo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderproduct.ProductTable,
			Columns: []string{orderproduct.ProductColumn},
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
	if ouo.mutation.CompanyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderproduct.CompanyTable,
			Columns: []string{orderproduct.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: company.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderproduct.CompanyTable,
			Columns: []string{orderproduct.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: company.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.TypeproductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderproduct.TypeproductTable,
			Columns: []string{orderproduct.TypeproductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: typeproduct.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.TypeproductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderproduct.TypeproductTable,
			Columns: []string{orderproduct.TypeproductColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ouo.mutation.ManagersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderproduct.ManagersTable,
			Columns: []string{orderproduct.ManagersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: manager.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.ManagersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderproduct.ManagersTable,
			Columns: []string{orderproduct.ManagersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: manager.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	o = &Orderproduct{config: ouo.config}
	_spec.Assign = o.assignValues
	_spec.ScanValues = o.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderproduct.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return o, nil
}
