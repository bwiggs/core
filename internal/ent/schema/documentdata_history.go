// Code generated by entx.history, DO NOT EDIT.
package schema

import (
	"context"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/theopenlane/core/internal/ent/generated"
	"github.com/theopenlane/core/internal/ent/generated/privacy"
	"github.com/theopenlane/core/internal/ent/interceptors"
	"github.com/theopenlane/entx"
	"github.com/theopenlane/entx/history"
	"github.com/theopenlane/iam/entfga"
)

// DocumentDataHistory holds the schema definition for the DocumentDataHistory entity.
type DocumentDataHistory struct {
	ent.Schema
}

// Annotations of the DocumentDataHistory.
func (DocumentDataHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.SchemaGenSkip(true),
		entsql.Annotation{
			Table: "document_data_history",
		},
		history.Annotations{
			IsHistory: true,
			Exclude:   true,
		},
		entgql.QueryField(),
		entgql.RelayConnection(),
		entfga.Annotations{
			ObjectType:   "organization",
			IDField:      "OwnerID",
			IncludeHooks: false,
		},
	}
}

// Fields of the DocumentDataHistory.
func (DocumentDataHistory) Fields() []ent.Field {
	historyFields := []ent.Field{
		field.Time("history_time").
			Default(time.Now).
			Immutable(),
		field.String("ref").
			Immutable().
			Optional(),
		field.Enum("operation").
			GoType(history.OpType("")).
			Immutable(),
	}

	// get the fields from the mixins
	// we only want to include mixin fields, not edges
	// so this prevents FKs back to the main tables
	mixins := DocumentData{}.Mixin()
	for _, mixin := range mixins {
		for _, field := range mixin.Fields() {
			// make sure the mixed in fields do not have unique constraints
			field.Descriptor().Unique = false

			// make sure the mixed in fields do not have validators
			field.Descriptor().Validators = nil

			// append the mixed in field to the history fields
			historyFields = append(historyFields, field)
		}
	}

	original := DocumentData{}
	for _, field := range original.Fields() {
		// make sure the fields do not have unique constraints
		field.Descriptor().Unique = false

		// make sure the mixed in fields do not have validators
		field.Descriptor().Validators = nil

		// append the field to the history fields
		historyFields = append(historyFields, field)
	}

	return historyFields
}

// Indexes of the DocumentDataHistory
func (DocumentDataHistory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("history_time"),
	}
}

// Interceptors of the DocumentDataHistory
func (DocumentDataHistory) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		interceptors.HistoryAccess("audit_log_viewer", true, false),
	}
}

// Policy of the DocumentDataHistory
func (DocumentDataHistory) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			privacy.DocumentDataHistoryQueryRuleFunc(func(ctx context.Context, q *generated.DocumentDataHistoryQuery) error {
				return q.CheckAccess(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}