# AGENTS.md

Conventions for agents working in this repository.

## Comments

Exported doc comments state the **contract**: what a caller needs to use the symbol.
They render on pkg.go.dev for every SDK consumer, so they carry only that.

```go
// JumpToTask jumps to a specific task in a running workflow.
```

One line usually carries a contract. Add a second only where a caller would otherwise
get it wrong — an argument that must be non-nil, a call that blocks, an error worth
handling specially.

Implementation detail goes in a body comment beside the code it explains: how a
request is shaped, a server quirk being worked around, why a value is sent twice.

```go
// Sent in the path and the query: the route needs the segment, the handler binds
// the value from the query parameter.
```

Evidence, measurements, and the reasoning behind a change go in the commit message,
where they stay with the change and out of the published API docs.

## Integration tests

`test/integration_tests` runs against a live Conductor server, configured by
`CONDUCTOR_SERVER_URL`, `CONDUCTOR_AUTH_KEY` and `CONDUCTOR_AUTH_SECRET`. Some tests
are gated on server version via `testdata.RequireAtLeast` and report as skipped below
it — check for `SKIPPED` before concluding a test passed.
