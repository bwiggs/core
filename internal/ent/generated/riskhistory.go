// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/theopenlane/core/internal/ent/generated/riskhistory"
	"github.com/theopenlane/core/pkg/enums"
	"github.com/theopenlane/entx/history"
)

// RiskHistory is the model entity for the RiskHistory schema.
type RiskHistory struct {
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
	// the name of the risk
	Name string `json:"name,omitempty"`
	// description of the risk
	Description string `json:"description,omitempty"`
	// status of the risk - mitigated or not, inflight, etc.
	Status string `json:"status,omitempty"`
	// type of the risk, e.g. strategic, operational, financial, external, etc.
	RiskType string `json:"risk_type,omitempty"`
	// business costs associated with the risk
	BusinessCosts string `json:"business_costs,omitempty"`
	// impact of the risk - high, medium, low
	Impact enums.RiskImpact `json:"impact,omitempty"`
	// likelihood of the risk occurring; unlikely, likely, highly likely
	Likelihood enums.RiskLikelihood `json:"likelihood,omitempty"`
	// mitigation for the risk
	Mitigation string `json:"mitigation,omitempty"`
	// which controls are satisfied by the risk
	Satisfies string `json:"satisfies,omitempty"`
	// json data for the risk document
	Details      map[string]interface{} `json:"details,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RiskHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case riskhistory.FieldTags, riskhistory.FieldDetails:
			values[i] = new([]byte)
		case riskhistory.FieldOperation:
			values[i] = new(history.OpType)
		case riskhistory.FieldID, riskhistory.FieldRef, riskhistory.FieldCreatedBy, riskhistory.FieldUpdatedBy, riskhistory.FieldDeletedBy, riskhistory.FieldMappingID, riskhistory.FieldName, riskhistory.FieldDescription, riskhistory.FieldStatus, riskhistory.FieldRiskType, riskhistory.FieldBusinessCosts, riskhistory.FieldImpact, riskhistory.FieldLikelihood, riskhistory.FieldMitigation, riskhistory.FieldSatisfies:
			values[i] = new(sql.NullString)
		case riskhistory.FieldHistoryTime, riskhistory.FieldCreatedAt, riskhistory.FieldUpdatedAt, riskhistory.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RiskHistory fields.
func (rh *RiskHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case riskhistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				rh.ID = value.String
			}
		case riskhistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				rh.HistoryTime = value.Time
			}
		case riskhistory.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				rh.Ref = value.String
			}
		case riskhistory.FieldOperation:
			if value, ok := values[i].(*history.OpType); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value != nil {
				rh.Operation = *value
			}
		case riskhistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rh.CreatedAt = value.Time
			}
		case riskhistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				rh.UpdatedAt = value.Time
			}
		case riskhistory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				rh.CreatedBy = value.String
			}
		case riskhistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				rh.UpdatedBy = value.String
			}
		case riskhistory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				rh.DeletedAt = value.Time
			}
		case riskhistory.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				rh.DeletedBy = value.String
			}
		case riskhistory.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				rh.MappingID = value.String
			}
		case riskhistory.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &rh.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case riskhistory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				rh.Name = value.String
			}
		case riskhistory.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				rh.Description = value.String
			}
		case riskhistory.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				rh.Status = value.String
			}
		case riskhistory.FieldRiskType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field risk_type", values[i])
			} else if value.Valid {
				rh.RiskType = value.String
			}
		case riskhistory.FieldBusinessCosts:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field business_costs", values[i])
			} else if value.Valid {
				rh.BusinessCosts = value.String
			}
		case riskhistory.FieldImpact:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field impact", values[i])
			} else if value.Valid {
				rh.Impact = enums.RiskImpact(value.String)
			}
		case riskhistory.FieldLikelihood:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field likelihood", values[i])
			} else if value.Valid {
				rh.Likelihood = enums.RiskLikelihood(value.String)
			}
		case riskhistory.FieldMitigation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mitigation", values[i])
			} else if value.Valid {
				rh.Mitigation = value.String
			}
		case riskhistory.FieldSatisfies:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field satisfies", values[i])
			} else if value.Valid {
				rh.Satisfies = value.String
			}
		case riskhistory.FieldDetails:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field details", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &rh.Details); err != nil {
					return fmt.Errorf("unmarshal field details: %w", err)
				}
			}
		default:
			rh.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RiskHistory.
// This includes values selected through modifiers, order, etc.
func (rh *RiskHistory) Value(name string) (ent.Value, error) {
	return rh.selectValues.Get(name)
}

// Update returns a builder for updating this RiskHistory.
// Note that you need to call RiskHistory.Unwrap() before calling this method if this RiskHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (rh *RiskHistory) Update() *RiskHistoryUpdateOne {
	return NewRiskHistoryClient(rh.config).UpdateOne(rh)
}

// Unwrap unwraps the RiskHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rh *RiskHistory) Unwrap() *RiskHistory {
	_tx, ok := rh.config.driver.(*txDriver)
	if !ok {
		panic("generated: RiskHistory is not a transactional entity")
	}
	rh.config.driver = _tx.drv
	return rh
}

// String implements the fmt.Stringer.
func (rh *RiskHistory) String() string {
	var builder strings.Builder
	builder.WriteString("RiskHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rh.ID))
	builder.WriteString("history_time=")
	builder.WriteString(rh.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(rh.Ref)
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", rh.Operation))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(rh.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(rh.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(rh.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(rh.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(rh.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(rh.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(rh.MappingID)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", rh.Tags))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(rh.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(rh.Description)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(rh.Status)
	builder.WriteString(", ")
	builder.WriteString("risk_type=")
	builder.WriteString(rh.RiskType)
	builder.WriteString(", ")
	builder.WriteString("business_costs=")
	builder.WriteString(rh.BusinessCosts)
	builder.WriteString(", ")
	builder.WriteString("impact=")
	builder.WriteString(fmt.Sprintf("%v", rh.Impact))
	builder.WriteString(", ")
	builder.WriteString("likelihood=")
	builder.WriteString(fmt.Sprintf("%v", rh.Likelihood))
	builder.WriteString(", ")
	builder.WriteString("mitigation=")
	builder.WriteString(rh.Mitigation)
	builder.WriteString(", ")
	builder.WriteString("satisfies=")
	builder.WriteString(rh.Satisfies)
	builder.WriteString(", ")
	builder.WriteString("details=")
	builder.WriteString(fmt.Sprintf("%v", rh.Details))
	builder.WriteByte(')')
	return builder.String()
}

// RiskHistories is a parsable slice of RiskHistory.
type RiskHistories []*RiskHistory