package integration_tests

import (
	"context"
	"testing"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTagsClient_WorkflowTags(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)
	tagsClient := testdata.TagsClient
	ctx := context.Background()

	// Create a simple workflow for testing
	workflowDef := workflow.NewConductorWorkflow(testdata.WorkflowExecutor).
		Name("TEST_GO_WORKFLOW_TAGS").
		Version(1).
		Add(testdata.TestSimpleTask)

	err := testdata.ValidateWorkflowRegistration(workflowDef)
	require.NoError(t, err)

	workflowName := workflowDef.GetName()

	// Test GetWorkflowTags - should return empty initially
	tags, resp, err := tagsClient.GetWorkflowTags(ctx, workflowName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, tags)
	assert.Empty(t, tags)

	// Test AddWorkflowTag
	tag := model.TagObject{
		Key:   "environment",
		Type:  "METADATA",
		Value: "test",
	}
	_, resp, err = tagsClient.AddWorkflowTag(ctx, tag, workflowName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test GetWorkflowTags - should now have the tag
	tags, resp, err = tagsClient.GetWorkflowTags(ctx, workflowName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.Equal(t, "environment", tags[0].Key)
	assert.Equal(t, "test", tags[0].Value)

	// Test SetWorkflowTags - replace all tags
	newTags := []model.TagObject{
		{Key: "team", Type: "METADATA", Value: "backend"},
		{Key: "priority", Type: "METADATA", Value: "high"},
	}
	_, resp, err = tagsClient.SetWorkflowTags(ctx, newTags, workflowName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify tags were replaced
	tags, resp, err = tagsClient.GetWorkflowTags(ctx, workflowName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, tags, 2)

	// Test DeleteWorkflowTag
	tagToDelete := model.TagObject{
		Key:   "team",
		Type:  "METADATA",
		Value: "backend",
	}
	_, resp, err = tagsClient.DeleteWorkflowTag(ctx, tagToDelete, workflowName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify tag was deleted
	tags, resp, err = tagsClient.GetWorkflowTags(ctx, workflowName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, tags, 1)
	assert.Equal(t, "priority", tags[0].Key)

	// Cleanup
	err = testdata.ValidateWorkflowDeletion(workflowDef)
	assert.Nil(t, err)
}

func TestTagsClient_TaskTags(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)
	tagsClient := testdata.TagsClient
	ctx := context.Background()

	// Register a task definition for testing
	taskDef := model.TaskDef{
		Name:           "TEST_GO_TASK_TAGS",
		Description:    "Test task for tags",
		RetryCount:     3,
		TimeoutSeconds: 60,
	}

	err := testdata.ValidateTaskRegistration(taskDef)
	require.NoError(t, err)

	taskName := taskDef.Name

	// Clean up any existing tags first
	existingTags, resp, err := tagsClient.GetTaskTags(ctx, taskName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, existingTags)
	_, resp, err = tagsClient.SetTaskTags(ctx, []model.TagObject{}, taskName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test GetTaskTags - should return empty initially
	tags, resp, err := tagsClient.GetTaskTags(ctx, taskName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, tags)
	assert.Empty(t, tags)

	// Test AddTaskTag
	tag := model.TagObject{
		Key:   "category",
		Type:  "METADATA",
		Value: "integration",
	}
	_, resp, err = tagsClient.AddTaskTag(ctx, tag, taskName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Test GetTaskTags - should now have the tag
	tags, resp, err = tagsClient.GetTaskTags(ctx, taskName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, tags)
	assert.Len(t, tags, 1)
	assert.Equal(t, "category", tags[0].Key)
	assert.Equal(t, "integration", tags[0].Value)

	// Test SetTaskTags - replace all tags
	newTags := []model.TagObject{
		{Key: "owner", Type: "METADATA", Value: "team-a"},
		{Key: "version", Type: "METADATA", Value: "1.0"},
	}
	_, resp, err = tagsClient.SetTaskTags(ctx, newTags, taskName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify tags were replaced
	tags, resp, err = tagsClient.GetTaskTags(ctx, taskName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, tags, 2)

	// Test DeleteTaskTag using TagString
	tagString := model.TagString{
		Key:   "owner",
		Type_: "METADATA",
		Value: "team-a",
	}
	_, resp, err = tagsClient.DeleteTaskTag(ctx, tagString, taskName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify tag was deleted
	tags, resp, err = tagsClient.GetTaskTags(ctx, taskName)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, tags, 1)
	assert.Equal(t, "version", tags[0].Key)
}

func TestTagsClient_GetAllTags(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV41)
	tagsClient := testdata.TagsClient
	ctx := context.Background()

	// Test GetTags - should return all tags in the system
	tags, resp, err := tagsClient.GetTags(ctx)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, tags)
	// Note: We don't assert on the length since there might be existing tags from other tests
}
