# Go Agents Parity Plan

One page, in the cross-SDK comparison format. Status: **built**; every golden fixture implemented; the e2e suite green against a live
server.

## Classes

### Definition (serialized to `agentConfig`)

| Python | Go | Notes |
|---|---|---|
| `Agent` | `ai.Agent` struct + `Validate()` | defaults substituted at serialization |
| `@tool` | `tool.Func(name, desc, fn)` | input **and output** schema by reflection |
| HTTP / human / agent / MCP tools | `tool.HTTP`, `tool.Human`, `tool.Agent`, `tool.MCP` | settings in `ToolDef.Config` |
| `Strategy` | `ai.Strategy` (9 constants) | emitted only with sub-agents |
| `OnToolResult`, `OnTextMention`, `OnCondition` | same names; `OnCondition.Condition` is `HandoffFunc(ctx, HandoffState)` | one `OnCondition` per target |
| regex / llm / custom guardrail | `RegexGuardrail`, `LLMGuardrail`, `CustomGuardrail` + `NewCustomGuardrail` | `GuardrailFunc(ctx, GuardrailInput)` |
| termination conditions, `and`/`or` | same names; `AndTermination`, `OrTermination` | sealed interface |
| `stop_when`, router fn | `StopWhenFunc`, `RouterFunc` | `{agent}_stop_when`, `{agent}_router_fn` |
| `output_type` | `OutputType any` (struct → JSON Schema) | |
| `code_execution`, `cli` | `CodeExecutionConfig`, `CLIConfig` | derived execution tools |
| `Plan` (static plans) | `ai.Plan`, `Step`, `Op`, `Ref`, `Generate`; `ai.WithPlan` | |
| `skill()`, `load_skills()` | `ai.LoadSkill`, `ai.LoadSkills` + `SkillOption`s | raw document normalized server-side; scripts and `read_skill_file` served as workers |

### Runtime + transport

| Python | Go |
|---|---|
| `AgentRuntime.run / start` | `Runtime.Run / Start` |
| `AgentRuntime.plan` | `Runtime.Plan` |
| `AgentHandle` | `AgentHandle`: `Events`, `Status`, `Waiting`, `Respond`, `Approve`, `Reject`, `Stop`, `Result` |
| `AgentClient` | `client.AgentClient` + `APIClient.StreamSSE` |
| forked worker processes | goroutines on one `worker.TaskRunner` |
| `EventType` (10) | `EventType` (same 10) |

## Examples

1. **Tools + guardrail** — `TestToolCall`, `TestCustomGuardrail*`
2. **Swarm handoff + credentialed tool** — `TestOnConditionHandoff`, `TestTeamWithSecret`
3. **Streaming + approval** — `TestStreamingWithApproval`

## Secrets

Declare (`tool.WithCredentials`) → server resolves → delivered on `Task.RuntimeMetadata` →
`ai.Secret(ctx, name)`, fail-closed (`ErrCredentialNotFound`). Subprocess: `ai.SecretsEnv`.
No env mutation. See `secrets-and-credentials.md`.

## Frameworks

Providers: native. Adapters (LangChain, LangGraph, Claude Agent SDK, ADK): **not ported**;
compositions covered by `tool.Agent`, `tool.MCP`, `tool.HTTP`, `tool.Human`, code execution, CLI.
See `framework-support.md`.

## Divergences from Python, on purpose

| | Python | Go | why |
|---|---|---|---|
| callback signatures | scalars, no ctx | `(ctx, State)` structs | cancellation; server sends more than one field |
| tool output schema | `{}` for any return type | reflected from the return struct | chained plan steps need field names |
| handoff worker | `{agent}_handoff_check` (retired) | `{agent}_handoff_{target}` | matches servers since 2026-07-20 |
| validation | little | `Validate` rejects unknown handoff targets, duplicate conditions, bad enums | fail before the server does |
| workers | forked processes | in-process goroutines | SDK convention; testable |
