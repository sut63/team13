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
	"github.com/team13/app/ent/predicate"
	"github.com/team13/app/ent/startwork"
)

// StartWorkQuery is the builder for querying StartWork entities.
type StartWorkQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.StartWork
	// eager-loading edges.
	withWhenwork *EmployeeWorkingHoursQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (swq *StartWorkQuery) Where(ps ...predicate.StartWork) *StartWorkQuery {
	swq.predicates = append(swq.predicates, ps...)
	return swq
}

// Limit adds a limit step to the query.
func (swq *StartWorkQuery) Limit(limit int) *StartWorkQuery {
	swq.limit = &limit
	return swq
}

// Offset adds an offset step to the query.
func (swq *StartWorkQuery) Offset(offset int) *StartWorkQuery {
	swq.offset = &offset
	return swq
}

// Order adds an order step to the query.
func (swq *StartWorkQuery) Order(o ...OrderFunc) *StartWorkQuery {
	swq.order = append(swq.order, o...)
	return swq
}

// QueryWhenwork chains the current query on the whenwork edge.
func (swq *StartWorkQuery) QueryWhenwork() *EmployeeWorkingHoursQuery {
	query := &EmployeeWorkingHoursQuery{config: swq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := swq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(startwork.Table, startwork.FieldID, swq.sqlQuery()),
			sqlgraph.To(employeeworkinghours.Table, employeeworkinghours.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, startwork.WhenworkTable, startwork.WhenworkColumn),
		)
		fromU = sqlgraph.SetNeighbors(swq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first StartWork entity in the query. Returns *NotFoundError when no startwork was found.
func (swq *StartWorkQuery) First(ctx context.Context) (*StartWork, error) {
	sws, err := swq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(sws) == 0 {
		return nil, &NotFoundError{startwork.Label}
	}
	return sws[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (swq *StartWorkQuery) FirstX(ctx context.Context) *StartWork {
	sw, err := swq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return sw
}

// FirstID returns the first StartWork id in the query. Returns *NotFoundError when no id was found.
func (swq *StartWorkQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = swq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{startwork.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (swq *StartWorkQuery) FirstXID(ctx context.Context) int {
	id, err := swq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only StartWork entity in the query, returns an error if not exactly one entity was returned.
func (swq *StartWorkQuery) Only(ctx context.Context) (*StartWork, error) {
	sws, err := swq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(sws) {
	case 1:
		return sws[0], nil
	case 0:
		return nil, &NotFoundError{startwork.Label}
	default:
		return nil, &NotSingularError{startwork.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (swq *StartWorkQuery) OnlyX(ctx context.Context) *StartWork {
	sw, err := swq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return sw
}

// OnlyID returns the only StartWork id in the query, returns an error if not exactly one id was returned.
func (swq *StartWorkQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = swq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{startwork.Label}
	default:
		err = &NotSingularError{startwork.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (swq *StartWorkQuery) OnlyIDX(ctx context.Context) int {
	id, err := swq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StartWorks.
func (swq *StartWorkQuery) All(ctx context.Context) ([]*StartWork, error) {
	if err := swq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return swq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (swq *StartWorkQuery) AllX(ctx context.Context) []*StartWork {
	sws, err := swq.All(ctx)
	if err != nil {
		panic(err)
	}
	return sws
}

// IDs executes the query and returns a list of StartWork ids.
func (swq *StartWorkQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := swq.Select(startwork.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (swq *StartWorkQuery) IDsX(ctx context.Context) []int {
	ids, err := swq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (swq *StartWorkQuery) Count(ctx context.Context) (int, error) {
	if err := swq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return swq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (swq *StartWorkQuery) CountX(ctx context.Context) int {
	count, err := swq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (swq *StartWorkQuery) Exist(ctx context.Context) (bool, error) {
	if err := swq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return swq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (swq *StartWorkQuery) ExistX(ctx context.Context) bool {
	exist, err := swq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (swq *StartWorkQuery) Clone() *StartWorkQuery {
	return &StartWorkQuery{
		config:     swq.config,
		limit:      swq.limit,
		offset:     swq.offset,
		order:      append([]OrderFunc{}, swq.order...),
		unique:     append([]string{}, swq.unique...),
		predicates: append([]predicate.StartWork{}, swq.predicates...),
		// clone intermediate query.
		sql:  swq.sql.Clone(),
		path: swq.path,
	}
}

//  WithWhenwork tells the query-builder to eager-loads the nodes that are connected to
// the "whenwork" edge. The optional arguments used to configure the query builder of the edge.
func (swq *StartWorkQuery) WithWhenwork(opts ...func(*EmployeeWorkingHoursQuery)) *StartWorkQuery {
	query := &EmployeeWorkingHoursQuery{config: swq.config}
	for _, opt := range opts {
		opt(query)
	}
	swq.withWhenwork = query
	return swq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		StartWork time.Time `json:"StartWork,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.StartWork.Query().
//		GroupBy(startwork.FieldStartWork).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (swq *StartWorkQuery) GroupBy(field string, fields ...string) *StartWorkGroupBy {
	group := &StartWorkGroupBy{config: swq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := swq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return swq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		StartWork time.Time `json:"StartWork,omitempty"`
//	}
//
//	client.StartWork.Query().
//		Select(startwork.FieldStartWork).
//		Scan(ctx, &v)
//
func (swq *StartWorkQuery) Select(field string, fields ...string) *StartWorkSelect {
	selector := &StartWorkSelect{config: swq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := swq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return swq.sqlQuery(), nil
	}
	return selector
}

func (swq *StartWorkQuery) prepareQuery(ctx context.Context) error {
	if swq.path != nil {
		prev, err := swq.path(ctx)
		if err != nil {
			return err
		}
		swq.sql = prev
	}
	return nil
}

func (swq *StartWorkQuery) sqlAll(ctx context.Context) ([]*StartWork, error) {
	var (
		nodes       = []*StartWork{}
		_spec       = swq.querySpec()
		loadedTypes = [1]bool{
			swq.withWhenwork != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &StartWork{config: swq.config}
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
	if err := sqlgraph.QueryNodes(ctx, swq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := swq.withWhenwork; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*StartWork)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.EmployeeWorkingHours(func(s *sql.Selector) {
			s.Where(sql.InValues(startwork.WhenworkColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.start_work_whenwork
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "start_work_whenwork" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "start_work_whenwork" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Whenwork = append(node.Edges.Whenwork, n)
		}
	}

	return nodes, nil
}

func (swq *StartWorkQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := swq.querySpec()
	return sqlgraph.CountNodes(ctx, swq.driver, _spec)
}

func (swq *StartWorkQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := swq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (swq *StartWorkQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   startwork.Table,
			Columns: startwork.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: startwork.FieldID,
			},
		},
		From:   swq.sql,
		Unique: true,
	}
	if ps := swq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := swq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := swq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := swq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (swq *StartWorkQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(swq.driver.Dialect())
	t1 := builder.Table(startwork.Table)
	selector := builder.Select(t1.Columns(startwork.Columns...)...).From(t1)
	if swq.sql != nil {
		selector = swq.sql
		selector.Select(selector.Columns(startwork.Columns...)...)
	}
	for _, p := range swq.predicates {
		p(selector)
	}
	for _, p := range swq.order {
		p(selector)
	}
	if offset := swq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := swq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StartWorkGroupBy is the builder for group-by StartWork entities.
type StartWorkGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (swgb *StartWorkGroupBy) Aggregate(fns ...AggregateFunc) *StartWorkGroupBy {
	swgb.fns = append(swgb.fns, fns...)
	return swgb
}

// Scan applies the group-by query and scan the result into the given value.
func (swgb *StartWorkGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := swgb.path(ctx)
	if err != nil {
		return err
	}
	swgb.sql = query
	return swgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (swgb *StartWorkGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := swgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (swgb *StartWorkGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(swgb.fields) > 1 {
		return nil, errors.New("ent: StartWorkGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := swgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (swgb *StartWorkGroupBy) StringsX(ctx context.Context) []string {
	v, err := swgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (swgb *StartWorkGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = swgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{startwork.Label}
	default:
		err = fmt.Errorf("ent: StartWorkGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (swgb *StartWorkGroupBy) StringX(ctx context.Context) string {
	v, err := swgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (swgb *StartWorkGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(swgb.fields) > 1 {
		return nil, errors.New("ent: StartWorkGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := swgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (swgb *StartWorkGroupBy) IntsX(ctx context.Context) []int {
	v, err := swgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (swgb *StartWorkGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = swgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{startwork.Label}
	default:
		err = fmt.Errorf("ent: StartWorkGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (swgb *StartWorkGroupBy) IntX(ctx context.Context) int {
	v, err := swgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (swgb *StartWorkGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(swgb.fields) > 1 {
		return nil, errors.New("ent: StartWorkGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := swgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (swgb *StartWorkGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := swgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (swgb *StartWorkGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = swgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{startwork.Label}
	default:
		err = fmt.Errorf("ent: StartWorkGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (swgb *StartWorkGroupBy) Float64X(ctx context.Context) float64 {
	v, err := swgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (swgb *StartWorkGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(swgb.fields) > 1 {
		return nil, errors.New("ent: StartWorkGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := swgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (swgb *StartWorkGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := swgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (swgb *StartWorkGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = swgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{startwork.Label}
	default:
		err = fmt.Errorf("ent: StartWorkGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (swgb *StartWorkGroupBy) BoolX(ctx context.Context) bool {
	v, err := swgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (swgb *StartWorkGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := swgb.sqlQuery().Query()
	if err := swgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (swgb *StartWorkGroupBy) sqlQuery() *sql.Selector {
	selector := swgb.sql
	columns := make([]string, 0, len(swgb.fields)+len(swgb.fns))
	columns = append(columns, swgb.fields...)
	for _, fn := range swgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(swgb.fields...)
}

// StartWorkSelect is the builder for select fields of StartWork entities.
type StartWorkSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (sws *StartWorkSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := sws.path(ctx)
	if err != nil {
		return err
	}
	sws.sql = query
	return sws.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sws *StartWorkSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sws.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (sws *StartWorkSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sws.fields) > 1 {
		return nil, errors.New("ent: StartWorkSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sws *StartWorkSelect) StringsX(ctx context.Context) []string {
	v, err := sws.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (sws *StartWorkSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sws.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{startwork.Label}
	default:
		err = fmt.Errorf("ent: StartWorkSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sws *StartWorkSelect) StringX(ctx context.Context) string {
	v, err := sws.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (sws *StartWorkSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sws.fields) > 1 {
		return nil, errors.New("ent: StartWorkSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sws *StartWorkSelect) IntsX(ctx context.Context) []int {
	v, err := sws.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (sws *StartWorkSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sws.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{startwork.Label}
	default:
		err = fmt.Errorf("ent: StartWorkSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sws *StartWorkSelect) IntX(ctx context.Context) int {
	v, err := sws.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (sws *StartWorkSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sws.fields) > 1 {
		return nil, errors.New("ent: StartWorkSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sws *StartWorkSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sws.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (sws *StartWorkSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sws.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{startwork.Label}
	default:
		err = fmt.Errorf("ent: StartWorkSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sws *StartWorkSelect) Float64X(ctx context.Context) float64 {
	v, err := sws.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (sws *StartWorkSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sws.fields) > 1 {
		return nil, errors.New("ent: StartWorkSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sws *StartWorkSelect) BoolsX(ctx context.Context) []bool {
	v, err := sws.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (sws *StartWorkSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sws.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{startwork.Label}
	default:
		err = fmt.Errorf("ent: StartWorkSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sws *StartWorkSelect) BoolX(ctx context.Context) bool {
	v, err := sws.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sws *StartWorkSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sws.sqlQuery().Query()
	if err := sws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sws *StartWorkSelect) sqlQuery() sql.Querier {
	selector := sws.sql
	selector.Select(selector.Columns(sws.fields...)...)
	return selector
}