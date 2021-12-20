// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"study_event_go/ent/garden"
	"study_event_go/ent/mentorshipsystem"
	"study_event_go/ent/predicate"
	"study_event_go/types"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GardenQuery is the builder for querying Garden entities.
type GardenQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Garden
	// eager-loading edges.
	withMentorshipSystem *MentorshipSystemQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GardenQuery builder.
func (gq *GardenQuery) Where(ps ...predicate.Garden) *GardenQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit adds a limit step to the query.
func (gq *GardenQuery) Limit(limit int) *GardenQuery {
	gq.limit = &limit
	return gq
}

// Offset adds an offset step to the query.
func (gq *GardenQuery) Offset(offset int) *GardenQuery {
	gq.offset = &offset
	return gq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gq *GardenQuery) Unique(unique bool) *GardenQuery {
	gq.unique = &unique
	return gq
}

// Order adds an order step to the query.
func (gq *GardenQuery) Order(o ...OrderFunc) *GardenQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// QueryMentorshipSystem chains the current query on the "mentorship_system" edge.
func (gq *GardenQuery) QueryMentorshipSystem() *MentorshipSystemQuery {
	query := &MentorshipSystemQuery{config: gq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(garden.Table, garden.FieldID, selector),
			sqlgraph.To(mentorshipsystem.Table, mentorshipsystem.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, garden.MentorshipSystemTable, garden.MentorshipSystemColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Garden entity from the query.
// Returns a *NotFoundError when no Garden was found.
func (gq *GardenQuery) First(ctx context.Context) (*Garden, error) {
	nodes, err := gq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{garden.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GardenQuery) FirstX(ctx context.Context) *Garden {
	node, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Garden ID from the query.
// Returns a *NotFoundError when no Garden ID was found.
func (gq *GardenQuery) FirstID(ctx context.Context) (id types.GardenID, err error) {
	var ids []types.GardenID
	if ids, err = gq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{garden.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gq *GardenQuery) FirstIDX(ctx context.Context) types.GardenID {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Garden entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Garden entity is not found.
// Returns a *NotFoundError when no Garden entities are found.
func (gq *GardenQuery) Only(ctx context.Context) (*Garden, error) {
	nodes, err := gq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{garden.Label}
	default:
		return nil, &NotSingularError{garden.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GardenQuery) OnlyX(ctx context.Context) *Garden {
	node, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Garden ID in the query.
// Returns a *NotSingularError when exactly one Garden ID is not found.
// Returns a *NotFoundError when no entities are found.
func (gq *GardenQuery) OnlyID(ctx context.Context) (id types.GardenID, err error) {
	var ids []types.GardenID
	if ids, err = gq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{garden.Label}
	default:
		err = &NotSingularError{garden.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gq *GardenQuery) OnlyIDX(ctx context.Context) types.GardenID {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Gardens.
func (gq *GardenQuery) All(ctx context.Context) ([]*Garden, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gq *GardenQuery) AllX(ctx context.Context) []*Garden {
	nodes, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Garden IDs.
func (gq *GardenQuery) IDs(ctx context.Context) ([]types.GardenID, error) {
	var ids []types.GardenID
	if err := gq.Select(garden.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GardenQuery) IDsX(ctx context.Context) []types.GardenID {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GardenQuery) Count(ctx context.Context) (int, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GardenQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GardenQuery) Exist(ctx context.Context) (bool, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GardenQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GardenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GardenQuery) Clone() *GardenQuery {
	if gq == nil {
		return nil
	}
	return &GardenQuery{
		config:               gq.config,
		limit:                gq.limit,
		offset:               gq.offset,
		order:                append([]OrderFunc{}, gq.order...),
		predicates:           append([]predicate.Garden{}, gq.predicates...),
		withMentorshipSystem: gq.withMentorshipSystem.Clone(),
		// clone intermediate query.
		sql:  gq.sql.Clone(),
		path: gq.path,
	}
}

// WithMentorshipSystem tells the query-builder to eager-load the nodes that are connected to
// the "mentorship_system" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GardenQuery) WithMentorshipSystem(opts ...func(*MentorshipSystemQuery)) *GardenQuery {
	query := &MentorshipSystemQuery{config: gq.config}
	for _, opt := range opts {
		opt(query)
	}
	gq.withMentorshipSystem = query
	return gq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Garden.Query().
//		GroupBy(garden.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gq *GardenQuery) GroupBy(field string, fields ...string) *GardenGroupBy {
	group := &GardenGroupBy{config: gq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Garden.Query().
//		Select(garden.FieldName).
//		Scan(ctx, &v)
//
func (gq *GardenQuery) Select(fields ...string) *GardenSelect {
	gq.fields = append(gq.fields, fields...)
	return &GardenSelect{GardenQuery: gq}
}

func (gq *GardenQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gq.fields {
		if !garden.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.sql = prev
	}
	return nil
}

func (gq *GardenQuery) sqlAll(ctx context.Context) ([]*Garden, error) {
	var (
		nodes       = []*Garden{}
		_spec       = gq.querySpec()
		loadedTypes = [1]bool{
			gq.withMentorshipSystem != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Garden{config: gq.config}
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
	if err := sqlgraph.QueryNodes(ctx, gq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := gq.withMentorshipSystem; query != nil {
		ids := make([]types.MentorshipSystemID, 0, len(nodes))
		nodeids := make(map[types.MentorshipSystemID][]*Garden)
		for i := range nodes {
			if nodes[i].MentorshipSystemID == nil {
				continue
			}
			fk := *nodes[i].MentorshipSystemID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(mentorshipsystem.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "mentorship_system_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.MentorshipSystem = n
			}
		}
	}

	return nodes, nil
}

func (gq *GardenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gq.querySpec()
	return sqlgraph.CountNodes(ctx, gq.driver, _spec)
}

func (gq *GardenQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (gq *GardenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   garden.Table,
			Columns: garden.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: garden.FieldID,
			},
		},
		From:   gq.sql,
		Unique: true,
	}
	if unique := gq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, garden.FieldID)
		for i := range fields {
			if fields[i] != garden.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gq *GardenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gq.driver.Dialect())
	t1 := builder.Table(garden.Table)
	columns := gq.fields
	if len(columns) == 0 {
		columns = garden.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gq.sql != nil {
		selector = gq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range gq.predicates {
		p(selector)
	}
	for _, p := range gq.order {
		p(selector)
	}
	if offset := gq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GardenGroupBy is the group-by builder for Garden entities.
type GardenGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GardenGroupBy) Aggregate(fns ...AggregateFunc) *GardenGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the group-by query and scans the result into the given value.
func (ggb *GardenGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ggb.path(ctx)
	if err != nil {
		return err
	}
	ggb.sql = query
	return ggb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ggb *GardenGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ggb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GardenGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GardenGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ggb *GardenGroupBy) StringsX(ctx context.Context) []string {
	v, err := ggb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GardenGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ggb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{garden.Label}
	default:
		err = fmt.Errorf("ent: GardenGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ggb *GardenGroupBy) StringX(ctx context.Context) string {
	v, err := ggb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GardenGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GardenGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ggb *GardenGroupBy) IntsX(ctx context.Context) []int {
	v, err := ggb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GardenGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ggb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{garden.Label}
	default:
		err = fmt.Errorf("ent: GardenGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ggb *GardenGroupBy) IntX(ctx context.Context) int {
	v, err := ggb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GardenGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GardenGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ggb *GardenGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ggb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GardenGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ggb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{garden.Label}
	default:
		err = fmt.Errorf("ent: GardenGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ggb *GardenGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ggb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GardenGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GardenGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ggb *GardenGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ggb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GardenGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ggb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{garden.Label}
	default:
		err = fmt.Errorf("ent: GardenGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ggb *GardenGroupBy) BoolX(ctx context.Context) bool {
	v, err := ggb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ggb *GardenGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ggb.fields {
		if !garden.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ggb *GardenGroupBy) sqlQuery() *sql.Selector {
	selector := ggb.sql.Select()
	aggregation := make([]string, 0, len(ggb.fns))
	for _, fn := range ggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ggb.fields)+len(ggb.fns))
		for _, f := range ggb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ggb.fields...)...)
}

// GardenSelect is the builder for selecting fields of Garden entities.
type GardenSelect struct {
	*GardenQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GardenSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	gs.sql = gs.GardenQuery.sqlQuery(ctx)
	return gs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gs *GardenSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gs *GardenSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GardenSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gs *GardenSelect) StringsX(ctx context.Context) []string {
	v, err := gs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gs *GardenSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{garden.Label}
	default:
		err = fmt.Errorf("ent: GardenSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gs *GardenSelect) StringX(ctx context.Context) string {
	v, err := gs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gs *GardenSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GardenSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gs *GardenSelect) IntsX(ctx context.Context) []int {
	v, err := gs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gs *GardenSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{garden.Label}
	default:
		err = fmt.Errorf("ent: GardenSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gs *GardenSelect) IntX(ctx context.Context) int {
	v, err := gs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gs *GardenSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GardenSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gs *GardenSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gs *GardenSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{garden.Label}
	default:
		err = fmt.Errorf("ent: GardenSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gs *GardenSelect) Float64X(ctx context.Context) float64 {
	v, err := gs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gs *GardenSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GardenSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gs *GardenSelect) BoolsX(ctx context.Context) []bool {
	v, err := gs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gs *GardenSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{garden.Label}
	default:
		err = fmt.Errorf("ent: GardenSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gs *GardenSelect) BoolX(ctx context.Context) bool {
	v, err := gs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gs *GardenSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gs.sql.Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
