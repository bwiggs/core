// Code generated by ent, DO NOT EDIT.

package programmembership

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/theopenlane/core/internal/ent/generated/predicate"
	"github.com/theopenlane/core/pkg/enums"

	"github.com/theopenlane/core/internal/ent/generated/internal"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldUpdatedBy, v))
}

// MappingID applies equality check predicate on the "mapping_id" field. It's identical to MappingIDEQ.
func MappingID(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldMappingID, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedBy applies equality check predicate on the "deleted_by" field. It's identical to DeletedByEQ.
func DeletedBy(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldDeletedBy, v))
}

// ProgramID applies equality check predicate on the "program_id" field. It's identical to ProgramIDEQ.
func ProgramID(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldProgramID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldUserID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNotNull(FieldUpdatedAt))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// MappingIDEQ applies the EQ predicate on the "mapping_id" field.
func MappingIDEQ(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldMappingID, v))
}

// MappingIDNEQ applies the NEQ predicate on the "mapping_id" field.
func MappingIDNEQ(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNEQ(FieldMappingID, v))
}

// MappingIDIn applies the In predicate on the "mapping_id" field.
func MappingIDIn(vs ...string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldIn(FieldMappingID, vs...))
}

// MappingIDNotIn applies the NotIn predicate on the "mapping_id" field.
func MappingIDNotIn(vs ...string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNotIn(FieldMappingID, vs...))
}

// MappingIDGT applies the GT predicate on the "mapping_id" field.
func MappingIDGT(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGT(FieldMappingID, v))
}

// MappingIDGTE applies the GTE predicate on the "mapping_id" field.
func MappingIDGTE(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGTE(FieldMappingID, v))
}

// MappingIDLT applies the LT predicate on the "mapping_id" field.
func MappingIDLT(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLT(FieldMappingID, v))
}

// MappingIDLTE applies the LTE predicate on the "mapping_id" field.
func MappingIDLTE(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLTE(FieldMappingID, v))
}

// MappingIDContains applies the Contains predicate on the "mapping_id" field.
func MappingIDContains(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldContains(FieldMappingID, v))
}

// MappingIDHasPrefix applies the HasPrefix predicate on the "mapping_id" field.
func MappingIDHasPrefix(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldHasPrefix(FieldMappingID, v))
}

// MappingIDHasSuffix applies the HasSuffix predicate on the "mapping_id" field.
func MappingIDHasSuffix(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldHasSuffix(FieldMappingID, v))
}

// MappingIDEqualFold applies the EqualFold predicate on the "mapping_id" field.
func MappingIDEqualFold(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEqualFold(FieldMappingID, v))
}

// MappingIDContainsFold applies the ContainsFold predicate on the "mapping_id" field.
func MappingIDContainsFold(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldContainsFold(FieldMappingID, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNotNull(FieldDeletedAt))
}

// DeletedByEQ applies the EQ predicate on the "deleted_by" field.
func DeletedByEQ(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "deleted_by" field.
func DeletedByNEQ(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "deleted_by" field.
func DeletedByIn(vs ...string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "deleted_by" field.
func DeletedByNotIn(vs ...string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "deleted_by" field.
func DeletedByGT(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "deleted_by" field.
func DeletedByGTE(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "deleted_by" field.
func DeletedByLT(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "deleted_by" field.
func DeletedByLTE(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLTE(FieldDeletedBy, v))
}

// DeletedByContains applies the Contains predicate on the "deleted_by" field.
func DeletedByContains(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldContains(FieldDeletedBy, v))
}

// DeletedByHasPrefix applies the HasPrefix predicate on the "deleted_by" field.
func DeletedByHasPrefix(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldHasPrefix(FieldDeletedBy, v))
}

// DeletedByHasSuffix applies the HasSuffix predicate on the "deleted_by" field.
func DeletedByHasSuffix(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldHasSuffix(FieldDeletedBy, v))
}

// DeletedByIsNil applies the IsNil predicate on the "deleted_by" field.
func DeletedByIsNil() predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldIsNull(FieldDeletedBy))
}

// DeletedByNotNil applies the NotNil predicate on the "deleted_by" field.
func DeletedByNotNil() predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNotNull(FieldDeletedBy))
}

// DeletedByEqualFold applies the EqualFold predicate on the "deleted_by" field.
func DeletedByEqualFold(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEqualFold(FieldDeletedBy, v))
}

// DeletedByContainsFold applies the ContainsFold predicate on the "deleted_by" field.
func DeletedByContainsFold(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldContainsFold(FieldDeletedBy, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v enums.Role) predicate.ProgramMembership {
	vc := v
	return predicate.ProgramMembership(sql.FieldEQ(FieldRole, vc))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v enums.Role) predicate.ProgramMembership {
	vc := v
	return predicate.ProgramMembership(sql.FieldNEQ(FieldRole, vc))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...enums.Role) predicate.ProgramMembership {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProgramMembership(sql.FieldIn(FieldRole, v...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...enums.Role) predicate.ProgramMembership {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProgramMembership(sql.FieldNotIn(FieldRole, v...))
}

// ProgramIDEQ applies the EQ predicate on the "program_id" field.
func ProgramIDEQ(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldProgramID, v))
}

// ProgramIDNEQ applies the NEQ predicate on the "program_id" field.
func ProgramIDNEQ(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNEQ(FieldProgramID, v))
}

// ProgramIDIn applies the In predicate on the "program_id" field.
func ProgramIDIn(vs ...string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldIn(FieldProgramID, vs...))
}

// ProgramIDNotIn applies the NotIn predicate on the "program_id" field.
func ProgramIDNotIn(vs ...string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNotIn(FieldProgramID, vs...))
}

// ProgramIDGT applies the GT predicate on the "program_id" field.
func ProgramIDGT(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGT(FieldProgramID, v))
}

// ProgramIDGTE applies the GTE predicate on the "program_id" field.
func ProgramIDGTE(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGTE(FieldProgramID, v))
}

// ProgramIDLT applies the LT predicate on the "program_id" field.
func ProgramIDLT(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLT(FieldProgramID, v))
}

// ProgramIDLTE applies the LTE predicate on the "program_id" field.
func ProgramIDLTE(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLTE(FieldProgramID, v))
}

// ProgramIDContains applies the Contains predicate on the "program_id" field.
func ProgramIDContains(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldContains(FieldProgramID, v))
}

// ProgramIDHasPrefix applies the HasPrefix predicate on the "program_id" field.
func ProgramIDHasPrefix(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldHasPrefix(FieldProgramID, v))
}

// ProgramIDHasSuffix applies the HasSuffix predicate on the "program_id" field.
func ProgramIDHasSuffix(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldHasSuffix(FieldProgramID, v))
}

// ProgramIDEqualFold applies the EqualFold predicate on the "program_id" field.
func ProgramIDEqualFold(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEqualFold(FieldProgramID, v))
}

// ProgramIDContainsFold applies the ContainsFold predicate on the "program_id" field.
func ProgramIDContainsFold(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldContainsFold(FieldProgramID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.FieldContainsFold(FieldUserID, v))
}

// HasProgram applies the HasEdge predicate on the "program" edge.
func HasProgram() predicate.ProgramMembership {
	return predicate.ProgramMembership(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ProgramTable, ProgramColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Program
		step.Edge.Schema = schemaConfig.ProgramMembership
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProgramWith applies the HasEdge predicate on the "program" edge with a given conditions (other predicates).
func HasProgramWith(preds ...predicate.Program) predicate.ProgramMembership {
	return predicate.ProgramMembership(func(s *sql.Selector) {
		step := newProgramStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Program
		step.Edge.Schema = schemaConfig.ProgramMembership
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.ProgramMembership {
	return predicate.ProgramMembership(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.ProgramMembership
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.ProgramMembership {
	return predicate.ProgramMembership(func(s *sql.Selector) {
		step := newUserStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.ProgramMembership
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProgramMembership) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProgramMembership) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProgramMembership) predicate.ProgramMembership {
	return predicate.ProgramMembership(sql.NotPredicates(p))
}