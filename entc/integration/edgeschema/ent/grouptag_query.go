// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/group"
	"entgo.io/ent/entc/integration/edgeschema/ent/grouptag"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/tag"
	"entgo.io/ent/schema/field"
)

// GroupTagQuery is the builder for querying GroupTag entities.
type GroupTagQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.GroupTag
	withTag    *TagQuery
	withGroup  *GroupQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroupTagQuery builder.
func (gtq *GroupTagQuery) Where(ps ...predicate.GroupTag) *GroupTagQuery {
	gtq.predicates = append(gtq.predicates, ps...)
	return gtq
}

// Limit the number of records to be returned by this query.
func (gtq *GroupTagQuery) Limit(limit int) *GroupTagQuery {
	gtq.limit = &limit
	return gtq
}

// Offset to start from.
func (gtq *GroupTagQuery) Offset(offset int) *GroupTagQuery {
	gtq.offset = &offset
	return gtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gtq *GroupTagQuery) Unique(unique bool) *GroupTagQuery {
	gtq.unique = &unique
	return gtq
}

// Order specifies how the records should be ordered.
func (gtq *GroupTagQuery) Order(o ...OrderFunc) *GroupTagQuery {
	gtq.order = append(gtq.order, o...)
	return gtq
}

// QueryTag chains the current query on the "tag" edge.
func (gtq *GroupTagQuery) QueryTag() *TagQuery {
	query := (&TagClient{config: gtq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(grouptag.Table, grouptag.FieldID, selector),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, grouptag.TagTable, grouptag.TagColumn),
		)
		fromU = sqlgraph.SetNeighbors(gtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGroup chains the current query on the "group" edge.
func (gtq *GroupTagQuery) QueryGroup() *GroupQuery {
	query := (&GroupClient{config: gtq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gtq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(grouptag.Table, grouptag.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, grouptag.GroupTable, grouptag.GroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(gtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GroupTag entity from the query.
// Returns a *NotFoundError when no GroupTag was found.
func (gtq *GroupTagQuery) First(ctx context.Context) (*GroupTag, error) {
	nodes, err := gtq.Limit(1).All(newQueryContext(ctx, TypeGroupTag, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{grouptag.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gtq *GroupTagQuery) FirstX(ctx context.Context) *GroupTag {
	node, err := gtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GroupTag ID from the query.
// Returns a *NotFoundError when no GroupTag ID was found.
func (gtq *GroupTagQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gtq.Limit(1).IDs(newQueryContext(ctx, TypeGroupTag, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{grouptag.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gtq *GroupTagQuery) FirstIDX(ctx context.Context) int {
	id, err := gtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GroupTag entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GroupTag entity is found.
// Returns a *NotFoundError when no GroupTag entities are found.
func (gtq *GroupTagQuery) Only(ctx context.Context) (*GroupTag, error) {
	nodes, err := gtq.Limit(2).All(newQueryContext(ctx, TypeGroupTag, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{grouptag.Label}
	default:
		return nil, &NotSingularError{grouptag.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gtq *GroupTagQuery) OnlyX(ctx context.Context) *GroupTag {
	node, err := gtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GroupTag ID in the query.
// Returns a *NotSingularError when more than one GroupTag ID is found.
// Returns a *NotFoundError when no entities are found.
func (gtq *GroupTagQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gtq.Limit(2).IDs(newQueryContext(ctx, TypeGroupTag, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{grouptag.Label}
	default:
		err = &NotSingularError{grouptag.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gtq *GroupTagQuery) OnlyIDX(ctx context.Context) int {
	id, err := gtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GroupTags.
func (gtq *GroupTagQuery) All(ctx context.Context) ([]*GroupTag, error) {
	ctx = newQueryContext(ctx, TypeGroupTag, "All")
	if err := gtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GroupTag, *GroupTagQuery]()
	return withInterceptors[[]*GroupTag](ctx, gtq, qr, gtq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gtq *GroupTagQuery) AllX(ctx context.Context) []*GroupTag {
	nodes, err := gtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GroupTag IDs.
func (gtq *GroupTagQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = newQueryContext(ctx, TypeGroupTag, "IDs")
	if err := gtq.Select(grouptag.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gtq *GroupTagQuery) IDsX(ctx context.Context) []int {
	ids, err := gtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gtq *GroupTagQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeGroupTag, "Count")
	if err := gtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gtq, querierCount[*GroupTagQuery](), gtq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gtq *GroupTagQuery) CountX(ctx context.Context) int {
	count, err := gtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gtq *GroupTagQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeGroupTag, "Exist")
	switch _, err := gtq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gtq *GroupTagQuery) ExistX(ctx context.Context) bool {
	exist, err := gtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupTagQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gtq *GroupTagQuery) Clone() *GroupTagQuery {
	if gtq == nil {
		return nil
	}
	return &GroupTagQuery{
		config:     gtq.config,
		limit:      gtq.limit,
		offset:     gtq.offset,
		order:      append([]OrderFunc{}, gtq.order...),
		predicates: append([]predicate.GroupTag{}, gtq.predicates...),
		withTag:    gtq.withTag.Clone(),
		withGroup:  gtq.withGroup.Clone(),
		// clone intermediate query.
		sql:    gtq.sql.Clone(),
		path:   gtq.path,
		unique: gtq.unique,
	}
}

// WithTag tells the query-builder to eager-load the nodes that are connected to
// the "tag" edge. The optional arguments are used to configure the query builder of the edge.
func (gtq *GroupTagQuery) WithTag(opts ...func(*TagQuery)) *GroupTagQuery {
	query := (&TagClient{config: gtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gtq.withTag = query
	return gtq
}

// WithGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (gtq *GroupTagQuery) WithGroup(opts ...func(*GroupQuery)) *GroupTagQuery {
	query := (&GroupClient{config: gtq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gtq.withGroup = query
	return gtq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TagID int `json:"tag_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GroupTag.Query().
//		GroupBy(grouptag.FieldTagID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gtq *GroupTagQuery) GroupBy(field string, fields ...string) *GroupTagGroupBy {
	gtq.fields = append([]string{field}, fields...)
	grbuild := &GroupTagGroupBy{build: gtq}
	grbuild.flds = &gtq.fields
	grbuild.label = grouptag.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TagID int `json:"tag_id,omitempty"`
//	}
//
//	client.GroupTag.Query().
//		Select(grouptag.FieldTagID).
//		Scan(ctx, &v)
func (gtq *GroupTagQuery) Select(fields ...string) *GroupTagSelect {
	gtq.fields = append(gtq.fields, fields...)
	sbuild := &GroupTagSelect{GroupTagQuery: gtq}
	sbuild.label = grouptag.Label
	sbuild.flds, sbuild.scan = &gtq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GroupTagSelect configured with the given aggregations.
func (gtq *GroupTagQuery) Aggregate(fns ...AggregateFunc) *GroupTagSelect {
	return gtq.Select().Aggregate(fns...)
}

func (gtq *GroupTagQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gtq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gtq); err != nil {
				return err
			}
		}
	}
	for _, f := range gtq.fields {
		if !grouptag.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gtq.path != nil {
		prev, err := gtq.path(ctx)
		if err != nil {
			return err
		}
		gtq.sql = prev
	}
	return nil
}

func (gtq *GroupTagQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GroupTag, error) {
	var (
		nodes       = []*GroupTag{}
		_spec       = gtq.querySpec()
		loadedTypes = [2]bool{
			gtq.withTag != nil,
			gtq.withGroup != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GroupTag).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GroupTag{config: gtq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gtq.withTag; query != nil {
		if err := gtq.loadTag(ctx, query, nodes, nil,
			func(n *GroupTag, e *Tag) { n.Edges.Tag = e }); err != nil {
			return nil, err
		}
	}
	if query := gtq.withGroup; query != nil {
		if err := gtq.loadGroup(ctx, query, nodes, nil,
			func(n *GroupTag, e *Group) { n.Edges.Group = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gtq *GroupTagQuery) loadTag(ctx context.Context, query *TagQuery, nodes []*GroupTag, init func(*GroupTag), assign func(*GroupTag, *Tag)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*GroupTag)
	for i := range nodes {
		fk := nodes[i].TagID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(tag.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "tag_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (gtq *GroupTagQuery) loadGroup(ctx context.Context, query *GroupQuery, nodes []*GroupTag, init func(*GroupTag), assign func(*GroupTag, *Group)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*GroupTag)
	for i := range nodes {
		fk := nodes[i].GroupID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(group.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "group_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (gtq *GroupTagQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gtq.querySpec()
	_spec.Node.Columns = gtq.fields
	if len(gtq.fields) > 0 {
		_spec.Unique = gtq.unique != nil && *gtq.unique
	}
	return sqlgraph.CountNodes(ctx, gtq.driver, _spec)
}

func (gtq *GroupTagQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   grouptag.Table,
			Columns: grouptag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: grouptag.FieldID,
			},
		},
		From:   gtq.sql,
		Unique: true,
	}
	if unique := gtq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gtq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, grouptag.FieldID)
		for i := range fields {
			if fields[i] != grouptag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gtq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gtq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gtq *GroupTagQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gtq.driver.Dialect())
	t1 := builder.Table(grouptag.Table)
	columns := gtq.fields
	if len(columns) == 0 {
		columns = grouptag.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gtq.sql != nil {
		selector = gtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gtq.unique != nil && *gtq.unique {
		selector.Distinct()
	}
	for _, p := range gtq.predicates {
		p(selector)
	}
	for _, p := range gtq.order {
		p(selector)
	}
	if offset := gtq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gtq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GroupTagGroupBy is the group-by builder for GroupTag entities.
type GroupTagGroupBy struct {
	selector
	build *GroupTagQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gtgb *GroupTagGroupBy) Aggregate(fns ...AggregateFunc) *GroupTagGroupBy {
	gtgb.fns = append(gtgb.fns, fns...)
	return gtgb
}

// Scan applies the selector query and scans the result into the given value.
func (gtgb *GroupTagGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeGroupTag, "GroupBy")
	if err := gtgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupTagQuery, *GroupTagGroupBy](ctx, gtgb.build, gtgb, gtgb.build.inters, v)
}

func (gtgb *GroupTagGroupBy) sqlScan(ctx context.Context, root *GroupTagQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gtgb.fns))
	for _, fn := range gtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gtgb.flds)+len(gtgb.fns))
		for _, f := range *gtgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gtgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gtgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GroupTagSelect is the builder for selecting fields of GroupTag entities.
type GroupTagSelect struct {
	*GroupTagQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gts *GroupTagSelect) Aggregate(fns ...AggregateFunc) *GroupTagSelect {
	gts.fns = append(gts.fns, fns...)
	return gts
}

// Scan applies the selector query and scans the result into the given value.
func (gts *GroupTagSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeGroupTag, "Select")
	if err := gts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupTagQuery, *GroupTagSelect](ctx, gts.GroupTagQuery, gts, gts.inters, v)
}

func (gts *GroupTagSelect) sqlScan(ctx context.Context, root *GroupTagQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gts.fns))
	for _, fn := range gts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
