# Conductor OSS Go SDK

[![Build Status](https://github.com/conductor-oss/go-sdk/actions/workflows/build.yml/badge.svg)](https://github.com/conductor-oss/go-sdk/actions/workflows/build.yml)

SDK for developing Go applications that create, manage and execute workflows, and run workers.

[Conductor](https://www.conductor-oss.org/) is the leading open-source orchestration platform allowing developers to build highly scalable distributed applications.

To learn more about Conductor checkout our [developer's guide](https://docs.conductor-oss.org/devguide/concepts/index.html) and give it a ⭐ to make it famous!

[![GitHub stars](https://img.shields.io/github/stars/conductor-oss/conductor.svg?style=social&label=Star&maxAge=)](https://GitHub.com/conductor-oss/conductor/)

# Content
<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Conductor OSS Go SDK](#conductor-oss-go-sdk)
- [Content](#content)
	- [Installation](#installation)
	- [Hello World](#hello-world)
		- [Step 1: Creating the workflow by code](#step-1-creating-the-workflow-by-code)
		- [Step 2: Creating the worker](#step-2-creating-the-worker)
		- [Step 3: Running the application](#step-3-running-the-application)
			- [Running the example with a local Conductor OSS server:](#running-the-example-with-a-local-conductor-oss-server)
			- [Running the example with an Orkes developer account.](#running-the-example-with-an-orkes-developer-account)
	- [AI Agents](#ai-agents)
		- [Step 1: Define and run an agent](#step-1-define-and-run-an-agent)
		- [Step 2: Add tools](#step-2-add-tools)
		- [Step 3: Run the examples](#step-3-run-the-examples)
		- [Deploy, Serve, and Run](#deploy-serve-and-run)
		- [Model providers and agent frameworks](#model-providers-and-agent-frameworks)
	- [Deprecated Methods](#deprecated-methods)
- [Further Reading](#further-reading)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Installation

1. Initialize your module. e.g.:

```shell
mkdir hello_world
cd hello_world
go mod init hello_world
```

2. Get the SDK:

```shell
go get github.com/conductor-sdk/conductor-go
```

> **Note:** The Go module path is `github.com/conductor-sdk/conductor-go` (historical). The source repository is at [conductor-oss/go-sdk](https://github.com/conductor-oss/go-sdk).

## Hello World

In this repo you will find a basic "Hello World" under [examples/hello_world](examples/hello_world/). 

Let's analyze the app in 3 steps.


> [!note]
> You will need an up & running Conductor Server. 
>
> For details on how to run Conductor take a look at [our documentation](https://docs.conductor-oss.org).
>
> The examples expect the server to be listening on http://localhost:8080.


### Step 1: Creating the workflow by code

The "greetings" workflow is going to be created by code and registered in Conductor. 

Check the `CreateWorkflow` function in [examples/hello_world/src/workflow.go](examples/hello_world/src/workflow.go).

```go
func CreateWorkflow(executor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {
	wf := workflow.NewConductorWorkflow(executor).
		Name("greetings").
		Version(1).
		Description("Greetings workflow - Greets a user by their name").
		TimeoutPolicy(workflow.TimeOutWorkflow, 600)

	greet := workflow.NewSimpleTask("greet", "greet_ref").
		Input("person_to_be_greeted", "${workflow.input.name}")

	wf.Add(greet)

	wf.OutputParameters(map[string]interface{}{
		"greetings": greet.OutputRef("hello"),
	})

	return wf
}
```

In the above code first we create a workflow by calling `workflow.NewConductorWorkflow(..)` and set its properties `Name`, `Version`, `Description` and `TimeoutPolicy`. 

Then we create a [Simple Task](https://orkes.io/content/reference-docs/worker-task) of type `"greet"` with reference name `"greet_ref"` and add it to the workflow. That task gets the workflow input `"name"` as an input with key `"person_to_be_greeted"`.

> [!note]
>`"person_to_be_greeted"` is too verbose! Why would you name it like that?
>
> It's just to make it clear that the workflow input is not passed automatically. 
>
> The worker will get the actual value of the workflow input because of this mapping  `Input("person_to_be_greeted", "${workflow.input.name}")` in the workflow definition. 
>
>Expressions like `"${workflow.input.name}"` will be replaced by their value during execution.

Last but not least, the output of the workflow is set by calling `wf.OutputParameters(..)`. 

The value of `"greetings"` is going to be whatever `"hello"` is in the output of the executed `"greet"` task, e.g.: if the task output is:
```
{
	"hello" : "Hello, John"
}
```

The expected workflow output will be:
```
{
	"greetings": "Hello, John"
}
```

The Go code translates to this JSON defininition. You can view this in your Conductor server after registering the workflow.

```json
{
  "schemaVersion": 2,
  "name": "greetings",
  "description": "Greetings workflow - Greets a user by their name",
  "version": 1,
  "tasks": [
    {
      "name": "greet",
      "taskReferenceName": "greet_ref",
      "type": "SIMPLE",
      "inputParameters": {
        "name": "${workflow.input.name}"
      }
    }
  ],
  "outputParameters": {
    "Greetings": "${greet_ref.output.greetings}"
  },
  "timeoutPolicy": "TIME_OUT_WF",
  "timeoutSeconds": 600
}
```

> [!note]
> Workflows can also be registered using the API. Using the JSON you can make the following request:
> ```shell
> curl -X POST -H "Content-Type:application/json" \
> http://localhost:8080/api/metadata/workflow -d @greetings_workflow.json
> ```

In [Step 3](#step-3-running-the-application) you will see how to create an instance of `executor.WorkflowExecutor`.


### Step 2: Creating the worker

A worker is a function with a specific task to perform.

In this example the worker just uses the input `person_to_be_greeted` to say hello, as you can see in [examples/hello_world/src/worker.go](examples/hello_world/src/worker.go).

```go
func Greet(task *model.Task) (interface{}, error) {
	return map[string]interface{}{
		"hello": "Hello, " + fmt.Sprintf("%v", task.InputData["person_to_be_greeted"]),
	}, nil
}
```

To learn more about workers take a look at [Writing Workers with the Go SDK](docs/workers_sdk.md).

> [!note]
> A single workflow can have task workers written in different languages and deployed anywhere, making your workflow polyglot and distributed!

### Step 3: Running the application

The application is going to start the Greet worker (to execute tasks of type "greet") and it will register the workflow created in [step 1](#step-1-creating-the-workflow-by-code).

To begin with, let's take a look at the variable declaration in [examples/hello_world/main.go](examples/hello_world/main.go).

```go

var (
	apiClient        = client.NewAPIClientFromEnv()
	taskRunner       = worker.NewTaskRunnerWithApiClient(apiClient)
	workflowExecutor = executor.NewWorkflowExecutor(apiClient)
)

```

First we create an `APIClient` instance. This is a REST client. 

We need to provide the correct settings to our client. In this example, `client.NewAPIClientFromEnv()` is used, which initializes a new client by reading the settings from the following environment variables: `CONDUCTOR_SERVER_URL`, `CONDUCTOR_AUTH_KEY`, and `CONDUCTOR_AUTH_SECRET`.
`CONDUCTOR_CLIENT_HTTP_TIMEOUT` lets you configure the HTTP timeout for our client, in seconds. If not set, defaults to 30 seconds.

> [!tip]
> For advanced configuration options and detailed examples see the [API Client Configuration Guide](docs/api_client/README.md).

Now let's take a look at the `main` function:

```go
func main() {
	// Start the Greet Worker. This worker will process "greet" tasks.
	taskRunner.StartWorker("greet", hello_world.Greet, 1, time.Millisecond*100)

	// This is used to register the Workflow, it's a one-time process. You can comment from here
	wf := hello_world.CreateWorkflow(workflowExecutor)
	err := wf.Register(true)
	if err != nil {
		log.Error(err.Error())
		return
	}
	// Till Here after registering the workflow

	// Start the greetings workflow 
	id, err := workflowExecutor.StartWorkflow(
		&model.StartWorkflowRequest{
			Name:    "greetings",
			Version: 1,
			Input: map[string]string{
				"name": "Gopher",
			},
		},
	)

	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Info("Started workflow with Id: ", id)

	// Get a channel to monitor the workflow execution -
	// Note: This is useful in case of short duration workflows that completes in few seconds.
	channel, _ := workflowExecutor.MonitorExecution(id)
	run := <-channel
	log.Info("Output of the workflow: ", run.Output)
}
```

The `taskRunner` uses the `apiClient` to poll for work and complete tasks. It also starts the worker and handles concurrency and polling intervals for us based on the configuration provided.

That simple line `taskRunner.StartWorker("greet", hello_world.Greet, 1, time.Millisecond*100)` is all that's needed to get our Greet worker up & running and processing tasks of type `"greet"`.

The `workflowExecutor` gives us an abstraction on top of the `apiClient` to manage workflows. It is used under the hood by `ConductorWorkflow` to register the workflow and it's also used to start and monitor the execution.

#### Running the example with a local Conductor OSS server:
```shell
export CONDUCTOR_SERVER_URL="http://localhost:8080/api"
cd examples
go run hello_world/main.go
```

#### Running the example with an [Orkes developer account](https://developer.orkescloud.com).
```shell
export CONDUCTOR_SERVER_URL="https://developer.orkescloud.com/api"
export CONDUCTOR_AUTH_KEY="..."
export CONDUCTOR_AUTH_SECRET="..."
cd examples
go run hello_world/main.go
```

> [!note]
> Orkes Conductor requires authentication. [Get a key and secret from the server](https://orkes.io/content/how-to-videos/access-key-and-secret) to set those variables.

The above commands should give an output similar to
```shell
INFO[0000] Updated poll interval for task: greet, to: 100ms 
INFO[0000] Started 1 worker(s) for taskName greet, polling in interval of 100 ms 
INFO[0000] Started workflow with Id:14a9fcc5-3d74-11ef-83dc-acde48001122 
INFO[0000] Output of the workflow:map[Greetings:Hello, Gopher] 
```

## AI Agents

An agent is defined in Go, compiled by the Conductor server into a workflow, and executed like any
other workflow. Each LLM turn, tool call and handoff becomes a task, so a run is durable, survives a
process restart, and is visible in the Conductor UI. Tools are ordinary Go functions that the SDK
registers as workers.

> [!note]
> Agents need a Conductor server with an LLM provider configured **on the server**. Your Go process
> never calls the model provider directly, so no provider API key is needed locally.

### Step 1: Define and run an agent

The smallest agent is a struct and one call. From [examples/agents/01_basic_agent.go](examples/agents/01_basic_agent.go):

```go
import (
	"context"
	"fmt"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

agent := &ai.Agent{
	Name:         "greeter",
	Model:        "openai/gpt-4o",
	Instructions: "You are a friendly assistant. Keep responses brief.",
}

// NewRuntime connects to the server named by CONDUCTOR_SERVER_URL.
//
// Open-source Conductor has no authentication, so that one variable is
// all it needs.
//
// Orkes Conductor requires authentication. In the Orkes UI, create an
// application and generate an access key for it; the key ID goes in
// CONDUCTOR_AUTH_KEY and the key secret in CONDUCTOR_AUTH_SECRET. The
// client then fetches a token and sends it with every request.
runtime := ai.NewRuntime(ai.Config{})
defer runtime.Shutdown()

result, err := runtime.Run(context.Background(), agent, "Say hello and tell me a fun fact about Go.")
if err != nil {
	return err
}
fmt.Println(result.Output)
```

`Name` becomes the workflow name the agent is registered under. `Model` is a `provider/model`
identifier. `Instructions` is the system prompt.

`Run` starts the execution and blocks until the server reports a terminal state, returning an
`ai.AgentResult`:

| Field | Meaning |
|---|---|
| `Output` | the agent's final answer |
| `Status` | how the run ended: `COMPLETED`, `FAILED`, `TERMINATED`, `TIMED_OUT`, or `WAITING` when it is blocked on a human |
| `ExecutionID` | the id to open in the Conductor UI |
| `FinishReason` | why the loop stopped, such as `STOP` or `MAX_TURNS` |
| `TokenUsage` | token counts, zero when the server does not report them |
| `Error` | the failure reason when `Status` is `FAILED` |

### Step 2: Add tools

A tool is a Go function. `tool.Func` derives the schema the model sees from the argument type: a
field is named by its json tag, or by its name in snake_case when it has none, so `AccountID` is
`account_id`. The runtime registers the function as a Conductor worker named after the tool, so
every call the model makes shows up as its own task with its inputs and outputs.

From [examples/agents/02a_simple_tools.go](examples/agents/02a_simple_tools.go):

```go
import "github.com/conductor-sdk/conductor-go/sdk/ai/tool"

type weatherIn struct {
	City string
}

func getWeather(ctx context.Context, in weatherIn) (map[string]any, error) {
	return map[string]any{"city": in.City, "temp_f": 72, "condition": "Sunny"}, nil
}

agent := &ai.Agent{
	Name:  "weather_stock_agent",
	Model: "openai/gpt-4o",
	Tools: ai.Tools(
		tool.Func("get_weather", getWeather, "Get the current weather for a city."),
	),
	Instructions: "You are a helpful assistant. Use tools to answer questions.",
}
```

The model picks the tool. The example registers a second tool for stock prices alongside this one,
and asking it about the weather in San Francisco calls `get_weather`, not `get_stock_price`.

### Step 3: Run the examples

[examples/agents](examples/agents/) is a package with one file per example, ports of the Python
SDK's `examples/agents`, and a command that runs one by name.

With a local Conductor OSS server:

```shell
export CONDUCTOR_SERVER_URL="http://localhost:8080/api"
export CONDUCTOR_AGENT_LLM_MODEL="openai/gpt-4o"
cd examples
go run ./agents/cmd 01_basic_agent
```

With an [Orkes developer account](https://developer.orkescloud.com):

```shell
export CONDUCTOR_SERVER_URL="https://developer.orkescloud.com/api"
export CONDUCTOR_AUTH_KEY="..."
export CONDUCTOR_AUTH_SECRET="..."
export CONDUCTOR_AGENT_LLM_MODEL="openai/gpt-4o"
cd examples
go run ./agents/cmd 01_basic_agent
```

`CONDUCTOR_AGENT_LLM_MODEL` is optional and defaults to `openai/gpt-4o-mini`. The example prints the
agent's answer, then its status and execution id. Open that execution id in the Conductor UI to see
each LLM turn and tool call as a task.

Beyond the basics, the examples cover tools, HTTP and MCP tools, handoffs, sequential and parallel
sub-agents, guardrails, human-in-the-loop approval, swarms, and plan-and-execute. See the table in
[examples/README.md](examples/README.md) for what each file demonstrates.

### Deploy, Serve, and Run

The runtime separates three concerns that Conductor has always kept apart: registering a definition
on the server, hosting the workers that answer its tasks, and starting an execution. Each call
covers a different subset.

| Call | Registers on server | Starts a run | Hosts tool workers | Blocks | Use when |
|---|---|---|---|---|---|
| `Plan(ctx, agent)` | no | no | no | no | You want to inspect or validate the compiled workflow, for example in CI. |
| `Deploy(ctx, agent)` | yes | no | no | no | A release step publishes the agent so anything can start it by name later. |
| `Serve(ctx, agents...)` | yes | no | yes, until `ctx` is cancelled | yes | A long-lived worker process answers tool calls for runs started elsewhere. |
| `Run(ctx, agent, prompt)` | yes, as a side effect | yes | yes | yes, until the run ends | One process defines the agent, runs it and needs the answer: scripts, CLIs, request handlers. |
| `Start(ctx, agent, prompt)` | yes, as a side effect | yes | yes | no, returns a handle | Same process, but you want to stream events, approve tool calls, stop early, or fan out several runs. |

`Run` and `Start` are the development path: one binary does everything, and changing the agent
means redeploying that binary. Both send the full definition with every start, and the server
registers it under the agent's name, so they also overwrite whatever an earlier `Deploy` stored
under that name.

The production path splits the work across three roles, each of which can be owned, deployed and
scaled on its own:

| Role | Call | Runs where | Notes |
|---|---|---|---|
| **Publish** the agent | `Deploy(ctx, agent)` | A release step, once per version | Compiles and stores the definition under the agent's name. Nothing runs yet. |
| **Serve** its tools | `Serve(ctx, agents...)` | A long-lived worker process | Hosts the tool workers and blocks. Add replicas to scale; Conductor spreads tasks across whoever is polling. Restarting one loses nothing: an in-flight run waits on its tool task until a worker polls again. |
| **Trigger** a run | Start by name | Wherever the trigger lives | A schedule, a webhook, a service in another language, or a plain `POST /agent/start` with the agent's `name`. The caller needs no knowledge of the tools. |

A run that is already in flight is driven from the runtime: the `AgentHandle` returned by `Start`
streams events and answers approvals, and `Signal`, `SendMessage`, `Pause` and `Resume` act on any
execution by ID.

None of this is specific to agents. The same lifecycle applies to every Conductor workflow, and the
agent runtime only bundles the calls that the workflow API leaves separate:

| Agent runtime | Plain workflow | Note |
|---|---|---|
| `Plan` | `ConductorWorkflow.ToWorkflowDef()` | The server compiles agents, so `Plan` makes a network call; a workflow is built locally. |
| `Deploy` | `ConductorWorkflow.Register(overwrite)` | Task definitions for a plain workflow are registered separately through the metadata client. |
| `Serve` | `TaskRunner.StartWorker(...)` per task, then `WaitWorkers()` | The classic worker process, as in [examples/hello_world](examples/hello_world/). |
| `Run` | `ExecuteWorkflow(request, waitUntilTask)` | The workflow form blocks only up to the server's synchronous timeout; the agent form polls to a terminal state. |
| `Start` | `StartWorkflow(request)`, then `MonitorExecution(id)` | Monitoring is a channel rather than a handle. |

Two differences matter when moving between the two. A plain workflow must be registered before it
is started, whereas an agent `Run` registers as a side effect. And a plain workflow does not start
workers for you: the runtime can, because it derives each worker's task name from the agent
definition, one per tool with a Go handler plus its guardrails, handoff predicates, gates and
callbacks.

### Model providers and agent frameworks

Any model provider the server supports works from Go with no SDK changes, since the server is what
talks to the provider. Set `Model` to its `provider/model` identifier.

The Go SDK deliberately ships no bridges to Python agent frameworks such as LangChain, LangGraph or
the OpenAI Agents SDK, because there are no officially maintained Go SDKs for them. MCP servers are
supported directly through `tool.MCP`, since MCP is a protocol the Conductor server speaks rather
than a dependency. See [framework support](docs/agents/framework-support.md) for the full policy.

## Deprecated Methods
Some methods in the SDK client interfaces are now deprecated. They’ve been replaced with newer methods that follow more consistent naming. Please refer to our [Migration Guide](docs/migration_guide.md) for detailed information on how to update your code.
# Further Reading

- [Writing Workers with the Go SDK](docs/workers_sdk.md)
- [Authoring Workflows with the Go SDK](docs/workflow_sdk.md)
- [AI Agents: design and parity docs](docs/agents/README.md) - Go agent types, worked examples, and how the port compares with the Python SDK
- [Agent framework support](docs/agents/framework-support.md) - which model providers and agent frameworks the Go SDK supports
- [Agent secrets and credentials](docs/agents/secrets-and-credentials.md) - delivering secrets to tools
- [Logging Configuration](docs/logger_sdk.md)
- [Migration Guide: Deprecated Methods](docs/migration_guide.md)
- [API Client Configuration](docs/api_client/README.md) - Complete guide to API client setup, authentication, and proxy configuration
- [TLS Configuration Guide](docs/api_client/tls_configuration.md) - TLS/SSL configuration for self-signed certificates and mTLS
