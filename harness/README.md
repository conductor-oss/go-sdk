# Go SDK Docker Harness

Two Docker targets built from the root `Dockerfile`: an **integration test runner** and a **long-running worker harness**.

## Integration Test Runner

Runs the Go integration test suite against a Conductor cluster and exits with 0 (pass) or 1 (fail).

```bash
docker build --target inttest -t go-sdk-tests \
  --build-arg CONDUCTOR_SERVER_URL=https://your-cluster.example.com/api \
  --build-arg CONDUCTOR_AUTH_KEY=$CONDUCTOR_AUTH_KEY \
  --build-arg CONDUCTOR_AUTH_SECRET=$CONDUCTOR_AUTH_SECRET \
  .
```

## Worker Harness

A self-feeding worker that runs indefinitely. On startup it registers five simulated tasks (`go_worker_0` through `go_worker_4`) and the `go_simulated_tasks_workflow`, then runs two background services:

- **WorkflowGovernor** -- starts a configurable number of `go_simulated_tasks_workflow` instances per second (default 2), indefinitely.
- **SimulatedTaskWorkers** -- five task handlers, each with a codename and a default sleep duration. Each worker supports configurable delay types, failure simulation, and output generation via task input parameters. The workflow chains them in sequence: quickpulse (1s) → whisperlink (2s) → shadowfetch (3s) → ironforge (4s) → deepcrawl (5s).

### Building Locally

```bash
docker build --target harness -t go-sdk-harness .
```

### Multiplatform Build and Push

To build for both `linux/amd64` and `linux/arm64` and push to GHCR:

```bash
# One-time: create a buildx builder if you don't have one
docker buildx create --name multiarch --use --bootstrap

# Build and push
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  --target harness \
  -t ghcr.io/conductor-oss/go-sdk/harness-worker:latest \
  --push .
```

> **Note:** Multi-platform builds require `docker buildx` and a builder that supports cross-compilation. On macOS this works out of the box with Docker Desktop. On Linux you may need to install QEMU user-space emulators:
>
> ```bash
> docker run --privileged --rm tonistiigi/binfmt --install all
> ```

### Running

```bash
docker run -d \
  -e CONDUCTOR_SERVER_URL=https://your-cluster.example.com/api \
  -e CONDUCTOR_AUTH_KEY=$CONDUCTOR_AUTH_KEY \
  -e CONDUCTOR_AUTH_SECRET=$CONDUCTOR_AUTH_SECRET \
  -e HARNESS_WORKFLOWS_PER_SEC=4 \
  go-sdk-harness
```

You can also run the harness locally without Docker:

```bash
export CONDUCTOR_SERVER_URL=https://your-cluster.example.com/api
export CONDUCTOR_AUTH_KEY=$CONDUCTOR_AUTH_KEY
export CONDUCTOR_AUTH_SECRET=$CONDUCTOR_AUTH_SECRET

go run ./harness
```

Override defaults with environment variables as needed:

```bash
HARNESS_WORKFLOWS_PER_SEC=4 HARNESS_BATCH_SIZE=10 go run ./harness
```

All resource names use a `go_` prefix so multiple SDK harnesses (C#, Python, Java, etc.) can coexist on the same cluster.

### Environment Variables

| Variable | Required | Default | Description |
|---|---|---|---|
| `CONDUCTOR_SERVER_URL` | yes | -- | Conductor API base URL |
| `CONDUCTOR_AUTH_KEY` | no | -- | Orkes auth key |
| `CONDUCTOR_AUTH_SECRET` | no | -- | Orkes auth secret |
| `HARNESS_WORKFLOWS_PER_SEC` | no | 2 | Workflows to start per second |
| `HARNESS_BATCH_SIZE` | no | 20 | Number of tasks each worker polls per batch |
| `HARNESS_POLL_INTERVAL_MS` | no | 100 | Milliseconds between poll cycles |
