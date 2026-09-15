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

### 🤖 **Basic Agent** (`agents/01_basic_agent.go`)
**The smallest possible AI agent**

Define an agent, run it, print the answer.

**Features:**
- Agent definition (`ai.Agent`) and `Runtime.Run`
- Reads `CONDUCTOR_AGENT_LLM_MODEL` from the environment (defaults to `openai/gpt-4o`)
- Comments show the production split: `Deploy` once, `Serve` in a worker process

**Requires** a Conductor server with LLM support.

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
# Required
export CONDUCTOR_SERVER_URL="http://localhost:8080/api"

# Optional (for authenticated environments)
export CONDUCTOR_AUTH_KEY="your_auth_key"
export CONDUCTOR_AUTH_SECRET="your_auth_secret"
```

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
   - Verify CONDUCTOR_AUTH_KEY and CONDUCTOR_AUTH_SECRET if using authenticated setup
