package integration_tests

import (
	"context"
	"testing"

	"github.com/antihax/optional"
	sdkclient "github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/model/human"
	"github.com/conductor-sdk/conductor-go/test/testdata"
	"github.com/stretchr/testify/assert"
)

func TestHumanTask_SearchWithRealTasks(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV52)
	client := testdata.HumanTaskClient
	ctx := context.Background()

	// Now test Search with real data
	search := human.HumanTaskSearch{
		SearchType: "ADMIN",
		Size:       10,
	}
	result, _, err := client.Search(ctx, search)
	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Greater(t, len(result.Results), 0)

	// Test GetTaskDisplayNames with real context
	displayNames, _, err := client.GetTaskDisplayNames(ctx, "INBOX")
	assert.Nil(t, err)
	assert.NotNil(t, displayNames)

	taskEntry, _, err := client.GetTask(ctx, result.Results[0].TaskId, nil)
	assert.Nil(t, err)

	assert.NotNil(t, taskEntry)

	// Verify that we got the correct task
	assert.Equal(t, result.Results[0].TaskId, taskEntry.TaskId)
}

func TestHumanTask_Templates_ListEmpty(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV52)
	client := testdata.HumanTaskClient
	ctx := context.Background()

	templates, _, err := client.GetAllTemplates(ctx, nil)

	assert.Nil(t, err)

	assert.NotNil(t, templates)
}

func TestHumanTask_SaveTemplate_Roundtrip(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV52)
	client := testdata.HumanTaskClient
	ctx := context.Background()

	// Save a minimal template, then fetch by name/version
	template := human.HumanTaskSearch{FullTextQuery: "ping"}

	// newVersion=false ensures idempotent behavior on re-runs
	opts := &sdkclient.HumanTaskApiSaveTemplateOpts{NewVersion: optional.NewBool(false)}
	_, _, err := client.SaveTemplate(ctx, template, opts)
	assert.Nil(t, err)

	got, _, err := client.GetTemplateByNameAndVersion(ctx, "search_template", int32(1))
	assert.Nil(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, template.FullTextQuery, got.FullTextQuery)

	// cleanup best-effort
	_, err = client.DeleteTemplatesByNameAndVersion(ctx, "search_template", int32(1))
	assert.Nil(t, err)
}

// TestHumanTask_TemplateOperations tests template-related operations
func TestHumanTask_TemplateOperations(t *testing.T) {
	testdata.RequireAtLeast(t, testdata.VersionResourceV52)
	client := testdata.HumanTaskClient
	ctx := context.Background()

	// Test SaveTemplates - save multiple templates
	templates := []human.HumanTaskSearch{
		{FullTextQuery: "template1", SearchType: "INBOX"},
		{FullTextQuery: "template2", SearchType: "ADMIN"},
	}
	opts := &sdkclient.HumanTaskApiSaveTemplatesOpts{NewVersion: optional.NewBool(false)}
	savedTemplates, _, err := client.SaveTemplates(ctx, templates, opts)
	assert.Nil(t, err)

	assert.NotNil(t, savedTemplates)
	assert.Equal(t, len(templates), len(savedTemplates))
	assert.Equal(t, templates[0].FullTextQuery, savedTemplates[0].FullTextQuery)
	assert.Equal(t, templates[1].FullTextQuery, savedTemplates[1].FullTextQuery)

	// Test DeleteTemplateByName
	resp, err := client.DeleteTemplateByName(ctx, "search_template")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}
