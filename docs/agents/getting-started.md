# Getting started

From an empty directory to a running agent, with and without tools.

## Prerequisites

- Go 1.23 or later.
- A Conductor server you can reach, with an LLM provider configured **on the server**. The Go
  process never calls the model provider, so no provider API key is needed locally.

## Install

```shell
mkdir my-agent
cd my-agent
go mod init my-agent
go get github.com/conductor-sdk/conductor-go
```

The module path is `github.com/conductor-sdk/conductor-go`; the repository is
[conductor-oss/go-sdk](https://github.com/conductor-oss/go-sdk).

## Configure the environment

```shell
export CONDUCTOR_SERVER_URL="http://localhost:8080/api"
export CONDUCTOR_AGENT_LLM_MODEL="openai/gpt-4o"
```

`CONDUCTOR_SERVER_URL` is the server's API base URL and is all an open-source Conductor needs.

For Orkes Conductor, create an application in the UI, generate an access key, and set the key ID and
secret:

```shell
export CONDUCTOR_SERVER_URL="https://developer.orkescloud.com/api"
export CONDUCTOR_AUTH_KEY="..."
export CONDUCTOR_AUTH_SECRET="..."
```

Leave `CONDUCTOR_AUTH_KEY` and `CONDUCTOR_AUTH_SECRET` unset for open-source Conductor. Set, the
client exchanges them for a token at an endpoint the open-source server does not have, and the
exchange and every request after it fails.

`CONDUCTOR_AGENT_LLM_MODEL` is read by the examples in this repository, not by `ai.NewRuntime`. In
your own code, `Model` is a plain field.

## The smallest agent

`main.go`, after [examples/agents/01_basic_agent.go](../../examples/agents/01_basic_agent.go):

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
)

func main() {
	agent := &ai.Agent{
		Name:         "greeter",
		Model:        "openai/gpt-4o",
		Instructions: "You are a friendly assistant. Keep responses brief.",
	}

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()

	result, err := runtime.Run(context.Background(), agent, "Say hello and tell me a fun fact about Go.")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	result.PrintResult()
}
```

`Name` becomes the workflow name the agent is registered under. `Model` is a `provider/model`
identifier the server supports. `Run` blocks until the run reaches a terminal state and returns an
`*ai.AgentResult`; `PrintResult` writes its output, status and execution ID to stdout. Open that
execution ID in the Conductor UI to see each LLM turn as a task.

```shell
go run .
```

## Add a tool

A tool is an ordinary Go function taking a `context.Context` and a struct. `tool.Func` derives the
schema the model sees from that struct: an untagged field is advertised in snake_case, so `City` is
`city` and `AccountID` is `account_id`. No `json` tag is needed.

After [examples/agents/02a_simple_tools.go](../../examples/agents/02a_simple_tools.go):

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

type weatherIn struct {
	City string
}

type stockIn struct {
	Symbol string
}

func getWeather(ctx context.Context, in weatherIn) (map[string]any, error) {
	return map[string]any{"city": in.City, "temp_f": 72, "condition": "Sunny"}, nil
}

func getStockPrice(ctx context.Context, in stockIn) (map[string]any, error) {
	return map[string]any{"symbol": in.Symbol, "price": 182.50, "change": "+1.2%"}, nil
}

func main() {
	agent := &ai.Agent{
		Name:  "weather_stock_agent",
		Model: "openai/gpt-4o",
		Tools: ai.Tools(
			tool.Func("get_weather", getWeather, "Get the current weather for a city."),
			tool.Func("get_stock_price", getStockPrice, "Get the current stock price for a ticker symbol."),
		),
		Instructions: "You are a helpful assistant. Use tools to answer questions.",
	}

	runtime := ai.NewRuntime(ai.Config{})
	defer runtime.Shutdown()

	result, err := runtime.Run(context.Background(), agent, "What's the weather like in San Francisco?")
	if err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
	result.PrintResult()
}
```

`tool.Func` takes the name, the function, then the description; `ai.Tools` collects them. The model
picks one: this prompt calls `get_weather`, not `get_stock_price`.

The runtime registers each function as a Conductor worker named after the tool and polls for it
while `Run` blocks, so every call the model makes is its own task in the UI. Keep the process alive
until the run finishes, or the tool task stays scheduled.

## Run the repository's examples

[examples/agents](../../examples/agents/) is a package with one file per example and a command
that runs one by name; run it with no argument for the list.

```shell
cd examples
go run ./agents/cmd 01_basic_agent
```

The examples default `Model` to `openai/gpt-4o-mini` when `CONDUCTOR_AGENT_LLM_MODEL` is unset.

## Where to go next

`Run` is the development path: one process defines the agent, runs it and waits. For production,
`Deploy` publishes the definition, `Serve` hosts the tool workers, and a trigger elsewhere starts
runs by name.

- [concepts/agents.md](concepts/agents.md) — agent fields, results, and the loop
- [concepts/tools.md](concepts/tools.md) — worker, HTTP and MCP tools, schemas, credentials
- [concepts/deploy-serve-run.md](concepts/deploy-serve-run.md) — `Plan`, `Deploy`, `Serve`, `Run`, `Start`
- [reference/runtime.md](reference/runtime.md) — `ai.Runtime` API and configuration
