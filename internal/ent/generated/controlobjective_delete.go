// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/predicate"

	"github.com/theopenlane/core/internal/ent/generated/controlobjective"
	"github.com/theopenlane/core/internal/ent/generated/internal"
)

// ControlObjectiveDelete is the builder for deleting a ControlObjective entity.
type ControlObjectiveDelete struct {
	config
	hooks    []Hook
	mutation *ControlObjectiveMutation
}

// Where appends a list predicates to the ControlObjectiveDelete builder.
func (cod *ControlObjectiveDelete) Where(ps ...predicate.ControlObjective) *ControlObjectiveDelete {
	cod.mutation.Where(ps...)
	return cod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cod *ControlObjectiveDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cod.sqlExec, cod.mutation, cod.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cod *ControlObjectiveDelete) ExecX(ctx context.Context) int {
	n, err := cod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cod *ControlObjectiveDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(controlobjective.Table, sqlgraph.NewFieldSpec(controlobjective.FieldID, field.TypeString))
	_spec.Node.Schema = cod.schemaConfig.ControlObjective
	ctx = internal.NewSchemaConfigContext(ctx, cod.schemaConfig)
	if ps := cod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cod.mutation.done = true
	return affected, err
}

// ControlObjectiveDeleteOne is the builder for deleting a single ControlObjective entity.
type ControlObjectiveDeleteOne struct {
	cod *ControlObjectiveDelete
}

// Where appends a list predicates to the ControlObjectiveDelete builder.
func (codo *ControlObjectiveDeleteOne) Where(ps ...predicate.ControlObjective) *ControlObjectiveDeleteOne {
	codo.cod.mutation.Where(ps...)
	return codo
}

// Exec executes the deletion query.
func (codo *ControlObjectiveDeleteOne) Exec(ctx context.Context) error {
	n, err := codo.cod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{controlobjective.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (codo *ControlObjectiveDeleteOne) ExecX(ctx context.Context) {
	if err := codo.Exec(ctx); err != nil {
		panic(err)
	}
}