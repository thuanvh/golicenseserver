// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/thuanvh/golicenseserver/ent/customer"
	"github.com/thuanvh/golicenseserver/ent/customerlicense"
	"github.com/thuanvh/golicenseserver/ent/license"
	"github.com/thuanvh/golicenseserver/ent/predicate"
)

// CustomerLicenseQuery is the builder for querying CustomerLicense entities.
type CustomerLicenseQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CustomerLicense
	// eager-loading edges.
	withCustomer *CustomerQuery
	withLicense  *LicenseQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CustomerLicenseQuery builder.
func (clq *CustomerLicenseQuery) Where(ps ...predicate.CustomerLicense) *CustomerLicenseQuery {
	clq.predicates = append(clq.predicates, ps...)
	return clq
}

// Limit adds a limit step to the query.
func (clq *CustomerLicenseQuery) Limit(limit int) *CustomerLicenseQuery {
	clq.limit = &limit
	return clq
}

// Offset adds an offset step to the query.
func (clq *CustomerLicenseQuery) Offset(offset int) *CustomerLicenseQuery {
	clq.offset = &offset
	return clq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (clq *CustomerLicenseQuery) Unique(unique bool) *CustomerLicenseQuery {
	clq.unique = &unique
	return clq
}

// Order adds an order step to the query.
func (clq *CustomerLicenseQuery) Order(o ...OrderFunc) *CustomerLicenseQuery {
	clq.order = append(clq.order, o...)
	return clq
}

// QueryCustomer chains the current query on the "customer" edge.
func (clq *CustomerLicenseQuery) QueryCustomer() *CustomerQuery {
	query := &CustomerQuery{config: clq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := clq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := clq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(customerlicense.Table, customerlicense.FieldID, selector),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, customerlicense.CustomerTable, customerlicense.CustomerColumn),
		)
		fromU = sqlgraph.SetNeighbors(clq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLicense chains the current query on the "license" edge.
func (clq *CustomerLicenseQuery) QueryLicense() *LicenseQuery {
	query := &LicenseQuery{config: clq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := clq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := clq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(customerlicense.Table, customerlicense.FieldID, selector),
			sqlgraph.To(license.Table, license.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, customerlicense.LicenseTable, customerlicense.LicenseColumn),
		)
		fromU = sqlgraph.SetNeighbors(clq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CustomerLicense entity from the query.
// Returns a *NotFoundError when no CustomerLicense was found.
func (clq *CustomerLicenseQuery) First(ctx context.Context) (*CustomerLicense, error) {
	nodes, err := clq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{customerlicense.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (clq *CustomerLicenseQuery) FirstX(ctx context.Context) *CustomerLicense {
	node, err := clq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CustomerLicense ID from the query.
// Returns a *NotFoundError when no CustomerLicense ID was found.
func (clq *CustomerLicenseQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = clq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{customerlicense.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (clq *CustomerLicenseQuery) FirstIDX(ctx context.Context) int {
	id, err := clq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CustomerLicense entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CustomerLicense entity is found.
// Returns a *NotFoundError when no CustomerLicense entities are found.
func (clq *CustomerLicenseQuery) Only(ctx context.Context) (*CustomerLicense, error) {
	nodes, err := clq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{customerlicense.Label}
	default:
		return nil, &NotSingularError{customerlicense.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (clq *CustomerLicenseQuery) OnlyX(ctx context.Context) *CustomerLicense {
	node, err := clq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CustomerLicense ID in the query.
// Returns a *NotSingularError when more than one CustomerLicense ID is found.
// Returns a *NotFoundError when no entities are found.
func (clq *CustomerLicenseQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = clq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{customerlicense.Label}
	default:
		err = &NotSingularError{customerlicense.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (clq *CustomerLicenseQuery) OnlyIDX(ctx context.Context) int {
	id, err := clq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CustomerLicenses.
func (clq *CustomerLicenseQuery) All(ctx context.Context) ([]*CustomerLicense, error) {
	if err := clq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return clq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (clq *CustomerLicenseQuery) AllX(ctx context.Context) []*CustomerLicense {
	nodes, err := clq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CustomerLicense IDs.
func (clq *CustomerLicenseQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := clq.Select(customerlicense.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (clq *CustomerLicenseQuery) IDsX(ctx context.Context) []int {
	ids, err := clq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (clq *CustomerLicenseQuery) Count(ctx context.Context) (int, error) {
	if err := clq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return clq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (clq *CustomerLicenseQuery) CountX(ctx context.Context) int {
	count, err := clq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (clq *CustomerLicenseQuery) Exist(ctx context.Context) (bool, error) {
	if err := clq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return clq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (clq *CustomerLicenseQuery) ExistX(ctx context.Context) bool {
	exist, err := clq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CustomerLicenseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (clq *CustomerLicenseQuery) Clone() *CustomerLicenseQuery {
	if clq == nil {
		return nil
	}
	return &CustomerLicenseQuery{
		config:       clq.config,
		limit:        clq.limit,
		offset:       clq.offset,
		order:        append([]OrderFunc{}, clq.order...),
		predicates:   append([]predicate.CustomerLicense{}, clq.predicates...),
		withCustomer: clq.withCustomer.Clone(),
		withLicense:  clq.withLicense.Clone(),
		// clone intermediate query.
		sql:    clq.sql.Clone(),
		path:   clq.path,
		unique: clq.unique,
	}
}

// WithCustomer tells the query-builder to eager-load the nodes that are connected to
// the "customer" edge. The optional arguments are used to configure the query builder of the edge.
func (clq *CustomerLicenseQuery) WithCustomer(opts ...func(*CustomerQuery)) *CustomerLicenseQuery {
	query := &CustomerQuery{config: clq.config}
	for _, opt := range opts {
		opt(query)
	}
	clq.withCustomer = query
	return clq
}

// WithLicense tells the query-builder to eager-load the nodes that are connected to
// the "license" edge. The optional arguments are used to configure the query builder of the edge.
func (clq *CustomerLicenseQuery) WithLicense(opts ...func(*LicenseQuery)) *CustomerLicenseQuery {
	query := &LicenseQuery{config: clq.config}
	for _, opt := range opts {
		opt(query)
	}
	clq.withLicense = query
	return clq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		LicenseCode string `json:"license_code,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CustomerLicense.Query().
//		GroupBy(customerlicense.FieldLicenseCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (clq *CustomerLicenseQuery) GroupBy(field string, fields ...string) *CustomerLicenseGroupBy {
	group := &CustomerLicenseGroupBy{config: clq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := clq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return clq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		LicenseCode string `json:"license_code,omitempty"`
//	}
//
//	client.CustomerLicense.Query().
//		Select(customerlicense.FieldLicenseCode).
//		Scan(ctx, &v)
//
func (clq *CustomerLicenseQuery) Select(fields ...string) *CustomerLicenseSelect {
	clq.fields = append(clq.fields, fields...)
	return &CustomerLicenseSelect{CustomerLicenseQuery: clq}
}

func (clq *CustomerLicenseQuery) prepareQuery(ctx context.Context) error {
	for _, f := range clq.fields {
		if !customerlicense.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if clq.path != nil {
		prev, err := clq.path(ctx)
		if err != nil {
			return err
		}
		clq.sql = prev
	}
	return nil
}

func (clq *CustomerLicenseQuery) sqlAll(ctx context.Context) ([]*CustomerLicense, error) {
	var (
		nodes       = []*CustomerLicense{}
		withFKs     = clq.withFKs
		_spec       = clq.querySpec()
		loadedTypes = [2]bool{
			clq.withCustomer != nil,
			clq.withLicense != nil,
		}
	)
	if clq.withCustomer != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, customerlicense.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CustomerLicense{config: clq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, clq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := clq.withCustomer; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*CustomerLicense)
		for i := range nodes {
			if nodes[i].customer_license_customer == nil {
				continue
			}
			fk := *nodes[i].customer_license_customer
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(customer.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "customer_license_customer" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Customer = n
			}
		}
	}

	if query := clq.withLicense; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*CustomerLicense)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.License = []*License{}
		}
		query.withFKs = true
		query.Where(predicate.License(func(s *sql.Selector) {
			s.Where(sql.InValues(customerlicense.LicenseColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.customer_license_license
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "customer_license_license" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "customer_license_license" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.License = append(node.Edges.License, n)
		}
	}

	return nodes, nil
}

func (clq *CustomerLicenseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := clq.querySpec()
	_spec.Node.Columns = clq.fields
	if len(clq.fields) > 0 {
		_spec.Unique = clq.unique != nil && *clq.unique
	}
	return sqlgraph.CountNodes(ctx, clq.driver, _spec)
}

func (clq *CustomerLicenseQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := clq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (clq *CustomerLicenseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customerlicense.Table,
			Columns: customerlicense.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customerlicense.FieldID,
			},
		},
		From:   clq.sql,
		Unique: true,
	}
	if unique := clq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := clq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customerlicense.FieldID)
		for i := range fields {
			if fields[i] != customerlicense.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := clq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := clq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := clq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := clq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (clq *CustomerLicenseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(clq.driver.Dialect())
	t1 := builder.Table(customerlicense.Table)
	columns := clq.fields
	if len(columns) == 0 {
		columns = customerlicense.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if clq.sql != nil {
		selector = clq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if clq.unique != nil && *clq.unique {
		selector.Distinct()
	}
	for _, p := range clq.predicates {
		p(selector)
	}
	for _, p := range clq.order {
		p(selector)
	}
	if offset := clq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := clq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CustomerLicenseGroupBy is the group-by builder for CustomerLicense entities.
type CustomerLicenseGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (clgb *CustomerLicenseGroupBy) Aggregate(fns ...AggregateFunc) *CustomerLicenseGroupBy {
	clgb.fns = append(clgb.fns, fns...)
	return clgb
}

// Scan applies the group-by query and scans the result into the given value.
func (clgb *CustomerLicenseGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := clgb.path(ctx)
	if err != nil {
		return err
	}
	clgb.sql = query
	return clgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (clgb *CustomerLicenseGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := clgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (clgb *CustomerLicenseGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(clgb.fields) > 1 {
		return nil, errors.New("ent: CustomerLicenseGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := clgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (clgb *CustomerLicenseGroupBy) StringsX(ctx context.Context) []string {
	v, err := clgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (clgb *CustomerLicenseGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = clgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{customerlicense.Label}
	default:
		err = fmt.Errorf("ent: CustomerLicenseGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (clgb *CustomerLicenseGroupBy) StringX(ctx context.Context) string {
	v, err := clgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (clgb *CustomerLicenseGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(clgb.fields) > 1 {
		return nil, errors.New("ent: CustomerLicenseGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := clgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (clgb *CustomerLicenseGroupBy) IntsX(ctx context.Context) []int {
	v, err := clgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (clgb *CustomerLicenseGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = clgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{customerlicense.Label}
	default:
		err = fmt.Errorf("ent: CustomerLicenseGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (clgb *CustomerLicenseGroupBy) IntX(ctx context.Context) int {
	v, err := clgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (clgb *CustomerLicenseGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(clgb.fields) > 1 {
		return nil, errors.New("ent: CustomerLicenseGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := clgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (clgb *CustomerLicenseGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := clgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (clgb *CustomerLicenseGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = clgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{customerlicense.Label}
	default:
		err = fmt.Errorf("ent: CustomerLicenseGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (clgb *CustomerLicenseGroupBy) Float64X(ctx context.Context) float64 {
	v, err := clgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (clgb *CustomerLicenseGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(clgb.fields) > 1 {
		return nil, errors.New("ent: CustomerLicenseGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := clgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (clgb *CustomerLicenseGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := clgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (clgb *CustomerLicenseGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = clgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{customerlicense.Label}
	default:
		err = fmt.Errorf("ent: CustomerLicenseGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (clgb *CustomerLicenseGroupBy) BoolX(ctx context.Context) bool {
	v, err := clgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (clgb *CustomerLicenseGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range clgb.fields {
		if !customerlicense.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := clgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := clgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (clgb *CustomerLicenseGroupBy) sqlQuery() *sql.Selector {
	selector := clgb.sql.Select()
	aggregation := make([]string, 0, len(clgb.fns))
	for _, fn := range clgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(clgb.fields)+len(clgb.fns))
		for _, f := range clgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(clgb.fields...)...)
}

// CustomerLicenseSelect is the builder for selecting fields of CustomerLicense entities.
type CustomerLicenseSelect struct {
	*CustomerLicenseQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cls *CustomerLicenseSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cls.prepareQuery(ctx); err != nil {
		return err
	}
	cls.sql = cls.CustomerLicenseQuery.sqlQuery(ctx)
	return cls.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cls *CustomerLicenseSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cls.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (cls *CustomerLicenseSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cls.fields) > 1 {
		return nil, errors.New("ent: CustomerLicenseSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cls *CustomerLicenseSelect) StringsX(ctx context.Context) []string {
	v, err := cls.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (cls *CustomerLicenseSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cls.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{customerlicense.Label}
	default:
		err = fmt.Errorf("ent: CustomerLicenseSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cls *CustomerLicenseSelect) StringX(ctx context.Context) string {
	v, err := cls.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (cls *CustomerLicenseSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cls.fields) > 1 {
		return nil, errors.New("ent: CustomerLicenseSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cls *CustomerLicenseSelect) IntsX(ctx context.Context) []int {
	v, err := cls.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (cls *CustomerLicenseSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cls.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{customerlicense.Label}
	default:
		err = fmt.Errorf("ent: CustomerLicenseSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cls *CustomerLicenseSelect) IntX(ctx context.Context) int {
	v, err := cls.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (cls *CustomerLicenseSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cls.fields) > 1 {
		return nil, errors.New("ent: CustomerLicenseSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cls *CustomerLicenseSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cls.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (cls *CustomerLicenseSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cls.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{customerlicense.Label}
	default:
		err = fmt.Errorf("ent: CustomerLicenseSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cls *CustomerLicenseSelect) Float64X(ctx context.Context) float64 {
	v, err := cls.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (cls *CustomerLicenseSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cls.fields) > 1 {
		return nil, errors.New("ent: CustomerLicenseSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cls *CustomerLicenseSelect) BoolsX(ctx context.Context) []bool {
	v, err := cls.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (cls *CustomerLicenseSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cls.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{customerlicense.Label}
	default:
		err = fmt.Errorf("ent: CustomerLicenseSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cls *CustomerLicenseSelect) BoolX(ctx context.Context) bool {
	v, err := cls.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cls *CustomerLicenseSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cls.sql.Query()
	if err := cls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
