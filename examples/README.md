# Conductor Go SDK Examples

This directory contains various examples demonstrating how to use the Conductor Go SDK for workflow orchestration.

## Prerequisites

Before running any example, you need to set up the Conductor server connection:

```bash
export CONDUCTOR_SERVER_URL="http://localhost:8080/api"
export CONDUCTOR_AUTH_KEY="your_auth_key"      # Optional, if authentication is enabled
export CONDUCTOR_AUTH_SECRET="your_auth_secret" # Optional, if authentication is enabled
```

## Examples

### 🤖 **Agent examples** (`agents/`)

Ports of the Python SDK's `examples/agents`, one standalone file each. They
need a Conductor server with LLM support and `CONDUCTOR_SERVER_URL`;
`CONDUCTOR_AGENT_LLM_MODEL` picks the model (default `openai/gpt-4o`). Each
has a matching integration test under `test/integration_tests/ai` that
replays the Python SDK's recording of the same example.

| Example | Shows |
|---|---|
| `01_basic_agent.go` | Define an agent, run it, print the answer. |
| `02a_simple_tools.go` | Two worker tools; the model picks the right one. |
| `02c_tool_retry_config.go` | Per-tool retry policy, count and delay (`tool.WithRetry`). |
| `04_http_and_mcp_tools.go` | Server-side HTTP and MCP tools mixed with a worker tool; needs `mcp-testkit` on port 3001 and two credentials on the server. |
| `05_handoffs.go` | A support agent hands off to billing, technical or sales sub-agents. |
| `06_sequential_pipeline.go` | Researcher, writer and editor run in order, each seeing the previous output. |
| `07_parallel_agents.go` | Three analysts examine the same topic at once. |
| `09_human_in_the_loop.go` | A transfer tool that pauses the run for approval at the terminal, with streamed events. |
| `09c_hitl_streaming.go` | Several tools, one needing approval, with streamed events. |
| `103_plan_and_compile.go` | A planner writes a plan over three tools; the server compiles and runs it. |
| `10_guardrails.go` | A custom output guardrail makes the model redact PII and revise. |
| `13_hierarchical_agents.go` | CEO routes to department leads, who route to specialists. |
| `16e_credentials_http_tool.go` | An HTTP tool whose `Authorization` header names a credential the server resolves; no worker runs. |
| `17_swarm_orchestration.go` | Front-line support transfers the conversation to a specialist and back. |
| `21_regex_guardrails.go` | Server-side regex guardrails block emails and SSNs. |
| `22_llm_guardrails.go` | A second model judges the answer against a policy; retries run out. |
| `33_external_workers.go` | Tools whose workers run in another service (`tool.External`), mixed with a local one. |
| `64_swarm_with_tools.go` | Swarm specialists that each carry their own domain tool. |
| `66_handoff_to_parallel.go` | A coordinator hands off to a single agent or to a parallel group. |

**Run:**
```bash
CONDUCTOR_AGENT_LLM_MODEL=openai/gpt-4o go run agents/01_basic_agent.go
```

---

### 🌟 **Hello World** (`hello_world/`)
**Basic introduction to Conductor workflows**

A simple example demonstrating the fundamental concepts of creating workflows and workers.

**Features:**
- Basic workflow definition
- Simple task worker
- Workflow registration and execution

**Run:**
```bash
cd hello_world
go run main.go
```

---

### **Dynamic Workflow Generation** (`workflow/dynamic_workflows/`)

Demonstrates how to create workflows dynamically at runtime by adding tasks programmatically.

**Features:**
- **Runtime Workflow Creation** - Workflows created programmatically at execution time
- **Task Chaining** - Input/output parameter references between tasks
- **Dynamic Task Definition** - Tasks defined with variable inputs
- **Real-time Monitoring** - `MonitorExecution()` for workflow completion tracking
- **Conditional Logic** - Different workflow structures based on runtime conditions

**Run:**
```bash
cd workflow/dynamic_workflows
go run .
```

---

### **Workflow Lifecycle Management** (`workflow/lifecycle/`)
**Complete workflow lifecycle operations and monitoring**

Comprehensive example showcasing all workflow management operations available in the Conductor Go SDK.

**Features:**
- **Workflow Execution** - Start and monitor workflows
- **Lifecycle Operations** - Pause, resume, retry, restart, rerun
- **Status Monitoring** - Real-time execution tracking with `MonitorExecution()`
- **Task Management** - Manual task updates and completions
- **Search & Query** - Find workflows by correlation ID
- **Error Handling** - Workflow failure scenarios and recovery

**Operations Demonstrated:**
- Start workflow
- Get execution status
- Pause/Resume workflow
- Retry failed workflow
- Restart workflow
- Rerun workflow
- Remove workflow
- Search by correlation ID
- Jump to specific task
- Real-time monitoring

**Run:**
```bash
cd workflow/lifecycle
go run .
```

---

## Common Setup

### 1. Install Dependencies
```bash
go mod download
```

### 2. Set Environment Variables
```bash
export CONDUCTOR_SERVER_URL=http://localhost:8080/api
export CONDUCTOR_AUTH_KEY=your_key        # Orkes Conductor only
export CONDUCTOR_AUTH_SECRET=your_secret  # Orkes Conductor only
```

On Orkes Conductor, the key and secret come from an application access key
(Access Control > Applications in the Orkes UI). On open-source Conductor,
leave both unset: a set pair makes the client call a token endpoint the
server does not have, and every request fails.

### 3. Run Examples
Navigate to any example directory and run:
```bash
go run .
# or
go run main.go
```

## Example Output

When running examples successfully, you'll see structured logs showing:
- Worker registration and startup
- Workflow registration
- Workflow execution with real-time monitoring
- Task completion and results
- Final workflow status

## Troubleshooting

### Common Issues:

1. **"Environment variable CONDUCTOR_SERVER_URL is not set"**
   - Set the required environment variable before running

2. **"unsupported protocol scheme"**
   - Check that CONDUCTOR_SERVER_URL includes the protocol (http:// or https://)

3. **Connection refused**
   - Ensure Conductor server is running on the specified URL

4. **Authentication errors**
   - On Orkes Conductor, check that CONDUCTOR_AUTH_KEY and CONDUCTOR_AUTH_SECRET
     hold the key ID and secret of an application access key
   - On open-source Conductor, make sure both variables are unset; a set pair
     makes the client call a token endpoint the server does not have
