package handlers

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	echo "github.com/theopenlane/echox"
	"github.com/theopenlane/iam/fgax"

	"github.com/theopenlane/utils/rout"

	"github.com/theopenlane/core/pkg/auth"
	"github.com/theopenlane/core/pkg/models"

	sliceutil "github.com/theopenlane/utils/slice"
)

// AccountRolesOrganizationHandler lists roles a subject has in relation to an organization
func (h *Handler) AccountRolesOrganizationHandler(ctx echo.Context) error {
	var in models.AccountRolesOrganizationRequest
	if err := ctx.Bind(&in); err != nil {
		return h.InvalidInput(ctx, err)
	}

	au, err := auth.GetAuthenticatedUserContext(ctx.Request().Context())
	if err != nil {
		h.Logger.Error("error getting authenticated user", "error", err)

		return h.InternalServerError(ctx, err)
	}

	in.ID, err = h.getOrganizationID(in, au)
	if err != nil {
		return h.BadRequest(ctx, err)
	}

	// validate the input
	if err := in.Validate(); err != nil {
		return h.BadRequest(ctx, err)
	}

	// determine the subject type
	subjectType := "user"
	if au.AuthenticationType == auth.APITokenAuthentication {
		subjectType = "service"
	}

	req := fgax.ListAccess{
		SubjectType: subjectType,
		SubjectID:   au.SubjectID,
		ObjectID:    in.ID,
		ObjectType:  fgax.Kind("organization"),
		Relations:   DefaultAllRelations,
	}

	roles, err := h.DBClient.Authz.ListRelations(ctx.Request().Context(), req)
	if err != nil {
		h.Logger.Errorw("error checking access", "error", err)

		return h.InternalServerError(ctx, err)
	}

	return h.Success(ctx, models.AccountRolesOrganizationReply{
		Reply:          rout.Reply{Success: true},
		Roles:          roles,
		OrganizationID: req.ObjectID,
	})
}

// getOrganizationID returns the organization ID to use for the request based on the input and authenticated user
func (h *Handler) getOrganizationID(in models.AccountRolesOrganizationRequest, au *auth.AuthenticatedUser) (string, error) {
	// if an ID is provided, check if the authenticated user has access to it
	if in.ID != "" {
		if !sliceutil.Contains(au.OrganizationIDs, in.ID) {
			return "", ErrInvalidInput
		}

		return in.ID, nil
	}

	// if no ID is provided, default to the authenticated organization
	if au.OrganizationID != "" {
		return au.OrganizationID, nil
	}

	// if it is still empty, and the personal access token only has one organization use that
	if len(au.OrganizationIDs) == 1 {
		return au.OrganizationIDs[0], nil
	}

	return "", nil
}

// BindAccountRolesOrganization returns the OpenAPI3 operation for accepting an account roles organization request
func (h *Handler) BindAccountRolesOrganization() *openapi3.Operation {
	orgRoles := openapi3.NewOperation()
	orgRoles.Description = "List roles a subject has in relation to the authenticated organization"
	orgRoles.OperationID = "AccountRolesOrganization"
	orgRoles.Security = &openapi3.SecurityRequirements{
		openapi3.SecurityRequirement{
			"bearerAuth": []string{},
		},
	}

	h.AddResponse("AccountRolesOrganizationReply", "success", models.ExampleAccountRolesOrganizationReply, orgRoles, http.StatusOK)
	orgRoles.AddResponse(http.StatusInternalServerError, internalServerError())
	orgRoles.AddResponse(http.StatusBadRequest, badRequest())
	orgRoles.AddResponse(http.StatusUnauthorized, unauthorized())

	return orgRoles
}

// BindAccountRolesOrganization returns the OpenAPI3 operation for accepting an account roles organization request
func (h *Handler) BindAccountRolesOrganizationWithParam() *openapi3.Operation {
	orgRoles := openapi3.NewOperation()
	orgRoles.Description = "List roles a subject has in relation to a specific organization id"
	orgRoles.OperationID = "AccountRolesOrganizationWithID"
	orgRoles.Security = &openapi3.SecurityRequirements{
		openapi3.SecurityRequirement{
			"bearerAuth": []string{},
		},
	}

	h.AddRequestBody("AccountRolesOrganizationRequest", models.ExampleAccountRolesOrganizationRequest, orgRoles)
	h.AddResponse("AccountRolesOrganizationReply", "success", models.ExampleAccountRolesOrganizationReply, orgRoles, http.StatusOK)
	orgRoles.AddResponse(http.StatusInternalServerError, internalServerError())
	orgRoles.AddResponse(http.StatusBadRequest, badRequest())
	orgRoles.AddResponse(http.StatusUnauthorized, unauthorized())

	return orgRoles
}