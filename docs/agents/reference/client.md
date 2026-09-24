# AgentClient reference

`client.AgentClient` is the lower-level `/agent` control plane: the raw calls that compile, deploy, start, inspect and steer an agent execution, without the worker hosting that [`Runtime`](./runtime.md) adds on top.

Reach for it when the process does not host Go tool workers — starting an agent that was already deployed and is served elsewhere, inspecting or stopping a run from an operator tool, listing executions. When the process does own the tools, use `Runtime`.

## Obtaining one

```go
// From a Runtime, sharing its APIClient:
agents := rt.AgentClient()

// Standalone:
apiClient := client.NewAPIClientFromEnv()
agents := client.NewAgentClient(apiClient)
```

`func NewAgentClient(apiClient *APIClient) AgentClient` reuses the `APIClient`'s auth, token refresh and retry. Paths below are relative to the `APIClient` base URL, which already ends in `/api`.

## Why the payloads are maps

`agentConfig` is generated from a schema shared by four SDKs, and the server tolerates unknown keys by design. Typing it in Go would mean re-deriving that schema and breaking on every field the schema grows. So every request and response here is `map[string]any`; the typed, field-checked surface is [`ai.Agent`](./agent-definition.md), which serializes to exactly this document.

## Methods

| Signature | Endpoint | Description |
|---|---|---|
| `Compile(ctx context.Context, payload map[string]any) (map[string]any, error)` | `POST /agent/compile` | Turns an agent config into a workflow definition without registering or running it. Returns `workflowDef` and `requiredWorkers`. |
| `Deploy(ctx context.Context, payload map[string]any) (map[string]any, error)` | `POST /agent/deploy` | Compiles and registers the workflow under the agent's name. Idempotent. |
| `Start(ctx context.Context, payload map[string]any) (map[string]any, error)` | `POST /agent/start` | Compiles (or resolves a deployed agent), registers and starts in one call. Returns `executionId`. |
| `Status(ctx context.Context, executionID string) (map[string]any, error)` | `GET /agent/{id}/status` | The cheap poll: state, output, finish reason, `isWaiting`. |
| `Execution(ctx context.Context, executionID string) (map[string]any, error)` | `GET /agent/execution/{id}` | The full execution record, including every task. |
| `Executions(ctx context.Context, params map[string]string) (map[string]any, error)` | `GET /agent/executions` | Lists runs matching the query parameters. |
| `Respond(ctx context.Context, executionID string, body map[string]any) error` | `POST /agent/{id}/respond` | Answers a run waiting on a human. Approval is `{"approved": true}`. |
| `Stop(ctx context.Context, executionID string) error` | `POST /agent/{id}/stop` | Terminates a run. Sends an empty body. |
| `Signal(ctx context.Context, executionID, message string) error` | `POST /agent/{id}/signal` | Sets the persistent signal prepended to the agent's next turn. Sends `{"message": message}`. |
| `SendMessage(ctx context.Context, executionID string, message map[string]any) error` | `POST /workflow/{id}/messages` | Pushes a message into the execution's workflow message queue, for an agent waiting on a `wait_for_message` tool. A workflow endpoint, not an `/agent` one; the server needs `conductor.workflow-message-queue.enabled=true`. |

There is no streaming method here: events come from `APIClient.StreamSSE(ctx, "/agent/stream/"+executionID, "")`, which is what [`ai.AgentHandle.Events`](./runtime.md) wraps.

## Starting a deployed agent by name

The main reason to use this client directly. Deploy once — from a release step, or with `Runtime.Deploy`/`Runtime.Serve` — then start runs by name from anywhere, with no agent definition and no local workers.

```go
started, err := agents.Start(ctx, map[string]any{
    "name":   "support_agent", // the deployed agent; no agentConfig
    "prompt": "Refund order 4471",
})
if err != nil {
    return err
}
executionID := started["executionId"].(string)
```

The three construction forms are mutually exclusive: a deployed `name` (with optional `version`), an inline `agentConfig`, or `framework` plus `rawConfig`/`skillRef`. `Runtime.Run` and `Runtime.Start` always send `agentConfig` (or `framework`/`rawConfig` for a skill), which is why by-name starts belong here.

### Start payload

| Key | Type | Description |
|---|---|---|
| `name` | `string` | A previously deployed agent definition. Also accepted as `agentName`. |
| `version` | `int` | Deployed agent version; latest when omitted. |
| `agentConfig` | `map[string]any` | Inline agent document, as `ai.Agent` serializes it. Use instead of `name`. |
| `framework` | `string` | Foreign-framework identifier for a non-native agent, used with `rawConfig` or `skillRef`. |
| `rawConfig` | `map[string]any` | The framework-specific agent document. |
| `skillRef` | `map[string]any` | Reference to a server-registered skill package, so the server resolves the config itself. |
| `prompt` | `string` | The user prompt for this run. |
| `model` | `string` | Per-call model override for a run started by `name`. The stored definition is never modified; ignored on inline construction, where the model belongs in the config. |
| `sessionId` | `string` | Conversation/session identifier. |
| `media` | `[]any` | Media paths or URLs. The server reads paths itself. |
| `context` | `map[string]any` | Extra context for the run. |
| `idempotencyKey` | `string` | Stable key for retrying one logical execution. |
| `timeoutSeconds` | `int` | Per-call timeout override, applied server-side. |
| `runId` | `string` | Per-execution isolation key for stateful agents; the server maps every worker tool task to this domain. `Runtime` generates and sends one on every run. |
| `static_plan` | `map[string]any` | Deterministic plan for a `plan_execute` agent, read ahead of the planner's output. Sent by `ai.WithPlan`. |

`Compile` and `Deploy` take the same construction keys and ignore the run-specific ones.

### Responses

| Call | Keys |
|---|---|
| `Compile` | `workflowDef`, `requiredWorkers` |
| `Deploy`, `Start` | `executionId` (start only), `agentName`, `requiredWorkers` |
| `Status` | `executionId`, `status`, `isComplete`, `isRunning`, `isWaiting`, `output`, `reasonForIncompletion`, `pendingTool`, `startTime`, `endTime` |

`requiredWorkers` names the SIMPLE tasks that need a worker polling for them. A by-name start leaves those to whatever process is serving the agent; without one, the run stalls at its first tool call.

### Execution query parameters

| Parameter | Default | Description |
|---|---|---|
| `start` | `0` | Offset. |
| `size` | `20` | Page size. |
| `sort` | `startTime:DESC` | Sort field and direction. |
| `freeText` | — | Free-text search. |
| `status` | — | Filter by execution status. |
| `agentName` | — | Filter by agent. |
| `sessionId` | — | Filter by session. |
| `classifier` | agent executions | Workflow classifier; defaults to agent runs. |

## Human input without a Runtime

`Respond`, `Stop` and `Signal` change a durable execution, so authorize the caller and make externally triggered calls idempotent. Unlike [`AgentHandle.Respond`](./runtime.md), `AgentClient.Respond` does not first confirm that the server reports the run waiting — poll `Status` for `isWaiting` before responding, or the response is dropped and the run hangs. See [streaming and human input](../concepts/streaming-hitl.md).
