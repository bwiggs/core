// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/event"
	"github.com/theopenlane/core/internal/ent/generated/hush"
	"github.com/theopenlane/core/internal/ent/generated/integration"
	"github.com/theopenlane/core/internal/ent/generated/organization"
	"github.com/theopenlane/core/internal/ent/generated/predicate"

	"github.com/theopenlane/core/internal/ent/generated/internal"
)

// HushUpdate is the builder for updating Hush entities.
type HushUpdate struct {
	config
	hooks    []Hook
	mutation *HushMutation
}

// Where appends a list predicates to the HushUpdate builder.
func (hu *HushUpdate) Where(ps ...predicate.Hush) *HushUpdate {
	hu.mutation.Where(ps...)
	return hu
}

// SetUpdatedAt sets the "updated_at" field.
func (hu *HushUpdate) SetUpdatedAt(t time.Time) *HushUpdate {
	hu.mutation.SetUpdatedAt(t)
	return hu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (hu *HushUpdate) ClearUpdatedAt() *HushUpdate {
	hu.mutation.ClearUpdatedAt()
	return hu
}

// SetUpdatedBy sets the "updated_by" field.
func (hu *HushUpdate) SetUpdatedBy(s string) *HushUpdate {
	hu.mutation.SetUpdatedBy(s)
	return hu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (hu *HushUpdate) SetNillableUpdatedBy(s *string) *HushUpdate {
	if s != nil {
		hu.SetUpdatedBy(*s)
	}
	return hu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (hu *HushUpdate) ClearUpdatedBy() *HushUpdate {
	hu.mutation.ClearUpdatedBy()
	return hu
}

// SetDeletedAt sets the "deleted_at" field.
func (hu *HushUpdate) SetDeletedAt(t time.Time) *HushUpdate {
	hu.mutation.SetDeletedAt(t)
	return hu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (hu *HushUpdate) SetNillableDeletedAt(t *time.Time) *HushUpdate {
	if t != nil {
		hu.SetDeletedAt(*t)
	}
	return hu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (hu *HushUpdate) ClearDeletedAt() *HushUpdate {
	hu.mutation.ClearDeletedAt()
	return hu
}

// SetDeletedBy sets the "deleted_by" field.
func (hu *HushUpdate) SetDeletedBy(s string) *HushUpdate {
	hu.mutation.SetDeletedBy(s)
	return hu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (hu *HushUpdate) SetNillableDeletedBy(s *string) *HushUpdate {
	if s != nil {
		hu.SetDeletedBy(*s)
	}
	return hu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (hu *HushUpdate) ClearDeletedBy() *HushUpdate {
	hu.mutation.ClearDeletedBy()
	return hu
}

// SetName sets the "name" field.
func (hu *HushUpdate) SetName(s string) *HushUpdate {
	hu.mutation.SetName(s)
	return hu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (hu *HushUpdate) SetNillableName(s *string) *HushUpdate {
	if s != nil {
		hu.SetName(*s)
	}
	return hu
}

// SetDescription sets the "description" field.
func (hu *HushUpdate) SetDescription(s string) *HushUpdate {
	hu.mutation.SetDescription(s)
	return hu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (hu *HushUpdate) SetNillableDescription(s *string) *HushUpdate {
	if s != nil {
		hu.SetDescription(*s)
	}
	return hu
}

// ClearDescription clears the value of the "description" field.
func (hu *HushUpdate) ClearDescription() *HushUpdate {
	hu.mutation.ClearDescription()
	return hu
}

// SetKind sets the "kind" field.
func (hu *HushUpdate) SetKind(s string) *HushUpdate {
	hu.mutation.SetKind(s)
	return hu
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (hu *HushUpdate) SetNillableKind(s *string) *HushUpdate {
	if s != nil {
		hu.SetKind(*s)
	}
	return hu
}

// ClearKind clears the value of the "kind" field.
func (hu *HushUpdate) ClearKind() *HushUpdate {
	hu.mutation.ClearKind()
	return hu
}

// AddIntegrationIDs adds the "integrations" edge to the Integration entity by IDs.
func (hu *HushUpdate) AddIntegrationIDs(ids ...string) *HushUpdate {
	hu.mutation.AddIntegrationIDs(ids...)
	return hu
}

// AddIntegrations adds the "integrations" edges to the Integration entity.
func (hu *HushUpdate) AddIntegrations(i ...*Integration) *HushUpdate {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return hu.AddIntegrationIDs(ids...)
}

// AddOrganizationIDs adds the "organization" edge to the Organization entity by IDs.
func (hu *HushUpdate) AddOrganizationIDs(ids ...string) *HushUpdate {
	hu.mutation.AddOrganizationIDs(ids...)
	return hu
}

// AddOrganization adds the "organization" edges to the Organization entity.
func (hu *HushUpdate) AddOrganization(o ...*Organization) *HushUpdate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return hu.AddOrganizationIDs(ids...)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (hu *HushUpdate) AddEventIDs(ids ...string) *HushUpdate {
	hu.mutation.AddEventIDs(ids...)
	return hu
}

// AddEvents adds the "events" edges to the Event entity.
func (hu *HushUpdate) AddEvents(e ...*Event) *HushUpdate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return hu.AddEventIDs(ids...)
}

// Mutation returns the HushMutation object of the builder.
func (hu *HushUpdate) Mutation() *HushMutation {
	return hu.mutation
}

// ClearIntegrations clears all "integrations" edges to the Integration entity.
func (hu *HushUpdate) ClearIntegrations() *HushUpdate {
	hu.mutation.ClearIntegrations()
	return hu
}

// RemoveIntegrationIDs removes the "integrations" edge to Integration entities by IDs.
func (hu *HushUpdate) RemoveIntegrationIDs(ids ...string) *HushUpdate {
	hu.mutation.RemoveIntegrationIDs(ids...)
	return hu
}

// RemoveIntegrations removes "integrations" edges to Integration entities.
func (hu *HushUpdate) RemoveIntegrations(i ...*Integration) *HushUpdate {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return hu.RemoveIntegrationIDs(ids...)
}

// ClearOrganization clears all "organization" edges to the Organization entity.
func (hu *HushUpdate) ClearOrganization() *HushUpdate {
	hu.mutation.ClearOrganization()
	return hu
}

// RemoveOrganizationIDs removes the "organization" edge to Organization entities by IDs.
func (hu *HushUpdate) RemoveOrganizationIDs(ids ...string) *HushUpdate {
	hu.mutation.RemoveOrganizationIDs(ids...)
	return hu
}

// RemoveOrganization removes "organization" edges to Organization entities.
func (hu *HushUpdate) RemoveOrganization(o ...*Organization) *HushUpdate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return hu.RemoveOrganizationIDs(ids...)
}

// ClearEvents clears all "events" edges to the Event entity.
func (hu *HushUpdate) ClearEvents() *HushUpdate {
	hu.mutation.ClearEvents()
	return hu
}

// RemoveEventIDs removes the "events" edge to Event entities by IDs.
func (hu *HushUpdate) RemoveEventIDs(ids ...string) *HushUpdate {
	hu.mutation.RemoveEventIDs(ids...)
	return hu
}

// RemoveEvents removes "events" edges to Event entities.
func (hu *HushUpdate) RemoveEvents(e ...*Event) *HushUpdate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return hu.RemoveEventIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HushUpdate) Save(ctx context.Context) (int, error) {
	if err := hu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, hu.sqlSave, hu.mutation, hu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HushUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HushUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HushUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hu *HushUpdate) defaults() error {
	if _, ok := hu.mutation.UpdatedAt(); !ok && !hu.mutation.UpdatedAtCleared() {
		if hush.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized hush.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := hush.UpdateDefaultUpdatedAt()
		hu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (hu *HushUpdate) check() error {
	if v, ok := hu.mutation.Name(); ok {
		if err := hush.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Hush.name": %w`, err)}
		}
	}
	return nil
}

func (hu *HushUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := hu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(hush.Table, hush.Columns, sqlgraph.NewFieldSpec(hush.FieldID, field.TypeString))
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if hu.mutation.CreatedAtCleared() {
		_spec.ClearField(hush.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := hu.mutation.UpdatedAt(); ok {
		_spec.SetField(hush.FieldUpdatedAt, field.TypeTime, value)
	}
	if hu.mutation.UpdatedAtCleared() {
		_spec.ClearField(hush.FieldUpdatedAt, field.TypeTime)
	}
	if hu.mutation.CreatedByCleared() {
		_spec.ClearField(hush.FieldCreatedBy, field.TypeString)
	}
	if value, ok := hu.mutation.UpdatedBy(); ok {
		_spec.SetField(hush.FieldUpdatedBy, field.TypeString, value)
	}
	if hu.mutation.UpdatedByCleared() {
		_spec.ClearField(hush.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := hu.mutation.DeletedAt(); ok {
		_spec.SetField(hush.FieldDeletedAt, field.TypeTime, value)
	}
	if hu.mutation.DeletedAtCleared() {
		_spec.ClearField(hush.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := hu.mutation.DeletedBy(); ok {
		_spec.SetField(hush.FieldDeletedBy, field.TypeString, value)
	}
	if hu.mutation.DeletedByCleared() {
		_spec.ClearField(hush.FieldDeletedBy, field.TypeString)
	}
	if value, ok := hu.mutation.Name(); ok {
		_spec.SetField(hush.FieldName, field.TypeString, value)
	}
	if value, ok := hu.mutation.Description(); ok {
		_spec.SetField(hush.FieldDescription, field.TypeString, value)
	}
	if hu.mutation.DescriptionCleared() {
		_spec.ClearField(hush.FieldDescription, field.TypeString)
	}
	if value, ok := hu.mutation.Kind(); ok {
		_spec.SetField(hush.FieldKind, field.TypeString, value)
	}
	if hu.mutation.KindCleared() {
		_spec.ClearField(hush.FieldKind, field.TypeString)
	}
	if hu.mutation.SecretNameCleared() {
		_spec.ClearField(hush.FieldSecretName, field.TypeString)
	}
	if hu.mutation.SecretValueCleared() {
		_spec.ClearField(hush.FieldSecretValue, field.TypeString)
	}
	if hu.mutation.IntegrationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hush.IntegrationsTable,
			Columns: hush.IntegrationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integration.FieldID, field.TypeString),
			},
		}
		edge.Schema = hu.schemaConfig.IntegrationSecrets
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.RemovedIntegrationsIDs(); len(nodes) > 0 && !hu.mutation.IntegrationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hush.IntegrationsTable,
			Columns: hush.IntegrationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integration.FieldID, field.TypeString),
			},
		}
		edge.Schema = hu.schemaConfig.IntegrationSecrets
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.IntegrationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hush.IntegrationsTable,
			Columns: hush.IntegrationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integration.FieldID, field.TypeString),
			},
		}
		edge.Schema = hu.schemaConfig.IntegrationSecrets
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hush.OrganizationTable,
			Columns: hush.OrganizationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = hu.schemaConfig.OrganizationSecrets
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.RemovedOrganizationIDs(); len(nodes) > 0 && !hu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hush.OrganizationTable,
			Columns: hush.OrganizationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = hu.schemaConfig.OrganizationSecrets
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hush.OrganizationTable,
			Columns: hush.OrganizationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = hu.schemaConfig.OrganizationSecrets
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if hu.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   hush.EventsTable,
			Columns: hush.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		edge.Schema = hu.schemaConfig.HushEvents
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.RemovedEventsIDs(); len(nodes) > 0 && !hu.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   hush.EventsTable,
			Columns: hush.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		edge.Schema = hu.schemaConfig.HushEvents
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hu.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   hush.EventsTable,
			Columns: hush.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		edge.Schema = hu.schemaConfig.HushEvents
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = hu.schemaConfig.Hush
	ctx = internal.NewSchemaConfigContext(ctx, hu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hush.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hu.mutation.done = true
	return n, nil
}

// HushUpdateOne is the builder for updating a single Hush entity.
type HushUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HushMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (huo *HushUpdateOne) SetUpdatedAt(t time.Time) *HushUpdateOne {
	huo.mutation.SetUpdatedAt(t)
	return huo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (huo *HushUpdateOne) ClearUpdatedAt() *HushUpdateOne {
	huo.mutation.ClearUpdatedAt()
	return huo
}

// SetUpdatedBy sets the "updated_by" field.
func (huo *HushUpdateOne) SetUpdatedBy(s string) *HushUpdateOne {
	huo.mutation.SetUpdatedBy(s)
	return huo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (huo *HushUpdateOne) SetNillableUpdatedBy(s *string) *HushUpdateOne {
	if s != nil {
		huo.SetUpdatedBy(*s)
	}
	return huo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (huo *HushUpdateOne) ClearUpdatedBy() *HushUpdateOne {
	huo.mutation.ClearUpdatedBy()
	return huo
}

// SetDeletedAt sets the "deleted_at" field.
func (huo *HushUpdateOne) SetDeletedAt(t time.Time) *HushUpdateOne {
	huo.mutation.SetDeletedAt(t)
	return huo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (huo *HushUpdateOne) SetNillableDeletedAt(t *time.Time) *HushUpdateOne {
	if t != nil {
		huo.SetDeletedAt(*t)
	}
	return huo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (huo *HushUpdateOne) ClearDeletedAt() *HushUpdateOne {
	huo.mutation.ClearDeletedAt()
	return huo
}

// SetDeletedBy sets the "deleted_by" field.
func (huo *HushUpdateOne) SetDeletedBy(s string) *HushUpdateOne {
	huo.mutation.SetDeletedBy(s)
	return huo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (huo *HushUpdateOne) SetNillableDeletedBy(s *string) *HushUpdateOne {
	if s != nil {
		huo.SetDeletedBy(*s)
	}
	return huo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (huo *HushUpdateOne) ClearDeletedBy() *HushUpdateOne {
	huo.mutation.ClearDeletedBy()
	return huo
}

// SetName sets the "name" field.
func (huo *HushUpdateOne) SetName(s string) *HushUpdateOne {
	huo.mutation.SetName(s)
	return huo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (huo *HushUpdateOne) SetNillableName(s *string) *HushUpdateOne {
	if s != nil {
		huo.SetName(*s)
	}
	return huo
}

// SetDescription sets the "description" field.
func (huo *HushUpdateOne) SetDescription(s string) *HushUpdateOne {
	huo.mutation.SetDescription(s)
	return huo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (huo *HushUpdateOne) SetNillableDescription(s *string) *HushUpdateOne {
	if s != nil {
		huo.SetDescription(*s)
	}
	return huo
}

// ClearDescription clears the value of the "description" field.
func (huo *HushUpdateOne) ClearDescription() *HushUpdateOne {
	huo.mutation.ClearDescription()
	return huo
}

// SetKind sets the "kind" field.
func (huo *HushUpdateOne) SetKind(s string) *HushUpdateOne {
	huo.mutation.SetKind(s)
	return huo
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (huo *HushUpdateOne) SetNillableKind(s *string) *HushUpdateOne {
	if s != nil {
		huo.SetKind(*s)
	}
	return huo
}

// ClearKind clears the value of the "kind" field.
func (huo *HushUpdateOne) ClearKind() *HushUpdateOne {
	huo.mutation.ClearKind()
	return huo
}

// AddIntegrationIDs adds the "integrations" edge to the Integration entity by IDs.
func (huo *HushUpdateOne) AddIntegrationIDs(ids ...string) *HushUpdateOne {
	huo.mutation.AddIntegrationIDs(ids...)
	return huo
}

// AddIntegrations adds the "integrations" edges to the Integration entity.
func (huo *HushUpdateOne) AddIntegrations(i ...*Integration) *HushUpdateOne {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return huo.AddIntegrationIDs(ids...)
}

// AddOrganizationIDs adds the "organization" edge to the Organization entity by IDs.
func (huo *HushUpdateOne) AddOrganizationIDs(ids ...string) *HushUpdateOne {
	huo.mutation.AddOrganizationIDs(ids...)
	return huo
}

// AddOrganization adds the "organization" edges to the Organization entity.
func (huo *HushUpdateOne) AddOrganization(o ...*Organization) *HushUpdateOne {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return huo.AddOrganizationIDs(ids...)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (huo *HushUpdateOne) AddEventIDs(ids ...string) *HushUpdateOne {
	huo.mutation.AddEventIDs(ids...)
	return huo
}

// AddEvents adds the "events" edges to the Event entity.
func (huo *HushUpdateOne) AddEvents(e ...*Event) *HushUpdateOne {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return huo.AddEventIDs(ids...)
}

// Mutation returns the HushMutation object of the builder.
func (huo *HushUpdateOne) Mutation() *HushMutation {
	return huo.mutation
}

// ClearIntegrations clears all "integrations" edges to the Integration entity.
func (huo *HushUpdateOne) ClearIntegrations() *HushUpdateOne {
	huo.mutation.ClearIntegrations()
	return huo
}

// RemoveIntegrationIDs removes the "integrations" edge to Integration entities by IDs.
func (huo *HushUpdateOne) RemoveIntegrationIDs(ids ...string) *HushUpdateOne {
	huo.mutation.RemoveIntegrationIDs(ids...)
	return huo
}

// RemoveIntegrations removes "integrations" edges to Integration entities.
func (huo *HushUpdateOne) RemoveIntegrations(i ...*Integration) *HushUpdateOne {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return huo.RemoveIntegrationIDs(ids...)
}

// ClearOrganization clears all "organization" edges to the Organization entity.
func (huo *HushUpdateOne) ClearOrganization() *HushUpdateOne {
	huo.mutation.ClearOrganization()
	return huo
}

// RemoveOrganizationIDs removes the "organization" edge to Organization entities by IDs.
func (huo *HushUpdateOne) RemoveOrganizationIDs(ids ...string) *HushUpdateOne {
	huo.mutation.RemoveOrganizationIDs(ids...)
	return huo
}

// RemoveOrganization removes "organization" edges to Organization entities.
func (huo *HushUpdateOne) RemoveOrganization(o ...*Organization) *HushUpdateOne {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return huo.RemoveOrganizationIDs(ids...)
}

// ClearEvents clears all "events" edges to the Event entity.
func (huo *HushUpdateOne) ClearEvents() *HushUpdateOne {
	huo.mutation.ClearEvents()
	return huo
}

// RemoveEventIDs removes the "events" edge to Event entities by IDs.
func (huo *HushUpdateOne) RemoveEventIDs(ids ...string) *HushUpdateOne {
	huo.mutation.RemoveEventIDs(ids...)
	return huo
}

// RemoveEvents removes "events" edges to Event entities.
func (huo *HushUpdateOne) RemoveEvents(e ...*Event) *HushUpdateOne {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return huo.RemoveEventIDs(ids...)
}

// Where appends a list predicates to the HushUpdate builder.
func (huo *HushUpdateOne) Where(ps ...predicate.Hush) *HushUpdateOne {
	huo.mutation.Where(ps...)
	return huo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (huo *HushUpdateOne) Select(field string, fields ...string) *HushUpdateOne {
	huo.fields = append([]string{field}, fields...)
	return huo
}

// Save executes the query and returns the updated Hush entity.
func (huo *HushUpdateOne) Save(ctx context.Context) (*Hush, error) {
	if err := huo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, huo.sqlSave, huo.mutation, huo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HushUpdateOne) SaveX(ctx context.Context) *Hush {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HushUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HushUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (huo *HushUpdateOne) defaults() error {
	if _, ok := huo.mutation.UpdatedAt(); !ok && !huo.mutation.UpdatedAtCleared() {
		if hush.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized hush.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := hush.UpdateDefaultUpdatedAt()
		huo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (huo *HushUpdateOne) check() error {
	if v, ok := huo.mutation.Name(); ok {
		if err := hush.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Hush.name": %w`, err)}
		}
	}
	return nil
}

func (huo *HushUpdateOne) sqlSave(ctx context.Context) (_node *Hush, err error) {
	if err := huo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(hush.Table, hush.Columns, sqlgraph.NewFieldSpec(hush.FieldID, field.TypeString))
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Hush.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := huo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hush.FieldID)
		for _, f := range fields {
			if !hush.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != hush.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := huo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if huo.mutation.CreatedAtCleared() {
		_spec.ClearField(hush.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := huo.mutation.UpdatedAt(); ok {
		_spec.SetField(hush.FieldUpdatedAt, field.TypeTime, value)
	}
	if huo.mutation.UpdatedAtCleared() {
		_spec.ClearField(hush.FieldUpdatedAt, field.TypeTime)
	}
	if huo.mutation.CreatedByCleared() {
		_spec.ClearField(hush.FieldCreatedBy, field.TypeString)
	}
	if value, ok := huo.mutation.UpdatedBy(); ok {
		_spec.SetField(hush.FieldUpdatedBy, field.TypeString, value)
	}
	if huo.mutation.UpdatedByCleared() {
		_spec.ClearField(hush.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := huo.mutation.DeletedAt(); ok {
		_spec.SetField(hush.FieldDeletedAt, field.TypeTime, value)
	}
	if huo.mutation.DeletedAtCleared() {
		_spec.ClearField(hush.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := huo.mutation.DeletedBy(); ok {
		_spec.SetField(hush.FieldDeletedBy, field.TypeString, value)
	}
	if huo.mutation.DeletedByCleared() {
		_spec.ClearField(hush.FieldDeletedBy, field.TypeString)
	}
	if value, ok := huo.mutation.Name(); ok {
		_spec.SetField(hush.FieldName, field.TypeString, value)
	}
	if value, ok := huo.mutation.Description(); ok {
		_spec.SetField(hush.FieldDescription, field.TypeString, value)
	}
	if huo.mutation.DescriptionCleared() {
		_spec.ClearField(hush.FieldDescription, field.TypeString)
	}
	if value, ok := huo.mutation.Kind(); ok {
		_spec.SetField(hush.FieldKind, field.TypeString, value)
	}
	if huo.mutation.KindCleared() {
		_spec.ClearField(hush.FieldKind, field.TypeString)
	}
	if huo.mutation.SecretNameCleared() {
		_spec.ClearField(hush.FieldSecretName, field.TypeString)
	}
	if huo.mutation.SecretValueCleared() {
		_spec.ClearField(hush.FieldSecretValue, field.TypeString)
	}
	if huo.mutation.IntegrationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hush.IntegrationsTable,
			Columns: hush.IntegrationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integration.FieldID, field.TypeString),
			},
		}
		edge.Schema = huo.schemaConfig.IntegrationSecrets
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.RemovedIntegrationsIDs(); len(nodes) > 0 && !huo.mutation.IntegrationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hush.IntegrationsTable,
			Columns: hush.IntegrationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integration.FieldID, field.TypeString),
			},
		}
		edge.Schema = huo.schemaConfig.IntegrationSecrets
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.IntegrationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hush.IntegrationsTable,
			Columns: hush.IntegrationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integration.FieldID, field.TypeString),
			},
		}
		edge.Schema = huo.schemaConfig.IntegrationSecrets
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if huo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hush.OrganizationTable,
			Columns: hush.OrganizationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = huo.schemaConfig.OrganizationSecrets
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.RemovedOrganizationIDs(); len(nodes) > 0 && !huo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hush.OrganizationTable,
			Columns: hush.OrganizationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = huo.schemaConfig.OrganizationSecrets
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   hush.OrganizationTable,
			Columns: hush.OrganizationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = huo.schemaConfig.OrganizationSecrets
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if huo.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   hush.EventsTable,
			Columns: hush.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		edge.Schema = huo.schemaConfig.HushEvents
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.RemovedEventsIDs(); len(nodes) > 0 && !huo.mutation.EventsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   hush.EventsTable,
			Columns: hush.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		edge.Schema = huo.schemaConfig.HushEvents
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := huo.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   hush.EventsTable,
			Columns: hush.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		edge.Schema = huo.schemaConfig.HushEvents
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = huo.schemaConfig.Hush
	ctx = internal.NewSchemaConfigContext(ctx, huo.schemaConfig)
	_node = &Hush{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hush.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	huo.mutation.done = true
	return _node, nil
}