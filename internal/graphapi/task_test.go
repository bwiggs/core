package graphapi_test

import (
	"context"
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	mock_fga "github.com/theopenlane/iam/fgax/mockery"
	"github.com/theopenlane/utils/ulids"

	"github.com/theopenlane/core/pkg/enums"
	"github.com/theopenlane/core/pkg/openlaneclient"
)

func (suite *GraphTestSuite) TestQueryTask() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	task := (&TaskBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name     string
		queryID  string
		client   *openlaneclient.OpenlaneClient
		ctx      context.Context
		errorMsg string
	}{
		{
			name:    "happy path",
			queryID: task.ID,
			client:  suite.client.api,
			ctx:     reqCtx,
		},
		{
			name:    "happy path using personal access token",
			queryID: task.ID,
			client:  suite.client.apiWithPAT,
			ctx:     context.Background(),
		},
		{
			name:     "not found",
			queryID:  "notfound",
			client:   suite.client.api,
			ctx:      reqCtx,
			errorMsg: "not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// mock the check and task tuples for the task
			mock_fga.CheckAny(t, suite.client.fga, true)
			mock_fga.ListOnce(t, suite.client.fga, []string{"task:" + task.ID}, nil)

			if tc.errorMsg == "" {
				// mock the user tuple for the task assigner for a successful query
				mock_fga.ListOnce(t, suite.client.fga, []string{"user:" + testUser.ID}, nil)
			}

			resp, err := tc.client.GetTaskByID(tc.ctx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, tc.queryID, resp.Task.ID)
			assert.NotEmpty(t, resp.Task.Title)
			assert.NotEmpty(t, resp.Task.Description)
			assert.NotEmpty(t, resp.Task.Status)
		})
	}
}

func (suite *GraphTestSuite) TestQueryTasks() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	task1 := (&TaskBuilder{client: suite.client}).MustNew(reqCtx, t)
	task2 := (&TaskBuilder{client: suite.client}).MustNew(reqCtx, t)

	otherUser := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)
	otherCtx, err := userContextWithID(otherUser.ID)
	require.NoError(t, err)

	testCases := []struct {
		name            string
		client          *openlaneclient.OpenlaneClient
		ctx             context.Context
		expectedResults int
	}{
		{
			name:            "happy path",
			client:          suite.client.api,
			ctx:             reqCtx,
			expectedResults: 2,
		},
		{
			name:            "happy path, using pat",
			client:          suite.client.apiWithPAT,
			ctx:             context.Background(),
			expectedResults: 2,
		},
		{
			name:            "another user, no entities should be returned",
			client:          suite.client.api,
			ctx:             otherCtx,
			expectedResults: 0,
		},
	}

	for _, tc := range testCases {
		t.Run("List "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// mock the task + user tuple for the task assigner for a successful query
			if tc.expectedResults > 0 {
				mock_fga.ListOnce(t, suite.client.fga, []string{"task:" + task1.ID, "task:" + task2.ID}, nil)
				mock_fga.ListOnce(t, suite.client.fga, []string{"user:" + testUser.ID}, nil)
			} else {
				// mock no access to any tasks
				mock_fga.ListOnce(t, suite.client.fga, []string{}, nil)
			}

			resp, err := tc.client.GetAllTasks(tc.ctx)
			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Len(t, resp.Tasks.Edges, tc.expectedResults)
		})
	}
}

func (suite *GraphTestSuite) TestMutationCreateTask() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	otherUser := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name        string
		request     openlaneclient.CreateTaskInput
		client      *openlaneclient.OpenlaneClient
		ctx         context.Context
		numWrites   int
		expectedErr string
	}{
		{
			name: "happy path, minimal input",
			request: openlaneclient.CreateTaskInput{
				Title: "test-task",
			},
			client:    suite.client.api,
			ctx:       reqCtx,
			numWrites: 1, // write permission for the task to the user
		},
		{
			name: "happy path, all input",
			request: openlaneclient.CreateTaskInput{
				Title:       "test-task",
				Description: lo.ToPtr("test description"),
				Status:      &enums.TaskStatusInProgress,
				Details: map[string]interface{}{
					"task": "do all the things for the thing",
				},
				Due:             lo.ToPtr(time.Now().Add(time.Hour * 24)),
				OrganizationIDs: []string{testOrgID}, // add the org to the task
				AssigneeID:      &otherUser.ID,       // assign the task to another user
			},
			client:    suite.client.api,
			ctx:       reqCtx,
			numWrites: 2, // assignee+assigner and organization write permissions
		},
		{
			name: "happy path, using pat",
			request: openlaneclient.CreateTaskInput{
				Title: "test-task",
			},
			client:    suite.client.apiWithPAT,
			ctx:       context.Background(),
			numWrites: 1, // write permission for the task to the user
		},
		{
			name: "missing title, but display name provided",
			request: openlaneclient.CreateTaskInput{
				Description: lo.ToPtr("makin' a list, checkin' it twice"),
			},
			client:      suite.client.api,
			ctx:         reqCtx,
			expectedErr: "value is less than the required length",
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			if tc.expectedErr == "" {
				// mock write permissions
				for range tc.numWrites {
					mock_fga.WriteOnce(t, suite.client.fga)
				}

				if tc.request.AssigneeID != nil {
					// mock the user tuple for the assignee
					// there is an assumption in this test that the user has access to the otherUser (this means that they should be in the org)
					mock_fga.ListUsersAny(t, suite.client.fga, []string{otherUser.ID, testUser.ID}, nil)
				}
			}

			resp, err := tc.client.CreateTask(tc.ctx, tc.request)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, tc.request.Title, resp.CreateTask.Task.Title)

			if tc.request.Description == nil {
				assert.Empty(t, resp.CreateTask.Task.Description)
			} else {
				assert.Equal(t, tc.request.Description, resp.CreateTask.Task.Description)
			}

			if tc.request.Status == nil {
				assert.Equal(t, enums.TaskStatusOpen, resp.CreateTask.Task.Status)
			} else {
				assert.Equal(t, *tc.request.Status, resp.CreateTask.Task.Status)
			}

			if tc.request.Details == nil {
				assert.Empty(t, resp.CreateTask.Task.Details)
			} else {
				assert.Equal(t, tc.request.Details, resp.CreateTask.Task.Details)
			}

			if tc.request.Due == nil {
				assert.Empty(t, resp.CreateTask.Task.Due)
			} else {
				assert.WithinDuration(t, *tc.request.Due, *resp.CreateTask.Task.Due, 10*time.Second)
			}
		})
	}
}

func (suite *GraphTestSuite) TestMutationUpdateTask() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	task := (&TaskBuilder{client: suite.client}).MustNew(reqCtx, t)
	group := (&GroupBuilder{client: suite.client}).MustNew(reqCtx, t)

	otherUser := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name        string
		request     openlaneclient.UpdateTaskInput
		client      *openlaneclient.OpenlaneClient
		ctx         context.Context
		numWrites   int
		allowed     bool
		expectedErr string
	}{
		{
			name: "happy path, update description",
			request: openlaneclient.UpdateTaskInput{
				Description: lo.ToPtr(("makin' a list, checkin' it twice")),
			},
			client:    suite.client.api,
			ctx:       reqCtx,
			numWrites: 0,
			allowed:   true,
		},
		{
			name: "update assignee",
			request: openlaneclient.UpdateTaskInput{
				AssigneeID: lo.ToPtr(otherUser.ID),
			},
			client:    suite.client.apiWithPAT,
			ctx:       context.Background(),
			numWrites: 1, // assignee write permission
			allowed:   true,
		},
		{
			name: "update status and details",
			request: openlaneclient.UpdateTaskInput{
				Status:  &enums.TaskStatusInProgress,
				Details: map[string]interface{}{"task": "do all the things for the thing"},
			},
			client:  suite.client.api,
			ctx:     reqCtx,
			allowed: true,
		},
		{
			name: "add to group",
			request: openlaneclient.UpdateTaskInput{
				AddGroupIDs: []string{group.ID},
			},
			client:    suite.client.api,
			ctx:       reqCtx,
			numWrites: 1, // group write permission
			allowed:   true,
		},
	}

	for _, tc := range testCases {
		t.Run("Update "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// mock the list objects request for the task
			mock_fga.ListAny(t, suite.client.fga, []string{"task:" + task.ID})
			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			if tc.expectedErr == "" {
				// mock write permissions
				for range tc.numWrites {
					mock_fga.WriteOnce(t, suite.client.fga)
				}

				if tc.request.AssigneeID != nil {
					// mock the user tuple for the assignee
					// there is an assumption in this test that the user has access to the otherUser (this means that they should be in the org)
					mock_fga.ListUsersAny(t, suite.client.fga, []string{otherUser.ID, testUser.ID}, nil)
				}
			}

			resp, err := tc.client.UpdateTask(tc.ctx, task.ID, tc.request)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			if tc.request.Description != nil {
				assert.Equal(t, *tc.request.Description, *resp.UpdateTask.Task.Description)
			}

			if tc.request.Status != nil {
				assert.Equal(t, *tc.request.Status, resp.UpdateTask.Task.Status)
			}

			if tc.request.Details != nil {
				assert.Equal(t, tc.request.Details, resp.UpdateTask.Task.Details)
			}
		})
	}
}

func (suite *GraphTestSuite) TestMutationDeleteTask() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	task1 := (&TaskBuilder{client: suite.client}).MustNew(reqCtx, t)
	task2 := (&TaskBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name        string
		idToDelete  string
		client      *openlaneclient.OpenlaneClient
		ctx         context.Context
		allowed     bool
		expectedErr string
	}{
		{
			name:       "happy path, delete task",
			idToDelete: task1.ID,
			client:     suite.client.api,
			ctx:        reqCtx,
			allowed:    true,
		},
		{
			name:        "task already deleted, not found",
			idToDelete:  task1.ID,
			client:      suite.client.api,
			ctx:         reqCtx,
			allowed:     true,
			expectedErr: "task not found",
		},
		{
			name:       "happy path, delete task using personal access token",
			idToDelete: task2.ID,
			client:     suite.client.apiWithPAT,
			ctx:        context.Background(),
			allowed:    true,
		},
		{
			name:        "unknown task, not found",
			idToDelete:  ulids.New().String(),
			client:      suite.client.api,
			ctx:         reqCtx,
			allowed:     true,
			expectedErr: "task not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Delete "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)
			mock_fga.ListAny(t, suite.client.fga, []string{"task:" + tc.idToDelete})

			resp, err := tc.client.DeleteTask(tc.ctx, tc.idToDelete)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, tc.idToDelete, resp.DeleteTask.DeletedID)
		})
	}
}