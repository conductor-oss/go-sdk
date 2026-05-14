package metrics

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathTemplate_RoundTrip(t *testing.T) {
	ctx := context.Background()
	ctx = WithPathTemplate(ctx, "/workflow/{workflowId}")
	assert.Equal(t, "/workflow/{workflowId}", PathTemplateFromContext(ctx))
}

func TestPathTemplate_EmptyContext(t *testing.T) {
	ctx := context.Background()
	assert.Equal(t, "", PathTemplateFromContext(ctx))
}

func TestPathTemplate_NilContext(t *testing.T) {
	// context.TODO() behaves like an empty context
	assert.Equal(t, "", PathTemplateFromContext(context.TODO()))
}

func TestPathTemplate_Overwrite(t *testing.T) {
	ctx := context.Background()
	ctx = WithPathTemplate(ctx, "/workflow/{workflowId}")
	ctx = WithPathTemplate(ctx, "/tasks/{taskId}")
	assert.Equal(t, "/tasks/{taskId}", PathTemplateFromContext(ctx))
}

func TestPathTemplate_EmptyString(t *testing.T) {
	ctx := WithPathTemplate(context.Background(), "")
	assert.Equal(t, "", PathTemplateFromContext(ctx))
}

func TestPathTemplate_DoesNotAffectParent(t *testing.T) {
	parent := context.Background()
	child := WithPathTemplate(parent, "/workflow/{workflowId}")
	assert.Equal(t, "", PathTemplateFromContext(parent))
	assert.Equal(t, "/workflow/{workflowId}", PathTemplateFromContext(child))
}
