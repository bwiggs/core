// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/theopenlane/core/internal/ent/generated/risk"
	"github.com/theopenlane/core/pkg/enums"
)

// Risk is the model entity for the Risk schema.
type Risk struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
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
	Details map[string]interface{} `json:"details,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RiskQuery when eager-loading is set.
	Edges                   RiskEdges `json:"edges"`
	control_objective_risks *string
	selectValues            sql.SelectValues
}

// RiskEdges holds the relations/edges for other nodes in the graph.
type RiskEdges struct {
	// Control holds the value of the control edge.
	Control []*Control `json:"control,omitempty"`
	// Procedure holds the value of the procedure edge.
	Procedure []*Procedure `json:"procedure,omitempty"`
	// Actionplans holds the value of the actionplans edge.
	Actionplans []*ActionPlan `json:"actionplans,omitempty"`
	// Program holds the value of the program edge.
	Program []*Program `json:"program,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
	// totalCount holds the count of the edges above.
	totalCount [4]map[string]int

	namedControl     map[string][]*Control
	namedProcedure   map[string][]*Procedure
	namedActionplans map[string][]*ActionPlan
	namedProgram     map[string][]*Program
}

// ControlOrErr returns the Control value or an error if the edge
// was not loaded in eager-loading.
func (e RiskEdges) ControlOrErr() ([]*Control, error) {
	if e.loadedTypes[0] {
		return e.Control, nil
	}
	return nil, &NotLoadedError{edge: "control"}
}

// ProcedureOrErr returns the Procedure value or an error if the edge
// was not loaded in eager-loading.
func (e RiskEdges) ProcedureOrErr() ([]*Procedure, error) {
	if e.loadedTypes[1] {
		return e.Procedure, nil
	}
	return nil, &NotLoadedError{edge: "procedure"}
}

// ActionplansOrErr returns the Actionplans value or an error if the edge
// was not loaded in eager-loading.
func (e RiskEdges) ActionplansOrErr() ([]*ActionPlan, error) {
	if e.loadedTypes[2] {
		return e.Actionplans, nil
	}
	return nil, &NotLoadedError{edge: "actionplans"}
}

// ProgramOrErr returns the Program value or an error if the edge
// was not loaded in eager-loading.
func (e RiskEdges) ProgramOrErr() ([]*Program, error) {
	if e.loadedTypes[3] {
		return e.Program, nil
	}
	return nil, &NotLoadedError{edge: "program"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Risk) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case risk.FieldTags, risk.FieldDetails:
			values[i] = new([]byte)
		case risk.FieldID, risk.FieldCreatedBy, risk.FieldUpdatedBy, risk.FieldDeletedBy, risk.FieldMappingID, risk.FieldName, risk.FieldDescription, risk.FieldStatus, risk.FieldRiskType, risk.FieldBusinessCosts, risk.FieldImpact, risk.FieldLikelihood, risk.FieldMitigation, risk.FieldSatisfies:
			values[i] = new(sql.NullString)
		case risk.FieldCreatedAt, risk.FieldUpdatedAt, risk.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case risk.ForeignKeys[0]: // control_objective_risks
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Risk fields.
func (r *Risk) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case risk.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				r.ID = value.String
			}
		case risk.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case risk.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case risk.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				r.CreatedBy = value.String
			}
		case risk.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				r.UpdatedBy = value.String
			}
		case risk.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				r.DeletedAt = value.Time
			}
		case risk.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				r.DeletedBy = value.String
			}
		case risk.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				r.MappingID = value.String
			}
		case risk.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case risk.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case risk.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				r.Description = value.String
			}
		case risk.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				r.Status = value.String
			}
		case risk.FieldRiskType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field risk_type", values[i])
			} else if value.Valid {
				r.RiskType = value.String
			}
		case risk.FieldBusinessCosts:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field business_costs", values[i])
			} else if value.Valid {
				r.BusinessCosts = value.String
			}
		case risk.FieldImpact:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field impact", values[i])
			} else if value.Valid {
				r.Impact = enums.RiskImpact(value.String)
			}
		case risk.FieldLikelihood:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field likelihood", values[i])
			} else if value.Valid {
				r.Likelihood = enums.RiskLikelihood(value.String)
			}
		case risk.FieldMitigation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mitigation", values[i])
			} else if value.Valid {
				r.Mitigation = value.String
			}
		case risk.FieldSatisfies:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field satisfies", values[i])
			} else if value.Valid {
				r.Satisfies = value.String
			}
		case risk.FieldDetails:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field details", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Details); err != nil {
					return fmt.Errorf("unmarshal field details: %w", err)
				}
			}
		case risk.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field control_objective_risks", values[i])
			} else if value.Valid {
				r.control_objective_risks = new(string)
				*r.control_objective_risks = value.String
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Risk.
// This includes values selected through modifiers, order, etc.
func (r *Risk) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryControl queries the "control" edge of the Risk entity.
func (r *Risk) QueryControl() *ControlQuery {
	return NewRiskClient(r.config).QueryControl(r)
}

// QueryProcedure queries the "procedure" edge of the Risk entity.
func (r *Risk) QueryProcedure() *ProcedureQuery {
	return NewRiskClient(r.config).QueryProcedure(r)
}

// QueryActionplans queries the "actionplans" edge of the Risk entity.
func (r *Risk) QueryActionplans() *ActionPlanQuery {
	return NewRiskClient(r.config).QueryActionplans(r)
}

// QueryProgram queries the "program" edge of the Risk entity.
func (r *Risk) QueryProgram() *ProgramQuery {
	return NewRiskClient(r.config).QueryProgram(r)
}

// Update returns a builder for updating this Risk.
// Note that you need to call Risk.Unwrap() before calling this method if this Risk
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Risk) Update() *RiskUpdateOne {
	return NewRiskClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Risk entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Risk) Unwrap() *Risk {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("generated: Risk is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Risk) String() string {
	var builder strings.Builder
	builder.WriteString("Risk(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(r.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(r.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(r.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(r.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(r.MappingID)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", r.Tags))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(r.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(r.Description)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(r.Status)
	builder.WriteString(", ")
	builder.WriteString("risk_type=")
	builder.WriteString(r.RiskType)
	builder.WriteString(", ")
	builder.WriteString("business_costs=")
	builder.WriteString(r.BusinessCosts)
	builder.WriteString(", ")
	builder.WriteString("impact=")
	builder.WriteString(fmt.Sprintf("%v", r.Impact))
	builder.WriteString(", ")
	builder.WriteString("likelihood=")
	builder.WriteString(fmt.Sprintf("%v", r.Likelihood))
	builder.WriteString(", ")
	builder.WriteString("mitigation=")
	builder.WriteString(r.Mitigation)
	builder.WriteString(", ")
	builder.WriteString("satisfies=")
	builder.WriteString(r.Satisfies)
	builder.WriteString(", ")
	builder.WriteString("details=")
	builder.WriteString(fmt.Sprintf("%v", r.Details))
	builder.WriteByte(')')
	return builder.String()
}

// NamedControl returns the Control named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Risk) NamedControl(name string) ([]*Control, error) {
	if r.Edges.namedControl == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedControl[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Risk) appendNamedControl(name string, edges ...*Control) {
	if r.Edges.namedControl == nil {
		r.Edges.namedControl = make(map[string][]*Control)
	}
	if len(edges) == 0 {
		r.Edges.namedControl[name] = []*Control{}
	} else {
		r.Edges.namedControl[name] = append(r.Edges.namedControl[name], edges...)
	}
}

// NamedProcedure returns the Procedure named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Risk) NamedProcedure(name string) ([]*Procedure, error) {
	if r.Edges.namedProcedure == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedProcedure[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Risk) appendNamedProcedure(name string, edges ...*Procedure) {
	if r.Edges.namedProcedure == nil {
		r.Edges.namedProcedure = make(map[string][]*Procedure)
	}
	if len(edges) == 0 {
		r.Edges.namedProcedure[name] = []*Procedure{}
	} else {
		r.Edges.namedProcedure[name] = append(r.Edges.namedProcedure[name], edges...)
	}
}

// NamedActionplans returns the Actionplans named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Risk) NamedActionplans(name string) ([]*ActionPlan, error) {
	if r.Edges.namedActionplans == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedActionplans[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Risk) appendNamedActionplans(name string, edges ...*ActionPlan) {
	if r.Edges.namedActionplans == nil {
		r.Edges.namedActionplans = make(map[string][]*ActionPlan)
	}
	if len(edges) == 0 {
		r.Edges.namedActionplans[name] = []*ActionPlan{}
	} else {
		r.Edges.namedActionplans[name] = append(r.Edges.namedActionplans[name], edges...)
	}
}

// NamedProgram returns the Program named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Risk) NamedProgram(name string) ([]*Program, error) {
	if r.Edges.namedProgram == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedProgram[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Risk) appendNamedProgram(name string, edges ...*Program) {
	if r.Edges.namedProgram == nil {
		r.Edges.namedProgram = make(map[string][]*Program)
	}
	if len(edges) == 0 {
		r.Edges.namedProgram[name] = []*Program{}
	} else {
		r.Edges.namedProgram[name] = append(r.Edges.namedProgram[name], edges...)
	}
}

// Risks is a parsable slice of Risk.
type Risks []*Risk
