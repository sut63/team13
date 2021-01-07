// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team13/app/ent/customer"
	"github.com/team13/app/ent/orderonline"
	"github.com/team13/app/ent/paymentchannel"
	"github.com/team13/app/ent/predicate"
	"github.com/team13/app/ent/product"
	"github.com/team13/app/ent/typeproduct"
)

// OrderonlineQuery is the builder for querying Orderonline entities.
type OrderonlineQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Orderonline
	// eager-loading edges.
	withProducton      *ProductQuery
	withPaymentchannel *PaymentchannelQuery
	withTypeproduct    *TypeproductQuery
	withCustomer       *CustomerQuery
	withFKs            bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (oq *OrderonlineQuery) Where(ps ...predicate.Orderonline) *OrderonlineQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit adds a limit step to the query.
func (oq *OrderonlineQuery) Limit(limit int) *OrderonlineQuery {
	oq.limit = &limit
	return oq
}

// Offset adds an offset step to the query.
func (oq *OrderonlineQuery) Offset(offset int) *OrderonlineQuery {
	oq.offset = &offset
	return oq
}

// Order adds an order step to the query.
func (oq *OrderonlineQuery) Order(o ...OrderFunc) *OrderonlineQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// QueryProducton chains the current query on the producton edge.
func (oq *OrderonlineQuery) QueryProducton() *ProductQuery {
	query := &ProductQuery{config: oq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderonline.Table, orderonline.FieldID, oq.sqlQuery()),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderonline.ProductonTable, orderonline.ProductonColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPaymentchannel chains the current query on the paymentchannel edge.
func (oq *OrderonlineQuery) QueryPaymentchannel() *PaymentchannelQuery {
	query := &PaymentchannelQuery{config: oq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderonline.Table, orderonline.FieldID, oq.sqlQuery()),
			sqlgraph.To(paymentchannel.Table, paymentchannel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderonline.PaymentchannelTable, orderonline.PaymentchannelColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTypeproduct chains the current query on the Typeproduct edge.
func (oq *OrderonlineQuery) QueryTypeproduct() *TypeproductQuery {
	query := &TypeproductQuery{config: oq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderonline.Table, orderonline.FieldID, oq.sqlQuery()),
			sqlgraph.To(typeproduct.Table, typeproduct.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderonline.TypeproductTable, orderonline.TypeproductColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCustomer chains the current query on the customer edge.
func (oq *OrderonlineQuery) QueryCustomer() *CustomerQuery {
	query := &CustomerQuery{config: oq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderonline.Table, orderonline.FieldID, oq.sqlQuery()),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderonline.CustomerTable, orderonline.CustomerColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Orderonline entity in the query. Returns *NotFoundError when no orderonline was found.
func (oq *OrderonlineQuery) First(ctx context.Context) (*Orderonline, error) {
	os, err := oq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(os) == 0 {
		return nil, &NotFoundError{orderonline.Label}
	}
	return os[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *OrderonlineQuery) FirstX(ctx context.Context) *Orderonline {
	o, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return o
}

// FirstID returns the first Orderonline id in the query. Returns *NotFoundError when no id was found.
func (oq *OrderonlineQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderonline.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (oq *OrderonlineQuery) FirstXID(ctx context.Context) int {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Orderonline entity in the query, returns an error if not exactly one entity was returned.
func (oq *OrderonlineQuery) Only(ctx context.Context) (*Orderonline, error) {
	os, err := oq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(os) {
	case 1:
		return os[0], nil
	case 0:
		return nil, &NotFoundError{orderonline.Label}
	default:
		return nil, &NotSingularError{orderonline.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *OrderonlineQuery) OnlyX(ctx context.Context) *Orderonline {
	o, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return o
}

// OnlyID returns the only Orderonline id in the query, returns an error if not exactly one id was returned.
func (oq *OrderonlineQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderonline.Label}
	default:
		err = &NotSingularError{orderonline.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *OrderonlineQuery) OnlyIDX(ctx context.Context) int {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Orderonlines.
func (oq *OrderonlineQuery) All(ctx context.Context) ([]*Orderonline, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oq *OrderonlineQuery) AllX(ctx context.Context) []*Orderonline {
	os, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return os
}

// IDs executes the query and returns a list of Orderonline ids.
func (oq *OrderonlineQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := oq.Select(orderonline.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *OrderonlineQuery) IDsX(ctx context.Context) []int {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *OrderonlineQuery) Count(ctx context.Context) (int, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oq *OrderonlineQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *OrderonlineQuery) Exist(ctx context.Context) (bool, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *OrderonlineQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *OrderonlineQuery) Clone() *OrderonlineQuery {
	return &OrderonlineQuery{
		config:     oq.config,
		limit:      oq.limit,
		offset:     oq.offset,
		order:      append([]OrderFunc{}, oq.order...),
		unique:     append([]string{}, oq.unique...),
		predicates: append([]predicate.Orderonline{}, oq.predicates...),
		// clone intermediate query.
		sql:  oq.sql.Clone(),
		path: oq.path,
	}
}

//  WithProducton tells the query-builder to eager-loads the nodes that are connected to
// the "producton" edge. The optional arguments used to configure the query builder of the edge.
func (oq *OrderonlineQuery) WithProducton(opts ...func(*ProductQuery)) *OrderonlineQuery {
	query := &ProductQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withProducton = query
	return oq
}

//  WithPaymentchannel tells the query-builder to eager-loads the nodes that are connected to
// the "paymentchannel" edge. The optional arguments used to configure the query builder of the edge.
func (oq *OrderonlineQuery) WithPaymentchannel(opts ...func(*PaymentchannelQuery)) *OrderonlineQuery {
	query := &PaymentchannelQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withPaymentchannel = query
	return oq
}

//  WithTypeproduct tells the query-builder to eager-loads the nodes that are connected to
// the "Typeproduct" edge. The optional arguments used to configure the query builder of the edge.
func (oq *OrderonlineQuery) WithTypeproduct(opts ...func(*TypeproductQuery)) *OrderonlineQuery {
	query := &TypeproductQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withTypeproduct = query
	return oq
}

//  WithCustomer tells the query-builder to eager-loads the nodes that are connected to
// the "customer" edge. The optional arguments used to configure the query builder of the edge.
func (oq *OrderonlineQuery) WithCustomer(opts ...func(*CustomerQuery)) *OrderonlineQuery {
	query := &CustomerQuery{config: oq.config}
	for _, opt := range opts {
		opt(query)
	}
	oq.withCustomer = query
	return oq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Addedtime time.Time `json:"addedtime,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Orderonline.Query().
//		GroupBy(orderonline.FieldAddedtime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (oq *OrderonlineQuery) GroupBy(field string, fields ...string) *OrderonlineGroupBy {
	group := &OrderonlineGroupBy{config: oq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Addedtime time.Time `json:"addedtime,omitempty"`
//	}
//
//	client.Orderonline.Query().
//		Select(orderonline.FieldAddedtime).
//		Scan(ctx, &v)
//
func (oq *OrderonlineQuery) Select(field string, fields ...string) *OrderonlineSelect {
	selector := &OrderonlineSelect{config: oq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oq.sqlQuery(), nil
	}
	return selector
}

func (oq *OrderonlineQuery) prepareQuery(ctx context.Context) error {
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	return nil
}

func (oq *OrderonlineQuery) sqlAll(ctx context.Context) ([]*Orderonline, error) {
	var (
		nodes       = []*Orderonline{}
		withFKs     = oq.withFKs
		_spec       = oq.querySpec()
		loadedTypes = [4]bool{
			oq.withProducton != nil,
			oq.withPaymentchannel != nil,
			oq.withTypeproduct != nil,
			oq.withCustomer != nil,
		}
	)
	if oq.withProducton != nil || oq.withPaymentchannel != nil || oq.withTypeproduct != nil || oq.withCustomer != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, orderonline.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Orderonline{config: oq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := oq.withProducton; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Orderonline)
		for i := range nodes {
			if fk := nodes[i].product_formproductonline; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(product.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "product_formproductonline" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Producton = n
			}
		}
	}

	if query := oq.withPaymentchannel; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Orderonline)
		for i := range nodes {
			if fk := nodes[i].paymentchannel_formpaymentchannel; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(paymentchannel.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "paymentchannel_formpaymentchannel" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Paymentchannel = n
			}
		}
	}

	if query := oq.withTypeproduct; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Orderonline)
		for i := range nodes {
			if fk := nodes[i].typeproduct_from_typeproductonline; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(typeproduct.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "typeproduct_from_typeproductonline" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Typeproduct = n
			}
		}
	}

	if query := oq.withCustomer; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Orderonline)
		for i := range nodes {
			if fk := nodes[i].customer_formcustomer; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(customer.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "customer_formcustomer" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Customer = n
			}
		}
	}

	return nodes, nil
}

func (oq *OrderonlineQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *OrderonlineQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (oq *OrderonlineQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderonline.Table,
			Columns: orderonline.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: orderonline.FieldID,
			},
		},
		From:   oq.sql,
		Unique: true,
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *OrderonlineQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(orderonline.Table)
	selector := builder.Select(t1.Columns(orderonline.Columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(orderonline.Columns...)...)
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderonlineGroupBy is the builder for group-by Orderonline entities.
type OrderonlineGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *OrderonlineGroupBy) Aggregate(fns ...AggregateFunc) *OrderonlineGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the group-by query and scan the result into the given value.
func (ogb *OrderonlineGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ogb.path(ctx)
	if err != nil {
		return err
	}
	ogb.sql = query
	return ogb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ogb *OrderonlineGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ogb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ogb *OrderonlineGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OrderonlineGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ogb *OrderonlineGroupBy) StringsX(ctx context.Context) []string {
	v, err := ogb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (ogb *OrderonlineGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ogb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderonline.Label}
	default:
		err = fmt.Errorf("ent: OrderonlineGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ogb *OrderonlineGroupBy) StringX(ctx context.Context) string {
	v, err := ogb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ogb *OrderonlineGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OrderonlineGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ogb *OrderonlineGroupBy) IntsX(ctx context.Context) []int {
	v, err := ogb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (ogb *OrderonlineGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ogb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderonline.Label}
	default:
		err = fmt.Errorf("ent: OrderonlineGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ogb *OrderonlineGroupBy) IntX(ctx context.Context) int {
	v, err := ogb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ogb *OrderonlineGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OrderonlineGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ogb *OrderonlineGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ogb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (ogb *OrderonlineGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ogb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderonline.Label}
	default:
		err = fmt.Errorf("ent: OrderonlineGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ogb *OrderonlineGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ogb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ogb *OrderonlineGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: OrderonlineGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ogb *OrderonlineGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ogb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (ogb *OrderonlineGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ogb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderonline.Label}
	default:
		err = fmt.Errorf("ent: OrderonlineGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ogb *OrderonlineGroupBy) BoolX(ctx context.Context) bool {
	v, err := ogb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ogb *OrderonlineGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ogb.sqlQuery().Query()
	if err := ogb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ogb *OrderonlineGroupBy) sqlQuery() *sql.Selector {
	selector := ogb.sql
	columns := make([]string, 0, len(ogb.fields)+len(ogb.fns))
	columns = append(columns, ogb.fields...)
	for _, fn := range ogb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(ogb.fields...)
}

// OrderonlineSelect is the builder for select fields of Orderonline entities.
type OrderonlineSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (os *OrderonlineSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := os.path(ctx)
	if err != nil {
		return err
	}
	os.sql = query
	return os.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (os *OrderonlineSelect) ScanX(ctx context.Context, v interface{}) {
	if err := os.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (os *OrderonlineSelect) Strings(ctx context.Context) ([]string, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OrderonlineSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (os *OrderonlineSelect) StringsX(ctx context.Context) []string {
	v, err := os.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (os *OrderonlineSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = os.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderonline.Label}
	default:
		err = fmt.Errorf("ent: OrderonlineSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (os *OrderonlineSelect) StringX(ctx context.Context) string {
	v, err := os.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (os *OrderonlineSelect) Ints(ctx context.Context) ([]int, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OrderonlineSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (os *OrderonlineSelect) IntsX(ctx context.Context) []int {
	v, err := os.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (os *OrderonlineSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = os.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderonline.Label}
	default:
		err = fmt.Errorf("ent: OrderonlineSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (os *OrderonlineSelect) IntX(ctx context.Context) int {
	v, err := os.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (os *OrderonlineSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OrderonlineSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (os *OrderonlineSelect) Float64sX(ctx context.Context) []float64 {
	v, err := os.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (os *OrderonlineSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = os.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderonline.Label}
	default:
		err = fmt.Errorf("ent: OrderonlineSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (os *OrderonlineSelect) Float64X(ctx context.Context) float64 {
	v, err := os.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (os *OrderonlineSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: OrderonlineSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (os *OrderonlineSelect) BoolsX(ctx context.Context) []bool {
	v, err := os.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (os *OrderonlineSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = os.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{orderonline.Label}
	default:
		err = fmt.Errorf("ent: OrderonlineSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (os *OrderonlineSelect) BoolX(ctx context.Context) bool {
	v, err := os.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (os *OrderonlineSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := os.sqlQuery().Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (os *OrderonlineSelect) sqlQuery() sql.Querier {
	selector := os.sql
	selector.Select(selector.Columns(os.fields...)...)
	return selector
}
