# Scheduling

Running a deployed agent on a cron cadence with `ai.Schedule` and the runtime's
agent-scoped schedule methods.

Schedules live in Conductor and survive server restarts. No cron daemon and no extra
process are involved: the scheduler starts the agent's workflow, and whatever is running
[`Serve`](./deploy-serve-run.md) answers its tool calls.

Deploy the agent first. A schedule starts the workflow stored under the agent's name, so
without a `Deploy` there is nothing to start.

```go
rt := ai.NewRuntime(ai.Config{})
defer rt.Shutdown()

if _, err := rt.Deploy(ctx, reportAgent); err != nil {
    return err
}

err := rt.SaveSchedule(ctx, reportAgent.Name, ai.Schedule{
    Name:     "nightly",
    Cron:     "0 0 * * *",
    Timezone: "America/New_York",
    Input:    map[string]any{"prompt": "Summarize yesterday's incidents."},
})
```

## ai.Schedule

| Field | Type | Meaning |
|---|---|---|
| `Name` | `string` | Identifies the schedule within its agent. Required. |
| `Cron` | `string` | The cron expression. Required. |
| `Timezone` | `string` | Zone the cron runs in. Empty means UTC. |
| `Input` | `map[string]any` | The input the scheduled run starts with; the agent's prompt goes under `"prompt"`. |
| `Catchup` | `bool` | Run occurrences missed while the scheduler was down. Off by default. |
| `Paused` | `bool` | Create the schedule paused. |
| `StartAt`, `EndAt` | `int64` | Epoch millis bounding when the schedule is active. Zero means unbounded. |
| `Description` | `string` | Free text. |

`Schedule.Validate` requires `Name` and `Cron`, and requires `StartAt` to be before `EndAt`
when both are set. `SaveSchedule` and `ReconcileSchedules` call it.

## Wire name

A schedule is stored under `<agent>-<Name>`, matching the Python SDK. The schedule above is
`daily_report-nightly` on the server and in the UI. The runtime's methods all take the agent
name and the short name and do the joining, so `Name` only has to be unique within its
agent. Reading a schedule back strips the prefix, so `Name` round-trips as the short name.

## Runtime methods

Every one is scoped to a single agent.

| Method | Effect |
|---|---|
| `SaveSchedule(ctx, agentName string, s Schedule) error` | Creates or updates one schedule. |
| `GetSchedule(ctx, agentName, name string) (*Schedule, error)` | Reads one by its short name. |
| `ListSchedules(ctx, agentName string) ([]Schedule, error)` | Every schedule whose wire name carries the agent's prefix. |
| `DeleteSchedule(ctx, agentName, name string) error` | Removes one by its short name. |
| `PauseSchedule(ctx, agentName, name string) error` | Stops it firing, keeping the definition. |
| `ResumeSchedule(ctx, agentName, name string) error` | The inverse of `PauseSchedule`. |
| `ReconcileSchedules(ctx, agentName string, desired []Schedule) error` | Makes the agent's schedules match `desired` exactly. |

```go
schedules, err := rt.ListSchedules(ctx, "daily_report")
err = rt.PauseSchedule(ctx, "daily_report", "nightly")
err = rt.ResumeSchedule(ctx, "daily_report", "nightly")
err = rt.DeleteSchedule(ctx, "daily_report", "nightly")
```

## Reconciling

`ReconcileSchedules` is the declarative form: it deletes every schedule of this agent that
`desired` does not name, then saves the ones it does. A nil `desired` is a no-op, so it
cannot delete by accident; an empty non-nil slice deletes all of the agent's schedules. It
validates the whole list first and rejects duplicate names.

```go
err := rt.ReconcileSchedules(ctx, reportAgent.Name, []ai.Schedule{
    {Name: "nightly", Cron: "0 0 * * *", Input: map[string]any{"prompt": "Nightly summary."}},
    {Name: "weekly", Cron: "0 8 * * MON", Timezone: "UTC",
        Input: map[string]any{"prompt": "Weekly digest."}, Description: "Monday digest"},
})
```

Deploying an agent does not reconcile its schedules. Call it from the same release step that
calls `Deploy` if you want the two to move together.

## Below the facade

`ai.Schedule` is a thin agent-scoped view of Conductor's workflow scheduler: it compiles to
a `model.SaveScheduleRequest` whose `StartWorkflowRequest` names the agent's workflow. For
anything the facade does not cover — tags, search, `GetNextFewSchedules`, pausing every
schedule on the server — build a `client.SchedulerClient` with
`client.NewSchedulerClient(apiClient)` and call it directly, sharing the same
`*client.APIClient` you passed to `ai.NewRuntimeWithClient`.

## Next steps

Continue with [deploy, serve, run](./deploy-serve-run.md), [stateful agents](./stateful.md), or the [runtime reference](../reference/runtime.md).
