// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/theopenlane/core/internal/ent/generated/standardhistory"
	"github.com/theopenlane/entx/history"
)

// StandardHistory is the model entity for the StandardHistory schema.
type StandardHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// HistoryTime holds the value of the "history_time" field.
	HistoryTime time.Time `json:"history_time,omitempty"`
	// Ref holds the value of the "ref" field.
	Ref string `json:"ref,omitempty"`
	// Operation holds the value of the "operation" field.
	Operation history.OpType `json:"operation,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// MappingID holds the value of the "mapping_id" field.
	MappingID string `json:"mapping_id,omitempty"`
	// tags associated with the object
	Tags []string `json:"tags,omitempty"`
	// the name of the standard body, e.g. TSC, NIST, SOC, HITRUST, FedRamp, etc.
	Name string `json:"name,omitempty"`
	// description of the standard
	Description string `json:"description,omitempty"`
	// family of the standard, e.g. 800-53, 800-171, 27001, etc.
	Family string `json:"family,omitempty"`
	// status of the standard - active, deprecated, etc.
	Status string `json:"status,omitempty"`
	// type of the standard - security, privacy, etc.
	StandardType string `json:"standard_type,omitempty"`
	// version of the standard
	Version string `json:"version,omitempty"`
	// purpose and scope
	PurposeAndScope string `json:"purpose_and_scope,omitempty"`
	// background of the standard
	Background string `json:"background,omitempty"`
	// which controls are satisfied by the standard
	Satisfies string `json:"satisfies,omitempty"`
	// json data with details of the standard
	Details      map[string]interface{} `json:"details,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*StandardHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case standardhistory.FieldTags, standardhistory.FieldDetails:
			values[i] = new([]byte)
		case standardhistory.FieldOperation:
			values[i] = new(history.OpType)
		case standardhistory.FieldID, standardhistory.FieldRef, standardhistory.FieldCreatedBy, standardhistory.FieldUpdatedBy, standardhistory.FieldDeletedBy, standardhistory.FieldMappingID, standardhistory.FieldName, standardhistory.FieldDescription, standardhistory.FieldFamily, standardhistory.FieldStatus, standardhistory.FieldStandardType, standardhistory.FieldVersion, standardhistory.FieldPurposeAndScope, standardhistory.FieldBackground, standardhistory.FieldSatisfies:
			values[i] = new(sql.NullString)
		case standardhistory.FieldHistoryTime, standardhistory.FieldCreatedAt, standardhistory.FieldUpdatedAt, standardhistory.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the StandardHistory fields.
func (sh *StandardHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case standardhistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				sh.ID = value.String
			}
		case standardhistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				sh.HistoryTime = value.Time
			}
		case standardhistory.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				sh.Ref = value.String
			}
		case standardhistory.FieldOperation:
			if value, ok := values[i].(*history.OpType); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value != nil {
				sh.Operation = *value
			}
		case standardhistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sh.CreatedAt = value.Time
			}
		case standardhistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sh.UpdatedAt = value.Time
			}
		case standardhistory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				sh.CreatedBy = value.String
			}
		case standardhistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				sh.UpdatedBy = value.String
			}
		case standardhistory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sh.DeletedAt = value.Time
			}
		case standardhistory.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				sh.DeletedBy = value.String
			}
		case standardhistory.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				sh.MappingID = value.String
			}
		case standardhistory.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &sh.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case standardhistory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sh.Name = value.String
			}
		case standardhistory.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				sh.Description = value.String
			}
		case standardhistory.FieldFamily:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field family", values[i])
			} else if value.Valid {
				sh.Family = value.String
			}
		case standardhistory.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				sh.Status = value.String
			}
		case standardhistory.FieldStandardType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field standard_type", values[i])
			} else if value.Valid {
				sh.StandardType = value.String
			}
		case standardhistory.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				sh.Version = value.String
			}
		case standardhistory.FieldPurposeAndScope:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field purpose_and_scope", values[i])
			} else if value.Valid {
				sh.PurposeAndScope = value.String
			}
		case standardhistory.FieldBackground:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field background", values[i])
			} else if value.Valid {
				sh.Background = value.String
			}
		case standardhistory.FieldSatisfies:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field satisfies", values[i])
			} else if value.Valid {
				sh.Satisfies = value.String
			}
		case standardhistory.FieldDetails:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field details", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &sh.Details); err != nil {
					return fmt.Errorf("unmarshal field details: %w", err)
				}
			}
		default:
			sh.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the StandardHistory.
// This includes values selected through modifiers, order, etc.
func (sh *StandardHistory) Value(name string) (ent.Value, error) {
	return sh.selectValues.Get(name)
}

// Update returns a builder for updating this StandardHistory.
// Note that you need to call StandardHistory.Unwrap() before calling this method if this StandardHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (sh *StandardHistory) Update() *StandardHistoryUpdateOne {
	return NewStandardHistoryClient(sh.config).UpdateOne(sh)
}

// Unwrap unwraps the StandardHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sh *StandardHistory) Unwrap() *StandardHistory {
	_tx, ok := sh.config.driver.(*txDriver)
	if !ok {
		panic("generated: StandardHistory is not a transactional entity")
	}
	sh.config.driver = _tx.drv
	return sh
}

// String implements the fmt.Stringer.
func (sh *StandardHistory) String() string {
	var builder strings.Builder
	builder.WriteString("StandardHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sh.ID))
	builder.WriteString("history_time=")
	builder.WriteString(sh.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(sh.Ref)
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", sh.Operation))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(sh.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sh.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(sh.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(sh.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(sh.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(sh.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(sh.MappingID)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", sh.Tags))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(sh.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(sh.Description)
	builder.WriteString(", ")
	builder.WriteString("family=")
	builder.WriteString(sh.Family)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(sh.Status)
	builder.WriteString(", ")
	builder.WriteString("standard_type=")
	builder.WriteString(sh.StandardType)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(sh.Version)
	builder.WriteString(", ")
	builder.WriteString("purpose_and_scope=")
	builder.WriteString(sh.PurposeAndScope)
	builder.WriteString(", ")
	builder.WriteString("background=")
	builder.WriteString(sh.Background)
	builder.WriteString(", ")
	builder.WriteString("satisfies=")
	builder.WriteString(sh.Satisfies)
	builder.WriteString(", ")
	builder.WriteString("details=")
	builder.WriteString(fmt.Sprintf("%v", sh.Details))
	builder.WriteByte(')')
	return builder.String()
}

// StandardHistories is a parsable slice of StandardHistory.
type StandardHistories []*StandardHistory