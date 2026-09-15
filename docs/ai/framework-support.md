# Framework Support

"Framework" covers two different things:

1. **Model providers** — `Model: "openai/gpt-4o-mini"`, `"anthropic/claude-sonnet-4-6"`. The
   server talks to the provider; every provider the server supports works from Go with no SDK
   code. Complete.
2. **Agent frameworks** — the five python-sdk bridges to: LangChain, LangGraph, the Claude Agent
   SDK, Google ADK, and the OpenAI Agents SDK. A bridge introspects a framework object and pulls
   its tools into a native agent.

Go ships the first and none of the second, by policy: **bridge only to an officially maintained
Go SDK**, since a bridge is per-framework work against a dependency this SDK does not control.

| Framework | Official Go SDK? | Go | Rationale |
|---|---|---|---|
| any provider | — | native, complete | server-side; no SDK work |
| Google ADK | **yes** — `google.golang.org/adk` | **not in this port**; demand-gated follow-on | the only bridge the policy would allow |
| Claude Agent SDK | no (the API SDK's `toolrunner` is a competing agent loop) | not ported | no Go agent SDK |
| OpenAI Agents SDK | no | not ported | no Go version |
| LangChain / LangGraph | community ports only | not ported | not officially maintained |
| MCP servers | n/a — a protocol the server speaks | `tool.MCP`, verified e2e | not a dependency |
