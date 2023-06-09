// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_schemas/ent/predicate"
	"_schemas/ent/website"
	"_schemas/ent/websiteevent"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WebsiteEventQuery is the builder for querying WebsiteEvent entities.
type WebsiteEventQuery struct {
	config
	ctx         *QueryContext
	order       []websiteevent.OrderOption
	inters      []Interceptor
	predicates  []predicate.WebsiteEvent
	withWebsite *WebsiteQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WebsiteEventQuery builder.
func (weq *WebsiteEventQuery) Where(ps ...predicate.WebsiteEvent) *WebsiteEventQuery {
	weq.predicates = append(weq.predicates, ps...)
	return weq
}

// Limit the number of records to be returned by this query.
func (weq *WebsiteEventQuery) Limit(limit int) *WebsiteEventQuery {
	weq.ctx.Limit = &limit
	return weq
}

// Offset to start from.
func (weq *WebsiteEventQuery) Offset(offset int) *WebsiteEventQuery {
	weq.ctx.Offset = &offset
	return weq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (weq *WebsiteEventQuery) Unique(unique bool) *WebsiteEventQuery {
	weq.ctx.Unique = &unique
	return weq
}

// Order specifies how the records should be ordered.
func (weq *WebsiteEventQuery) Order(o ...websiteevent.OrderOption) *WebsiteEventQuery {
	weq.order = append(weq.order, o...)
	return weq
}

// QueryWebsite chains the current query on the "website" edge.
func (weq *WebsiteEventQuery) QueryWebsite() *WebsiteQuery {
	query := (&WebsiteClient{config: weq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := weq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := weq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(websiteevent.Table, websiteevent.FieldID, selector),
			sqlgraph.To(website.Table, website.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, websiteevent.WebsiteTable, websiteevent.WebsiteColumn),
		)
		fromU = sqlgraph.SetNeighbors(weq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WebsiteEvent entity from the query.
// Returns a *NotFoundError when no WebsiteEvent was found.
func (weq *WebsiteEventQuery) First(ctx context.Context) (*WebsiteEvent, error) {
	nodes, err := weq.Limit(1).All(setContextOp(ctx, weq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{websiteevent.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (weq *WebsiteEventQuery) FirstX(ctx context.Context) *WebsiteEvent {
	node, err := weq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WebsiteEvent ID from the query.
// Returns a *NotFoundError when no WebsiteEvent ID was found.
func (weq *WebsiteEventQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = weq.Limit(1).IDs(setContextOp(ctx, weq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{websiteevent.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (weq *WebsiteEventQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := weq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WebsiteEvent entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WebsiteEvent entity is found.
// Returns a *NotFoundError when no WebsiteEvent entities are found.
func (weq *WebsiteEventQuery) Only(ctx context.Context) (*WebsiteEvent, error) {
	nodes, err := weq.Limit(2).All(setContextOp(ctx, weq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{websiteevent.Label}
	default:
		return nil, &NotSingularError{websiteevent.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (weq *WebsiteEventQuery) OnlyX(ctx context.Context) *WebsiteEvent {
	node, err := weq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WebsiteEvent ID in the query.
// Returns a *NotSingularError when more than one WebsiteEvent ID is found.
// Returns a *NotFoundError when no entities are found.
func (weq *WebsiteEventQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = weq.Limit(2).IDs(setContextOp(ctx, weq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{websiteevent.Label}
	default:
		err = &NotSingularError{websiteevent.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (weq *WebsiteEventQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := weq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WebsiteEvents.
func (weq *WebsiteEventQuery) All(ctx context.Context) ([]*WebsiteEvent, error) {
	ctx = setContextOp(ctx, weq.ctx, "All")
	if err := weq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WebsiteEvent, *WebsiteEventQuery]()
	return withInterceptors[[]*WebsiteEvent](ctx, weq, qr, weq.inters)
}

// AllX is like All, but panics if an error occurs.
func (weq *WebsiteEventQuery) AllX(ctx context.Context) []*WebsiteEvent {
	nodes, err := weq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WebsiteEvent IDs.
func (weq *WebsiteEventQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if weq.ctx.Unique == nil && weq.path != nil {
		weq.Unique(true)
	}
	ctx = setContextOp(ctx, weq.ctx, "IDs")
	if err = weq.Select(websiteevent.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (weq *WebsiteEventQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := weq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (weq *WebsiteEventQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, weq.ctx, "Count")
	if err := weq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, weq, querierCount[*WebsiteEventQuery](), weq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (weq *WebsiteEventQuery) CountX(ctx context.Context) int {
	count, err := weq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (weq *WebsiteEventQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, weq.ctx, "Exist")
	switch _, err := weq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (weq *WebsiteEventQuery) ExistX(ctx context.Context) bool {
	exist, err := weq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WebsiteEventQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (weq *WebsiteEventQuery) Clone() *WebsiteEventQuery {
	if weq == nil {
		return nil
	}
	return &WebsiteEventQuery{
		config:      weq.config,
		ctx:         weq.ctx.Clone(),
		order:       append([]websiteevent.OrderOption{}, weq.order...),
		inters:      append([]Interceptor{}, weq.inters...),
		predicates:  append([]predicate.WebsiteEvent{}, weq.predicates...),
		withWebsite: weq.withWebsite.Clone(),
		// clone intermediate query.
		sql:  weq.sql.Clone(),
		path: weq.path,
	}
}

// WithWebsite tells the query-builder to eager-load the nodes that are connected to
// the "website" edge. The optional arguments are used to configure the query builder of the edge.
func (weq *WebsiteEventQuery) WithWebsite(opts ...func(*WebsiteQuery)) *WebsiteEventQuery {
	query := (&WebsiteClient{config: weq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	weq.withWebsite = query
	return weq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EventName string `json:"event_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WebsiteEvent.Query().
//		GroupBy(websiteevent.FieldEventName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (weq *WebsiteEventQuery) GroupBy(field string, fields ...string) *WebsiteEventGroupBy {
	weq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WebsiteEventGroupBy{build: weq}
	grbuild.flds = &weq.ctx.Fields
	grbuild.label = websiteevent.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EventName string `json:"event_name,omitempty"`
//	}
//
//	client.WebsiteEvent.Query().
//		Select(websiteevent.FieldEventName).
//		Scan(ctx, &v)
func (weq *WebsiteEventQuery) Select(fields ...string) *WebsiteEventSelect {
	weq.ctx.Fields = append(weq.ctx.Fields, fields...)
	sbuild := &WebsiteEventSelect{WebsiteEventQuery: weq}
	sbuild.label = websiteevent.Label
	sbuild.flds, sbuild.scan = &weq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WebsiteEventSelect configured with the given aggregations.
func (weq *WebsiteEventQuery) Aggregate(fns ...AggregateFunc) *WebsiteEventSelect {
	return weq.Select().Aggregate(fns...)
}

func (weq *WebsiteEventQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range weq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, weq); err != nil {
				return err
			}
		}
	}
	for _, f := range weq.ctx.Fields {
		if !websiteevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if weq.path != nil {
		prev, err := weq.path(ctx)
		if err != nil {
			return err
		}
		weq.sql = prev
	}
	return nil
}

func (weq *WebsiteEventQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WebsiteEvent, error) {
	var (
		nodes       = []*WebsiteEvent{}
		withFKs     = weq.withFKs
		_spec       = weq.querySpec()
		loadedTypes = [1]bool{
			weq.withWebsite != nil,
		}
	)
	if weq.withWebsite != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, websiteevent.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WebsiteEvent).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WebsiteEvent{config: weq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, weq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := weq.withWebsite; query != nil {
		if err := weq.loadWebsite(ctx, query, nodes, nil,
			func(n *WebsiteEvent, e *Website) { n.Edges.Website = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (weq *WebsiteEventQuery) loadWebsite(ctx context.Context, query *WebsiteQuery, nodes []*WebsiteEvent, init func(*WebsiteEvent), assign func(*WebsiteEvent, *Website)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*WebsiteEvent)
	for i := range nodes {
		if nodes[i].website_id == nil {
			continue
		}
		fk := *nodes[i].website_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(website.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "website_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (weq *WebsiteEventQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := weq.querySpec()
	_spec.Node.Columns = weq.ctx.Fields
	if len(weq.ctx.Fields) > 0 {
		_spec.Unique = weq.ctx.Unique != nil && *weq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, weq.driver, _spec)
}

func (weq *WebsiteEventQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(websiteevent.Table, websiteevent.Columns, sqlgraph.NewFieldSpec(websiteevent.FieldID, field.TypeUUID))
	_spec.From = weq.sql
	if unique := weq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if weq.path != nil {
		_spec.Unique = true
	}
	if fields := weq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, websiteevent.FieldID)
		for i := range fields {
			if fields[i] != websiteevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := weq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := weq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := weq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := weq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (weq *WebsiteEventQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(weq.driver.Dialect())
	t1 := builder.Table(websiteevent.Table)
	columns := weq.ctx.Fields
	if len(columns) == 0 {
		columns = websiteevent.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if weq.sql != nil {
		selector = weq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if weq.ctx.Unique != nil && *weq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range weq.predicates {
		p(selector)
	}
	for _, p := range weq.order {
		p(selector)
	}
	if offset := weq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := weq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WebsiteEventGroupBy is the group-by builder for WebsiteEvent entities.
type WebsiteEventGroupBy struct {
	selector
	build *WebsiteEventQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wegb *WebsiteEventGroupBy) Aggregate(fns ...AggregateFunc) *WebsiteEventGroupBy {
	wegb.fns = append(wegb.fns, fns...)
	return wegb
}

// Scan applies the selector query and scans the result into the given value.
func (wegb *WebsiteEventGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wegb.build.ctx, "GroupBy")
	if err := wegb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WebsiteEventQuery, *WebsiteEventGroupBy](ctx, wegb.build, wegb, wegb.build.inters, v)
}

func (wegb *WebsiteEventGroupBy) sqlScan(ctx context.Context, root *WebsiteEventQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wegb.fns))
	for _, fn := range wegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wegb.flds)+len(wegb.fns))
		for _, f := range *wegb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wegb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wegb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WebsiteEventSelect is the builder for selecting fields of WebsiteEvent entities.
type WebsiteEventSelect struct {
	*WebsiteEventQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wes *WebsiteEventSelect) Aggregate(fns ...AggregateFunc) *WebsiteEventSelect {
	wes.fns = append(wes.fns, fns...)
	return wes
}

// Scan applies the selector query and scans the result into the given value.
func (wes *WebsiteEventSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wes.ctx, "Select")
	if err := wes.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WebsiteEventQuery, *WebsiteEventSelect](ctx, wes.WebsiteEventQuery, wes, wes.inters, v)
}

func (wes *WebsiteEventSelect) sqlScan(ctx context.Context, root *WebsiteEventQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wes.fns))
	for _, fn := range wes.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wes.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
