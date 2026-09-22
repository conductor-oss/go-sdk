# Tools

Tools are what an agent can do. This page covers the `tool` package: worker tools written as Go
functions, the schema derived from their types, the tool types the server runs itself, and the
options every tool shares.

Every tool call runs as a Conductor task — durable, retryable and visible in the execution
history. Only **worker** tools are dispatched back to your process; every other tool type is
executed by the server, so no Go code runs and nothing needs to be polling.

## Worker tools

`tool.Func` turns a Go function into a tool. The runtime registers the function as a Conductor
worker under `name`, which is both what the model calls and the task name workers poll for.

```go
import (
	"context"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/ai/tool"
)

type orderIn struct {
	OrderID string
}

type orderOut struct {
	Status string
	Total  float64
}

func lookupOrder(ctx context.Context, in orderIn) (orderOut, error) {
	return orderOut{Status: "shipped", Total: 42.50}, nil
}

agent := &ai.Agent{
	Name:  "support_agent",
	Model: "openai/gpt-4o",
	Tools: ai.Tools(
		tool.Func("lookup_order", lookupOrder, "Look up an order by its id."),
	),
}
```

The handler must be `func(context.Context, In) (Out, error)`. `In` and `Out` may be any JSON-able
types; `map[string]any` and `string` are common when the shape is loose.

## Input and output schemas

The model sees JSON Schema derived by reflection from `In`, and the same is done for `Out`.

- A field with a `json` tag uses the tag name.
- An **untagged** field is advertised in snake_case: `AccountID` becomes `account_id`,
  `HTTPStatus` becomes `http_status`, `TempF` becomes `temp_f`. This is the recommended style —
  it reads like a Python tool without a tag on every field.
- A field is **required** unless it is a pointer or carries `omitempty`.
- `json:"-"` and unexported fields are skipped; an embedded struct's fields are lifted, as
  `encoding/json` does.
- Properties keep declaration order, because `required` is a positional array the server copies
  into the text the model reads.
- `[]byte` is a string, `map[string]T` is an object with `additionalProperties`, `any` is an open
  schema.

```go
type transferIn struct {
	FromAcct string              // from_acct, required
	ToAcct   string              // to_acct, required
	Amount   float64             // amount, required
	Memo     string  `json:"note,omitempty"` // note, optional
	DryRun   *bool                           // dry_run, optional
}
```

Only types and required-ness are emitted — no descriptions, titles or formats — so a Go tool puts
the same document on the wire as the Python and Java SDKs. Put the prose in the tool's
description instead.

The mapping is applied in both directions: arguments arrive under the schema names and are bound
back onto the Go field names, and the return value is re-keyed to the names the schema promised.

## Server-side tools

These need no worker. Each takes the name first and the description after the type-specific
arguments.

```go
search := tool.HTTP("search_products", "https://api.example.com/search",
	"Search the product catalog.",
	tool.WithMethod("POST"),
	tool.WithHeaders(map[string]string{"Authorization": "Bearer ${API_TOKEN}"}),
	tool.WithCredentials("API_TOKEN"))

files := tool.MCP("filesystem", "http://localhost:3001", "Access files over MCP.")

stripe := tool.API("stripe", "https://api.example.com/openapi.json", "Stripe operations.")
```

| Constructor | What the server does |
|---|---|
| `tool.HTTP(name, url, description, opts...)` | Calls the endpoint itself. Defaults to `GET`, `application/json` for both `accept` and `contentType`, and no headers. Override with `tool.WithMethod`, `tool.WithHeaders`, `tool.WithAccept`, `tool.WithContentType`. |
| `tool.MCP(name, serverURL, description, opts...)` | Lists the MCP server's tools at compile time and expands this definition into one tool per entry, each call running as a `CALL_MCP_TOOL` task. The input schema is empty because the discovered tools carry their own. Empty name and description default to `"mcp_tools"` and `"MCP tools from <serverURL>"`. |
| `tool.API(name, url, description, opts...)` | Fetches an OpenAPI 3.x, Swagger 2.0 or Postman document and expands it into a tool per operation. Empty name and description default to `"api_tools"` and `"API tools from <url>"`. |

`tool.WithMaxTools(n)` caps how many discovered tools MCP and API expand to (default 64) before
the server has the model pick a subset per turn. `tool.WithToolNames(...)` is sent for parity with
the Python SDK; current servers do not filter on it.

A `${NAME}` placeholder in a header is resolved server-side from the credential store and never
passes through your process. The same name must appear in `tool.WithCredentials`, or `Validate`
rejects the tool.

## Human approval and human input

`tool.RequiresApproval()` is an option on **any** tool: the run pauses before the call is
dispatched and waits for a person.

```go
transfer := tool.Func("transfer_funds", transferFunds,
	"Transfer money between accounts.", tool.RequiresApproval())
```

`tool.Human(name, description, opts...)` is a tool whose whole purpose is to ask a person; it
compiles to a Conductor HUMAN task. Its default schema is a single `question` string — use
`tool.WithInputSchema` to collect something structured.

Drive either from the handle returned by `Start`:

```go
handle, _ := runtime.Start(ctx, agent, "Transfer $500 from ACC-789 to ACC-456.")
events, _ := handle.Events(ctx)
for event := range events {
	if event.Type == ai.EventWaiting {
		_ = handle.Approve(ctx)               // or handle.Reject(ctx, "declined")
	}
}
```

`handle.Respond(ctx, map[string]any{...})` answers a `tool.Human` call. The run is durable, so it
can wait for as long as the person needs. See [Streaming and human input](./streaming-hitl.md).

## Media and PDF

The server generates the media against the provider in the config; the model supplies only the
dynamic parameters.

```go
img := tool.Image("generate_image", "openai", "dall-e-3", "Generate an image from a prompt.")
tts := tool.Audio("speak", "openai", "tts-1", "Convert text to speech.")
clip := tool.Video("generate_video", "openai", "sora", "Generate a short video.")
report := tool.PDF("build_report", "Render Markdown to a PDF.")
```

`tool.PDF` takes no provider — it renders Markdown. The default input schemas match the Python
SDK field for field; replace one with `tool.WithInputSchema` or add fixed generation parameters
with `tool.WithConfig`.

## RAG: index and search

Both run against a vector database configured on the server.

```go
index := tool.Index("index_doc", "pgvectordb", "docs", "openai", "text-embedding-3-small",
	"Add a document to the knowledge base.",
	tool.WithChunking(1000, 100))

search := tool.Search("search_docs", "pgvectordb", "docs", "openai", "text-embedding-3-small",
	"Search the knowledge base.",
	tool.WithMaxResults(3))
```

`tool.Index` takes `text`, `docId` and optional `metadata` from the model; `tool.Search` takes
`query` and returns up to `WithMaxResults` matches (default 5). Both default to the `default_ns`
namespace — change it with `tool.WithNamespace`. `tool.WithDimensions` overrides the embedding
size.

## Waiting for messages

`tool.WaitForMessage(name, description, opts...)` pauses the agent until a message is sent into
the execution with `runtime.Signal`, `runtime.SendMessage` or the handle's equivalents. It blocks
for one message unless `tool.WithBatchSize(n)` or `tool.NonBlocking()` says otherwise, and needs
`conductor.workflow-message-queue.enabled=true` on the server.

## Agent as a tool

`tool.Agent(name, agent, description, opts...)` exposes another agent as a callable tool. The
sub-agent is serialized into this tool's config and needs no registration of its own; the parent's
LLM calls it inline with a single `request` string and gets its output back.

```go
researcher := &ai.Agent{Name: "researcher", Model: "openai/gpt-4o",
	Instructions: "Research a topic and return a summary."}

manager := &ai.Agent{
	Name:         "manager",
	Model:        "openai/gpt-4o",
	Instructions: "Use the researcher tool to gather information.",
	Tools:        ai.Tools(tool.Agent("", researcher, "")),
}
```

An empty name takes the sub-agent's name and an empty description is generated. Putting the agent
in the parent's `Agents` instead delegates control rather than calling inline; see
[Multi-agent](./multi-agent.md).

## External workers

`tool.External[In, Out](name, description, opts...)` declares a worker tool whose worker runs
somewhere else — another service, another language, an existing task definition. The type
parameters give the model the same schema a `tool.Func` handler's types would, but no worker
starts here, so Conductor dispatches each call to whatever is polling for `name`.

```go
var processOrder = tool.External[orderIn, map[string]any]("process_order",
	"Process a customer order. Actions: refund, cancel, update.")
```

## Credentials

Go cannot inspect a function body to discover which secrets it touches, so declare them:

```go
createIssue := tool.Func("create_issue", createIssueFn,
	"Create a GitHub issue.", tool.WithCredentials("GITHUB_TOKEN"))
```

The names travel in the tool's task definition. A secured server resolves them when the worker
polls and attaches the values to the task; nothing is read from the environment and no extra
network call is made. Read them inside the handler:

```go
func createIssueFn(ctx context.Context, in issueIn) (string, error) {
	token, err := ai.Secret(ctx, "GITHUB_TOKEN")
	if err != nil {
		return "", err   // ai.ErrCredentialNotFound when it was not delivered
	}
	// ...
}
```

`ai.SecretsEnv(ctx, names...)` returns `KEY=VALUE` strings for a child process, for a tool that
shells out — `os.Setenv` would be process-wide and racy, since workers are goroutines in one
process.

Only names declared with `tool.WithCredentials` are delivered. `Agent.Credentials` is the
agent-wide fallback for tools that cannot declare their own.

## Options

Options apply to any tool unless noted.

| Option | Effect |
|---|---|
| `tool.WithCredentials(names...)` | Secret names this tool may read. |
| `tool.RequiresApproval()` | Pause for a human before the call is dispatched. |
| `tool.WithTimeout(seconds)` | Bound one call. |
| `tool.WithRetry(count, delaySeconds, policy)` | Configure the registered task definition. `policy` is `ai.RetryFixed`, `ai.RetryLinearBackoff` or `ai.RetryExponentialBackoff`; the default is 2 retries, 2 seconds apart, linear. |
| `tool.WithMaxCalls(n)` | Cap how many times the agent may call this tool in a run. |
| `tool.Stateful()` | Route the tool to a per-execution worker domain, so one run's calls reach one process. |
| `tool.WithInputSchema(doc)` | Replace the schema the model is shown. |
| `tool.WithConfig(key, value)` | Set a type-specific config key directly. |
| `tool.WithMethod`, `tool.WithHeaders`, `tool.WithAccept`, `tool.WithContentType` | HTTP. |
| `tool.WithMaxTools`, `tool.WithToolNames` | MCP and API. |
| `tool.WithNamespace`, `tool.WithChunking`, `tool.WithDimensions`, `tool.WithMaxResults` | Index and Search. |
| `tool.WithBatchSize`, `tool.NonBlocking` | WaitForMessage. |

`ToolDef.Validate` rejects an empty or malformed name, an unknown tool type or retry policy, an
MCP tool with no `server_url`, an API tool with no `url`, and a `${NAME}` header placeholder that
is not declared in the tool's credentials. `Agent.Validate` runs it for every tool.

See [Agents](./agents.md) for how tools attach to an agent, and
[Deploy · Serve · Run](./deploy-serve-run.md) for where the workers run.

## Next steps

Continue with [guardrails](./guardrails.md), [structured output](./structured-output.md), or [multi-agent](./multi-agent.md).
