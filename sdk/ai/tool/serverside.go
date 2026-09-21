//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package tool

import (
	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/internal/schema"
)

// More tools the server runs itself. Defaults and config keys match the Python
// SDK's api_tool, index_tool, search_tool and wait_for_message_tool.

const (
	defaultAPIName          = "api_tools"
	defaultAPIMaxTools      = 64
	defaultRAGNamespace     = "default_ns"
	defaultSearchMaxResults = 5
	defaultWaitBatchSize    = 1
)

// API exposes the operations of an OpenAPI 3.x spec, a Swagger 2.0 spec or a
// Postman collection at url as tools. Like MCP it is entirely server-side: the
// server fetches the spec at compile time and expands this definition into a
// tool per operation, filtering with an LLM beyond WithMaxTools. A ${NAME}
// credential in a header needs the same name in WithCredentials or Validate
// rejects the tool. Empty name or description: "api_tools", "API tools from <url>".
func API(name, description, url string, opts ...Option) ai.ToolDef {
	if name == "" {
		name = defaultAPIName
	}
	if description == "" {
		description = "API tools from " + url
	}
	td := ai.ToolDef{
		Name:        name,
		Description: description,
		InputSchema: map[string]any{},
		ToolType:    ai.ToolTypeAPI,
		Config:      map[string]any{"url": url, "max_tools": defaultAPIMaxTools},
	}
	for _, o := range opts {
		o(&td)
	}
	return td
}

// Index adds text to a vector index, Conductor's LLM_INDEX_TEXT task: the model
// supplies text and a document id, the rest is fixed here, namespace "default_ns".
func Index(name, description, vectorDB, index, embeddingProvider, embeddingModel string, opts ...Option) ai.ToolDef {
	td := ai.ToolDef{
		Name:        name,
		Description: description,
		InputSchema: indexSchema(),
		ToolType:    ai.ToolTypeRAGIndex,
		Config: map[string]any{
			"taskType":               "LLM_INDEX_TEXT",
			"vectorDB":               vectorDB,
			"namespace":              defaultRAGNamespace,
			"index":                  index,
			"embeddingModelProvider": embeddingProvider,
			"embeddingModel":         embeddingModel,
		},
	}
	for _, o := range opts {
		o(&td)
	}
	return td
}

// Search queries a vector index, Conductor's LLM_SEARCH_INDEX task, returning
// up to WithMaxResults matches (default 5). Other settings are as for Index.
func Search(name, description, vectorDB, index, embeddingProvider, embeddingModel string, opts ...Option) ai.ToolDef {
	td := ai.ToolDef{
		Name:        name,
		Description: description,
		InputSchema: searchSchema(),
		ToolType:    ai.ToolTypeRAGSearch,
		Config: map[string]any{
			"taskType":               "LLM_SEARCH_INDEX",
			"vectorDB":               vectorDB,
			"namespace":              defaultRAGNamespace,
			"index":                  index,
			"embeddingModelProvider": embeddingProvider,
			"embeddingModel":         embeddingModel,
			"maxResults":             defaultSearchMaxResults,
		},
	}
	for _, o := range opts {
		o(&td)
	}
	return td
}

// WaitForMessage waits for messages sent into the execution with
// AgentClient.Signal or the workflow message queue, blocking for one unless
// WithBatchSize or NonBlocking says otherwise. Needs the server's
// conductor.workflow-message-queue.enabled=true.
func WaitForMessage(name, description string, opts ...Option) ai.ToolDef {
	td := ai.ToolDef{
		Name:        name,
		Description: description,
		InputSchema: schema.EmptyObject(),
		ToolType:    ai.ToolTypePullWorkflowMessages,
		Config:      map[string]any{"batchSize": defaultWaitBatchSize},
	}
	for _, o := range opts {
		o(&td)
	}
	return td
}

// WithNamespace sets the partition within the vector index for Index and Search.
func WithNamespace(namespace string) Option {
	return func(t *ai.ToolDef) { setConfig(t, "namespace", namespace) }
}

// WithChunking splits Index's text into chunks of size characters overlapping by overlap.
func WithChunking(size, overlap int) Option {
	return func(t *ai.ToolDef) {
		setConfig(t, "chunkSize", size)
		setConfig(t, "chunkOverlap", overlap)
	}
}

// WithDimensions overrides the embedding size for Index and Search.
func WithDimensions(n int) Option {
	return func(t *ai.ToolDef) { setConfig(t, "dimensions", n) }
}

// WithMaxResults caps how many matches Search returns.
func WithMaxResults(n int) Option {
	return func(t *ai.ToolDef) { setConfig(t, "maxResults", n) }
}

// WithBatchSize is how many messages WaitForMessage takes per call.
func WithBatchSize(n int) Option {
	return func(t *ai.ToolDef) { setConfig(t, "batchSize", n) }
}

// NonBlocking makes WaitForMessage return at once, empty if nothing is queued.
func NonBlocking() Option {
	return func(t *ai.ToolDef) { setConfig(t, "blocking", false) }
}

func indexSchema() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"text":     map[string]any{"type": "string", "description": "The text content to index."},
			"docId":    map[string]any{"type": "string", "description": "Unique document identifier."},
			"metadata": map[string]any{"type": "object", "description": "Optional metadata to store with the document."},
		},
		"required": []string{"text", "docId"},
	}
}

func searchSchema() map[string]any {
	return map[string]any{
		"type": "object",
		"properties": map[string]any{
			"query": map[string]any{"type": "string", "description": "The search query."},
		},
		"required": []string{"query"},
	}
}
