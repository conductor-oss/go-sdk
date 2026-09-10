//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package client

import (
	"context"
	"fmt"
	"net/url"
)

// AgentClient drives the server's /agent control plane.
//
// The SDK does not run the agent loop: it sends an agentConfig document and the
// server compiles it into a Conductor workflow and executes the turns. These
// calls start, inspect and steer that execution; the tool calls the server
// makes come back as ordinary Conductor tasks, served by workers.
//
// Payloads are maps rather than typed structs because agentConfig is generated
// from a schema shared by four SDKs, and the server tolerates unknown keys by
// design. Typing it here would mean re-deriving that schema in Go and breaking
// every time it grows.
type AgentClient interface {
	// Compile turns an agentConfig into a workflow definition without running it.
	Compile(ctx context.Context, payload map[string]any) (map[string]any, error)
	// Deploy registers the compiled workflow under the agent's name.
	Deploy(ctx context.Context, payload map[string]any) (map[string]any, error)
	// Start compiles, registers and starts in one call, returning executionId
	// and the worker task names the run needs.
	Start(ctx context.Context, payload map[string]any) (map[string]any, error)

	// Status is the cheap poll: state, output and finish reason.
	Status(ctx context.Context, executionID string) (map[string]any, error)
	// Execution is the full record, including every task.
	Execution(ctx context.Context, executionID string) (map[string]any, error)
	// Executions lists runs matching params.
	Executions(ctx context.Context, params map[string]string) (map[string]any, error)

	// Respond answers a run waiting on a human.
	Respond(ctx context.Context, executionID string, body map[string]any) error
	// Stop terminates a run.
	Stop(ctx context.Context, executionID string) error
	// Signal sends a message into a running execution.
	Signal(ctx context.Context, executionID, message string) error
}

type agentClient struct {
	apiClient *APIClient
}

// NewAgentClient creates an AgentClient over the given APIClient, reusing its
// auth, token refresh and retry.
func NewAgentClient(apiClient *APIClient) AgentClient {
	return &agentClient{apiClient: apiClient}
}

func (c *agentClient) post(ctx context.Context, path string, body, out any) error {
	_, err := c.apiClient.Post(ctx, path, body, out)
	return err
}

func (c *agentClient) Compile(ctx context.Context, payload map[string]any) (map[string]any, error) {
	var out map[string]any
	return out, c.post(ctx, "/agent/compile", payload, &out)
}

func (c *agentClient) Deploy(ctx context.Context, payload map[string]any) (map[string]any, error) {
	var out map[string]any
	return out, c.post(ctx, "/agent/deploy", payload, &out)
}

func (c *agentClient) Start(ctx context.Context, payload map[string]any) (map[string]any, error) {
	var out map[string]any
	return out, c.post(ctx, "/agent/start", payload, &out)
}

func (c *agentClient) Status(ctx context.Context, executionID string) (map[string]any, error) {
	var out map[string]any
	_, err := c.apiClient.Get(ctx, fmt.Sprintf("/agent/%s/status", executionID), nil, &out)
	return out, err
}

func (c *agentClient) Execution(ctx context.Context, executionID string) (map[string]any, error) {
	var out map[string]any
	_, err := c.apiClient.Get(ctx, fmt.Sprintf("/agent/execution/%s", executionID), nil, &out)
	return out, err
}

func (c *agentClient) Executions(ctx context.Context, params map[string]string) (map[string]any, error) {
	q := url.Values{}
	for k, v := range params {
		q.Set(k, v)
	}
	var out map[string]any
	_, err := c.apiClient.Get(ctx, "/agent/executions", q, &out)
	return out, err
}

func (c *agentClient) Respond(ctx context.Context, executionID string, body map[string]any) error {
	return c.post(ctx, fmt.Sprintf("/agent/%s/respond", executionID), body, nil)
}

func (c *agentClient) Stop(ctx context.Context, executionID string) error {
	return c.post(ctx, fmt.Sprintf("/agent/%s/stop", executionID), map[string]any{}, nil)
}

func (c *agentClient) Signal(ctx context.Context, executionID, message string) error {
	return c.post(ctx, fmt.Sprintf("/agent/%s/signal", executionID),
		map[string]any{"message": message}, nil)
}
