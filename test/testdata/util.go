//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package testdata

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"

	"github.com/conductor-sdk/conductor-go/sdk/authentication"
	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/worker"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"

	"github.com/conductor-sdk/conductor-go/sdk/log"
	"go.uber.org/zap"
)

const (
	VersionResourceV41 = "4.1"
	VersionResourceV52 = "5.2"
)

var (
	apiClient = client.NewAPIClientWithTokenExpiration(
		client.NewAuthenticationSettingsFromEnv(),
		client.NewHttpSettingsFromEnv(),
		authentication.NewTokenExpiration(3*time.Second, 30*time.Second),
	)
	MetadataClient = client.MetadataResourceApiService{
		APIClient: apiClient,
	}
	TaskClient = client.TaskResourceApiService{
		APIClient: apiClient,
	}
	WorkflowClient = client.WorkflowResourceApiService{
		APIClient: apiClient,
	}
	EventClient = client.EventResourceApiService{
		APIClient: apiClient,
	}
	TagsClient = client.TagsApiService{
		APIClient: apiClient,
	}
	ApplicationClient     = client.NewApplicationClient(apiClient)
	AuthorizationClient   = client.NewAuthorizationClient(apiClient)
	EnvironmentClient     = client.NewEnvironmentClient(apiClient)
	HumanTaskClient       = client.NewHumanTaskClient(apiClient)
	IntegrationClient     = client.NewIntegrationClient(apiClient)
	PromptClient          = client.NewPromptClient(apiClient)
	UserClient            = client.NewUserClient(apiClient)
	GroupClient           = client.NewGroupClient(apiClient)
	SchedulerClient       = client.NewSchedulerClient(apiClient)
	SecretClient          = client.NewSecretsClient(apiClient)
	WebhookClient         = client.NewWebhooksConfigClient(apiClient)
	ServiceRegistryClient = client.NewServiceRegistryClient(apiClient)
	VersionResourceClient = client.NewVersionResourceClient(apiClient)
	WorkflowBulkClient    = client.NewWorkflowBulkClient(apiClient)
)

var TaskRunner = worker.NewTaskRunnerWithApiClient(apiClient)

var WorkflowExecutor = executor.NewWorkflowExecutor(apiClient)

// VersionResource is the version of the Conductor server
var VersionResource string

func init() {
	log.SetLogger(log.NewZap(zap.Must(zap.NewProduction())))

	version, _, err := VersionResourceClient.GetVersion(context.Background())
	if err != nil {
		log.Fatal("Failed to get version", "error", err)
	}

	VersionResource = parseVersion(version)
}

func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ValidateWorkflowDaemon(waitTime time.Duration, outputChannel chan error, workflowId string, expectedOutput map[string]interface{}, expectedStatus model.WorkflowStatus) {
	time.Sleep(waitTime)
	workflow, _, err := WorkflowClient.GetExecutionStatus(
		context.Background(),
		workflowId,
		nil,
	)
	if err != nil {
		outputChannel <- err
		return
	}
	if !isWorkflowCompleted(&workflow, expectedStatus) {
		outputChannel <- fmt.Errorf(
			"workflow status different than expected, workflowId: %s, workflowStatus: %s",
			workflow.WorkflowId, workflow.Status,
		)
		return
	}
	if !reflect.DeepEqual(workflow.Output, expectedOutput) {
		outputChannel <- fmt.Errorf(
			"workflow output is different than expected, workflowId: %s, output: %+v",
			workflow.WorkflowId, workflow.Output,
		)
		return
	}
	outputChannel <- nil
}

func StartWorkflows(workflowQty int, workflowName string) ([]string, error) {
	workflowIdList := make([]string, workflowQty)
	for i := 0; i < workflowQty; i += 1 {
		workflowId, _, err := WorkflowClient.StartWorkflow(
			context.Background(),
			make(map[string]interface{}),
			workflowName,
			nil,
		)
		if err != nil {
			return nil, err
		}
		log.Debug(
			"Started workflow",
			"workflow_name", workflowName,
			"workflow_id", workflowId,
		)
		workflowIdList[i] = workflowId
	}
	return workflowIdList, nil
}

func ValidateWorkflow(conductorWorkflow *workflow.ConductorWorkflow, timeout time.Duration, expectedStatus model.WorkflowStatus) error {
	err := ValidateWorkflowRegistration(conductorWorkflow)
	if err != nil {
		return err
	}
	workflowId, err := conductorWorkflow.StartWorkflowWithInput(make(map[string]interface{}))
	if err != nil {
		return err
	}
	log.Debug("Started workflowId", "workflow_id", workflowId)
	workflowExecutionChannel, err := WorkflowExecutor.MonitorExecution(workflowId)
	if err != nil {
		return err
	}
	log.Debug("Generated workflowExecutionChannel for workflowId", "workflow_id", workflowId)
	workflow, err := executor.WaitForWorkflowCompletionUntilTimeout(
		workflowExecutionChannel,
		timeout,
	)
	if err != nil {
		return err
	}
	log.Debug("Workflow completed", "workflow_id", workflowId)
	if !isWorkflowCompleted(workflow, expectedStatus) {
		return fmt.Errorf("workflow finished with unexpected status: %s", workflow.Status)
	}
	return nil
}

func ValidateWorkflowBulk(conductorWorkflow *workflow.ConductorWorkflow, timeout time.Duration, amount int) error {
	err := ValidateWorkflowRegistration(conductorWorkflow)
	if err != nil {
		return err
	}
	version := conductorWorkflow.GetVersion()
	startWorkflowRequests := make([]*model.StartWorkflowRequest, amount)
	for i := 0; i < amount; i += 1 {
		startWorkflowRequests[i] = model.NewStartWorkflowRequest(
			conductorWorkflow.GetName(),
			version,
			"",
			make(map[string]interface{}),
		)
	}
	runningWorkflows := WorkflowExecutor.StartWorkflows(true, startWorkflowRequests...)
	WorkflowExecutor.WaitForRunningWorkflowsUntilTimeout(timeout, runningWorkflows...)
	for _, runningWorkflow := range runningWorkflows {
		if runningWorkflow.Err != nil {
			return err
		}
		if runningWorkflow.CompletedWorkflow == nil {
			return fmt.Errorf("invalid completed workflows")
		}
		if !isWorkflowCompleted(runningWorkflow.CompletedWorkflow, model.CompletedWorkflow) {
			return fmt.Errorf("workflow finished with status: %s", runningWorkflow.CompletedWorkflow.Status)
		}
	}
	return nil
}

func ValidateTaskRegistration(taskDefs ...model.TaskDef) error {
	response, err := MetadataClient.RegisterTaskDef(
		context.Background(),
		taskDefs,
	)
	if err != nil {
		log.Debug(
			"Failed to validate task registration",
			"reason", os.ErrClosed,
			"response", *response,
		)
		return err
	}
	return nil
}

func ValidateWorkflowRegistration(workflow *workflow.ConductorWorkflow) error {
	for attempt := 0; attempt < 5; attempt += 1 {
		err := workflow.Register(true)
		if err != nil {
			time.Sleep(time.Duration(attempt+2) * time.Second)
			fmt.Println("Failed to validate workflow registration, reason: " + err.Error())
			continue
		}
		return nil
	}
	return fmt.Errorf("exhausted retries")
}

func ValidateWorkflowDeletion(workflow *workflow.ConductorWorkflow) error {
	for attempt := 0; attempt < 5; attempt += 1 {
		err := workflow.UnRegister()
		if err != nil {
			time.Sleep(time.Duration(attempt+2) * time.Second)
			fmt.Println("Failed to validate workflow deletion, reason: " + err.Error())
			continue
		}
		return nil
	}
	return fmt.Errorf("exhausted retries")
}

func CreateNewUser(ctx context.Context) (rbac.ConductorUser, error) {
	// Generate random suffix for username and ID
	randomSuffix := fmt.Sprintf("%d", time.Now().UnixNano())

	body := rbac.UpsertUserRequest{
		Name:  fmt.Sprintf("testuser-%s", randomSuffix),
		Roles: []string{"ADMIN", "USER"},
	}
	id := "testUser"

	user, _, err := UserClient.UpsertUser(ctx, body, id)
	if err != nil {
		fmt.Printf("Unable to create new user. %v", err)
		return rbac.ConductorUser{}, err
	}
	return *user, nil
}

func isWorkflowCompleted(workflow *model.Workflow, expectedStatus model.WorkflowStatus) bool {
	return workflow.Status == expectedStatus
}

// WaitForWorkflowStatus waits for a workflow to reach any of the specified statuses
func WaitForWorkflowStatus(workflowId string, expectedStatuses []model.WorkflowStatus, timeout time.Duration) (*model.Workflow, error) {
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		var workflow *model.Workflow
		err := RetryTimeout(2, 500*time.Millisecond, func() error {
			var getErr error
			workflow, getErr = WorkflowExecutor.GetWorkflow(workflowId, true)
			return getErr
		})
		if err != nil {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		// Check if workflow has reached any of the expected statuses
		for _, expectedStatus := range expectedStatuses {
			if workflow.Status == expectedStatus {
				return workflow, nil
			}
		}

		// Check for terminal failure states if not explicitly expected
		isFailureExpected := false
		for _, expectedStatus := range expectedStatuses {
			if expectedStatus == model.FailedWorkflow || expectedStatus == model.TerminatedWorkflow {
				isFailureExpected = true
				break
			}
		}

		if !isFailureExpected && (workflow.Status == model.FailedWorkflow || workflow.Status == model.TerminatedWorkflow) {
			return workflow, fmt.Errorf("workflow %s failed with unexpected status: %s", workflowId, workflow.Status)
		}

		time.Sleep(500 * time.Millisecond)
	}

	return nil, fmt.Errorf("workflow %s didn't reach any of the expected statuses %v within %v", workflowId, expectedStatuses, timeout)
}

// WaitForWorkflowCompletion waits specifically for workflow completion
func WaitForWorkflowCompletion(workflowId string, timeout time.Duration) (*model.Workflow, error) {
	return WaitForWorkflowStatus(workflowId, []model.WorkflowStatus{model.CompletedWorkflow}, timeout)
}

// WaitForWorkflowRunning waits for workflow to reach RUNNING status
func WaitForWorkflowRunning(workflowId string, timeout time.Duration) (*model.Workflow, error) {
	return WaitForWorkflowStatus(workflowId, []model.WorkflowStatus{model.RunningWorkflow}, timeout)
}

// WaitForMultipleWorkflowsCompletion waits for multiple workflows to complete
func WaitForMultipleWorkflowsCompletion(workflowIds []string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		allCompleted := true

		for _, workflowId := range workflowIds {
			// Use RetryWithBackoff for executing request with retries
			var workflow *model.Workflow
			err := RetryTimeout(2, 500*time.Millisecond, func() error {
				var getErr error
				workflow, getErr = WorkflowExecutor.GetWorkflow(workflowId, false)
				return getErr
			})
			if err != nil {
				// Continue retrying on error
				allCompleted = false
				break
			}

			if workflow.Status == model.FailedWorkflow || workflow.Status == model.TerminatedWorkflow {
				return fmt.Errorf("workflow %s failed with status: %s", workflowId, workflow.Status)
			}

			if workflow.Status != model.CompletedWorkflow {
				allCompleted = false
				break
			}
		}

		if allCompleted {
			return nil
		}

		time.Sleep(1 * time.Second)
	}

	return fmt.Errorf("not all workflows completed within %v", timeout)
}

// WaitForMultipleWorkflowsStatus waits for multiple workflows to reach any of the expected statuses
func WaitForMultipleWorkflowsStatus(workflowIds []string, expectedStatuses []model.WorkflowStatus, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		allReachedExpectedStatus := true

		for _, workflowId := range workflowIds {
			// Use RetryWithBackoff for executing request with retries
			var workflow *model.Workflow
			err := RetryTimeout(2, 500*time.Millisecond, func() error {
				var getErr error
				workflow, getErr = WorkflowExecutor.GetWorkflow(workflowId, false)
				return getErr
			})

			if err != nil {
				// Log error and continue
				allReachedExpectedStatus = false
				break
			}

			// Check if current workflow status is in the expected statuses
			statusMatched := false
			for _, expectedStatus := range expectedStatuses {
				if workflow.Status == expectedStatus {
					statusMatched = true
					break
				}
			}

			// If workflow status is not in expected statuses
			if !statusMatched {
				// Check if failure is expected
				isFailureExpected := false
				for _, expectedStatus := range expectedStatuses {
					if expectedStatus == model.FailedWorkflow || expectedStatus == model.TerminatedWorkflow {
						isFailureExpected = true
						break
					}
				}

				// If failure is not expected but workflow failed
				if !isFailureExpected && (workflow.Status == model.FailedWorkflow || workflow.Status == model.TerminatedWorkflow) {
					return fmt.Errorf("workflow %s failed with unexpected status: %s", workflowId, workflow.Status)
				}

				allReachedExpectedStatus = false
				break
			}
		}

		if allReachedExpectedStatus {
			return nil
		}

		time.Sleep(1 * time.Second)
	}

	return fmt.Errorf("not all workflows reached expected statuses %v within %v", expectedStatuses, timeout)
}

// WaitForTaskInWorkflow waits for a specific task within a workflow to reach a certain status
func WaitForTaskInWorkflow(workflowId, taskReferenceName string, expectedTaskStatus model.TaskResultStatus, timeout time.Duration) (*model.Task, error) {
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		workflow, err := WorkflowExecutor.GetWorkflow(workflowId, true) // Include tasks
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		// Look for the specific task
		for _, task := range workflow.Tasks {
			if task.ReferenceTaskName == taskReferenceName && task.Status == expectedTaskStatus {
				return &task, nil
			}
		}

		time.Sleep(1 * time.Second)
	}

	return nil, fmt.Errorf("task %s with status %s not found in workflow %s within %v",
		taskReferenceName, expectedTaskStatus, workflowId, timeout)
}

// Common Test Tasks
const (
	WorkflowValidationTimeout = 7 * time.Second
	ExtendedValidationTimeout = 15 * time.Second
	WorkflowBulkQty           = 10
)

var (
	TestHttpTask = workflow.NewHttpTask(
		"TEST_GO_TASK_HTTP",
		&workflow.HttpInput{
			Uri: "https://orkes-api-tester.orkesconductor.com/get",
		},
	)

	TestSimpleTask = workflow.NewSimpleTask(
		"TEST_GO_TASK_SIMPLE", "TEST_GO_TASK_SIMPLE",
	)

	TestTerminateTask = workflow.NewTerminateTask(
		"TEST_GO_TASK_TERMINATE",
		model.FailedWorkflow,
		"Task used to mark workflow as failed",
	)

	TestSwitchTask = workflow.NewSwitchTask(
		"TEST_GO_TASK_SWITCH",
		"switchCaseValue",
	).
		Input("switchCaseValue", "${workflow.input.service}").
		UseJavascript(true).
		SwitchCase(
			"REQUEST",
			TestHttpTask,
		).
		SwitchCase(
			"STOP",
			TestTerminateTask,
		)

	TestInlineTask = workflow.NewInlineTask(
		"TEST_GO_TASK_INLINE",
		"function e() { if ($.value == 1){return {\"result\": true}} else { return {\"result\": false}}} e();",
	)

	TestKafkaPublishTask = workflow.NewKafkaPublishTask(
		"TEST_GO_TASK_KAFKA_PUBLISH",
		&workflow.KafkaPublishTaskInput{
			Topic:            "userTopic",
			Value:            "Message to publish",
			BootStrapServers: "localhost:9092",
			Headers: map[string]interface{}{
				"x-Auth": "Auth-key",
			},
			Key:           "123",
			KeySerializer: "org.apache.kafka.common.serialization.IntegerSerializer",
		},
	)

	TestSqsEventTask = workflow.NewSqsEventTask(
		"TEST_GO_TASK_EVENT_SQS",
		"QUEUE",
	)

	TestConductorEventTask = workflow.NewConductorEventTask(
		"TEST_GO_TASK_EVENT_CONDUCTOR",
		"EVENT_NAME",
	)
)

func ValidateWorkflowWithOutput(conductorWorkflow *workflow.ConductorWorkflow, timeout time.Duration, expectedStatus model.WorkflowStatus, expectedOutput map[string]interface{}) error {
	err := ValidateWorkflowRegistration(conductorWorkflow)
	if err != nil {
		return err
	}
	workflowId, err := conductorWorkflow.StartWorkflowWithInput(make(map[string]interface{}))
	if err != nil {
		return err
	}
	log.Debug("Started workflowId: ", workflowId)

	workflowExecutionChannel, err := WorkflowExecutor.MonitorExecution(workflowId)
	if err != nil {
		return err
	}
	completed, err := executor.WaitForWorkflowCompletionUntilTimeout(workflowExecutionChannel, timeout)
	if err != nil {
		return err
	}
	if !isWorkflowCompleted(completed, expectedStatus) {
		return fmt.Errorf("workflow finished with unexpected status: %s", completed.Status)
	}

	wf, _, err := WorkflowClient.GetExecutionStatus(context.Background(), workflowId, nil)
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(wf.Output, expectedOutput) {
		return fmt.Errorf("workflow output is different than expected, workflowId: %s, output: %+v", workflowId, wf.Output)
	}
	return nil
}

func parseVersion(version string) string {
	parts := strings.Split(version, ".")
	if len(parts) >= 2 {
		return parts[0] + "." + parts[1]
	}
	return version
}

func RequireAtLeast(t *testing.T, min string) {
	t.Helper()
	have := VersionResource

	if !isVersionAtLeast(have, min) {
		t.Skipf("skip: requires >= %s, have %s", min, have)
	}
}

func isVersionAtLeast(have, min string) bool {
	haveParts := strings.Split(have, ".")
	minParts := strings.Split(min, ".")

	if len(haveParts) > 0 && len(minParts) > 0 {
		haveMajor, _ := strconv.Atoi(haveParts[0])
		minMajor, _ := strconv.Atoi(minParts[0])
		if haveMajor < minMajor {
			return false
		}
		if haveMajor > minMajor {
			return true
		}
	}

	if len(haveParts) > 1 && len(minParts) > 1 {
		haveMinor, _ := strconv.Atoi(haveParts[1])
		minMinor, _ := strconv.Atoi(minParts[1])
		return haveMinor >= minMinor
	}

	return true
}
