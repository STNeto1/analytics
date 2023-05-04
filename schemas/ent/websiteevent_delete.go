// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_schemas/ent/predicate"
	"_schemas/ent/websiteevent"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WebsiteEventDelete is the builder for deleting a WebsiteEvent entity.
type WebsiteEventDelete struct {
	config
	hooks    []Hook
	mutation *WebsiteEventMutation
}

// Where appends a list predicates to the WebsiteEventDelete builder.
func (wed *WebsiteEventDelete) Where(ps ...predicate.WebsiteEvent) *WebsiteEventDelete {
	wed.mutation.Where(ps...)
	return wed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wed *WebsiteEventDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wed.sqlExec, wed.mutation, wed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wed *WebsiteEventDelete) ExecX(ctx context.Context) int {
	n, err := wed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wed *WebsiteEventDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(websiteevent.Table, sqlgraph.NewFieldSpec(websiteevent.FieldID, field.TypeUUID))
	if ps := wed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wed.mutation.done = true
	return affected, err
}

// WebsiteEventDeleteOne is the builder for deleting a single WebsiteEvent entity.
type WebsiteEventDeleteOne struct {
	wed *WebsiteEventDelete
}

// Where appends a list predicates to the WebsiteEventDelete builder.
func (wedo *WebsiteEventDeleteOne) Where(ps ...predicate.WebsiteEvent) *WebsiteEventDeleteOne {
	wedo.wed.mutation.Where(ps...)
	return wedo
}

// Exec executes the deletion query.
func (wedo *WebsiteEventDeleteOne) Exec(ctx context.Context) error {
	n, err := wedo.wed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{websiteevent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wedo *WebsiteEventDeleteOne) ExecX(ctx context.Context) {
	if err := wedo.Exec(ctx); err != nil {
		panic(err)
	}
}
