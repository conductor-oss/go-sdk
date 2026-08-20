# Conductor OSS Go SDK

[![Build Status](https://github.com/conductor-oss/go-sdk/actions/workflows/build.yml/badge.svg)](https://github.com/conductor-oss/go-sdk/actions/workflows/build.yml)

SDK for developing Go applications that create, manage and execute workflows, and run workers.

[Conductor](https://www.conductor-oss.org/) is the leading open-source orchestration platform allowing developers to build highly scalable distributed applications.

To learn more about Conductor checkout our [developer's guide](https://docs.conductor-oss.org/devguide/concepts/index.html) and give it a ⭐ to make it famous!

[![GitHub stars](https://img.shields.io/github/stars/conductor-oss/conductor.svg?style=social&label=Star&maxAge=)](https://GitHub.com/conductor-oss/conductor/)

## Requirements and compatibility

Go 1.23+. For the exercised server matrix, the published module, and how this SDK's
capabilities compare with the Java and Python SDKs, see [compatibility](docs/compatibility.md).

## Install the SDK

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

## Workflow and worker quickstart

Run the maintained Hello World example from this repository.

> [!note]
> You will need an up & running Conductor Server.
>
> For details on how to run Conductor take a look at [our documentation](https://docs.conductor-oss.org).
>
> The examples expect the server to be listening on http://localhost:8080.

```shell
export CONDUCTOR_SERVER_URL="http://localhost:8080/api"
cd examples
go run hello_world/main.go
```

With an [Orkes developer account](https://developer.orkescloud.com) instead:

```shell
export CONDUCTOR_SERVER_URL="https://developer.orkescloud.com/api"
export CONDUCTOR_AUTH_KEY="..."
export CONDUCTOR_AUTH_SECRET="..."
cd examples
go run hello_world/main.go
```

The above should give an output similar to

```shell
INFO[0000] Started 1 worker(s) for taskName greet, polling in interval of 100 ms 
INFO[0000] Started workflow with Id:14a9fcc5-3d74-11ef-83dc-acde48001122 
INFO[0000] Output of the workflow:map[Greetings:Hello, Gopher] 
```

For the full walkthrough — creating the workflow in code, writing the worker, and
running the application — see the [core quickstart](docs/core-quickstart.md).

## Documentation

Start at the [documentation hub](docs/README.md).

- [Core quickstart](docs/core-quickstart.md) — the end-to-end Hello World
- [Authoring workflows](docs/workflows.md)
- [Writing workers](docs/workers.md)
- [Connection and authentication](docs/connection-authentication.md) — client setup, environment variables, and proxies
- [Security](docs/security.md) — TLS/SSL for self-signed certificates and mTLS
- [Observability](docs/observability.md) — Prometheus metrics and logging
- [Upgrading](docs/upgrading.md) — deprecated methods and their replacements
- [Compatibility](docs/compatibility.md) — supported baselines and capability differences

The Go SDK does not provide the Conductor agent runtime available in the Java and
Python SDKs. It does expose the LLM prompt and integration clients. See
[documentation parity](docs/documentation-parity.md) for the full map.

## Deprecated Methods

Some methods in the SDK client interfaces are now deprecated. They’ve been replaced with newer methods that follow more consistent naming. Please refer to our [Migration Guide](docs/upgrading.md) for detailed information on how to update your code.

## License

Apache License 2.0 — see [LICENSE](LICENSE).
