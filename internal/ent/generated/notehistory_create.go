// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/notehistory"
	"github.com/theopenlane/entx/history"
)

// NoteHistoryCreate is the builder for creating a NoteHistory entity.
type NoteHistoryCreate struct {
	config
	mutation *NoteHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (nhc *NoteHistoryCreate) SetHistoryTime(t time.Time) *NoteHistoryCreate {
	nhc.mutation.SetHistoryTime(t)
	return nhc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (nhc *NoteHistoryCreate) SetNillableHistoryTime(t *time.Time) *NoteHistoryCreate {
	if t != nil {
		nhc.SetHistoryTime(*t)
	}
	return nhc
}

// SetRef sets the "ref" field.
func (nhc *NoteHistoryCreate) SetRef(s string) *NoteHistoryCreate {
	nhc.mutation.SetRef(s)
	return nhc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (nhc *NoteHistoryCreate) SetNillableRef(s *string) *NoteHistoryCreate {
	if s != nil {
		nhc.SetRef(*s)
	}
	return nhc
}

// SetOperation sets the "operation" field.
func (nhc *NoteHistoryCreate) SetOperation(ht history.OpType) *NoteHistoryCreate {
	nhc.mutation.SetOperation(ht)
	return nhc
}

// SetCreatedAt sets the "created_at" field.
func (nhc *NoteHistoryCreate) SetCreatedAt(t time.Time) *NoteHistoryCreate {
	nhc.mutation.SetCreatedAt(t)
	return nhc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nhc *NoteHistoryCreate) SetNillableCreatedAt(t *time.Time) *NoteHistoryCreate {
	if t != nil {
		nhc.SetCreatedAt(*t)
	}
	return nhc
}

// SetUpdatedAt sets the "updated_at" field.
func (nhc *NoteHistoryCreate) SetUpdatedAt(t time.Time) *NoteHistoryCreate {
	nhc.mutation.SetUpdatedAt(t)
	return nhc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nhc *NoteHistoryCreate) SetNillableUpdatedAt(t *time.Time) *NoteHistoryCreate {
	if t != nil {
		nhc.SetUpdatedAt(*t)
	}
	return nhc
}

// SetCreatedBy sets the "created_by" field.
func (nhc *NoteHistoryCreate) SetCreatedBy(s string) *NoteHistoryCreate {
	nhc.mutation.SetCreatedBy(s)
	return nhc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (nhc *NoteHistoryCreate) SetNillableCreatedBy(s *string) *NoteHistoryCreate {
	if s != nil {
		nhc.SetCreatedBy(*s)
	}
	return nhc
}

// SetUpdatedBy sets the "updated_by" field.
func (nhc *NoteHistoryCreate) SetUpdatedBy(s string) *NoteHistoryCreate {
	nhc.mutation.SetUpdatedBy(s)
	return nhc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (nhc *NoteHistoryCreate) SetNillableUpdatedBy(s *string) *NoteHistoryCreate {
	if s != nil {
		nhc.SetUpdatedBy(*s)
	}
	return nhc
}

// SetMappingID sets the "mapping_id" field.
func (nhc *NoteHistoryCreate) SetMappingID(s string) *NoteHistoryCreate {
	nhc.mutation.SetMappingID(s)
	return nhc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (nhc *NoteHistoryCreate) SetNillableMappingID(s *string) *NoteHistoryCreate {
	if s != nil {
		nhc.SetMappingID(*s)
	}
	return nhc
}

// SetDeletedAt sets the "deleted_at" field.
func (nhc *NoteHistoryCreate) SetDeletedAt(t time.Time) *NoteHistoryCreate {
	nhc.mutation.SetDeletedAt(t)
	return nhc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (nhc *NoteHistoryCreate) SetNillableDeletedAt(t *time.Time) *NoteHistoryCreate {
	if t != nil {
		nhc.SetDeletedAt(*t)
	}
	return nhc
}

// SetDeletedBy sets the "deleted_by" field.
func (nhc *NoteHistoryCreate) SetDeletedBy(s string) *NoteHistoryCreate {
	nhc.mutation.SetDeletedBy(s)
	return nhc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (nhc *NoteHistoryCreate) SetNillableDeletedBy(s *string) *NoteHistoryCreate {
	if s != nil {
		nhc.SetDeletedBy(*s)
	}
	return nhc
}

// SetTags sets the "tags" field.
func (nhc *NoteHistoryCreate) SetTags(s []string) *NoteHistoryCreate {
	nhc.mutation.SetTags(s)
	return nhc
}

// SetOwnerID sets the "owner_id" field.
func (nhc *NoteHistoryCreate) SetOwnerID(s string) *NoteHistoryCreate {
	nhc.mutation.SetOwnerID(s)
	return nhc
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (nhc *NoteHistoryCreate) SetNillableOwnerID(s *string) *NoteHistoryCreate {
	if s != nil {
		nhc.SetOwnerID(*s)
	}
	return nhc
}

// SetText sets the "text" field.
func (nhc *NoteHistoryCreate) SetText(s string) *NoteHistoryCreate {
	nhc.mutation.SetText(s)
	return nhc
}

// SetID sets the "id" field.
func (nhc *NoteHistoryCreate) SetID(s string) *NoteHistoryCreate {
	nhc.mutation.SetID(s)
	return nhc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (nhc *NoteHistoryCreate) SetNillableID(s *string) *NoteHistoryCreate {
	if s != nil {
		nhc.SetID(*s)
	}
	return nhc
}

// Mutation returns the NoteHistoryMutation object of the builder.
func (nhc *NoteHistoryCreate) Mutation() *NoteHistoryMutation {
	return nhc.mutation
}

// Save creates the NoteHistory in the database.
func (nhc *NoteHistoryCreate) Save(ctx context.Context) (*NoteHistory, error) {
	if err := nhc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, nhc.sqlSave, nhc.mutation, nhc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nhc *NoteHistoryCreate) SaveX(ctx context.Context) *NoteHistory {
	v, err := nhc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nhc *NoteHistoryCreate) Exec(ctx context.Context) error {
	_, err := nhc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nhc *NoteHistoryCreate) ExecX(ctx context.Context) {
	if err := nhc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nhc *NoteHistoryCreate) defaults() error {
	if _, ok := nhc.mutation.HistoryTime(); !ok {
		if notehistory.DefaultHistoryTime == nil {
			return fmt.Errorf("generated: uninitialized notehistory.DefaultHistoryTime (forgotten import generated/runtime?)")
		}
		v := notehistory.DefaultHistoryTime()
		nhc.mutation.SetHistoryTime(v)
	}
	if _, ok := nhc.mutation.CreatedAt(); !ok {
		if notehistory.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized notehistory.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := notehistory.DefaultCreatedAt()
		nhc.mutation.SetCreatedAt(v)
	}
	if _, ok := nhc.mutation.UpdatedAt(); !ok {
		if notehistory.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized notehistory.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := notehistory.DefaultUpdatedAt()
		nhc.mutation.SetUpdatedAt(v)
	}
	if _, ok := nhc.mutation.MappingID(); !ok {
		if notehistory.DefaultMappingID == nil {
			return fmt.Errorf("generated: uninitialized notehistory.DefaultMappingID (forgotten import generated/runtime?)")
		}
		v := notehistory.DefaultMappingID()
		nhc.mutation.SetMappingID(v)
	}
	if _, ok := nhc.mutation.Tags(); !ok {
		v := notehistory.DefaultTags
		nhc.mutation.SetTags(v)
	}
	if _, ok := nhc.mutation.ID(); !ok {
		if notehistory.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized notehistory.DefaultID (forgotten import generated/runtime?)")
		}
		v := notehistory.DefaultID()
		nhc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (nhc *NoteHistoryCreate) check() error {
	if _, ok := nhc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`generated: missing required field "NoteHistory.history_time"`)}
	}
	if _, ok := nhc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`generated: missing required field "NoteHistory.operation"`)}
	}
	if v, ok := nhc.mutation.Operation(); ok {
		if err := notehistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`generated: validator failed for field "NoteHistory.operation": %w`, err)}
		}
	}
	if _, ok := nhc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "NoteHistory.mapping_id"`)}
	}
	if _, ok := nhc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`generated: missing required field "NoteHistory.text"`)}
	}
	return nil
}

func (nhc *NoteHistoryCreate) sqlSave(ctx context.Context) (*NoteHistory, error) {
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
			return nil, fmt.Errorf("unexpected NoteHistory.ID type: %T", _spec.ID.Value)
		}
	}
	nhc.mutation.id = &_node.ID
	nhc.mutation.done = true
	return _node, nil
}

func (nhc *NoteHistoryCreate) createSpec() (*NoteHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &NoteHistory{config: nhc.config}
		_spec = sqlgraph.NewCreateSpec(notehistory.Table, sqlgraph.NewFieldSpec(notehistory.FieldID, field.TypeString))
	)
	_spec.Schema = nhc.schemaConfig.NoteHistory
	if id, ok := nhc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := nhc.mutation.HistoryTime(); ok {
		_spec.SetField(notehistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := nhc.mutation.Ref(); ok {
		_spec.SetField(notehistory.FieldRef, field.TypeString, value)
		_node.Ref = value
	}
	if value, ok := nhc.mutation.Operation(); ok {
		_spec.SetField(notehistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := nhc.mutation.CreatedAt(); ok {
		_spec.SetField(notehistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := nhc.mutation.UpdatedAt(); ok {
		_spec.SetField(notehistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := nhc.mutation.CreatedBy(); ok {
		_spec.SetField(notehistory.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := nhc.mutation.UpdatedBy(); ok {
		_spec.SetField(notehistory.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := nhc.mutation.MappingID(); ok {
		_spec.SetField(notehistory.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := nhc.mutation.DeletedAt(); ok {
		_spec.SetField(notehistory.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := nhc.mutation.DeletedBy(); ok {
		_spec.SetField(notehistory.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := nhc.mutation.Tags(); ok {
		_spec.SetField(notehistory.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := nhc.mutation.OwnerID(); ok {
		_spec.SetField(notehistory.FieldOwnerID, field.TypeString, value)
		_node.OwnerID = value
	}
	if value, ok := nhc.mutation.Text(); ok {
		_spec.SetField(notehistory.FieldText, field.TypeString, value)
		_node.Text = value
	}
	return _node, _spec
}

// NoteHistoryCreateBulk is the builder for creating many NoteHistory entities in bulk.
type NoteHistoryCreateBulk struct {
	config
	err      error
	builders []*NoteHistoryCreate
}

// Save creates the NoteHistory entities in the database.
func (nhcb *NoteHistoryCreateBulk) Save(ctx context.Context) ([]*NoteHistory, error) {
	if nhcb.err != nil {
		return nil, nhcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(nhcb.builders))
	nodes := make([]*NoteHistory, len(nhcb.builders))
	mutators := make([]Mutator, len(nhcb.builders))
	for i := range nhcb.builders {
		func(i int, root context.Context) {
			builder := nhcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NoteHistoryMutation)
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
func (nhcb *NoteHistoryCreateBulk) SaveX(ctx context.Context) []*NoteHistory {
	v, err := nhcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nhcb *NoteHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := nhcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nhcb *NoteHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := nhcb.Exec(ctx); err != nil {
		panic(err)
	}
}