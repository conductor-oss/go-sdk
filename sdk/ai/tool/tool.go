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
// A worker tool is an ordinary Go function. Its input and output schemas are
// derived by reflection over the argument and return types, so the model sees
// the same contract the function actually accepts:
//
//	type WeatherIn struct {
//	    City string `json:"city"`
//	    Days int    `json:"days,omitempty"`
//	}
//
//	func getWeather(ctx context.Context, in WeatherIn) (map[string]any, error) {
//	    return map[string]any{"city": in.City, "tempF": 72}, nil
//	}
//
//	tool.Func("get_weather", "Get the current weather for a city", getWeather)
//
// Fields carry `json` tags because those names are what the model is shown and
// what it sends back. A field is required unless it is a pointer or carries
// omitempty.
package tool

import (
	"context"
	"reflect"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/internal/schema"
)

// Option configures a tool at construction.
type Option func(*ai.ToolDef)

// Func builds a worker tool from a Go function.
//
// The runtime registers fn as a Conductor worker under name, so name is both
// what the model calls and the task name workers poll for.
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

// WithCredentials declares the secret names this tool may read.
//
// The names reach the server in the tool's task definition; the server resolves
// them at poll time and delivers the values with the task, where ai.Secret
// reads them. Go cannot see which credentials a function body touches, so this
// declaration is what tells the server what to resolve.
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

// WithMaxCalls caps how many times the agent may call this tool in a run.
func WithMaxCalls(n int) Option {
	return func(t *ai.ToolDef) { t.MaxCalls = &n }
}

// Stateful routes this tool to a per-execution worker domain, so calls within
// one run reach the same process.
func Stateful() Option {
	return func(t *ai.ToolDef) { t.Stateful = true }
}

// WithConfig sets a type-specific config key. Credentials have their own
// option; this is for the settings that vary by tool type.
func WithConfig(key string, value any) Option {
	return func(t *ai.ToolDef) {
		if t.Config == nil {
			t.Config = map[string]any{}
		}
		t.Config[key] = value
	}
}
