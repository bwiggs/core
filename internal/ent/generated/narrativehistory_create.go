// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/narrativehistory"
	"github.com/theopenlane/entx/history"
)

// NarrativeHistoryCreate is the builder for creating a NarrativeHistory entity.
type NarrativeHistoryCreate struct {
	config
	mutation *NarrativeHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (nhc *NarrativeHistoryCreate) SetHistoryTime(t time.Time) *NarrativeHistoryCreate {
	nhc.mutation.SetHistoryTime(t)
	return nhc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (nhc *NarrativeHistoryCreate) SetNillableHistoryTime(t *time.Time) *NarrativeHistoryCreate {
	if t != nil {
		nhc.SetHistoryTime(*t)
	}
	return nhc
}

// SetRef sets the "ref" field.
func (nhc *NarrativeHistoryCreate) SetRef(s string) *NarrativeHistoryCreate {
	nhc.mutation.SetRef(s)
	return nhc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (nhc *NarrativeHistoryCreate) SetNillableRef(s *string) *NarrativeHistoryCreate {
	if s != nil {
		nhc.SetRef(*s)
	}
	return nhc
}

// SetOperation sets the "operation" field.
func (nhc *NarrativeHistoryCreate) SetOperation(ht history.OpType) *NarrativeHistoryCreate {
	nhc.mutation.SetOperation(ht)
	return nhc
}

// SetCreatedAt sets the "created_at" field.
func (nhc *NarrativeHistoryCreate) SetCreatedAt(t time.Time) *NarrativeHistoryCreate {
	nhc.mutation.SetCreatedAt(t)
	return nhc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nhc *NarrativeHistoryCreate) SetNillableCreatedAt(t *time.Time) *NarrativeHistoryCreate {
	if t != nil {
		nhc.SetCreatedAt(*t)
	}
	return nhc
}

// SetUpdatedAt sets the "updated_at" field.
func (nhc *NarrativeHistoryCreate) SetUpdatedAt(t time.Time) *NarrativeHistoryCreate {
	nhc.mutation.SetUpdatedAt(t)
	return nhc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nhc *NarrativeHistoryCreate) SetNillableUpdatedAt(t *time.Time) *NarrativeHistoryCreate {
	if t != nil {
		nhc.SetUpdatedAt(*t)
	}
	return nhc
}

// SetCreatedBy sets the "created_by" field.
func (nhc *NarrativeHistoryCreate) SetCreatedBy(s string) *NarrativeHistoryCreate {
	nhc.mutation.SetCreatedBy(s)
	return nhc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (nhc *NarrativeHistoryCreate) SetNillableCreatedBy(s *string) *NarrativeHistoryCreate {
	if s != nil {
		nhc.SetCreatedBy(*s)
	}
	return nhc
}

// SetUpdatedBy sets the "updated_by" field.
func (nhc *NarrativeHistoryCreate) SetUpdatedBy(s string) *NarrativeHistoryCreate {
	nhc.mutation.SetUpdatedBy(s)
	return nhc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (nhc *NarrativeHistoryCreate) SetNillableUpdatedBy(s *string) *NarrativeHistoryCreate {
	if s != nil {
		nhc.SetUpdatedBy(*s)
	}
	return nhc
}

// SetDeletedAt sets the "deleted_at" field.
func (nhc *NarrativeHistoryCreate) SetDeletedAt(t time.Time) *NarrativeHistoryCreate {
	nhc.mutation.SetDeletedAt(t)
	return nhc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (nhc *NarrativeHistoryCreate) SetNillableDeletedAt(t *time.Time) *NarrativeHistoryCreate {
	if t != nil {
		nhc.SetDeletedAt(*t)
	}
	return nhc
}

// SetDeletedBy sets the "deleted_by" field.
func (nhc *NarrativeHistoryCreate) SetDeletedBy(s string) *NarrativeHistoryCreate {
	nhc.mutation.SetDeletedBy(s)
	return nhc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (nhc *NarrativeHistoryCreate) SetNillableDeletedBy(s *string) *NarrativeHistoryCreate {
	if s != nil {
		nhc.SetDeletedBy(*s)
	}
	return nhc
}

// SetMappingID sets the "mapping_id" field.
func (nhc *NarrativeHistoryCreate) SetMappingID(s string) *NarrativeHistoryCreate {
	nhc.mutation.SetMappingID(s)
	return nhc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (nhc *NarrativeHistoryCreate) SetNillableMappingID(s *string) *NarrativeHistoryCreate {
	if s != nil {
		nhc.SetMappingID(*s)
	}
	return nhc
}

// SetTags sets the "tags" field.
func (nhc *NarrativeHistoryCreate) SetTags(s []string) *NarrativeHistoryCreate {
	nhc.mutation.SetTags(s)
	return nhc
}

// SetName sets the "name" field.
func (nhc *NarrativeHistoryCreate) SetName(s string) *NarrativeHistoryCreate {
	nhc.mutation.SetName(s)
	return nhc
}

// SetDescription sets the "description" field.
func (nhc *NarrativeHistoryCreate) SetDescription(s string) *NarrativeHistoryCreate {
	nhc.mutation.SetDescription(s)
	return nhc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (nhc *NarrativeHistoryCreate) SetNillableDescription(s *string) *NarrativeHistoryCreate {
	if s != nil {
		nhc.SetDescription(*s)
	}
	return nhc
}

// SetSatisfies sets the "satisfies" field.
func (nhc *NarrativeHistoryCreate) SetSatisfies(s string) *NarrativeHistoryCreate {
	nhc.mutation.SetSatisfies(s)
	return nhc
}

// SetNillableSatisfies sets the "satisfies" field if the given value is not nil.
func (nhc *NarrativeHistoryCreate) SetNillableSatisfies(s *string) *NarrativeHistoryCreate {
	if s != nil {
		nhc.SetSatisfies(*s)
	}
	return nhc
}

// SetDetails sets the "details" field.
func (nhc *NarrativeHistoryCreate) SetDetails(m map[string]interface{}) *NarrativeHistoryCreate {
	nhc.mutation.SetDetails(m)
	return nhc
}

// SetID sets the "id" field.
func (nhc *NarrativeHistoryCreate) SetID(s string) *NarrativeHistoryCreate {
	nhc.mutation.SetID(s)
	return nhc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (nhc *NarrativeHistoryCreate) SetNillableID(s *string) *NarrativeHistoryCreate {
	if s != nil {
		nhc.SetID(*s)
	}
	return nhc
}

// Mutation returns the NarrativeHistoryMutation object of the builder.
func (nhc *NarrativeHistoryCreate) Mutation() *NarrativeHistoryMutation {
	return nhc.mutation
}

// Save creates the NarrativeHistory in the database.
func (nhc *NarrativeHistoryCreate) Save(ctx context.Context) (*NarrativeHistory, error) {
	nhc.defaults()
	return withHooks(ctx, nhc.sqlSave, nhc.mutation, nhc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nhc *NarrativeHistoryCreate) SaveX(ctx context.Context) *NarrativeHistory {
	v, err := nhc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nhc *NarrativeHistoryCreate) Exec(ctx context.Context) error {
	_, err := nhc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nhc *NarrativeHistoryCreate) ExecX(ctx context.Context) {
	if err := nhc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nhc *NarrativeHistoryCreate) defaults() {
	if _, ok := nhc.mutation.HistoryTime(); !ok {
		v := narrativehistory.DefaultHistoryTime()
		nhc.mutation.SetHistoryTime(v)
	}
	if _, ok := nhc.mutation.CreatedAt(); !ok {
		v := narrativehistory.DefaultCreatedAt()
		nhc.mutation.SetCreatedAt(v)
	}
	if _, ok := nhc.mutation.UpdatedAt(); !ok {
		v := narrativehistory.DefaultUpdatedAt()
		nhc.mutation.SetUpdatedAt(v)
	}
	if _, ok := nhc.mutation.MappingID(); !ok {
		v := narrativehistory.DefaultMappingID()
		nhc.mutation.SetMappingID(v)
	}
	if _, ok := nhc.mutation.Tags(); !ok {
		v := narrativehistory.DefaultTags
		nhc.mutation.SetTags(v)
	}
	if _, ok := nhc.mutation.ID(); !ok {
		v := narrativehistory.DefaultID()
		nhc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nhc *NarrativeHistoryCreate) check() error {
	if _, ok := nhc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`generated: missing required field "NarrativeHistory.history_time"`)}
	}
	if _, ok := nhc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`generated: missing required field "NarrativeHistory.operation"`)}
	}
	if v, ok := nhc.mutation.Operation(); ok {
		if err := narrativehistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`generated: validator failed for field "NarrativeHistory.operation": %w`, err)}
		}
	}
	if _, ok := nhc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "NarrativeHistory.mapping_id"`)}
	}
	if _, ok := nhc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "NarrativeHistory.name"`)}
	}
	return nil
}

func (nhc *NarrativeHistoryCreate) sqlSave(ctx context.Context) (*NarrativeHistory, error) {
	if err := nhc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nhc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nhc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected NarrativeHistory.ID type: %T", _spec.ID.Value)
		}
	}
	nhc.mutation.id = &_node.ID
	nhc.mutation.done = true
	return _node, nil
}

func (nhc *NarrativeHistoryCreate) createSpec() (*NarrativeHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &NarrativeHistory{config: nhc.config}
		_spec = sqlgraph.NewCreateSpec(narrativehistory.Table, sqlgraph.NewFieldSpec(narrativehistory.FieldID, field.TypeString))
	)
	_spec.Schema = nhc.schemaConfig.NarrativeHistory
	if id, ok := nhc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := nhc.mutation.HistoryTime(); ok {
		_spec.SetField(narrativehistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := nhc.mutation.Ref(); ok {
		_spec.SetField(narrativehistory.FieldRef, field.TypeString, value)
		_node.Ref = value
	}
	if value, ok := nhc.mutation.Operation(); ok {
		_spec.SetField(narrativehistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := nhc.mutation.CreatedAt(); ok {
		_spec.SetField(narrativehistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := nhc.mutation.UpdatedAt(); ok {
		_spec.SetField(narrativehistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := nhc.mutation.CreatedBy(); ok {
		_spec.SetField(narrativehistory.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := nhc.mutation.UpdatedBy(); ok {
		_spec.SetField(narrativehistory.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := nhc.mutation.DeletedAt(); ok {
		_spec.SetField(narrativehistory.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := nhc.mutation.DeletedBy(); ok {
		_spec.SetField(narrativehistory.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := nhc.mutation.MappingID(); ok {
		_spec.SetField(narrativehistory.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := nhc.mutation.Tags(); ok {
		_spec.SetField(narrativehistory.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := nhc.mutation.Name(); ok {
		_spec.SetField(narrativehistory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := nhc.mutation.Description(); ok {
		_spec.SetField(narrativehistory.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := nhc.mutation.Satisfies(); ok {
		_spec.SetField(narrativehistory.FieldSatisfies, field.TypeString, value)
		_node.Satisfies = value
	}
	if value, ok := nhc.mutation.Details(); ok {
		_spec.SetField(narrativehistory.FieldDetails, field.TypeJSON, value)
		_node.Details = value
	}
	return _node, _spec
}

// NarrativeHistoryCreateBulk is the builder for creating many NarrativeHistory entities in bulk.
type NarrativeHistoryCreateBulk struct {
	config
	err      error
	builders []*NarrativeHistoryCreate
}

// Save creates the NarrativeHistory entities in the database.
func (nhcb *NarrativeHistoryCreateBulk) Save(ctx context.Context) ([]*NarrativeHistory, error) {
	if nhcb.err != nil {
		return nil, nhcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(nhcb.builders))
	nodes := make([]*NarrativeHistory, len(nhcb.builders))
	mutators := make([]Mutator, len(nhcb.builders))
	for i := range nhcb.builders {
		func(i int, root context.Context) {
			builder := nhcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NarrativeHistoryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, nhcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, nhcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, nhcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (nhcb *NarrativeHistoryCreateBulk) SaveX(ctx context.Context) []*NarrativeHistory {
	v, err := nhcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nhcb *NarrativeHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := nhcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nhcb *NarrativeHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := nhcb.Exec(ctx); err != nil {
		panic(err)
	}
}