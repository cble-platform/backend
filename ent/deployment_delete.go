// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cble-platform/backend/ent/deployment"
	"github.com/cble-platform/backend/ent/predicate"
)

// DeploymentDelete is the builder for deleting a Deployment entity.
type DeploymentDelete struct {
	config
	hooks    []Hook
	mutation *DeploymentMutation
}

// Where appends a list predicates to the DeploymentDelete builder.
func (dd *DeploymentDelete) Where(ps ...predicate.Deployment) *DeploymentDelete {
	dd.mutation.Where(ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DeploymentDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dd.sqlExec, dd.mutation, dd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DeploymentDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DeploymentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(deployment.Table, sqlgraph.NewFieldSpec(deployment.FieldID, field.TypeUUID))
	if ps := dd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dd.mutation.done = true
	return affected, err
}

// DeploymentDeleteOne is the builder for deleting a single Deployment entity.
type DeploymentDeleteOne struct {
	dd *DeploymentDelete
}

// Where appends a list predicates to the DeploymentDelete builder.
func (ddo *DeploymentDeleteOne) Where(ps ...predicate.Deployment) *DeploymentDeleteOne {
	ddo.dd.mutation.Where(ps...)
	return ddo
}

// Exec executes the deletion query.
func (ddo *DeploymentDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{deployment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DeploymentDeleteOne) ExecX(ctx context.Context) {
	if err := ddo.Exec(ctx); err != nil {
		panic(err)
	}
}
