// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team13/app/ent/manager"
	"github.com/team13/app/ent/orderproduct"
	"github.com/team13/app/ent/predicate"
)

// ManagerQuery is the builder for querying Manager entities.
type ManagerQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Manager
	// eager-loading edges.
	withManagers *OrderproductQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (mq *ManagerQuery) Where(ps ...predicate.Manager) *ManagerQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit adds a limit step to the query.
func (mq *ManagerQuery) Limit(limit int) *ManagerQuery {
	mq.limit = &limit
	return mq
}

// Offset adds an offset step to the query.
func (mq *ManagerQuery) Offset(offset int) *ManagerQuery {
	mq.offset = &offset
	return mq
}

// Order adds an order step to the query.
func (mq *ManagerQuery) Order(o ...OrderFunc) *ManagerQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryManagers chains the current query on the managers edge.
func (mq *ManagerQuery) QueryManagers() *OrderproductQuery {
	query := &OrderproductQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(manager.Table, manager.FieldID, mq.sqlQuery()),
			sqlgraph.To(orderproduct.Table, orderproduct.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, manager.ManagersTable, manager.ManagersColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Manager entity in the query. Returns *NotFoundError when no manager was found.
func (mq *ManagerQuery) First(ctx context.Context) (*Manager, error) {
	ms, err := mq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ms) == 0 {
		return nil, &NotFoundError{manager.Label}
	}
	return ms[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *ManagerQuery) FirstX(ctx context.Context) *Manager {
	m, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return m
}

// FirstID returns the first Manager id in the query. Returns *NotFoundError when no id was found.
func (mq *ManagerQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{manager.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (mq *ManagerQuery) FirstXID(ctx context.Context) int {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Manager entity in the query, returns an error if not exactly one entity was returned.
func (mq *ManagerQuery) Only(ctx context.Context) (*Manager, error) {
	ms, err := mq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ms) {
	case 1:
		return ms[0], nil
	case 0:
		return nil, &NotFoundError{manager.Label}
	default:
		return nil, &NotSingularError{manager.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *ManagerQuery) OnlyX(ctx context.Context) *Manager {
	m, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return m
}

// OnlyID returns the only Manager id in the query, returns an error if not exactly one id was returned.
func (mq *ManagerQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{manager.Label}
	default:
		err = &NotSingularError{manager.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *ManagerQuery) OnlyIDX(ctx context.Context) int {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Managers.
func (mq *ManagerQuery) All(ctx context.Context) ([]*Manager, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mq *ManagerQuery) AllX(ctx context.Context) []*Manager {
	ms, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ms
}

// IDs executes the query and returns a list of Manager ids.
func (mq *ManagerQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mq.Select(manager.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *ManagerQuery) IDsX(ctx context.Context) []int {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *ManagerQuery) Count(ctx context.Context) (int, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mq *ManagerQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *ManagerQuery) Exist(ctx context.Context) (bool, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *ManagerQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *ManagerQuery) Clone() *ManagerQuery {
	return &ManagerQuery{
		config:     mq.config,
		limit:      mq.limit,
		offset:     mq.offset,
		order:      append([]OrderFunc{}, mq.order...),
		unique:     append([]string{}, mq.unique...),
		predicates: append([]predicate.Manager{}, mq.predicates...),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

//  WithManagers tells the query-builder to eager-loads the nodes that are connected to
// the "managers" edge. The optional arguments used to configure the query builder of the edge.
func (mq *ManagerQuery) WithManagers(opts ...func(*OrderproductQuery)) *ManagerQuery {
	query := &OrderproductQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withManagers = query
	return mq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Manager.Query().
//		GroupBy(manager.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mq *ManagerQuery) GroupBy(field string, fields ...string) *ManagerGroupBy {
	group := &ManagerGroupBy{config: mq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Manager.Query().
//		Select(manager.FieldName).
//		Scan(ctx, &v)
//
func (mq *ManagerQuery) Select(field string, fields ...string) *ManagerSelect {
	selector := &ManagerSelect{config: mq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mq.sqlQuery(), nil
	}
	return selector
}

func (mq *ManagerQuery) prepareQuery(ctx context.Context) error {
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *ManagerQuery) sqlAll(ctx context.Context) ([]*Manager, error) {
	var (
		nodes       = []*Manager{}
		_spec       = mq.querySpec()
		loadedTypes = [1]bool{
			mq.withManagers != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Manager{config: mq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
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
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := mq.withManagers; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Manager)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Orderproduct(func(s *sql.Selector) {
			s.Where(sql.InValues(manager.ManagersColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.manager_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "manager_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "manager_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Managers = append(node.Edges.Managers, n)
		}
	}

	return nodes, nil
}

func (mq *ManagerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *ManagerQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (mq *ManagerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   manager.Table,
			Columns: manager.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: manager.FieldID,
			},
		},
		From:   mq.sql,
		Unique: true,
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *ManagerQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(manager.Table)
	selector := builder.Select(t1.Columns(manager.Columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(manager.Columns...)...)
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ManagerGroupBy is the builder for group-by Manager entities.
type ManagerGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *ManagerGroupBy) Aggregate(fns ...AggregateFunc) *ManagerGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the group-by query and scan the result into the given value.
func (mgb *ManagerGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mgb.path(ctx)
	if err != nil {
		return err
	}
	mgb.sql = query
	return mgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mgb *ManagerGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (mgb *ManagerGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: ManagerGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mgb *ManagerGroupBy) StringsX(ctx context.Context) []string {
	v, err := mgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (mgb *ManagerGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{manager.Label}
	default:
		err = fmt.Errorf("ent: ManagerGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mgb *ManagerGroupBy) StringX(ctx context.Context) string {
	v, err := mgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (mgb *ManagerGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: ManagerGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mgb *ManagerGroupBy) IntsX(ctx context.Context) []int {
	v, err := mgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (mgb *ManagerGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{manager.Label}
	default:
		err = fmt.Errorf("ent: ManagerGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mgb *ManagerGroupBy) IntX(ctx context.Context) int {
	v, err := mgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (mgb *ManagerGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: ManagerGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mgb *ManagerGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (mgb *ManagerGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{manager.Label}
	default:
		err = fmt.Errorf("ent: ManagerGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mgb *ManagerGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (mgb *ManagerGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: ManagerGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mgb *ManagerGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (mgb *ManagerGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{manager.Label}
	default:
		err = fmt.Errorf("ent: ManagerGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mgb *ManagerGroupBy) BoolX(ctx context.Context) bool {
	v, err := mgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mgb *ManagerGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mgb.sqlQuery().Query()
	if err := mgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mgb *ManagerGroupBy) sqlQuery() *sql.Selector {
	selector := mgb.sql
	columns := make([]string, 0, len(mgb.fields)+len(mgb.fns))
	columns = append(columns, mgb.fields...)
	for _, fn := range mgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(mgb.fields...)
}

// ManagerSelect is the builder for select fields of Manager entities.
type ManagerSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ms *ManagerSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ms.path(ctx)
	if err != nil {
		return err
	}
	ms.sql = query
	return ms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ms *ManagerSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ms *ManagerSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: ManagerSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ms *ManagerSelect) StringsX(ctx context.Context) []string {
	v, err := ms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ms *ManagerSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{manager.Label}
	default:
		err = fmt.Errorf("ent: ManagerSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ms *ManagerSelect) StringX(ctx context.Context) string {
	v, err := ms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ms *ManagerSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: ManagerSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ms *ManagerSelect) IntsX(ctx context.Context) []int {
	v, err := ms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ms *ManagerSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{manager.Label}
	default:
		err = fmt.Errorf("ent: ManagerSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ms *ManagerSelect) IntX(ctx context.Context) int {
	v, err := ms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ms *ManagerSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: ManagerSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ms *ManagerSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ms *ManagerSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{manager.Label}
	default:
		err = fmt.Errorf("ent: ManagerSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ms *ManagerSelect) Float64X(ctx context.Context) float64 {
	v, err := ms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ms *ManagerSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: ManagerSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ms *ManagerSelect) BoolsX(ctx context.Context) []bool {
	v, err := ms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ms *ManagerSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{manager.Label}
	default:
		err = fmt.Errorf("ent: ManagerSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ms *ManagerSelect) BoolX(ctx context.Context) bool {
	v, err := ms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ms *ManagerSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ms.sqlQuery().Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ms *ManagerSelect) sqlQuery() sql.Querier {
	selector := ms.sql
	selector.Select(selector.Columns(ms.fields...)...)
	return selector
}
