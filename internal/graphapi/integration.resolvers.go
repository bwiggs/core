package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/theopenlane/core/internal/ent/generated"
	"github.com/theopenlane/utils/rout"
)

// CreateIntegration is the resolver for the createIntegration field.
func (r *mutationResolver) CreateIntegration(ctx context.Context, input generated.CreateIntegrationInput) (*IntegrationCreatePayload, error) {
	// set the organization in the auth context if its not done for us
	if err := setOrganizationInAuthContext(ctx, input.OwnerID); err != nil {
		r.logger.Errorw("failed to set organization in auth context", "error", err)

		return nil, rout.NewMissingRequiredFieldError("owner_id")
	}

	res, err := withTransactionalMutation(ctx).Integration.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "integration"}, r.logger)
	}

	return &IntegrationCreatePayload{
		Integration: res,
	}, nil
}

// CreateBulkIntegration is the resolver for the createBulkIntegration field.
func (r *mutationResolver) CreateBulkIntegration(ctx context.Context, input []*generated.CreateIntegrationInput) (*IntegrationBulkCreatePayload, error) {
	return r.bulkCreateIntegration(ctx, input)
}

// CreateBulkCSVIntegration is the resolver for the createBulkCSVIntegration field.
func (r *mutationResolver) CreateBulkCSVIntegration(ctx context.Context, input graphql.Upload) (*IntegrationBulkCreatePayload, error) {
	data, err := unmarshalBulkData[generated.CreateIntegrationInput](input)
	if err != nil {
		r.logger.Errorw("failed to unmarshal bulk data", "error", err)

		return nil, err
	}

	return r.bulkCreateIntegration(ctx, data)
}

// UpdateIntegration is the resolver for the updateIntegration field.
func (r *mutationResolver) UpdateIntegration(ctx context.Context, id string, input generated.UpdateIntegrationInput) (*IntegrationUpdatePayload, error) {
	res, err := withTransactionalMutation(ctx).Integration.Get(ctx, id)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionUpdate, object: "integration"}, r.logger)
	}
	// set the organization in the auth context if its not done for us
	if err := setOrganizationInAuthContext(ctx, &res.OwnerID); err != nil {
		r.logger.Errorw("failed to set organization in auth context", "error", err)

		return nil, ErrPermissionDenied
	}

	// setup update request
	req := res.Update().SetInput(input).AppendTags(input.AppendTags)

	res, err = req.Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionUpdate, object: "integration"}, r.logger)
	}

	return &IntegrationUpdatePayload{
		Integration: res,
	}, nil
}

// DeleteIntegration is the resolver for the deleteIntegration field.
func (r *mutationResolver) DeleteIntegration(ctx context.Context, id string) (*IntegrationDeletePayload, error) {
	if err := withTransactionalMutation(ctx).Integration.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, parseRequestError(err, action{action: ActionDelete, object: "integration"}, r.logger)
	}

	if err := generated.IntegrationEdgeCleanup(ctx, id); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	return &IntegrationDeletePayload{
		DeletedID: id,
	}, nil
}

// Integration is the resolver for the integration field.
func (r *queryResolver) Integration(ctx context.Context, id string) (*generated.Integration, error) {
	res, err := withTransactionalMutation(ctx).Integration.Get(ctx, id)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionGet, object: "integration"}, r.logger)
	}

	return res, nil
}