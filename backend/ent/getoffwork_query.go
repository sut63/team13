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
	"github.com/team13/app/ent/getoffwork"
	"github.com/team13/app/ent/predicate"
)

// GetOffWorkQuery is the builder for querying GetOffWork entities.
type GetOffWorkQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.GetOffWork
	// eager-loading edges.
	withWhenendwork *EmployeeWorkingHoursQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (gowq *GetOffWorkQuery) Where(ps ...predicate.GetOffWork) *GetOffWorkQuery {
	gowq.predicates = append(gowq.predicates, ps...)
	return gowq
}

// Limit adds a limit step to the query.
func (gowq *GetOffWorkQuery) Limit(limit int) *GetOffWorkQuery {
	gowq.limit = &limit
	return gowq
}

// Offset adds an offset step to the query.
func (gowq *GetOffWorkQuery) Offset(offset int) *GetOffWorkQuery {
	gowq.offset = &offset
	return gowq
}

// Order adds an order step to the query.
func (gowq *GetOffWorkQuery) Order(o ...OrderFunc) *GetOffWorkQuery {
	gowq.order = append(gowq.order, o...)
	return gowq
}

// QueryWhenendwork chains the current query on the whenendwork edge.
func (gowq *GetOffWorkQuery) QueryWhenendwork() *EmployeeWorkingHoursQuery {
	query := &EmployeeWorkingHoursQuery{config: gowq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gowq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(getoffwork.Table, getoffwork.FieldID, gowq.sqlQuery()),
			sqlgraph.To(employeeworkinghours.Table, employeeworkinghours.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, getoffwork.WhenendworkTable, getoffwork.WhenendworkColumn),
		)
		fromU = sqlgraph.SetNeighbors(gowq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GetOffWork entity in the query. Returns *NotFoundError when no getoffwork was found.
func (gowq *GetOffWorkQuery) First(ctx context.Context) (*GetOffWork, error) {
	gows, err := gowq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(gows) == 0 {
		return nil, &NotFoundError{getoffwork.Label}
	}
	return gows[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gowq *GetOffWorkQuery) FirstX(ctx context.Context) *GetOffWork {
	gow, err := gowq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return gow
}

// FirstID returns the first GetOffWork id in the query. Returns *NotFoundError when no id was found.
func (gowq *GetOffWorkQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gowq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{getoffwork.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (gowq *GetOffWorkQuery) FirstXID(ctx context.Context) int {
	id, err := gowq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only GetOffWork entity in the query, returns an error if not exactly one entity was returned.
func (gowq *GetOffWorkQuery) Only(ctx context.Context) (*GetOffWork, error) {
	gows, err := gowq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(gows) {
	case 1:
		return gows[0], nil
	case 0:
		return nil, &NotFoundError{getoffwork.Label}
	default:
		return nil, &NotSingularError{getoffwork.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gowq *GetOffWorkQuery) OnlyX(ctx context.Context) *GetOffWork {
	gow, err := gowq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return gow
}

// OnlyID returns the only GetOffWork id in the query, returns an error if not exactly one id was returned.
func (gowq *GetOffWorkQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gowq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{getoffwork.Label}
	default:
		err = &NotSingularError{getoffwork.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gowq *GetOffWorkQuery) OnlyIDX(ctx context.Context) int {
	id, err := gowq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GetOffWorks.
func (gowq *GetOffWorkQuery) All(ctx context.Context) ([]*GetOffWork, error) {
	if err := gowq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gowq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gowq *GetOffWorkQuery) AllX(ctx context.Context) []*GetOffWork {
	gows, err := gowq.All(ctx)
	if err != nil {
		panic(err)
	}
	return gows
}

// IDs executes the query and returns a list of GetOffWork ids.
func (gowq *GetOffWorkQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := gowq.Select(getoffwork.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gowq *GetOffWorkQuery) IDsX(ctx context.Context) []int {
	ids, err := gowq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gowq *GetOffWorkQuery) Count(ctx context.Context) (int, error) {
	if err := gowq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gowq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gowq *GetOffWorkQuery) CountX(ctx context.Context) int {
	count, err := gowq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gowq *GetOffWorkQuery) Exist(ctx context.Context) (bool, error) {
	if err := gowq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gowq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gowq *GetOffWorkQuery) ExistX(ctx context.Context) bool {
	exist, err := gowq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gowq *GetOffWorkQuery) Clone() *GetOffWorkQuery {
	return &GetOffWorkQuery{
		config:     gowq.config,
		limit:      gowq.limit,
		offset:     gowq.offset,
		order:      append([]OrderFunc{}, gowq.order...),
		unique:     append([]string{}, gowq.unique...),
		predicates: append([]predicate.GetOffWork{}, gowq.predicates...),
		// clone intermediate query.
		sql:  gowq.sql.Clone(),
		path: gowq.path,
	}
}

//  WithWhenendwork tells the query-builder to eager-loads the nodes that are connected to
// the "whenendwork" edge. The optional arguments used to configure the query builder of the edge.
func (gowq *GetOffWorkQuery) WithWhenendwork(opts ...func(*EmployeeWorkingHoursQuery)) *GetOffWorkQuery {
	query := &EmployeeWorkingHoursQuery{config: gowq.config}
	for _, opt := range opts {
		opt(query)
	}
	gowq.withWhenendwork = query
	return gowq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		GetOffWork time.Time `json:"GetOffWork,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GetOffWork.Query().
//		GroupBy(getoffwork.FieldGetOffWork).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gowq *GetOffWorkQuery) GroupBy(field string, fields ...string) *GetOffWorkGroupBy {
	group := &GetOffWorkGroupBy{config: gowq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gowq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gowq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		GetOffWork time.Time `json:"GetOffWork,omitempty"`
//	}
//
//	client.GetOffWork.Query().
//		Select(getoffwork.FieldGetOffWork).
//		Scan(ctx, &v)
//
func (gowq *GetOffWorkQuery) Select(field string, fields ...string) *GetOffWorkSelect {
	selector := &GetOffWorkSelect{config: gowq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gowq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gowq.sqlQuery(), nil
	}
	return selector
}

func (gowq *GetOffWorkQuery) prepareQuery(ctx context.Context) error {
	if gowq.path != nil {
		prev, err := gowq.path(ctx)
		if err != nil {
			return err
		}
		gowq.sql = prev
	}
	return nil
}

func (gowq *GetOffWorkQuery) sqlAll(ctx context.Context) ([]*GetOffWork, error) {
	var (
		nodes       = []*GetOffWork{}
		_spec       = gowq.querySpec()
		loadedTypes = [1]bool{
			gowq.withWhenendwork != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &GetOffWork{config: gowq.config}
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
	if err := sqlgraph.QueryNodes(ctx, gowq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := gowq.withWhenendwork; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*GetOffWork)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.EmployeeWorkingHours(func(s *sql.Selector) {
			s.Where(sql.InValues(getoffwork.WhenendworkColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.get_off_work_whenendwork
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "get_off_work_whenendwork" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "get_off_work_whenendwork" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Whenendwork = append(node.Edges.Whenendwork, n)
		}
	}

	return nodes, nil
}

func (gowq *GetOffWorkQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gowq.querySpec()
	return sqlgraph.CountNodes(ctx, gowq.driver, _spec)
}

func (gowq *GetOffWorkQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gowq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (gowq *GetOffWorkQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   getoffwork.Table,
			Columns: getoffwork.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: getoffwork.FieldID,
			},
		},
		From:   gowq.sql,
		Unique: true,
	}
	if ps := gowq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gowq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gowq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gowq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gowq *GetOffWorkQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(gowq.driver.Dialect())
	t1 := builder.Table(getoffwork.Table)
	selector := builder.Select(t1.Columns(getoffwork.Columns...)...).From(t1)
	if gowq.sql != nil {
		selector = gowq.sql
		selector.Select(selector.Columns(getoffwork.Columns...)...)
	}
	for _, p := range gowq.predicates {
		p(selector)
	}
	for _, p := range gowq.order {
		p(selector)
	}
	if offset := gowq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gowq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GetOffWorkGroupBy is the builder for group-by GetOffWork entities.
type GetOffWorkGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gowgb *GetOffWorkGroupBy) Aggregate(fns ...AggregateFunc) *GetOffWorkGroupBy {
	gowgb.fns = append(gowgb.fns, fns...)
	return gowgb
}

// Scan applies the group-by query and scan the result into the given value.
func (gowgb *GetOffWorkGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gowgb.path(ctx)
	if err != nil {
		return err
	}
	gowgb.sql = query
	return gowgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gowgb *GetOffWorkGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gowgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (gowgb *GetOffWorkGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gowgb.fields) > 1 {
		return nil, errors.New("ent: GetOffWorkGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gowgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gowgb *GetOffWorkGroupBy) StringsX(ctx context.Context) []string {
	v, err := gowgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (gowgb *GetOffWorkGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gowgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{getoffwork.Label}
	default:
		err = fmt.Errorf("ent: GetOffWorkGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gowgb *GetOffWorkGroupBy) StringX(ctx context.Context) string {
	v, err := gowgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (gowgb *GetOffWorkGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gowgb.fields) > 1 {
		return nil, errors.New("ent: GetOffWorkGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gowgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gowgb *GetOffWorkGroupBy) IntsX(ctx context.Context) []int {
	v, err := gowgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (gowgb *GetOffWorkGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gowgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{getoffwork.Label}
	default:
		err = fmt.Errorf("ent: GetOffWorkGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gowgb *GetOffWorkGroupBy) IntX(ctx context.Context) int {
	v, err := gowgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (gowgb *GetOffWorkGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gowgb.fields) > 1 {
		return nil, errors.New("ent: GetOffWorkGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gowgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gowgb *GetOffWorkGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gowgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (gowgb *GetOffWorkGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gowgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{getoffwork.Label}
	default:
		err = fmt.Errorf("ent: GetOffWorkGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gowgb *GetOffWorkGroupBy) Float64X(ctx context.Context) float64 {
	v, err := gowgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (gowgb *GetOffWorkGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gowgb.fields) > 1 {
		return nil, errors.New("ent: GetOffWorkGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gowgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gowgb *GetOffWorkGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gowgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (gowgb *GetOffWorkGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gowgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{getoffwork.Label}
	default:
		err = fmt.Errorf("ent: GetOffWorkGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gowgb *GetOffWorkGroupBy) BoolX(ctx context.Context) bool {
	v, err := gowgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gowgb *GetOffWorkGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gowgb.sqlQuery().Query()
	if err := gowgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gowgb *GetOffWorkGroupBy) sqlQuery() *sql.Selector {
	selector := gowgb.sql
	columns := make([]string, 0, len(gowgb.fields)+len(gowgb.fns))
	columns = append(columns, gowgb.fields...)
	for _, fn := range gowgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(gowgb.fields...)
}

// GetOffWorkSelect is the builder for select fields of GetOffWork entities.
type GetOffWorkSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (gows *GetOffWorkSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := gows.path(ctx)
	if err != nil {
		return err
	}
	gows.sql = query
	return gows.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gows *GetOffWorkSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gows.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (gows *GetOffWorkSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gows.fields) > 1 {
		return nil, errors.New("ent: GetOffWorkSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gows.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gows *GetOffWorkSelect) StringsX(ctx context.Context) []string {
	v, err := gows.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (gows *GetOffWorkSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gows.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{getoffwork.Label}
	default:
		err = fmt.Errorf("ent: GetOffWorkSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gows *GetOffWorkSelect) StringX(ctx context.Context) string {
	v, err := gows.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (gows *GetOffWorkSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gows.fields) > 1 {
		return nil, errors.New("ent: GetOffWorkSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gows.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gows *GetOffWorkSelect) IntsX(ctx context.Context) []int {
	v, err := gows.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (gows *GetOffWorkSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gows.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{getoffwork.Label}
	default:
		err = fmt.Errorf("ent: GetOffWorkSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gows *GetOffWorkSelect) IntX(ctx context.Context) int {
	v, err := gows.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (gows *GetOffWorkSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gows.fields) > 1 {
		return nil, errors.New("ent: GetOffWorkSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gows.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gows *GetOffWorkSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gows.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (gows *GetOffWorkSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gows.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{getoffwork.Label}
	default:
		err = fmt.Errorf("ent: GetOffWorkSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gows *GetOffWorkSelect) Float64X(ctx context.Context) float64 {
	v, err := gows.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (gows *GetOffWorkSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gows.fields) > 1 {
		return nil, errors.New("ent: GetOffWorkSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gows.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gows *GetOffWorkSelect) BoolsX(ctx context.Context) []bool {
	v, err := gows.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (gows *GetOffWorkSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gows.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{getoffwork.Label}
	default:
		err = fmt.Errorf("ent: GetOffWorkSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gows *GetOffWorkSelect) BoolX(ctx context.Context) bool {
	v, err := gows.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gows *GetOffWorkSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gows.sqlQuery().Query()
	if err := gows.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gows *GetOffWorkSelect) sqlQuery() sql.Querier {
	selector := gows.sql
	selector.Select(selector.Columns(gows.fields...)...)
	return selector
}
