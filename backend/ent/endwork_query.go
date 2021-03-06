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
	"github.com/team13/app/ent/employeeworkinghours"
	"github.com/team13/app/ent/endwork"
	"github.com/team13/app/ent/predicate"
)

// EndWorkQuery is the builder for querying EndWork entities.
type EndWorkQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.EndWork
	// eager-loading edges.
	withWhenendwork *EmployeeWorkingHoursQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (ewq *EndWorkQuery) Where(ps ...predicate.EndWork) *EndWorkQuery {
	ewq.predicates = append(ewq.predicates, ps...)
	return ewq
}

// Limit adds a limit step to the query.
func (ewq *EndWorkQuery) Limit(limit int) *EndWorkQuery {
	ewq.limit = &limit
	return ewq
}

// Offset adds an offset step to the query.
func (ewq *EndWorkQuery) Offset(offset int) *EndWorkQuery {
	ewq.offset = &offset
	return ewq
}

// Order adds an order step to the query.
func (ewq *EndWorkQuery) Order(o ...OrderFunc) *EndWorkQuery {
	ewq.order = append(ewq.order, o...)
	return ewq
}

// QueryWhenendwork chains the current query on the whenendwork edge.
func (ewq *EndWorkQuery) QueryWhenendwork() *EmployeeWorkingHoursQuery {
	query := &EmployeeWorkingHoursQuery{config: ewq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ewq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(endwork.Table, endwork.FieldID, ewq.sqlQuery()),
			sqlgraph.To(employeeworkinghours.Table, employeeworkinghours.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, endwork.WhenendworkTable, endwork.WhenendworkColumn),
		)
		fromU = sqlgraph.SetNeighbors(ewq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EndWork entity in the query. Returns *NotFoundError when no endwork was found.
func (ewq *EndWorkQuery) First(ctx context.Context) (*EndWork, error) {
	ews, err := ewq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ews) == 0 {
		return nil, &NotFoundError{endwork.Label}
	}
	return ews[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ewq *EndWorkQuery) FirstX(ctx context.Context) *EndWork {
	ew, err := ewq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return ew
}

// FirstID returns the first EndWork id in the query. Returns *NotFoundError when no id was found.
func (ewq *EndWorkQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ewq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{endwork.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (ewq *EndWorkQuery) FirstXID(ctx context.Context) int {
	id, err := ewq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only EndWork entity in the query, returns an error if not exactly one entity was returned.
func (ewq *EndWorkQuery) Only(ctx context.Context) (*EndWork, error) {
	ews, err := ewq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ews) {
	case 1:
		return ews[0], nil
	case 0:
		return nil, &NotFoundError{endwork.Label}
	default:
		return nil, &NotSingularError{endwork.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ewq *EndWorkQuery) OnlyX(ctx context.Context) *EndWork {
	ew, err := ewq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return ew
}

// OnlyID returns the only EndWork id in the query, returns an error if not exactly one id was returned.
func (ewq *EndWorkQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ewq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{endwork.Label}
	default:
		err = &NotSingularError{endwork.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ewq *EndWorkQuery) OnlyIDX(ctx context.Context) int {
	id, err := ewq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EndWorks.
func (ewq *EndWorkQuery) All(ctx context.Context) ([]*EndWork, error) {
	if err := ewq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ewq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ewq *EndWorkQuery) AllX(ctx context.Context) []*EndWork {
	ews, err := ewq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ews
}

// IDs executes the query and returns a list of EndWork ids.
func (ewq *EndWorkQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ewq.Select(endwork.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ewq *EndWorkQuery) IDsX(ctx context.Context) []int {
	ids, err := ewq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ewq *EndWorkQuery) Count(ctx context.Context) (int, error) {
	if err := ewq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ewq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ewq *EndWorkQuery) CountX(ctx context.Context) int {
	count, err := ewq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ewq *EndWorkQuery) Exist(ctx context.Context) (bool, error) {
	if err := ewq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ewq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ewq *EndWorkQuery) ExistX(ctx context.Context) bool {
	exist, err := ewq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ewq *EndWorkQuery) Clone() *EndWorkQuery {
	return &EndWorkQuery{
		config:     ewq.config,
		limit:      ewq.limit,
		offset:     ewq.offset,
		order:      append([]OrderFunc{}, ewq.order...),
		unique:     append([]string{}, ewq.unique...),
		predicates: append([]predicate.EndWork{}, ewq.predicates...),
		// clone intermediate query.
		sql:  ewq.sql.Clone(),
		path: ewq.path,
	}
}

//  WithWhenendwork tells the query-builder to eager-loads the nodes that are connected to
// the "whenendwork" edge. The optional arguments used to configure the query builder of the edge.
func (ewq *EndWorkQuery) WithWhenendwork(opts ...func(*EmployeeWorkingHoursQuery)) *EndWorkQuery {
	query := &EmployeeWorkingHoursQuery{config: ewq.config}
	for _, opt := range opts {
		opt(query)
	}
	ewq.withWhenendwork = query
	return ewq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EndWork time.Time `json:"EndWork,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EndWork.Query().
//		GroupBy(endwork.FieldEndWork).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ewq *EndWorkQuery) GroupBy(field string, fields ...string) *EndWorkGroupBy {
	group := &EndWorkGroupBy{config: ewq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ewq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ewq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		EndWork time.Time `json:"EndWork,omitempty"`
//	}
//
//	client.EndWork.Query().
//		Select(endwork.FieldEndWork).
//		Scan(ctx, &v)
//
func (ewq *EndWorkQuery) Select(field string, fields ...string) *EndWorkSelect {
	selector := &EndWorkSelect{config: ewq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ewq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ewq.sqlQuery(), nil
	}
	return selector
}

func (ewq *EndWorkQuery) prepareQuery(ctx context.Context) error {
	if ewq.path != nil {
		prev, err := ewq.path(ctx)
		if err != nil {
			return err
		}
		ewq.sql = prev
	}
	return nil
}

func (ewq *EndWorkQuery) sqlAll(ctx context.Context) ([]*EndWork, error) {
	var (
		nodes       = []*EndWork{}
		_spec       = ewq.querySpec()
		loadedTypes = [1]bool{
			ewq.withWhenendwork != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &EndWork{config: ewq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ewq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ewq.withWhenendwork; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*EndWork)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.EmployeeWorkingHours(func(s *sql.Selector) {
			s.Where(sql.InValues(endwork.WhenendworkColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.end_work_whenendwork
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "end_work_whenendwork" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "end_work_whenendwork" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Whenendwork = append(node.Edges.Whenendwork, n)
		}
	}

	return nodes, nil
}

func (ewq *EndWorkQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ewq.querySpec()
	return sqlgraph.CountNodes(ctx, ewq.driver, _spec)
}

func (ewq *EndWorkQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ewq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (ewq *EndWorkQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   endwork.Table,
			Columns: endwork.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: endwork.FieldID,
			},
		},
		From:   ewq.sql,
		Unique: true,
	}
	if ps := ewq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ewq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ewq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ewq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ewq *EndWorkQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(ewq.driver.Dialect())
	t1 := builder.Table(endwork.Table)
	selector := builder.Select(t1.Columns(endwork.Columns...)...).From(t1)
	if ewq.sql != nil {
		selector = ewq.sql
		selector.Select(selector.Columns(endwork.Columns...)...)
	}
	for _, p := range ewq.predicates {
		p(selector)
	}
	for _, p := range ewq.order {
		p(selector)
	}
	if offset := ewq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ewq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EndWorkGroupBy is the builder for group-by EndWork entities.
type EndWorkGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ewgb *EndWorkGroupBy) Aggregate(fns ...AggregateFunc) *EndWorkGroupBy {
	ewgb.fns = append(ewgb.fns, fns...)
	return ewgb
}

// Scan applies the group-by query and scan the result into the given value.
func (ewgb *EndWorkGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ewgb.path(ctx)
	if err != nil {
		return err
	}
	ewgb.sql = query
	return ewgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ewgb *EndWorkGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ewgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ewgb *EndWorkGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ewgb.fields) > 1 {
		return nil, errors.New("ent: EndWorkGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ewgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ewgb *EndWorkGroupBy) StringsX(ctx context.Context) []string {
	v, err := ewgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (ewgb *EndWorkGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ewgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{endwork.Label}
	default:
		err = fmt.Errorf("ent: EndWorkGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ewgb *EndWorkGroupBy) StringX(ctx context.Context) string {
	v, err := ewgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ewgb *EndWorkGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ewgb.fields) > 1 {
		return nil, errors.New("ent: EndWorkGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ewgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ewgb *EndWorkGroupBy) IntsX(ctx context.Context) []int {
	v, err := ewgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (ewgb *EndWorkGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ewgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{endwork.Label}
	default:
		err = fmt.Errorf("ent: EndWorkGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ewgb *EndWorkGroupBy) IntX(ctx context.Context) int {
	v, err := ewgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ewgb *EndWorkGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ewgb.fields) > 1 {
		return nil, errors.New("ent: EndWorkGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ewgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ewgb *EndWorkGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ewgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (ewgb *EndWorkGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ewgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{endwork.Label}
	default:
		err = fmt.Errorf("ent: EndWorkGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ewgb *EndWorkGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ewgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ewgb *EndWorkGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ewgb.fields) > 1 {
		return nil, errors.New("ent: EndWorkGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ewgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ewgb *EndWorkGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ewgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (ewgb *EndWorkGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ewgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{endwork.Label}
	default:
		err = fmt.Errorf("ent: EndWorkGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ewgb *EndWorkGroupBy) BoolX(ctx context.Context) bool {
	v, err := ewgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ewgb *EndWorkGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ewgb.sqlQuery().Query()
	if err := ewgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ewgb *EndWorkGroupBy) sqlQuery() *sql.Selector {
	selector := ewgb.sql
	columns := make([]string, 0, len(ewgb.fields)+len(ewgb.fns))
	columns = append(columns, ewgb.fields...)
	for _, fn := range ewgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(ewgb.fields...)
}

// EndWorkSelect is the builder for select fields of EndWork entities.
type EndWorkSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ews *EndWorkSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ews.path(ctx)
	if err != nil {
		return err
	}
	ews.sql = query
	return ews.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ews *EndWorkSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ews.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ews *EndWorkSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ews.fields) > 1 {
		return nil, errors.New("ent: EndWorkSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ews.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ews *EndWorkSelect) StringsX(ctx context.Context) []string {
	v, err := ews.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ews *EndWorkSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ews.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{endwork.Label}
	default:
		err = fmt.Errorf("ent: EndWorkSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ews *EndWorkSelect) StringX(ctx context.Context) string {
	v, err := ews.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ews *EndWorkSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ews.fields) > 1 {
		return nil, errors.New("ent: EndWorkSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ews.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ews *EndWorkSelect) IntsX(ctx context.Context) []int {
	v, err := ews.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ews *EndWorkSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ews.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{endwork.Label}
	default:
		err = fmt.Errorf("ent: EndWorkSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ews *EndWorkSelect) IntX(ctx context.Context) int {
	v, err := ews.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ews *EndWorkSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ews.fields) > 1 {
		return nil, errors.New("ent: EndWorkSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ews.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ews *EndWorkSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ews.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ews *EndWorkSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ews.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{endwork.Label}
	default:
		err = fmt.Errorf("ent: EndWorkSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ews *EndWorkSelect) Float64X(ctx context.Context) float64 {
	v, err := ews.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ews *EndWorkSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ews.fields) > 1 {
		return nil, errors.New("ent: EndWorkSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ews.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ews *EndWorkSelect) BoolsX(ctx context.Context) []bool {
	v, err := ews.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ews *EndWorkSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ews.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{endwork.Label}
	default:
		err = fmt.Errorf("ent: EndWorkSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ews *EndWorkSelect) BoolX(ctx context.Context) bool {
	v, err := ews.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ews *EndWorkSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ews.sqlQuery().Query()
	if err := ews.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ews *EndWorkSelect) sqlQuery() sql.Querier {
	selector := ews.sql
	selector.Select(selector.Columns(ews.fields...)...)
	return selector
}
