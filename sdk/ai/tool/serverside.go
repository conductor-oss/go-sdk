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

// More tools the server runs itself, so none registers a Go function: API
// specs expanded into tools, the two RAG tools, and the workflow message
// queue. Defaults and config keys match the Python SDK's api_tool, index_tool,
// search_tool and wait_for_message_tool.

const (
	defaultAPIName          = "api_tools"
	defaultAPIMaxTools      = 64
	defaultRAGNamespace     = "default_ns"
	defaultSearchMaxResults = 5
	defaultWaitBatchSize    = 1
)

// API exposes the operations of an OpenAPI 3.x spec, a Swagger 2.0 spec or a
// Postman collection at url as tools. Like MCP it is entirely server-side:
// the server fetches the spec when the agent is compiled and expands this one
// definition into a tool per operation, filtering with an LLM when there are
// more than WithMaxTools allows.
//
// An empty name or description takes Python's default: "api_tools" and
// "API tools from <url>". Headers may reference a credential as ${NAME};
// declare the same names with WithCredentials or Validate rejects the tool.
//
//	tool.API("stripe", "Stripe API", "https://api.stripe.com/openapi.json",
//	    tool.WithHeaders(map[string]string{"Authorization": "Bearer ${STRIPE_KEY}"}),
//	    tool.WithCredentials("STRIPE_KEY"), tool.WithMaxTools(20))
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

// Index adds text to a vector index, Conductor's LLM_INDEX_TEXT task. The
// model supplies the text and a document id; the vector database, index and
// embedding model are fixed here. Namespace defaults to "default_ns"; set it
// with WithNamespace, chunking with WithChunking, and the embedding size with
// WithDimensions.
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
// up to WithMaxResults matches (default 5). The other settings are as for
// Index.
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

// WaitForMessage lets the agent wait for messages sent into its execution
// with AgentClient.Signal or the workflow message queue. By default it blocks
// until one message arrives; WithBatchSize takes more per call and NonBlocking
// returns at once with whatever is queued. The server needs
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

// WithChunking splits indexed text into chunks of size characters that overlap
// by overlap. For Index.
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

// NonBlocking makes WaitForMessage return at once, with an empty result when
// nothing is queued, instead of waiting for a message.
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
