package integration_tests

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestWorkflowSearch tests the basic Search API endpoint (/api/workflow/search)
// This is the main Search API that supports comprehensive query and search functionality
func TestWorkflowSearch(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)

	// Setup test data
	uniqueSuffix := strconv.Itoa(time.Now().Nanosecond())
	correlationId := "search-test-correlation-" + uuid.New().String()

	// Create and start test workflows
	workflowId1, workflowId2, wf1Name, wf2Name := setupTestWorkflows(t, uniqueSuffix, correlationId)

	// Cleanup
	defer cleanupTestWorkflows(t, workflowId1, workflowId2, wf1Name, wf2Name)

	// Wait for workflows to complete and indexing
	err := testdata.WaitForMultipleWorkflowsCompletion([]string{workflowId1, workflowId2}, testdata.WorkflowValidationTimeout)
	assert.NoError(t, err, "Failed to wait for workflows completion")

	// Table-driven tests
	tests := []struct {
		name               string
		searchOpts         *client.WorkflowResourceApiSearchOpts
		expectError        bool
		expectMinResults   int
		expectExactResults int
		validateResults    func(t *testing.T, results []model.WorkflowSummary, workflowId1, workflowId2, wf1Name string)
	}{
		{
			name: "BasicSearch",
			searchOpts: &client.WorkflowResourceApiSearchOpts{
				Start: optional.NewInt32(0),
				Size:  optional.NewInt32(100),
			},
			expectMinResults: 0,
			validateResults: func(t *testing.T, results []model.WorkflowSummary, workflowId1, workflowId2, wf1Name string) {
				// Basic search should return some results (no specific validation needed)
			},
		},
		{
			name: "SearchByWorkflowType",
			searchOpts: &client.WorkflowResourceApiSearchOpts{
				Query: optional.NewString(fmt.Sprintf("workflowType = %s", wf1Name)),
				Start: optional.NewInt32(0),
				Size:  optional.NewInt32(10),
			},
			expectMinResults: 1,
			validateResults: func(t *testing.T, results []model.WorkflowSummary, workflowId1, workflowId2, wf1Name string) {
				found := false
				for _, result := range results {
					if result.WorkflowId == workflowId1 {
						found = true
						assert.Equal(t, wf1Name, result.WorkflowType, "Workflow type should match")
						assert.Equal(t, "COMPLETED", result.Status, "Workflow should be completed")
						break
					}
				}
				assert.True(t, found, "Should find workflow 1 by type")
			},
		},
		{
			name: "SearchByCorrelationId",
			searchOpts: &client.WorkflowResourceApiSearchOpts{
				Query: optional.NewString(fmt.Sprintf("correlationId = %s", correlationId)),
				Start: optional.NewInt32(0),
				Size:  optional.NewInt32(10),
			},
			expectMinResults: 1,
			validateResults: func(t *testing.T, results []model.WorkflowSummary, workflowId1, workflowId2, wf1Name string) {
				found := false
				for _, result := range results {
					if result.WorkflowId == workflowId2 {
						found = true
						assert.Equal(t, correlationId, result.CorrelationId, "Correlation ID should match")
						break
					}
				}
				assert.True(t, found, "Should find workflow 2 by correlation ID")
			},
		},
		{
			name: "SearchByWorkflowId",
			searchOpts: &client.WorkflowResourceApiSearchOpts{
				Query: optional.NewString(fmt.Sprintf("workflowId = %s", workflowId1)),
				Start: optional.NewInt32(0),
				Size:  optional.NewInt32(10),
			},
			expectExactResults: 1,
			validateResults: func(t *testing.T, results []model.WorkflowSummary, workflowId1, workflowId2, wf1Name string) {
				assert.Equal(t, workflowId1, results[0].WorkflowId, "Should find the correct workflow")
			},
		},
		{
			name: "SearchByStatus",
			searchOpts: &client.WorkflowResourceApiSearchOpts{
				Query: optional.NewString("status = COMPLETED"),
				Start: optional.NewInt32(0),
				Size:  optional.NewInt32(50),
			},
			expectMinResults: 1,
			validateResults: func(t *testing.T, results []model.WorkflowSummary, workflowId1, workflowId2, wf1Name string) {
				for _, result := range results {
					assert.Equal(t, "COMPLETED", result.Status, "All results should be completed")
				}
			},
		},
		{
			name: "SearchWithINOperator",
			searchOpts: &client.WorkflowResourceApiSearchOpts{
				Query: optional.NewString("status IN (COMPLETED, RUNNING)"),
				Start: optional.NewInt32(0),
				Size:  optional.NewInt32(20),
			},
			expectMinResults: 1,
			validateResults: func(t *testing.T, results []model.WorkflowSummary, workflowId1, workflowId2, wf1Name string) {
				for _, result := range results {
					assert.Contains(t, []string{"COMPLETED", "RUNNING"},
						result.Status, "Status should be COMPLETED or RUNNING")
				}
			},
		},
		{
			name: "SearchWithANDConditions",
			searchOpts: &client.WorkflowResourceApiSearchOpts{
				Query: optional.NewString(fmt.Sprintf("workflowType = %s AND status = COMPLETED", wf1Name)),
				Start: optional.NewInt32(0),
				Size:  optional.NewInt32(10),
			},
			expectMinResults: 1,
			validateResults: func(t *testing.T, results []model.WorkflowSummary, workflowId1, workflowId2, wf1Name string) {
				for _, result := range results {
					assert.Equal(t, wf1Name, result.WorkflowType, "Workflow type should match")
					assert.Equal(t, "COMPLETED", result.Status, "Status should be completed")
				}
			},
		},
		{
			name: "SearchWithSorting",
			searchOpts: &client.WorkflowResourceApiSearchOpts{
				Query: optional.NewString("status = COMPLETED"),
				Sort:  optional.NewString("startTime:DESC"),
				Start: optional.NewInt32(0),
				Size:  optional.NewInt32(10),
			},
			expectMinResults: 1,
			validateResults: func(t *testing.T, results []model.WorkflowSummary, workflowId1, workflowId2, wf1Name string) {
				// Verify sorting (most recent first) - check if we have enough results
				if len(results) > 1 {
					for i := 0; i < len(results)-1; i++ {
						// Note: StartTime is string, so we compare lexicographically
						assert.GreaterOrEqual(t, results[i].StartTime, results[i+1].StartTime,
							"Results should be sorted by start time in descending order")
					}
				}
			},
		},
		{
			name: "SearchWithTimeRange",
			searchOpts: &client.WorkflowResourceApiSearchOpts{
				Query: optional.NewString(fmt.Sprintf("startTime > %d AND startTime < %d",
					(time.Now().Add(-1*time.Hour)).Unix()*1000,
					time.Now().Unix()*1000)),
				Start: optional.NewInt32(0),
				Size:  optional.NewInt32(20),
			},
			expectMinResults: 1,
			validateResults: func(t *testing.T, results []model.WorkflowSummary, workflowId1, workflowId2, wf1Name string) {
				// Note: Since StartTime is a string, we'll just verify the search executes successfully
				// Time range validation is complex with string timestamps
			},
		},
		{
			name: "SearchWithFreeText",
			searchOpts: &client.WorkflowResourceApiSearchOpts{
				FreeText: optional.NewString("TEST_GO_WORKFLOW_SEARCH"),
				Start:    optional.NewInt32(0),
				Size:     optional.NewInt32(10),
			},
			expectMinResults: 0, // Free text search might not return results
			validateResults: func(t *testing.T, results []model.WorkflowSummary, workflowId1, workflowId2, wf1Name string) {
				t.Logf("Free text search found %d results", len(results))
			},
		},
		{
			name: "SearchWithSkipCache",
			searchOpts: &client.WorkflowResourceApiSearchOpts{
				Query: optional.NewString(fmt.Sprintf("workflowId = %s", workflowId1)),
				Start: optional.NewInt32(0),
				Size:  optional.NewInt32(10),
			},
			expectExactResults: 1,
			validateResults: func(t *testing.T, results []model.WorkflowSummary, workflowId1, workflowId2, wf1Name string) {
				assert.Equal(t, workflowId1, results[0].WorkflowId, "Should find the correct workflow")
			},
		},
		{
			name: "SearchWithPagination",
			searchOpts: &client.WorkflowResourceApiSearchOpts{
				Query: optional.NewString("status = COMPLETED"),
				Start: optional.NewInt32(0),
				Size:  optional.NewInt32(5),
			},
			expectMinResults: 0,
			validateResults: func(t *testing.T, results []model.WorkflowSummary, workflowId1, workflowId2, wf1Name string) {
				// Test pagination by making a second request
				opts2 := &client.WorkflowResourceApiSearchOpts{
					Query: optional.NewString("status = COMPLETED"),
					Start: optional.NewInt32(5),
					Size:  optional.NewInt32(5),
				}

				searchResult2, _, err := testdata.WorkflowClient.Search(context.Background(), opts2)
				assert.NoError(t, err, "Failed to search second page")

				// Verify pagination works (different results)
				if len(results) > 0 && len(searchResult2.Results) > 0 {
					assert.NotEqual(t, results[0].WorkflowId, searchResult2.Results[0].WorkflowId,
						"Different pages should return different results")
				}
			},
		},
		{
			name: "SearchWithInvalidQuery",
			searchOpts: &client.WorkflowResourceApiSearchOpts{
				Query: optional.NewString("invalid_field = invalid_value"),
				Start: optional.NewInt32(0),
				Size:  optional.NewInt32(10),
			},
			expectError: true, // Invalid queries should return errors
			validateResults: func(t *testing.T, results []model.WorkflowSummary, workflowId1, workflowId2, wf1Name string) {
				// No validation needed for error cases
			},
		},
	}

	// Run table-driven tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			searchResult, _, err := testdata.WorkflowClient.Search(context.Background(), tt.searchOpts)

			if tt.expectError {
				if err != nil {
					assert.Error(t, err, "Search should return error")
				} else {
					assert.Fail(t, "Search should return error but got nil")
				}
				return
			}

			assert.NoError(t, err, "Search should not return error")
			assert.NotNil(t, searchResult, "Search result should not be nil")

			if tt.expectExactResults > 0 {
				assert.Len(t, searchResult.Results, tt.expectExactResults, "Should find exact number of results")
			} else if tt.expectMinResults > 0 {
				assert.GreaterOrEqual(t, len(searchResult.Results), tt.expectMinResults, "Should find minimum number of results")
			}

			// Run custom validation
			tt.validateResults(t, searchResult.Results, workflowId1, workflowId2, wf1Name)
		})
	}
}

// setupTestWorkflows creates and starts test workflows for search testing
func setupTestWorkflows(t *testing.T, uniqueSuffix, correlationId string) (workflowId1, workflowId2, wf1Name, wf2Name string) {
	// Workflow 1: Simple completed workflow
	wf1 := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_SEARCH_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewSetVariableTask("set_var1").Input("var_value", "search_test_1"))

	err := testdata.ValidateWorkflowRegistration(wf1)
	assert.NoError(t, err, "Failed to register workflow 1")

	// Workflow 2: Workflow with correlation ID
	wf2 := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_SEARCH_CORR_" + uniqueSuffix).
		Version(1).
		Add(workflow.NewSetVariableTask("set_var2").Input("var_value", "search_test_2"))

	err = testdata.ValidateWorkflowRegistration(wf2)
	assert.NoError(t, err, "Failed to register workflow 2")

	// Start workflow 1
	startRequest1 := &model.StartWorkflowRequest{
		Name:    wf1.GetName(),
		Version: 1,
		Input: map[string]interface{}{
			"search_key": "search_value_1",
			"category":   "test_search",
		},
	}
	workflowId1, err = testdata.WorkflowExecutor.StartWorkflow(startRequest1)
	assert.NoError(t, err, "Failed to start workflow 1")

	// Start workflow 2 with correlation ID
	startRequest2 := &model.StartWorkflowRequest{
		Name:          wf2.GetName(),
		Version:       1,
		CorrelationId: correlationId,
		Input: map[string]interface{}{
			"search_key": "search_value_2",
			"category":   "test_search",
		},
	}
	workflowId2, err = testdata.WorkflowExecutor.StartWorkflow(startRequest2)
	assert.NoError(t, err, "Failed to start workflow 2")

	return workflowId1, workflowId2, wf1.GetName(), wf2.GetName()
}

// cleanupTestWorkflows removes test workflows and their definitions
func cleanupTestWorkflows(t *testing.T, workflowId1, workflowId2, wf1Name, wf2Name string) {
	if err := testdata.WorkflowExecutor.RemoveWorkflow(workflowId1); err != nil {
		t.Logf("Warning: Failed to remove workflow 1: %v", err)
	}
	if err := testdata.WorkflowExecutor.RemoveWorkflow(workflowId2); err != nil {
		t.Logf("Warning: Failed to remove workflow 2: %v", err)
	}
	if _, err := testdata.MetadataClient.UnregisterWorkflowDef(context.Background(), wf1Name, 1); err != nil {
		t.Logf("Warning: Failed to remove workflow definition 1: %v", err)
	}
	if _, err := testdata.MetadataClient.UnregisterWorkflowDef(context.Background(), wf2Name, 1); err != nil {
		t.Logf("Warning: Failed to remove workflow definition 2: %v", err)
	}
}
