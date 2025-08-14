## Logging Configuration

The SDK uses a configurable logging system that allows you to customize the logging behavior to fit your application's needs.

### Using Custom Logger

You can configure a custom logger by calling the `SetLogger` function with any implementation of the `Logger` interface:

```go
import "github.com/conductor-sdk/conductor-go/sdk/log"

// Set your custom logger
log.SetLogger(yourCustomLogger)

// Reset to default logger
log.SetLogger(nil)
```

### Zap Logger Integration

The SDK provides built-in support for the popular [Zap logger](https://github.com/uber-go/zap). You can easily integrate Zap with the SDK:

```go
import (
    "go.uber.org/zap"
    "github.com/conductor-sdk/conductor-go/sdk/log"
)

// Create a Zap logger
zapLogger, _ := zap.NewProduction()
defer zapLogger.Sync()

// Set it as the SDK logger
log.SetLogger(log.NewZap(zapLogger))
```

The Zap adapter automatically handles structured logging and converts the SDK's logging calls to Zap's structured format.