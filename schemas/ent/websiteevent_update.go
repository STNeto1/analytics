// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_schemas/ent/predicate"
	"_schemas/ent/website"
	"_schemas/ent/websiteevent"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WebsiteEventUpdate is the builder for updating WebsiteEvent entities.
type WebsiteEventUpdate struct {
	config
	hooks    []Hook
	mutation *WebsiteEventMutation
}

// Where appends a list predicates to the WebsiteEventUpdate builder.
func (weu *WebsiteEventUpdate) Where(ps ...predicate.WebsiteEvent) *WebsiteEventUpdate {
	weu.mutation.Where(ps...)
	return weu
}

// SetEventName sets the "event_name" field.
func (weu *WebsiteEventUpdate) SetEventName(s string) *WebsiteEventUpdate {
	weu.mutation.SetEventName(s)
	return weu
}

// SetURLPath sets the "url_path" field.
func (weu *WebsiteEventUpdate) SetURLPath(s string) *WebsiteEventUpdate {
	weu.mutation.SetURLPath(s)
	return weu
}

// SetNillableURLPath sets the "url_path" field if the given value is not nil.
func (weu *WebsiteEventUpdate) SetNillableURLPath(s *string) *WebsiteEventUpdate {
	if s != nil {
		weu.SetURLPath(*s)
	}
	return weu
}

// ClearURLPath clears the value of the "url_path" field.
func (weu *WebsiteEventUpdate) ClearURLPath() *WebsiteEventUpdate {
	weu.mutation.ClearURLPath()
	return weu
}

// SetURLQuery sets the "url_query" field.
func (weu *WebsiteEventUpdate) SetURLQuery(s string) *WebsiteEventUpdate {
	weu.mutation.SetURLQuery(s)
	return weu
}

// SetNillableURLQuery sets the "url_query" field if the given value is not nil.
func (weu *WebsiteEventUpdate) SetNillableURLQuery(s *string) *WebsiteEventUpdate {
	if s != nil {
		weu.SetURLQuery(*s)
	}
	return weu
}

// ClearURLQuery clears the value of the "url_query" field.
func (weu *WebsiteEventUpdate) ClearURLQuery() *WebsiteEventUpdate {
	weu.mutation.ClearURLQuery()
	return weu
}

// SetReferrerPath sets the "referrer_path" field.
func (weu *WebsiteEventUpdate) SetReferrerPath(s string) *WebsiteEventUpdate {
	weu.mutation.SetReferrerPath(s)
	return weu
}

// SetNillableReferrerPath sets the "referrer_path" field if the given value is not nil.
func (weu *WebsiteEventUpdate) SetNillableReferrerPath(s *string) *WebsiteEventUpdate {
	if s != nil {
		weu.SetReferrerPath(*s)
	}
	return weu
}

// ClearReferrerPath clears the value of the "referrer_path" field.
func (weu *WebsiteEventUpdate) ClearReferrerPath() *WebsiteEventUpdate {
	weu.mutation.ClearReferrerPath()
	return weu
}

// SetReferrerQuery sets the "referrer_query" field.
func (weu *WebsiteEventUpdate) SetReferrerQuery(s string) *WebsiteEventUpdate {
	weu.mutation.SetReferrerQuery(s)
	return weu
}

// SetNillableReferrerQuery sets the "referrer_query" field if the given value is not nil.
func (weu *WebsiteEventUpdate) SetNillableReferrerQuery(s *string) *WebsiteEventUpdate {
	if s != nil {
		weu.SetReferrerQuery(*s)
	}
	return weu
}

// ClearReferrerQuery clears the value of the "referrer_query" field.
func (weu *WebsiteEventUpdate) ClearReferrerQuery() *WebsiteEventUpdate {
	weu.mutation.ClearReferrerQuery()
	return weu
}

// SetReferrerDomain sets the "referrer_domain" field.
func (weu *WebsiteEventUpdate) SetReferrerDomain(s string) *WebsiteEventUpdate {
	weu.mutation.SetReferrerDomain(s)
	return weu
}

// SetNillableReferrerDomain sets the "referrer_domain" field if the given value is not nil.
func (weu *WebsiteEventUpdate) SetNillableReferrerDomain(s *string) *WebsiteEventUpdate {
	if s != nil {
		weu.SetReferrerDomain(*s)
	}
	return weu
}

// ClearReferrerDomain clears the value of the "referrer_domain" field.
func (weu *WebsiteEventUpdate) ClearReferrerDomain() *WebsiteEventUpdate {
	weu.mutation.ClearReferrerDomain()
	return weu
}

// SetPageTitle sets the "page_title" field.
func (weu *WebsiteEventUpdate) SetPageTitle(s string) *WebsiteEventUpdate {
	weu.mutation.SetPageTitle(s)
	return weu
}

// SetNillablePageTitle sets the "page_title" field if the given value is not nil.
func (weu *WebsiteEventUpdate) SetNillablePageTitle(s *string) *WebsiteEventUpdate {
	if s != nil {
		weu.SetPageTitle(*s)
	}
	return weu
}

// ClearPageTitle clears the value of the "page_title" field.
func (weu *WebsiteEventUpdate) ClearPageTitle() *WebsiteEventUpdate {
	weu.mutation.ClearPageTitle()
	return weu
}

// SetPageData sets the "page_data" field.
func (weu *WebsiteEventUpdate) SetPageData(s string) *WebsiteEventUpdate {
	weu.mutation.SetPageData(s)
	return weu
}

// SetNillablePageData sets the "page_data" field if the given value is not nil.
func (weu *WebsiteEventUpdate) SetNillablePageData(s *string) *WebsiteEventUpdate {
	if s != nil {
		weu.SetPageData(*s)
	}
	return weu
}

// ClearPageData clears the value of the "page_data" field.
func (weu *WebsiteEventUpdate) ClearPageData() *WebsiteEventUpdate {
	weu.mutation.ClearPageData()
	return weu
}

// SetCreatedAt sets the "created_at" field.
func (weu *WebsiteEventUpdate) SetCreatedAt(t time.Time) *WebsiteEventUpdate {
	weu.mutation.SetCreatedAt(t)
	return weu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (weu *WebsiteEventUpdate) SetNillableCreatedAt(t *time.Time) *WebsiteEventUpdate {
	if t != nil {
		weu.SetCreatedAt(*t)
	}
	return weu
}

// SetWebsiteID sets the "website" edge to the Website entity by ID.
func (weu *WebsiteEventUpdate) SetWebsiteID(id uuid.UUID) *WebsiteEventUpdate {
	weu.mutation.SetWebsiteID(id)
	return weu
}

// SetNillableWebsiteID sets the "website" edge to the Website entity by ID if the given value is not nil.
func (weu *WebsiteEventUpdate) SetNillableWebsiteID(id *uuid.UUID) *WebsiteEventUpdate {
	if id != nil {
		weu = weu.SetWebsiteID(*id)
	}
	return weu
}

// SetWebsite sets the "website" edge to the Website entity.
func (weu *WebsiteEventUpdate) SetWebsite(w *Website) *WebsiteEventUpdate {
	return weu.SetWebsiteID(w.ID)
}

// Mutation returns the WebsiteEventMutation object of the builder.
func (weu *WebsiteEventUpdate) Mutation() *WebsiteEventMutation {
	return weu.mutation
}

// ClearWebsite clears the "website" edge to the Website entity.
func (weu *WebsiteEventUpdate) ClearWebsite() *WebsiteEventUpdate {
	weu.mutation.ClearWebsite()
	return weu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (weu *WebsiteEventUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, weu.sqlSave, weu.mutation, weu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (weu *WebsiteEventUpdate) SaveX(ctx context.Context) int {
	affected, err := weu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (weu *WebsiteEventUpdate) Exec(ctx context.Context) error {
	_, err := weu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (weu *WebsiteEventUpdate) ExecX(ctx context.Context) {
	if err := weu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (weu *WebsiteEventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(websiteevent.Table, websiteevent.Columns, sqlgraph.NewFieldSpec(websiteevent.FieldID, field.TypeUUID))
	if ps := weu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := weu.mutation.EventName(); ok {
		_spec.SetField(websiteevent.FieldEventName, field.TypeString, value)
	}
	if value, ok := weu.mutation.URLPath(); ok {
		_spec.SetField(websiteevent.FieldURLPath, field.TypeString, value)
	}
	if weu.mutation.URLPathCleared() {
		_spec.ClearField(websiteevent.FieldURLPath, field.TypeString)
	}
	if value, ok := weu.mutation.URLQuery(); ok {
		_spec.SetField(websiteevent.FieldURLQuery, field.TypeString, value)
	}
	if weu.mutation.URLQueryCleared() {
		_spec.ClearField(websiteevent.FieldURLQuery, field.TypeString)
	}
	if value, ok := weu.mutation.ReferrerPath(); ok {
		_spec.SetField(websiteevent.FieldReferrerPath, field.TypeString, value)
	}
	if weu.mutation.ReferrerPathCleared() {
		_spec.ClearField(websiteevent.FieldReferrerPath, field.TypeString)
	}
	if value, ok := weu.mutation.ReferrerQuery(); ok {
		_spec.SetField(websiteevent.FieldReferrerQuery, field.TypeString, value)
	}
	if weu.mutation.ReferrerQueryCleared() {
		_spec.ClearField(websiteevent.FieldReferrerQuery, field.TypeString)
	}
	if value, ok := weu.mutation.ReferrerDomain(); ok {
		_spec.SetField(websiteevent.FieldReferrerDomain, field.TypeString, value)
	}
	if weu.mutation.ReferrerDomainCleared() {
		_spec.ClearField(websiteevent.FieldReferrerDomain, field.TypeString)
	}
	if value, ok := weu.mutation.PageTitle(); ok {
		_spec.SetField(websiteevent.FieldPageTitle, field.TypeString, value)
	}
	if weu.mutation.PageTitleCleared() {
		_spec.ClearField(websiteevent.FieldPageTitle, field.TypeString)
	}
	if value, ok := weu.mutation.PageData(); ok {
		_spec.SetField(websiteevent.FieldPageData, field.TypeString, value)
	}
	if weu.mutation.PageDataCleared() {
		_spec.ClearField(websiteevent.FieldPageData, field.TypeString)
	}
	if value, ok := weu.mutation.CreatedAt(); ok {
		_spec.SetField(websiteevent.FieldCreatedAt, field.TypeTime, value)
	}
	if weu.mutation.WebsiteCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   websiteevent.WebsiteTable,
			Columns: []string{websiteevent.WebsiteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(website.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := weu.mutation.WebsiteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   websiteevent.WebsiteTable,
			Columns: []string{websiteevent.WebsiteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(website.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, weu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{websiteevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	weu.mutation.done = true
	return n, nil
}

// WebsiteEventUpdateOne is the builder for updating a single WebsiteEvent entity.
type WebsiteEventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WebsiteEventMutation
}

// SetEventName sets the "event_name" field.
func (weuo *WebsiteEventUpdateOne) SetEventName(s string) *WebsiteEventUpdateOne {
	weuo.mutation.SetEventName(s)
	return weuo
}

// SetURLPath sets the "url_path" field.
func (weuo *WebsiteEventUpdateOne) SetURLPath(s string) *WebsiteEventUpdateOne {
	weuo.mutation.SetURLPath(s)
	return weuo
}

// SetNillableURLPath sets the "url_path" field if the given value is not nil.
func (weuo *WebsiteEventUpdateOne) SetNillableURLPath(s *string) *WebsiteEventUpdateOne {
	if s != nil {
		weuo.SetURLPath(*s)
	}
	return weuo
}

// ClearURLPath clears the value of the "url_path" field.
func (weuo *WebsiteEventUpdateOne) ClearURLPath() *WebsiteEventUpdateOne {
	weuo.mutation.ClearURLPath()
	return weuo
}

// SetURLQuery sets the "url_query" field.
func (weuo *WebsiteEventUpdateOne) SetURLQuery(s string) *WebsiteEventUpdateOne {
	weuo.mutation.SetURLQuery(s)
	return weuo
}

// SetNillableURLQuery sets the "url_query" field if the given value is not nil.
func (weuo *WebsiteEventUpdateOne) SetNillableURLQuery(s *string) *WebsiteEventUpdateOne {
	if s != nil {
		weuo.SetURLQuery(*s)
	}
	return weuo
}

// ClearURLQuery clears the value of the "url_query" field.
func (weuo *WebsiteEventUpdateOne) ClearURLQuery() *WebsiteEventUpdateOne {
	weuo.mutation.ClearURLQuery()
	return weuo
}

// SetReferrerPath sets the "referrer_path" field.
func (weuo *WebsiteEventUpdateOne) SetReferrerPath(s string) *WebsiteEventUpdateOne {
	weuo.mutation.SetReferrerPath(s)
	return weuo
}

// SetNillableReferrerPath sets the "referrer_path" field if the given value is not nil.
func (weuo *WebsiteEventUpdateOne) SetNillableReferrerPath(s *string) *WebsiteEventUpdateOne {
	if s != nil {
		weuo.SetReferrerPath(*s)
	}
	return weuo
}

// ClearReferrerPath clears the value of the "referrer_path" field.
func (weuo *WebsiteEventUpdateOne) ClearReferrerPath() *WebsiteEventUpdateOne {
	weuo.mutation.ClearReferrerPath()
	return weuo
}

// SetReferrerQuery sets the "referrer_query" field.
func (weuo *WebsiteEventUpdateOne) SetReferrerQuery(s string) *WebsiteEventUpdateOne {
	weuo.mutation.SetReferrerQuery(s)
	return weuo
}

// SetNillableReferrerQuery sets the "referrer_query" field if the given value is not nil.
func (weuo *WebsiteEventUpdateOne) SetNillableReferrerQuery(s *string) *WebsiteEventUpdateOne {
	if s != nil {
		weuo.SetReferrerQuery(*s)
	}
	return weuo
}

// ClearReferrerQuery clears the value of the "referrer_query" field.
func (weuo *WebsiteEventUpdateOne) ClearReferrerQuery() *WebsiteEventUpdateOne {
	weuo.mutation.ClearReferrerQuery()
	return weuo
}

// SetReferrerDomain sets the "referrer_domain" field.
func (weuo *WebsiteEventUpdateOne) SetReferrerDomain(s string) *WebsiteEventUpdateOne {
	weuo.mutation.SetReferrerDomain(s)
	return weuo
}

// SetNillableReferrerDomain sets the "referrer_domain" field if the given value is not nil.
func (weuo *WebsiteEventUpdateOne) SetNillableReferrerDomain(s *string) *WebsiteEventUpdateOne {
	if s != nil {
		weuo.SetReferrerDomain(*s)
	}
	return weuo
}

// ClearReferrerDomain clears the value of the "referrer_domain" field.
func (weuo *WebsiteEventUpdateOne) ClearReferrerDomain() *WebsiteEventUpdateOne {
	weuo.mutation.ClearReferrerDomain()
	return weuo
}

// SetPageTitle sets the "page_title" field.
func (weuo *WebsiteEventUpdateOne) SetPageTitle(s string) *WebsiteEventUpdateOne {
	weuo.mutation.SetPageTitle(s)
	return weuo
}

// SetNillablePageTitle sets the "page_title" field if the given value is not nil.
func (weuo *WebsiteEventUpdateOne) SetNillablePageTitle(s *string) *WebsiteEventUpdateOne {
	if s != nil {
		weuo.SetPageTitle(*s)
	}
	return weuo
}

// ClearPageTitle clears the value of the "page_title" field.
func (weuo *WebsiteEventUpdateOne) ClearPageTitle() *WebsiteEventUpdateOne {
	weuo.mutation.ClearPageTitle()
	return weuo
}

// SetPageData sets the "page_data" field.
func (weuo *WebsiteEventUpdateOne) SetPageData(s string) *WebsiteEventUpdateOne {
	weuo.mutation.SetPageData(s)
	return weuo
}

// SetNillablePageData sets the "page_data" field if the given value is not nil.
func (weuo *WebsiteEventUpdateOne) SetNillablePageData(s *string) *WebsiteEventUpdateOne {
	if s != nil {
		weuo.SetPageData(*s)
	}
	return weuo
}

// ClearPageData clears the value of the "page_data" field.
func (weuo *WebsiteEventUpdateOne) ClearPageData() *WebsiteEventUpdateOne {
	weuo.mutation.ClearPageData()
	return weuo
}

// SetCreatedAt sets the "created_at" field.
func (weuo *WebsiteEventUpdateOne) SetCreatedAt(t time.Time) *WebsiteEventUpdateOne {
	weuo.mutation.SetCreatedAt(t)
	return weuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (weuo *WebsiteEventUpdateOne) SetNillableCreatedAt(t *time.Time) *WebsiteEventUpdateOne {
	if t != nil {
		weuo.SetCreatedAt(*t)
	}
	return weuo
}

// SetWebsiteID sets the "website" edge to the Website entity by ID.
func (weuo *WebsiteEventUpdateOne) SetWebsiteID(id uuid.UUID) *WebsiteEventUpdateOne {
	weuo.mutation.SetWebsiteID(id)
	return weuo
}

// SetNillableWebsiteID sets the "website" edge to the Website entity by ID if the given value is not nil.
func (weuo *WebsiteEventUpdateOne) SetNillableWebsiteID(id *uuid.UUID) *WebsiteEventUpdateOne {
	if id != nil {
		weuo = weuo.SetWebsiteID(*id)
	}
	return weuo
}

// SetWebsite sets the "website" edge to the Website entity.
func (weuo *WebsiteEventUpdateOne) SetWebsite(w *Website) *WebsiteEventUpdateOne {
	return weuo.SetWebsiteID(w.ID)
}

// Mutation returns the WebsiteEventMutation object of the builder.
func (weuo *WebsiteEventUpdateOne) Mutation() *WebsiteEventMutation {
	return weuo.mutation
}

// ClearWebsite clears the "website" edge to the Website entity.
func (weuo *WebsiteEventUpdateOne) ClearWebsite() *WebsiteEventUpdateOne {
	weuo.mutation.ClearWebsite()
	return weuo
}

// Where appends a list predicates to the WebsiteEventUpdate builder.
func (weuo *WebsiteEventUpdateOne) Where(ps ...predicate.WebsiteEvent) *WebsiteEventUpdateOne {
	weuo.mutation.Where(ps...)
	return weuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (weuo *WebsiteEventUpdateOne) Select(field string, fields ...string) *WebsiteEventUpdateOne {
	weuo.fields = append([]string{field}, fields...)
	return weuo
}

// Save executes the query and returns the updated WebsiteEvent entity.
func (weuo *WebsiteEventUpdateOne) Save(ctx context.Context) (*WebsiteEvent, error) {
	return withHooks(ctx, weuo.sqlSave, weuo.mutation, weuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (weuo *WebsiteEventUpdateOne) SaveX(ctx context.Context) *WebsiteEvent {
	node, err := weuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (weuo *WebsiteEventUpdateOne) Exec(ctx context.Context) error {
	_, err := weuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (weuo *WebsiteEventUpdateOne) ExecX(ctx context.Context) {
	if err := weuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (weuo *WebsiteEventUpdateOne) sqlSave(ctx context.Context) (_node *WebsiteEvent, err error) {
	_spec := sqlgraph.NewUpdateSpec(websiteevent.Table, websiteevent.Columns, sqlgraph.NewFieldSpec(websiteevent.FieldID, field.TypeUUID))
	id, ok := weuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WebsiteEvent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := weuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, websiteevent.FieldID)
		for _, f := range fields {
			if !websiteevent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != websiteevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := weuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := weuo.mutation.EventName(); ok {
		_spec.SetField(websiteevent.FieldEventName, field.TypeString, value)
	}
	if value, ok := weuo.mutation.URLPath(); ok {
		_spec.SetField(websiteevent.FieldURLPath, field.TypeString, value)
	}
	if weuo.mutation.URLPathCleared() {
		_spec.ClearField(websiteevent.FieldURLPath, field.TypeString)
	}
	if value, ok := weuo.mutation.URLQuery(); ok {
		_spec.SetField(websiteevent.FieldURLQuery, field.TypeString, value)
	}
	if weuo.mutation.URLQueryCleared() {
		_spec.ClearField(websiteevent.FieldURLQuery, field.TypeString)
	}
	if value, ok := weuo.mutation.ReferrerPath(); ok {
		_spec.SetField(websiteevent.FieldReferrerPath, field.TypeString, value)
	}
	if weuo.mutation.ReferrerPathCleared() {
		_spec.ClearField(websiteevent.FieldReferrerPath, field.TypeString)
	}
	if value, ok := weuo.mutation.ReferrerQuery(); ok {
		_spec.SetField(websiteevent.FieldReferrerQuery, field.TypeString, value)
	}
	if weuo.mutation.ReferrerQueryCleared() {
		_spec.ClearField(websiteevent.FieldReferrerQuery, field.TypeString)
	}
	if value, ok := weuo.mutation.ReferrerDomain(); ok {
		_spec.SetField(websiteevent.FieldReferrerDomain, field.TypeString, value)
	}
	if weuo.mutation.ReferrerDomainCleared() {
		_spec.ClearField(websiteevent.FieldReferrerDomain, field.TypeString)
	}
	if value, ok := weuo.mutation.PageTitle(); ok {
		_spec.SetField(websiteevent.FieldPageTitle, field.TypeString, value)
	}
	if weuo.mutation.PageTitleCleared() {
		_spec.ClearField(websiteevent.FieldPageTitle, field.TypeString)
	}
	if value, ok := weuo.mutation.PageData(); ok {
		_spec.SetField(websiteevent.FieldPageData, field.TypeString, value)
	}
	if weuo.mutation.PageDataCleared() {
		_spec.ClearField(websiteevent.FieldPageData, field.TypeString)
	}
	if value, ok := weuo.mutation.CreatedAt(); ok {
		_spec.SetField(websiteevent.FieldCreatedAt, field.TypeTime, value)
	}
	if weuo.mutation.WebsiteCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   websiteevent.WebsiteTable,
			Columns: []string{websiteevent.WebsiteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(website.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := weuo.mutation.WebsiteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   websiteevent.WebsiteTable,
			Columns: []string{websiteevent.WebsiteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(website.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &WebsiteEvent{config: weuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, weuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{websiteevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	weuo.mutation.done = true
	return _node, nil
}
