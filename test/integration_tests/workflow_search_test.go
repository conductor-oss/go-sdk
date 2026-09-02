// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.
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
	"github.com/stretchr/testify/require"
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
		// ossSkipReason, when set, skips this case against plain OSS
		// Conductor -- see oss_gaps_test.go.
		ossSkipReason string
	}{
		{
			name: "BasicSearch",
			searchOpts: &client.WorkflowResourceApiSearchOpts{
				Start: optional.NewInt32(0),
				Size:  optional.NewInt32(10),
			},
			expectMinResults: 2,
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
				require.NotEmpty(t, results, "Should find at least one result")
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
				// Define the time range used in the search query
				startTimeRange := time.Now().Add(-1 * time.Hour)
				endTimeRange := time.Now()

				t.Logf("Search time range (local): %s to %s", startTimeRange.Format(time.RFC3339), endTimeRange.Format(time.RFC3339))
				t.Logf("Search time range (UTC):   %s to %s", startTimeRange.UTC().Format(time.RFC3339), endTimeRange.UTC().Format(time.RFC3339))
				t.Logf("Search time range (ms):    %d to %d", startTimeRange.Unix()*1000, endTimeRange.Unix()*1000)
				t.Logf("Found %d results:\n", len(results))

				for _, result := range results {
					// Every result should have a StartTime
					require.NotEmpty(t, result.StartTime, "Result should have a StartTime")

					// Parse the ISO timestamp and validate it's within range
					parsedTime, err := time.Parse(time.RFC3339, result.StartTime)
					require.NoError(t, err, "Failed to parse StartTime '%s'", result.StartTime)

					t.Logf("Result StartTime (raw):   %s", result.StartTime)
					t.Logf("Result StartTime (UTC):   %s", parsedTime.UTC().Format(time.RFC3339))
					t.Logf("Result StartTime (ms):    %d", parsedTime.Unix()*1000)

					// Validate the timestamp is within the search range
					require.False(t, parsedTime.Before(startTimeRange),
						"StartTime %s (UTC) should not be before search range start %s (UTC)",
						parsedTime.UTC().Format(time.RFC3339), startTimeRange.UTC().Format(time.RFC3339))
					require.False(t, parsedTime.After(endTimeRange),
						"StartTime %s (UTC) should not be after search range end %s (UTC)",
						parsedTime.UTC().Format(time.RFC3339), endTimeRange.UTC().Format(time.RFC3339))

					t.Logf("✓ StartTime is within range")
				}
			},
		},
		{
			name: "SearchWithFreeText_NoExactMatch",
			searchOpts: &client.WorkflowResourceApiSearchOpts{
				FreeText: optional.NewString("TEST_GO_WORKFLOW_SEARCH"),
				Start:    optional.NewInt32(0),
				Size:     optional.NewInt32(10),
			},
			expectMinResults: 0,
			validateResults: func(t *testing.T, results []model.WorkflowSummary, workflowId1, workflowId2, wf1Name string) {
				//Free text search only searches for exact matches.
			},
		},
		{
			name: "SearchWithFreeText_WorkflowInputValues",
			searchOpts: &client.WorkflowResourceApiSearchOpts{
				FreeText: optional.NewString("search_value_1" + uniqueSuffix), // Search in workflow input values
				Start:    optional.NewInt32(0),
				Size:     optional.NewInt32(10),
			},
			expectMinResults: 1, // Should find at least one workflow with this input value
			validateResults: func(t *testing.T, results []model.WorkflowSummary, workflowId1, workflowId2, wf1Name string) {
				// Verify we found the correct workflow
				found := false
				for _, result := range results {
					if result.WorkflowId == workflowId1 {
						found = true
						break
					}
				}
				assert.True(t, found, "Should find workflow by FreeText search in workflow input values")
			},
		},
		{
			name: "SearchWithPagination",
			searchOpts: &client.WorkflowResourceApiSearchOpts{
				Query: optional.NewString(fmt.Sprintf("status = COMPLETED AND workflowType IN (%s, %s)", wf1Name, wf2Name)),
				Start: optional.NewInt32(0),
				Size:  optional.NewInt32(1),
			},
			expectMinResults: 1,
			validateResults: func(t *testing.T, results []model.WorkflowSummary, workflowId1, workflowId2, wf1Name string) {
				// Test pagination with page size 1 - first page should have one result
				assert.Len(t, results, 1, "First page should have exactly 1 result")

				// Verify the first page contains one of our test workflows
				firstPageWorkflowId := results[0].WorkflowId
				assert.True(t, firstPageWorkflowId == workflowId1 || firstPageWorkflowId == workflowId2,
					"First page should contain one of our test workflows")

				// Test second page with size 1. Only the start>0 fetch is
				// gated: the first-page assertions above (page size honored,
				// correct workflow returned) do pass against plain OSS
				// Conductor. See ossGapSearchPaging.
				if testdata.OSSGapSkipped() {
					t.Logf("skipping second-page assertions: %s", ossGapSearchPaging)
					return
				}

				opts2 := &client.WorkflowResourceApiSearchOpts{
					Query: optional.NewString(fmt.Sprintf("status = COMPLETED AND workflowType IN (%s, %s)", wf1Name, wf2Name)),
					Start: optional.NewInt32(1),
					Size:  optional.NewInt32(1),
				}

				searchResult2, _, err := testdata.WorkflowClient.Search(context.Background(), opts2)
				assert.NoError(t, err, "Failed to search second page")
				assert.Len(t, searchResult2.Results, 1, "Second page should have exactly 1 result")

				// Verify the second page contains the other test workflow
				secondPageWorkflowId := searchResult2.Results[0].WorkflowId
				assert.True(t, secondPageWorkflowId == workflowId1 || secondPageWorkflowId == workflowId2,
					"Second page should contain one of our test workflows")

				// Verify different workflows on different pages
				assert.NotEqual(t, firstPageWorkflowId, secondPageWorkflowId,
					"First and second pages should contain different workflows")
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
			if tt.ossSkipReason != "" {
				testdata.SkipIfOSS(t, tt.ossSkipReason)
			}

			var searchResult model.SearchResultWorkflowSummary
			var err error

			// For tests with minimum results requirement, use retry mechanism
			if tt.expectMinResults > 0 && !tt.expectError {
				err = testdata.RetryCondition(3, time.Second,
					func() error {
						// Execute search
						searchResult, _, err = testdata.WorkflowClient.Search(context.Background(), tt.searchOpts)
						return err
					},
					func() bool {
						return len(searchResult.Results) >= tt.expectMinResults
					})

				// Verify retry results
				assert.NoError(t, err, "Search should eventually return enough results")
				assert.GreaterOrEqual(t, len(searchResult.Results), tt.expectMinResults,
					"Should find minimum number of results after retries")
			} else {
				// For other tests, use standard search without retries
				searchResult, _, err = testdata.WorkflowClient.Search(context.Background(), tt.searchOpts)

				if tt.expectError {
					assert.Error(t, err, "Search should return error for invalid query")
					return
				}

				assert.NoError(t, err, "Search should not return error")
				assert.NotNil(t, searchResult, "Search result should not be nil")

				if tt.expectExactResults > 0 {
					assert.Len(t, searchResult.Results, tt.expectExactResults, "Should find exact number of results")
				}
			}

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
			"search_key": "search_value_1" + uniqueSuffix,
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
	err := testdata.WorkflowExecutor.RemoveWorkflow(workflowId1)
	assert.NoError(t, err, "Failed to remove workflow 1")
	err = testdata.WorkflowExecutor.RemoveWorkflow(workflowId2)
	assert.NoError(t, err, "Failed to remove workflow 2")

	_, err = testdata.MetadataClient.UnregisterWorkflowDef(context.Background(), wf1Name, 1)
	assert.NoError(t, err, "Failed to remove workflow definition 1")
	_, err = testdata.MetadataClient.UnregisterWorkflowDef(context.Background(), wf2Name, 1)
	assert.NoError(t, err, "Failed to remove workflow definition 2")
}
