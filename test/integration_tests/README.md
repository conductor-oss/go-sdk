# Workflow API Integration Tests

This directory contains integration tests for the Conductor Workflow API endpoints. Each endpoint is tested to ensure proper functionality and compatibility with the Go SDK.

## Workflow API Endpoints and Tests

| HTTP Method | URL | Description | Test Function |
|-------------|-----|-------------|--------------|
| `PUT` | `/api/workflow/decide/{workflowId}` | Starts the decision task for a workflow | `TestRetryWorkflowWithDecide` |
| `DELETE` | `/api/workflow/{workflowId}/remove` | Removes the workflow from the system | `TestRemoveWorkflow` |
| `GET` | `/api/workflow/{workflowId}` | Gets the workflow by workflow id | Multiple tests |
| `GET` | `/api/workflow/{workflowId}/status` | Gets the workflow state including variables and output | `TestGetWorkflowState` |
| `GET` | `/api/workflow/running/{name}` | Retrieve all running workflows | `TestGetRunningWorkflow` |
| `POST` | `/api/workflow/{name}/correlated` | Lists workflows for given correlation id list | `TestGetWorkflows` |
| `POST` | `/api/workflow/correlated/batch` | Lists workflows for multiple correlation ids | `TestGetWorkflowsBatch` |
| `GET` | `/api/workflow/{name}/correlated/{correlationId}` | Lists workflows for a specific correlation id | `TestGetWorkflowsByCorrelationId` |
| `PUT` | `/api/workflow/{workflowId}/pause` | Pauses the workflow | `TestPauseResumeWorkflow` |
| `POST` | `/api/workflow/{workflowId}/rerun` | Reruns the workflow from a specific task | `TestRerunWorkflow` |
| `POST` | `/api/workflow/{workflowId}/resetcallbacks` | Resets callback times of non-terminal tasks | `TestResetWorkflow` |
| `POST` | `/api/workflow/{workflowId}/restart` | Restarts a completed workflow | `TestRestartWorkflow` |
| `PUT` | `/api/workflow/{workflowId}/resume` | Resumes the workflow | `TestPauseResumeWorkflow`|
| `POST` | `/api/workflow/{workflowId}/retry` | Retries the last failed task | `TestRetryWorkflow`, `TestRetryWorkflowWithDecide` |
| `GET` | `/api/workflow/search` | Search for workflows based on payload and parameters | `TestWorkflowSearch` |
| `PUT` | `/api/workflow/{workflowId}/skiptask/{taskReferenceName}` | Skips a given task from a running workflow | `TestSkipTaskFromWorkflow` |
| `POST` | `/api/workflow/{name}` | Start a new workflow | Multiple tests |
| `POST` | `/api/workflow/execute/{name}/{version}` | Execute a workflow and wait for completion | `TestExecuteWorkflow`, `TestExecuteWorkflowSync` |
| `POST` | `/api/workflow` | Start a workflow with a StartWorkflowRequest | Multiple tests `TestStartWorkflowWithRequest` |
| `POST` | `/api/workflow/execute/{name}/{version}` | Execute workflow with specified return strategy | `TestSignal_AllStrategies_Comprehensive` |
| `POST` | `/api/workflow/{workflowId}/jump/{taskReferenceName}` | Jump workflow execution to a given task | `TestJumpToTask` |
| `POST` | `/api/workflow/{workflowId}/state` | Update workflow variables, tasks and trigger evaluation | `TestUpdateWorkflowAndTaskState` |
| `POST` | `/api/workflow/{workflowId}/upgrade` | Upgrade running workflow to newer version | `TestUpgradeRunningWorkflowToVersion` |
| `POST` | `/api/workflow/test` | Test workflow execution using mock data | `TestWorkflowTest` |
| `GET` | `/api/workflow/{workflowId}/tasks` | Gets workflow tasks by workflow execution id | `TestGetExecutionStatusTaskList` |
| `POST` | `/api/workflow/{workflowId}/variables` | Updates workflow variables and triggers evaluation | `TestUpdateWorkflowState` |
| `DELETE` | `/api/workflow/{workflowId}` | Terminate workflow execution | `TestTerminateWorkflow`, `TestTerminateWorkflowWithFailure` |

## Workflow Bulk API Endpoints and Tests

| HTTP Method | URL | Description | Test Function |
|-------------|-----|-------------|--------------|
| `POST` | `/api/workflow/bulk/delete` | Permanently remove workflows from the system | `TestWorkflowBulkDelete` |
| `PUT` | `/api/workflow/bulk/pause` | Pause the list of workflows | `TestWorkflowBulkPause` |
| `POST` | `/api/workflow/bulk/restart` | Restart the list of completed workflow | `TestWorkflowBulkRestart` |
| `PUT` | `/api/workflow/bulk/resume` | Resume the list of workflows | `TestWorkflowBulkResume` |
| `POST` | `/api/workflow/bulk/retry` | Retry the last failed task for each workflow from the list | `TestWorkflowBulkRetry` |
| `POST` | `/api/workflow/bulk/terminate` | Terminate workflows execution | `TestWorkflowBulkTerminate` |

## Running Tests

To run all integration tests:

```bash
cd test/integration_tests
go test -v
```

To run a specific test:

```bash
go test -v -run TestTerminateWorkflow
```

## Test Requirements

These tests require a running Conductor server. The tests use the configuration defined in the `testdata` package to connect to the server.

## Test Coverage

Most of the Workflow API endpoints are covered by integration tests. Some endpoints that are not directly tested may be indirectly tested as part of other test flows.

## Adding New Tests

When adding new tests:

1. Follow the naming convention: `Test<EndpointName>`
2. Add proper cleanup to remove workflows and workflow definitions
3. Use the helper functions in the `testdata` package
4. Update this README to document the new test