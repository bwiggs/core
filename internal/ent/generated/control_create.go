// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/actionplan"
	"github.com/theopenlane/core/internal/ent/generated/control"
	"github.com/theopenlane/core/internal/ent/generated/controlobjective"
	"github.com/theopenlane/core/internal/ent/generated/narrative"
	"github.com/theopenlane/core/internal/ent/generated/procedure"
	"github.com/theopenlane/core/internal/ent/generated/risk"
	"github.com/theopenlane/core/internal/ent/generated/standard"
	"github.com/theopenlane/core/internal/ent/generated/subcontrol"
)

// ControlCreate is the builder for creating a Control entity.
type ControlCreate struct {
	config
	mutation *ControlMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *ControlCreate) SetCreatedAt(t time.Time) *ControlCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *ControlCreate) SetNillableCreatedAt(t *time.Time) *ControlCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *ControlCreate) SetUpdatedAt(t time.Time) *ControlCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *ControlCreate) SetNillableUpdatedAt(t *time.Time) *ControlCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetCreatedBy sets the "created_by" field.
func (cc *ControlCreate) SetCreatedBy(s string) *ControlCreate {
	cc.mutation.SetCreatedBy(s)
	return cc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (cc *ControlCreate) SetNillableCreatedBy(s *string) *ControlCreate {
	if s != nil {
		cc.SetCreatedBy(*s)
	}
	return cc
}

// SetUpdatedBy sets the "updated_by" field.
func (cc *ControlCreate) SetUpdatedBy(s string) *ControlCreate {
	cc.mutation.SetUpdatedBy(s)
	return cc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cc *ControlCreate) SetNillableUpdatedBy(s *string) *ControlCreate {
	if s != nil {
		cc.SetUpdatedBy(*s)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *ControlCreate) SetDeletedAt(t time.Time) *ControlCreate {
	cc.mutation.SetDeletedAt(t)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *ControlCreate) SetNillableDeletedAt(t *time.Time) *ControlCreate {
	if t != nil {
		cc.SetDeletedAt(*t)
	}
	return cc
}

// SetDeletedBy sets the "deleted_by" field.
func (cc *ControlCreate) SetDeletedBy(s string) *ControlCreate {
	cc.mutation.SetDeletedBy(s)
	return cc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (cc *ControlCreate) SetNillableDeletedBy(s *string) *ControlCreate {
	if s != nil {
		cc.SetDeletedBy(*s)
	}
	return cc
}

// SetMappingID sets the "mapping_id" field.
func (cc *ControlCreate) SetMappingID(s string) *ControlCreate {
	cc.mutation.SetMappingID(s)
	return cc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (cc *ControlCreate) SetNillableMappingID(s *string) *ControlCreate {
	if s != nil {
		cc.SetMappingID(*s)
	}
	return cc
}

// SetTags sets the "tags" field.
func (cc *ControlCreate) SetTags(s []string) *ControlCreate {
	cc.mutation.SetTags(s)
	return cc
}

// SetName sets the "name" field.
func (cc *ControlCreate) SetName(s string) *ControlCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetDescription sets the "description" field.
func (cc *ControlCreate) SetDescription(s string) *ControlCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cc *ControlCreate) SetNillableDescription(s *string) *ControlCreate {
	if s != nil {
		cc.SetDescription(*s)
	}
	return cc
}

// SetStatus sets the "status" field.
func (cc *ControlCreate) SetStatus(s string) *ControlCreate {
	cc.mutation.SetStatus(s)
	return cc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cc *ControlCreate) SetNillableStatus(s *string) *ControlCreate {
	if s != nil {
		cc.SetStatus(*s)
	}
	return cc
}

// SetControlType sets the "control_type" field.
func (cc *ControlCreate) SetControlType(s string) *ControlCreate {
	cc.mutation.SetControlType(s)
	return cc
}

// SetNillableControlType sets the "control_type" field if the given value is not nil.
func (cc *ControlCreate) SetNillableControlType(s *string) *ControlCreate {
	if s != nil {
		cc.SetControlType(*s)
	}
	return cc
}

// SetVersion sets the "version" field.
func (cc *ControlCreate) SetVersion(s string) *ControlCreate {
	cc.mutation.SetVersion(s)
	return cc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (cc *ControlCreate) SetNillableVersion(s *string) *ControlCreate {
	if s != nil {
		cc.SetVersion(*s)
	}
	return cc
}

// SetControlNumber sets the "control_number" field.
func (cc *ControlCreate) SetControlNumber(s string) *ControlCreate {
	cc.mutation.SetControlNumber(s)
	return cc
}

// SetNillableControlNumber sets the "control_number" field if the given value is not nil.
func (cc *ControlCreate) SetNillableControlNumber(s *string) *ControlCreate {
	if s != nil {
		cc.SetControlNumber(*s)
	}
	return cc
}

// SetFamily sets the "family" field.
func (cc *ControlCreate) SetFamily(s string) *ControlCreate {
	cc.mutation.SetFamily(s)
	return cc
}

// SetNillableFamily sets the "family" field if the given value is not nil.
func (cc *ControlCreate) SetNillableFamily(s *string) *ControlCreate {
	if s != nil {
		cc.SetFamily(*s)
	}
	return cc
}

// SetClass sets the "class" field.
func (cc *ControlCreate) SetClass(s string) *ControlCreate {
	cc.mutation.SetClass(s)
	return cc
}

// SetNillableClass sets the "class" field if the given value is not nil.
func (cc *ControlCreate) SetNillableClass(s *string) *ControlCreate {
	if s != nil {
		cc.SetClass(*s)
	}
	return cc
}

// SetSource sets the "source" field.
func (cc *ControlCreate) SetSource(s string) *ControlCreate {
	cc.mutation.SetSource(s)
	return cc
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (cc *ControlCreate) SetNillableSource(s *string) *ControlCreate {
	if s != nil {
		cc.SetSource(*s)
	}
	return cc
}

// SetSatisfies sets the "satisfies" field.
func (cc *ControlCreate) SetSatisfies(s string) *ControlCreate {
	cc.mutation.SetSatisfies(s)
	return cc
}

// SetNillableSatisfies sets the "satisfies" field if the given value is not nil.
func (cc *ControlCreate) SetNillableSatisfies(s *string) *ControlCreate {
	if s != nil {
		cc.SetSatisfies(*s)
	}
	return cc
}

// SetMappedFrameworks sets the "mapped_frameworks" field.
func (cc *ControlCreate) SetMappedFrameworks(s string) *ControlCreate {
	cc.mutation.SetMappedFrameworks(s)
	return cc
}

// SetNillableMappedFrameworks sets the "mapped_frameworks" field if the given value is not nil.
func (cc *ControlCreate) SetNillableMappedFrameworks(s *string) *ControlCreate {
	if s != nil {
		cc.SetMappedFrameworks(*s)
	}
	return cc
}

// SetDetails sets the "details" field.
func (cc *ControlCreate) SetDetails(m map[string]interface{}) *ControlCreate {
	cc.mutation.SetDetails(m)
	return cc
}

// SetID sets the "id" field.
func (cc *ControlCreate) SetID(s string) *ControlCreate {
	cc.mutation.SetID(s)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *ControlCreate) SetNillableID(s *string) *ControlCreate {
	if s != nil {
		cc.SetID(*s)
	}
	return cc
}

// AddProcedureIDs adds the "procedures" edge to the Procedure entity by IDs.
func (cc *ControlCreate) AddProcedureIDs(ids ...string) *ControlCreate {
	cc.mutation.AddProcedureIDs(ids...)
	return cc
}

// AddProcedures adds the "procedures" edges to the Procedure entity.
func (cc *ControlCreate) AddProcedures(p ...*Procedure) *ControlCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddProcedureIDs(ids...)
}

// AddSubcontrolIDs adds the "subcontrols" edge to the Subcontrol entity by IDs.
func (cc *ControlCreate) AddSubcontrolIDs(ids ...string) *ControlCreate {
	cc.mutation.AddSubcontrolIDs(ids...)
	return cc
}

// AddSubcontrols adds the "subcontrols" edges to the Subcontrol entity.
func (cc *ControlCreate) AddSubcontrols(s ...*Subcontrol) *ControlCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cc.AddSubcontrolIDs(ids...)
}

// AddControlobjectiveIDs adds the "controlobjectives" edge to the ControlObjective entity by IDs.
func (cc *ControlCreate) AddControlobjectiveIDs(ids ...string) *ControlCreate {
	cc.mutation.AddControlobjectiveIDs(ids...)
	return cc
}

// AddControlobjectives adds the "controlobjectives" edges to the ControlObjective entity.
func (cc *ControlCreate) AddControlobjectives(c ...*ControlObjective) *ControlCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddControlobjectiveIDs(ids...)
}

// AddStandardIDs adds the "standard" edge to the Standard entity by IDs.
func (cc *ControlCreate) AddStandardIDs(ids ...string) *ControlCreate {
	cc.mutation.AddStandardIDs(ids...)
	return cc
}

// AddStandard adds the "standard" edges to the Standard entity.
func (cc *ControlCreate) AddStandard(s ...*Standard) *ControlCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cc.AddStandardIDs(ids...)
}

// AddNarrativeIDs adds the "narratives" edge to the Narrative entity by IDs.
func (cc *ControlCreate) AddNarrativeIDs(ids ...string) *ControlCreate {
	cc.mutation.AddNarrativeIDs(ids...)
	return cc
}

// AddNarratives adds the "narratives" edges to the Narrative entity.
func (cc *ControlCreate) AddNarratives(n ...*Narrative) *ControlCreate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return cc.AddNarrativeIDs(ids...)
}

// AddRiskIDs adds the "risks" edge to the Risk entity by IDs.
func (cc *ControlCreate) AddRiskIDs(ids ...string) *ControlCreate {
	cc.mutation.AddRiskIDs(ids...)
	return cc
}

// AddRisks adds the "risks" edges to the Risk entity.
func (cc *ControlCreate) AddRisks(r ...*Risk) *ControlCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cc.AddRiskIDs(ids...)
}

// AddActionplanIDs adds the "actionplans" edge to the ActionPlan entity by IDs.
func (cc *ControlCreate) AddActionplanIDs(ids ...string) *ControlCreate {
	cc.mutation.AddActionplanIDs(ids...)
	return cc
}

// AddActionplans adds the "actionplans" edges to the ActionPlan entity.
func (cc *ControlCreate) AddActionplans(a ...*ActionPlan) *ControlCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cc.AddActionplanIDs(ids...)
}

// Mutation returns the ControlMutation object of the builder.
func (cc *ControlCreate) Mutation() *ControlMutation {
	return cc.mutation
}

// Save creates the Control in the database.
func (cc *ControlCreate) Save(ctx context.Context) (*Control, error) {
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ControlCreate) SaveX(ctx context.Context) *Control {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ControlCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ControlCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ControlCreate) defaults() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		if control.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized control.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := control.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		if control.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized control.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := control.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.MappingID(); !ok {
		if control.DefaultMappingID == nil {
			return fmt.Errorf("generated: uninitialized control.DefaultMappingID (forgotten import generated/runtime?)")
		}
		v := control.DefaultMappingID()
		cc.mutation.SetMappingID(v)
	}
	if _, ok := cc.mutation.Tags(); !ok {
		v := control.DefaultTags
		cc.mutation.SetTags(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		if control.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized control.DefaultID (forgotten import generated/runtime?)")
		}
		v := control.DefaultID()
		cc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *ControlCreate) check() error {
	if _, ok := cc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "Control.mapping_id"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "Control.name"`)}
	}
	return nil
}

func (cc *ControlCreate) sqlSave(ctx context.Context) (*Control, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Control.ID type: %T", _spec.ID.Value)
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ControlCreate) createSpec() (*Control, *sqlgraph.CreateSpec) {
	var (
		_node = &Control{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(control.Table, sqlgraph.NewFieldSpec(control.FieldID, field.TypeString))
	)
	_spec.Schema = cc.schemaConfig.Control
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(control.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(control.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.CreatedBy(); ok {
		_spec.SetField(control.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := cc.mutation.UpdatedBy(); ok {
		_spec.SetField(control.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.SetField(control.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := cc.mutation.DeletedBy(); ok {
		_spec.SetField(control.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := cc.mutation.MappingID(); ok {
		_spec.SetField(control.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := cc.mutation.Tags(); ok {
		_spec.SetField(control.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(control.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.SetField(control.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.SetField(control.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := cc.mutation.ControlType(); ok {
		_spec.SetField(control.FieldControlType, field.TypeString, value)
		_node.ControlType = value
	}
	if value, ok := cc.mutation.Version(); ok {
		_spec.SetField(control.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := cc.mutation.ControlNumber(); ok {
		_spec.SetField(control.FieldControlNumber, field.TypeString, value)
		_node.ControlNumber = value
	}
	if value, ok := cc.mutation.Family(); ok {
		_spec.SetField(control.FieldFamily, field.TypeString, value)
		_node.Family = value
	}
	if value, ok := cc.mutation.Class(); ok {
		_spec.SetField(control.FieldClass, field.TypeString, value)
		_node.Class = value
	}
	if value, ok := cc.mutation.Source(); ok {
		_spec.SetField(control.FieldSource, field.TypeString, value)
		_node.Source = value
	}
	if value, ok := cc.mutation.Satisfies(); ok {
		_spec.SetField(control.FieldSatisfies, field.TypeString, value)
		_node.Satisfies = value
	}
	if value, ok := cc.mutation.MappedFrameworks(); ok {
		_spec.SetField(control.FieldMappedFrameworks, field.TypeString, value)
		_node.MappedFrameworks = value
	}
	if value, ok := cc.mutation.Details(); ok {
		_spec.SetField(control.FieldDetails, field.TypeJSON, value)
		_node.Details = value
	}
	if nodes := cc.mutation.ProceduresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   control.ProceduresTable,
			Columns: control.ProceduresPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(procedure.FieldID, field.TypeString),
			},
		}
		edge.Schema = cc.schemaConfig.ControlProcedures
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.SubcontrolsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   control.SubcontrolsTable,
			Columns: control.SubcontrolsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subcontrol.FieldID, field.TypeString),
			},
		}
		edge.Schema = cc.schemaConfig.ControlSubcontrols
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ControlobjectivesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   control.ControlobjectivesTable,
			Columns: []string{control.ControlobjectivesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(controlobjective.FieldID, field.TypeString),
			},
		}
		edge.Schema = cc.schemaConfig.ControlObjective
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.StandardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   control.StandardTable,
			Columns: control.StandardPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(standard.FieldID, field.TypeString),
			},
		}
		edge.Schema = cc.schemaConfig.StandardControls
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.NarrativesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   control.NarrativesTable,
			Columns: control.NarrativesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(narrative.FieldID, field.TypeString),
			},
		}
		edge.Schema = cc.schemaConfig.ControlNarratives
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.RisksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   control.RisksTable,
			Columns: control.RisksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(risk.FieldID, field.TypeString),
			},
		}
		edge.Schema = cc.schemaConfig.ControlRisks
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ActionplansIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   control.ActionplansTable,
			Columns: control.ActionplansPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(actionplan.FieldID, field.TypeString),
			},
		}
		edge.Schema = cc.schemaConfig.ControlActionplans
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ControlCreateBulk is the builder for creating many Control entities in bulk.
type ControlCreateBulk struct {
	config
	err      error
	builders []*ControlCreate
}

// Save creates the Control entities in the database.
func (ccb *ControlCreateBulk) Save(ctx context.Context) ([]*Control, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Control, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ControlMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ControlCreateBulk) SaveX(ctx context.Context) []*Control {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ControlCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ControlCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}