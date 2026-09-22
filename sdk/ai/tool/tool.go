//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

// Package tool builds the tools an agent can call.
//
// A worker tool is an ordinary Go function; reflection over its argument and
// return types gives the input and output schemas. A field is shown to the
// model under its `json` tag, or under its name in snake_case when it has
// none, so AccountID is account_id; it is required unless it is a pointer or
// carries omitempty.
package tool

import (
	"context"
	"reflect"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/internal/schema"
)

// Option configures a tool at construction.
type Option func(*ai.ToolDef)

// Func builds a worker tool from a Go function. The runtime registers fn as a
// Conductor worker under name, so name is both what the model calls and the
// task name workers poll for.
func Func[In, Out any](name, description string,
	fn func(context.Context, In) (Out, error), opts ...Option) ai.ToolDef {

	var in In
	var out Out
	td := ai.ToolDef{
		Name:         name,
		Description:  description,
		InputSchema:  schema.Of(reflect.TypeOf(&in).Elem()),
		OutputSchema: schema.Of(reflect.TypeOf(&out).Elem()),
		ToolType:     ai.ToolTypeWorker,
		Handler:      fn,
	}
	for _, o := range opts {
		o(&td)
	}
	return td
}

// External declares a worker tool whose worker runs in another process, the
// counterpart of the Python SDK's @tool(external=True). In and Out give the
// model the schema a Func handler's types would; no worker is started here, so
// Conductor dispatches each call to whatever is polling for name.
func External[In, Out any](name, description string, opts ...Option) ai.ToolDef {
	var in In
	var out Out
	td := ai.ToolDef{
		Name:         name,
		Description:  description,
		InputSchema:  schema.Of(reflect.TypeOf(&in).Elem()),
		OutputSchema: schema.Of(reflect.TypeOf(&out).Elem()),
		ToolType:     ai.ToolTypeWorker,
	}
	for _, o := range opts {
		o(&td)
	}
	return td
}

// WithCredentials declares the secret names this tool may read; Go cannot see
// which ones a function body touches. They travel in the tool's task
// definition, and the server resolves them at poll time and delivers the values
// with the task, where ai.Secret reads them.
func WithCredentials(names ...string) Option {
	return func(t *ai.ToolDef) { t.Credentials = append(t.Credentials, names...) }
}

// RequiresApproval pauses the run for a human before this tool is dispatched.
func RequiresApproval() Option {
	return func(t *ai.ToolDef) { t.ApprovalRequired = true }
}

// WithTimeout bounds one call to this tool.
func WithTimeout(seconds int) Option {
	return func(t *ai.ToolDef) { t.TimeoutSeconds = &seconds }
}

// WithRetry retries a failed call count times, delaySeconds apart, per policy.
// It configures the task definition the runtime registers, not the agent
// document; the default is Python's 2 retries, 2 seconds apart, linear.
func WithRetry(count, delaySeconds int, policy ai.RetryPolicy) Option {
	return func(t *ai.ToolDef) {
		t.RetryCount = ai.Ptr(count)
		t.RetryDelaySeconds = ai.Ptr(delaySeconds)
		t.RetryPolicy = policy
	}
}

// WithMaxCalls caps how many times the agent may call this tool in a run.
func WithMaxCalls(n int) Option {
	return func(t *ai.ToolDef) { t.MaxCalls = &n }
}

// Stateful routes this tool to a per-execution worker domain, so one run's calls reach one process.
func Stateful() Option {
	return func(t *ai.ToolDef) { t.Stateful = true }
}

// WithConfig sets a type-specific config key. Credentials have their own option.
func WithConfig(key string, value any) Option {
	return func(t *ai.ToolDef) {
		if t.Config == nil {
			t.Config = map[string]any{}
		}
		t.Config[key] = value
	}
}
