// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tanapon395/playlist-video/ent/customer"
	"github.com/tanapon395/playlist-video/ent/orderonline"
	"github.com/tanapon395/playlist-video/ent/paymentchannel"
	"github.com/tanapon395/playlist-video/ent/predicate"
	"github.com/tanapon395/playlist-video/ent/product"
	"github.com/tanapon395/playlist-video/ent/typeproduct"
)

// OrderonlineUpdate is the builder for updating Orderonline entities.
type OrderonlineUpdate struct {
	config
	hooks      []Hook
	mutation   *OrderonlineMutation
	predicates []predicate.Orderonline
}

// Where adds a new predicate for the builder.
func (ou *OrderonlineUpdate) Where(ps ...predicate.Orderonline) *OrderonlineUpdate {
	ou.predicates = append(ou.predicates, ps...)
	return ou
}

// SetAddedtime sets the addedtime field.
func (ou *OrderonlineUpdate) SetAddedtime(t time.Time) *OrderonlineUpdate {
	ou.mutation.SetAddedtime(t)
	return ou
}

// SetStock sets the stock field.
func (ou *OrderonlineUpdate) SetStock(i int) *OrderonlineUpdate {
	ou.mutation.ResetStock()
	ou.mutation.SetStock(i)
	return ou
}

// AddStock adds i to stock.
func (ou *OrderonlineUpdate) AddStock(i int) *OrderonlineUpdate {
	ou.mutation.AddStock(i)
	return ou
}

// SetProductonID sets the producton edge to Product by id.
func (ou *OrderonlineUpdate) SetProductonID(id int) *OrderonlineUpdate {
	ou.mutation.SetProductonID(id)
	return ou
}

// SetNillableProductonID sets the producton edge to Product by id if the given value is not nil.
func (ou *OrderonlineUpdate) SetNillableProductonID(id *int) *OrderonlineUpdate {
	if id != nil {
		ou = ou.SetProductonID(*id)
	}
	return ou
}

// SetProducton sets the producton edge to Product.
func (ou *OrderonlineUpdate) SetProducton(p *Product) *OrderonlineUpdate {
	return ou.SetProductonID(p.ID)
}

// SetPaymentchannelID sets the paymentchannel edge to Paymentchannel by id.
func (ou *OrderonlineUpdate) SetPaymentchannelID(id int) *OrderonlineUpdate {
	ou.mutation.SetPaymentchannelID(id)
	return ou
}

// SetNillablePaymentchannelID sets the paymentchannel edge to Paymentchannel by id if the given value is not nil.
func (ou *OrderonlineUpdate) SetNillablePaymentchannelID(id *int) *OrderonlineUpdate {
	if id != nil {
		ou = ou.SetPaymentchannelID(*id)
	}
	return ou
}

// SetPaymentchannel sets the paymentchannel edge to Paymentchannel.
func (ou *OrderonlineUpdate) SetPaymentchannel(p *Paymentchannel) *OrderonlineUpdate {
	return ou.SetPaymentchannelID(p.ID)
}

// SetTypeproductID sets the Typeproduct edge to Typeproduct by id.
func (ou *OrderonlineUpdate) SetTypeproductID(id int) *OrderonlineUpdate {
	ou.mutation.SetTypeproductID(id)
	return ou
}

// SetNillableTypeproductID sets the Typeproduct edge to Typeproduct by id if the given value is not nil.
func (ou *OrderonlineUpdate) SetNillableTypeproductID(id *int) *OrderonlineUpdate {
	if id != nil {
		ou = ou.SetTypeproductID(*id)
	}
	return ou
}

// SetTypeproduct sets the Typeproduct edge to Typeproduct.
func (ou *OrderonlineUpdate) SetTypeproduct(t *Typeproduct) *OrderonlineUpdate {
	return ou.SetTypeproductID(t.ID)
}

// SetCustomerID sets the customer edge to Customer by id.
func (ou *OrderonlineUpdate) SetCustomerID(id int) *OrderonlineUpdate {
	ou.mutation.SetCustomerID(id)
	return ou
}

// SetNillableCustomerID sets the customer edge to Customer by id if the given value is not nil.
func (ou *OrderonlineUpdate) SetNillableCustomerID(id *int) *OrderonlineUpdate {
	if id != nil {
		ou = ou.SetCustomerID(*id)
	}
	return ou
}

// SetCustomer sets the customer edge to Customer.
func (ou *OrderonlineUpdate) SetCustomer(c *Customer) *OrderonlineUpdate {
	return ou.SetCustomerID(c.ID)
}

// Mutation returns the OrderonlineMutation object of the builder.
func (ou *OrderonlineUpdate) Mutation() *OrderonlineMutation {
	return ou.mutation
}

// ClearProducton clears the producton edge to Product.
func (ou *OrderonlineUpdate) ClearProducton() *OrderonlineUpdate {
	ou.mutation.ClearProducton()
	return ou
}

// ClearPaymentchannel clears the paymentchannel edge to Paymentchannel.
func (ou *OrderonlineUpdate) ClearPaymentchannel() *OrderonlineUpdate {
	ou.mutation.ClearPaymentchannel()
	return ou
}

// ClearTypeproduct clears the Typeproduct edge to Typeproduct.
func (ou *OrderonlineUpdate) ClearTypeproduct() *OrderonlineUpdate {
	ou.mutation.ClearTypeproduct()
	return ou
}

// ClearCustomer clears the customer edge to Customer.
func (ou *OrderonlineUpdate) ClearCustomer() *OrderonlineUpdate {
	ou.mutation.ClearCustomer()
	return ou
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ou *OrderonlineUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(ou.hooks) == 0 {
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderonlineMutation)
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
func (ou *OrderonlineUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderonlineUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderonlineUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ou *OrderonlineUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderonline.Table,
			Columns: orderonline.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderonline.FieldID,
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
			Column: orderonline.FieldAddedtime,
		})
	}
	if value, ok := ou.mutation.Stock(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderonline.FieldStock,
		})
	}
	if value, ok := ou.mutation.AddedStock(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderonline.FieldStock,
		})
	}
	if ou.mutation.ProductonCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderonline.ProductonTable,
			Columns: []string{orderonline.ProductonColumn},
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
	if nodes := ou.mutation.ProductonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderonline.ProductonTable,
			Columns: []string{orderonline.ProductonColumn},
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
	if ou.mutation.PaymentchannelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderonline.PaymentchannelTable,
			Columns: []string{orderonline.PaymentchannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: paymentchannel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.PaymentchannelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderonline.PaymentchannelTable,
			Columns: []string{orderonline.PaymentchannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: paymentchannel.FieldID,
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
			Table:   orderonline.TypeproductTable,
			Columns: []string{orderonline.TypeproductColumn},
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
			Table:   orderonline.TypeproductTable,
			Columns: []string{orderonline.TypeproductColumn},
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
	if ou.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderonline.CustomerTable,
			Columns: []string{orderonline.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderonline.CustomerTable,
			Columns: []string{orderonline.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
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
			err = &NotFoundError{orderonline.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// OrderonlineUpdateOne is the builder for updating a single Orderonline entity.
type OrderonlineUpdateOne struct {
	config
	hooks    []Hook
	mutation *OrderonlineMutation
}

// SetAddedtime sets the addedtime field.
func (ouo *OrderonlineUpdateOne) SetAddedtime(t time.Time) *OrderonlineUpdateOne {
	ouo.mutation.SetAddedtime(t)
	return ouo
}

// SetStock sets the stock field.
func (ouo *OrderonlineUpdateOne) SetStock(i int) *OrderonlineUpdateOne {
	ouo.mutation.ResetStock()
	ouo.mutation.SetStock(i)
	return ouo
}

// AddStock adds i to stock.
func (ouo *OrderonlineUpdateOne) AddStock(i int) *OrderonlineUpdateOne {
	ouo.mutation.AddStock(i)
	return ouo
}

// SetProductonID sets the producton edge to Product by id.
func (ouo *OrderonlineUpdateOne) SetProductonID(id int) *OrderonlineUpdateOne {
	ouo.mutation.SetProductonID(id)
	return ouo
}

// SetNillableProductonID sets the producton edge to Product by id if the given value is not nil.
func (ouo *OrderonlineUpdateOne) SetNillableProductonID(id *int) *OrderonlineUpdateOne {
	if id != nil {
		ouo = ouo.SetProductonID(*id)
	}
	return ouo
}

// SetProducton sets the producton edge to Product.
func (ouo *OrderonlineUpdateOne) SetProducton(p *Product) *OrderonlineUpdateOne {
	return ouo.SetProductonID(p.ID)
}

// SetPaymentchannelID sets the paymentchannel edge to Paymentchannel by id.
func (ouo *OrderonlineUpdateOne) SetPaymentchannelID(id int) *OrderonlineUpdateOne {
	ouo.mutation.SetPaymentchannelID(id)
	return ouo
}

// SetNillablePaymentchannelID sets the paymentchannel edge to Paymentchannel by id if the given value is not nil.
func (ouo *OrderonlineUpdateOne) SetNillablePaymentchannelID(id *int) *OrderonlineUpdateOne {
	if id != nil {
		ouo = ouo.SetPaymentchannelID(*id)
	}
	return ouo
}

// SetPaymentchannel sets the paymentchannel edge to Paymentchannel.
func (ouo *OrderonlineUpdateOne) SetPaymentchannel(p *Paymentchannel) *OrderonlineUpdateOne {
	return ouo.SetPaymentchannelID(p.ID)
}

// SetTypeproductID sets the Typeproduct edge to Typeproduct by id.
func (ouo *OrderonlineUpdateOne) SetTypeproductID(id int) *OrderonlineUpdateOne {
	ouo.mutation.SetTypeproductID(id)
	return ouo
}

// SetNillableTypeproductID sets the Typeproduct edge to Typeproduct by id if the given value is not nil.
func (ouo *OrderonlineUpdateOne) SetNillableTypeproductID(id *int) *OrderonlineUpdateOne {
	if id != nil {
		ouo = ouo.SetTypeproductID(*id)
	}
	return ouo
}

// SetTypeproduct sets the Typeproduct edge to Typeproduct.
func (ouo *OrderonlineUpdateOne) SetTypeproduct(t *Typeproduct) *OrderonlineUpdateOne {
	return ouo.SetTypeproductID(t.ID)
}

// SetCustomerID sets the customer edge to Customer by id.
func (ouo *OrderonlineUpdateOne) SetCustomerID(id int) *OrderonlineUpdateOne {
	ouo.mutation.SetCustomerID(id)
	return ouo
}

// SetNillableCustomerID sets the customer edge to Customer by id if the given value is not nil.
func (ouo *OrderonlineUpdateOne) SetNillableCustomerID(id *int) *OrderonlineUpdateOne {
	if id != nil {
		ouo = ouo.SetCustomerID(*id)
	}
	return ouo
}

// SetCustomer sets the customer edge to Customer.
func (ouo *OrderonlineUpdateOne) SetCustomer(c *Customer) *OrderonlineUpdateOne {
	return ouo.SetCustomerID(c.ID)
}

// Mutation returns the OrderonlineMutation object of the builder.
func (ouo *OrderonlineUpdateOne) Mutation() *OrderonlineMutation {
	return ouo.mutation
}

// ClearProducton clears the producton edge to Product.
func (ouo *OrderonlineUpdateOne) ClearProducton() *OrderonlineUpdateOne {
	ouo.mutation.ClearProducton()
	return ouo
}

// ClearPaymentchannel clears the paymentchannel edge to Paymentchannel.
func (ouo *OrderonlineUpdateOne) ClearPaymentchannel() *OrderonlineUpdateOne {
	ouo.mutation.ClearPaymentchannel()
	return ouo
}

// ClearTypeproduct clears the Typeproduct edge to Typeproduct.
func (ouo *OrderonlineUpdateOne) ClearTypeproduct() *OrderonlineUpdateOne {
	ouo.mutation.ClearTypeproduct()
	return ouo
}

// ClearCustomer clears the customer edge to Customer.
func (ouo *OrderonlineUpdateOne) ClearCustomer() *OrderonlineUpdateOne {
	ouo.mutation.ClearCustomer()
	return ouo
}

// Save executes the query and returns the updated entity.
func (ouo *OrderonlineUpdateOne) Save(ctx context.Context) (*Orderonline, error) {

	var (
		err  error
		node *Orderonline
	)
	if len(ouo.hooks) == 0 {
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderonlineMutation)
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
func (ouo *OrderonlineUpdateOne) SaveX(ctx context.Context) *Orderonline {
	o, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return o
}

// Exec executes the query on the entity.
func (ouo *OrderonlineUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderonlineUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouo *OrderonlineUpdateOne) sqlSave(ctx context.Context) (o *Orderonline, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderonline.Table,
			Columns: orderonline.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderonline.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Orderonline.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ouo.mutation.Addedtime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: orderonline.FieldAddedtime,
		})
	}
	if value, ok := ouo.mutation.Stock(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderonline.FieldStock,
		})
	}
	if value, ok := ouo.mutation.AddedStock(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: orderonline.FieldStock,
		})
	}
	if ouo.mutation.ProductonCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderonline.ProductonTable,
			Columns: []string{orderonline.ProductonColumn},
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
	if nodes := ouo.mutation.ProductonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderonline.ProductonTable,
			Columns: []string{orderonline.ProductonColumn},
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
	if ouo.mutation.PaymentchannelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderonline.PaymentchannelTable,
			Columns: []string{orderonline.PaymentchannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: paymentchannel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.PaymentchannelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderonline.PaymentchannelTable,
			Columns: []string{orderonline.PaymentchannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: paymentchannel.FieldID,
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
			Table:   orderonline.TypeproductTable,
			Columns: []string{orderonline.TypeproductColumn},
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
			Table:   orderonline.TypeproductTable,
			Columns: []string{orderonline.TypeproductColumn},
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
	if ouo.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderonline.CustomerTable,
			Columns: []string{orderonline.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderonline.CustomerTable,
			Columns: []string{orderonline.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	o = &Orderonline{config: ouo.config}
	_spec.Assign = o.assignValues
	_spec.ScanValues = o.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderonline.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return o, nil
}