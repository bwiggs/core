// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/orgmembershiphistory"
	"github.com/theopenlane/core/pkg/enums"
	"github.com/theopenlane/entx/history"
)

// OrgMembershipHistoryCreate is the builder for creating a OrgMembershipHistory entity.
type OrgMembershipHistoryCreate struct {
	config
	mutation *OrgMembershipHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (omhc *OrgMembershipHistoryCreate) SetHistoryTime(t time.Time) *OrgMembershipHistoryCreate {
	omhc.mutation.SetHistoryTime(t)
	return omhc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (omhc *OrgMembershipHistoryCreate) SetNillableHistoryTime(t *time.Time) *OrgMembershipHistoryCreate {
	if t != nil {
		omhc.SetHistoryTime(*t)
	}
	return omhc
}

// SetRef sets the "ref" field.
func (omhc *OrgMembershipHistoryCreate) SetRef(s string) *OrgMembershipHistoryCreate {
	omhc.mutation.SetRef(s)
	return omhc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (omhc *OrgMembershipHistoryCreate) SetNillableRef(s *string) *OrgMembershipHistoryCreate {
	if s != nil {
		omhc.SetRef(*s)
	}
	return omhc
}

// SetOperation sets the "operation" field.
func (omhc *OrgMembershipHistoryCreate) SetOperation(ht history.OpType) *OrgMembershipHistoryCreate {
	omhc.mutation.SetOperation(ht)
	return omhc
}

// SetCreatedAt sets the "created_at" field.
func (omhc *OrgMembershipHistoryCreate) SetCreatedAt(t time.Time) *OrgMembershipHistoryCreate {
	omhc.mutation.SetCreatedAt(t)
	return omhc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (omhc *OrgMembershipHistoryCreate) SetNillableCreatedAt(t *time.Time) *OrgMembershipHistoryCreate {
	if t != nil {
		omhc.SetCreatedAt(*t)
	}
	return omhc
}

// SetUpdatedAt sets the "updated_at" field.
func (omhc *OrgMembershipHistoryCreate) SetUpdatedAt(t time.Time) *OrgMembershipHistoryCreate {
	omhc.mutation.SetUpdatedAt(t)
	return omhc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (omhc *OrgMembershipHistoryCreate) SetNillableUpdatedAt(t *time.Time) *OrgMembershipHistoryCreate {
	if t != nil {
		omhc.SetUpdatedAt(*t)
	}
	return omhc
}

// SetCreatedBy sets the "created_by" field.
func (omhc *OrgMembershipHistoryCreate) SetCreatedBy(s string) *OrgMembershipHistoryCreate {
	omhc.mutation.SetCreatedBy(s)
	return omhc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (omhc *OrgMembershipHistoryCreate) SetNillableCreatedBy(s *string) *OrgMembershipHistoryCreate {
	if s != nil {
		omhc.SetCreatedBy(*s)
	}
	return omhc
}

// SetUpdatedBy sets the "updated_by" field.
func (omhc *OrgMembershipHistoryCreate) SetUpdatedBy(s string) *OrgMembershipHistoryCreate {
	omhc.mutation.SetUpdatedBy(s)
	return omhc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (omhc *OrgMembershipHistoryCreate) SetNillableUpdatedBy(s *string) *OrgMembershipHistoryCreate {
	if s != nil {
		omhc.SetUpdatedBy(*s)
	}
	return omhc
}

// SetMappingID sets the "mapping_id" field.
func (omhc *OrgMembershipHistoryCreate) SetMappingID(s string) *OrgMembershipHistoryCreate {
	omhc.mutation.SetMappingID(s)
	return omhc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (omhc *OrgMembershipHistoryCreate) SetNillableMappingID(s *string) *OrgMembershipHistoryCreate {
	if s != nil {
		omhc.SetMappingID(*s)
	}
	return omhc
}

// SetDeletedAt sets the "deleted_at" field.
func (omhc *OrgMembershipHistoryCreate) SetDeletedAt(t time.Time) *OrgMembershipHistoryCreate {
	omhc.mutation.SetDeletedAt(t)
	return omhc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (omhc *OrgMembershipHistoryCreate) SetNillableDeletedAt(t *time.Time) *OrgMembershipHistoryCreate {
	if t != nil {
		omhc.SetDeletedAt(*t)
	}
	return omhc
}

// SetDeletedBy sets the "deleted_by" field.
func (omhc *OrgMembershipHistoryCreate) SetDeletedBy(s string) *OrgMembershipHistoryCreate {
	omhc.mutation.SetDeletedBy(s)
	return omhc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (omhc *OrgMembershipHistoryCreate) SetNillableDeletedBy(s *string) *OrgMembershipHistoryCreate {
	if s != nil {
		omhc.SetDeletedBy(*s)
	}
	return omhc
}

// SetRole sets the "role" field.
func (omhc *OrgMembershipHistoryCreate) SetRole(e enums.Role) *OrgMembershipHistoryCreate {
	omhc.mutation.SetRole(e)
	return omhc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (omhc *OrgMembershipHistoryCreate) SetNillableRole(e *enums.Role) *OrgMembershipHistoryCreate {
	if e != nil {
		omhc.SetRole(*e)
	}
	return omhc
}

// SetOrganizationID sets the "organization_id" field.
func (omhc *OrgMembershipHistoryCreate) SetOrganizationID(s string) *OrgMembershipHistoryCreate {
	omhc.mutation.SetOrganizationID(s)
	return omhc
}

// SetUserID sets the "user_id" field.
func (omhc *OrgMembershipHistoryCreate) SetUserID(s string) *OrgMembershipHistoryCreate {
	omhc.mutation.SetUserID(s)
	return omhc
}

// SetID sets the "id" field.
func (omhc *OrgMembershipHistoryCreate) SetID(s string) *OrgMembershipHistoryCreate {
	omhc.mutation.SetID(s)
	return omhc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (omhc *OrgMembershipHistoryCreate) SetNillableID(s *string) *OrgMembershipHistoryCreate {
	if s != nil {
		omhc.SetID(*s)
	}
	return omhc
}

// Mutation returns the OrgMembershipHistoryMutation object of the builder.
func (omhc *OrgMembershipHistoryCreate) Mutation() *OrgMembershipHistoryMutation {
	return omhc.mutation
}

// Save creates the OrgMembershipHistory in the database.
func (omhc *OrgMembershipHistoryCreate) Save(ctx context.Context) (*OrgMembershipHistory, error) {
	if err := omhc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, omhc.sqlSave, omhc.mutation, omhc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (omhc *OrgMembershipHistoryCreate) SaveX(ctx context.Context) *OrgMembershipHistory {
	v, err := omhc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (omhc *OrgMembershipHistoryCreate) Exec(ctx context.Context) error {
	_, err := omhc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (omhc *OrgMembershipHistoryCreate) ExecX(ctx context.Context) {
	if err := omhc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (omhc *OrgMembershipHistoryCreate) defaults() error {
	if _, ok := omhc.mutation.HistoryTime(); !ok {
		if orgmembershiphistory.DefaultHistoryTime == nil {
			return fmt.Errorf("generated: uninitialized orgmembershiphistory.DefaultHistoryTime (forgotten import generated/runtime?)")
		}
		v := orgmembershiphistory.DefaultHistoryTime()
		omhc.mutation.SetHistoryTime(v)
	}
	if _, ok := omhc.mutation.CreatedAt(); !ok {
		if orgmembershiphistory.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized orgmembershiphistory.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := orgmembershiphistory.DefaultCreatedAt()
		omhc.mutation.SetCreatedAt(v)
	}
	if _, ok := omhc.mutation.UpdatedAt(); !ok {
		if orgmembershiphistory.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized orgmembershiphistory.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := orgmembershiphistory.DefaultUpdatedAt()
		omhc.mutation.SetUpdatedAt(v)
	}
	if _, ok := omhc.mutation.MappingID(); !ok {
		if orgmembershiphistory.DefaultMappingID == nil {
			return fmt.Errorf("generated: uninitialized orgmembershiphistory.DefaultMappingID (forgotten import generated/runtime?)")
		}
		v := orgmembershiphistory.DefaultMappingID()
		omhc.mutation.SetMappingID(v)
	}
	if _, ok := omhc.mutation.Role(); !ok {
		v := orgmembershiphistory.DefaultRole
		omhc.mutation.SetRole(v)
	}
	if _, ok := omhc.mutation.ID(); !ok {
		if orgmembershiphistory.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized orgmembershiphistory.DefaultID (forgotten import generated/runtime?)")
		}
		v := orgmembershiphistory.DefaultID()
		omhc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (omhc *OrgMembershipHistoryCreate) check() error {
	if _, ok := omhc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`generated: missing required field "OrgMembershipHistory.history_time"`)}
	}
	if _, ok := omhc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`generated: missing required field "OrgMembershipHistory.operation"`)}
	}
	if v, ok := omhc.mutation.Operation(); ok {
		if err := orgmembershiphistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`generated: validator failed for field "OrgMembershipHistory.operation": %w`, err)}
		}
	}
	if _, ok := omhc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "OrgMembershipHistory.mapping_id"`)}
	}
	if _, ok := omhc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`generated: missing required field "OrgMembershipHistory.role"`)}
	}
	if v, ok := omhc.mutation.Role(); ok {
		if err := orgmembershiphistory.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`generated: validator failed for field "OrgMembershipHistory.role": %w`, err)}
		}
	}
	if _, ok := omhc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`generated: missing required field "OrgMembershipHistory.organization_id"`)}
	}
	if _, ok := omhc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`generated: missing required field "OrgMembershipHistory.user_id"`)}
	}
	return nil
}

func (omhc *OrgMembershipHistoryCreate) sqlSave(ctx context.Context) (*OrgMembershipHistory, error) {
	if err := omhc.check(); err != nil {
		return nil, err
	}
	_node, _spec := omhc.createSpec()
	if err := sqlgraph.CreateNode(ctx, omhc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected OrgMembershipHistory.ID type: %T", _spec.ID.Value)
		}
	}
	omhc.mutation.id = &_node.ID
	omhc.mutation.done = true
	return _node, nil
}

func (omhc *OrgMembershipHistoryCreate) createSpec() (*OrgMembershipHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &OrgMembershipHistory{config: omhc.config}
		_spec = sqlgraph.NewCreateSpec(orgmembershiphistory.Table, sqlgraph.NewFieldSpec(orgmembershiphistory.FieldID, field.TypeString))
	)
	_spec.Schema = omhc.schemaConfig.OrgMembershipHistory
	if id, ok := omhc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := omhc.mutation.HistoryTime(); ok {
		_spec.SetField(orgmembershiphistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := omhc.mutation.Ref(); ok {
		_spec.SetField(orgmembershiphistory.FieldRef, field.TypeString, value)
		_node.Ref = value
	}
	if value, ok := omhc.mutation.Operation(); ok {
		_spec.SetField(orgmembershiphistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := omhc.mutation.CreatedAt(); ok {
		_spec.SetField(orgmembershiphistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := omhc.mutation.UpdatedAt(); ok {
		_spec.SetField(orgmembershiphistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := omhc.mutation.CreatedBy(); ok {
		_spec.SetField(orgmembershiphistory.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := omhc.mutation.UpdatedBy(); ok {
		_spec.SetField(orgmembershiphistory.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := omhc.mutation.MappingID(); ok {
		_spec.SetField(orgmembershiphistory.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := omhc.mutation.DeletedAt(); ok {
		_spec.SetField(orgmembershiphistory.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := omhc.mutation.DeletedBy(); ok {
		_spec.SetField(orgmembershiphistory.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := omhc.mutation.Role(); ok {
		_spec.SetField(orgmembershiphistory.FieldRole, field.TypeEnum, value)
		_node.Role = value
	}
	if value, ok := omhc.mutation.OrganizationID(); ok {
		_spec.SetField(orgmembershiphistory.FieldOrganizationID, field.TypeString, value)
		_node.OrganizationID = value
	}
	if value, ok := omhc.mutation.UserID(); ok {
		_spec.SetField(orgmembershiphistory.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	return _node, _spec
}

// OrgMembershipHistoryCreateBulk is the builder for creating many OrgMembershipHistory entities in bulk.
type OrgMembershipHistoryCreateBulk struct {
	config
	err      error
	builders []*OrgMembershipHistoryCreate
}

// Save creates the OrgMembershipHistory entities in the database.
func (omhcb *OrgMembershipHistoryCreateBulk) Save(ctx context.Context) ([]*OrgMembershipHistory, error) {
	if omhcb.err != nil {
		return nil, omhcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(omhcb.builders))
	nodes := make([]*OrgMembershipHistory, len(omhcb.builders))
	mutators := make([]Mutator, len(omhcb.builders))
	for i := range omhcb.builders {
		func(i int, root context.Context) {
			builder := omhcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrgMembershipHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, omhcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, omhcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, omhcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (omhcb *OrgMembershipHistoryCreateBulk) SaveX(ctx context.Context) []*OrgMembershipHistory {
	v, err := omhcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (omhcb *OrgMembershipHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := omhcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (omhcb *OrgMembershipHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := omhcb.Exec(ctx); err != nil {
		panic(err)
	}
}