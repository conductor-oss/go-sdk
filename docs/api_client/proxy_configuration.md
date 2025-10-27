# Proxy Configuration

The Conductor Go SDK supports comprehensive proxy configuration for both HTTP and HTTPS traffic. This is essential when your application needs to route traffic through corporate firewalls, load balancers, or other network intermediaries.

## Table of Contents

- [Supported Proxy Types](#supported-proxy-types)
- [How Proxy Works](#how-proxy-works)
- [Basic Proxy Configuration](#basic-proxy-configuration)
- [Environment Variable Configuration](#environment-variable-configuration)
- [Troubleshooting](#troubleshooting)

## Supported Proxy Types

- **HTTP Proxy**: `http://proxy.example.com:8080`
- **HTTPS Proxy**: `https://proxy.example.com:8443`
- **Proxy with Authentication**: `http://username:password@proxy.example.com:8080`

## How Proxy Works

The Conductor Go SDK uses a **fallback mechanism** for proxy configuration:

### 1. **Custom Proxy Configuration**
If a custom proxy is configured via `CONDUCTOR_PROXY` or settings options:
- **Valid proxy URL**: Uses the custom proxy with credentials support
- **Invalid/empty proxy URL**: **Critical error** - `log.Fatalf("failed to load proxy from environment: %v", err)`

### 2. **System Proxy Fallback**
If no custom proxy is configured, the SDK automatically falls back to `http.ProxyFromEnvironment`, which parses these **system environment variables**:

| Variable | Description | Example |
|----------|-------------|---------|
| `HTTP_PROXY` | HTTP proxy for HTTP requests | `http://proxy.company.com:8080` |
| `HTTPS_PROXY` | HTTP proxy for HTTPS requests | `http://proxy.company.com:8080` |
| `NO_PROXY` | Comma-separated list of hosts to bypass proxy | `localhost,127.0.0.1,.local` |

### 3. **Proxy Headers Support**
For custom proxy headers (like `Proxy-Authorization`, `X-Proxy-Client`, etc.), use the `WithHTTPHeaders()` option:

```go
clientSettings := settings.NewClientSettings(
    settings.WithProxyURL("http://proxy.company.com:8080"),
    settings.WithHTTPHeaders(map[string]string{
        "Proxy-Authorization": "Basic dXNlcm5hbWU6cGFzc3dvcmQ=",
        "X-Proxy-Client": "Conductor-Go-SDK/1.0",
        "User-Agent": "MyApp/1.0",
    }),
)
```

## Basic Proxy Configuration

### Using Settings Options

```go
import (
    "github.com/conductor-sdk/conductor-go/sdk/client"
    "github.com/conductor-sdk/conductor-go/sdk/settings"
)

// Basic HTTP proxy configuration
clientSettings := settings.NewClientSettings(
    settings.WithServerURL("https://api.orkes.io/api"),
    settings.WithAuthCredentials("your_key", "your_secret"),
    settings.WithProxyURL("http://proxy.company.com:8080"),
)

apiClient := client.NewAPIClientFromSettings(clientSettings)

// Or using complete proxy settings
proxySettings := settings.NewProxySettings()
proxySettings.URL, _ = url.Parse("http://proxy.company.com:8080")
proxySettings.Username = "username"
proxySettings.Password = "password"

clientSettings := settings.NewClientSettings(
    settings.WithServerURL("https://api.orkes.io/api"),
    settings.WithAuthCredentials("your_key", "your_secret"),
    settings.WithProxySettings(proxySettings),
)

apiClient := client.NewAPIClientFromSettings(clientSettings)
```


## Environment Variable Configuration

You can configure proxy settings using Conductor-specific environment variables:

```bash
# Basic Conductor configuration
export CONDUCTOR_SERVER_URL="https://api.orkes.io/api"
export CONDUCTOR_AUTH_KEY="your_key"
export CONDUCTOR_AUTH_SECRET="your_secret"

# Basic proxy configuration
export CONDUCTOR_PROXY=http://proxy.company.com:8080

# Proxy with authentication
export CONDUCTOR_PROXY=http://username:password@proxy.company.com:8080

# HTTPS proxy
export CONDUCTOR_PROXY=https://secure-proxy.company.com:8443
```

**Priority Order:**
1. **Explicit proxy parameters** in settings options (highest priority)
2. **`CONDUCTOR_PROXY`** environment variable (medium priority)
3. **System proxy environment variables** (`HTTP_PROXY`, `HTTPS_PROXY`, etc.) (fallback)

### Example Usage with Environment Variables

```go
import (
    "os"
    "github.com/conductor-sdk/conductor-go/sdk/client"
    "github.com/conductor-sdk/conductor-go/sdk/settings"
)

func main() {
    // Set Conductor environment variables
    os.Setenv("CONDUCTOR_SERVER_URL", "https://api.orkes.io/api")
    os.Setenv("CONDUCTOR_AUTH_KEY", "your_key")
    os.Setenv("CONDUCTOR_AUTH_SECRET", "your_secret")
    
    // Set proxy environment variable
    os.Setenv("CONDUCTOR_PROXY", "http://proxy.company.com:8080")
    
    // Configuration will automatically use proxy from environment
    apiClient := client.NewAPIClientFromEnv()
}
```

### Proxy with Custom Headers

```go
import (
    "github.com/conductor-sdk/conductor-go/sdk/client"
    "github.com/conductor-sdk/conductor-go/sdk/settings"
)

func main() {
    // Configure proxy with custom headers
    clientSettings := settings.NewClientSettings(
        settings.WithServerURL("https://api.orkes.io/api"),
        settings.WithAuthCredentials("your_key", "your_secret"),
        settings.WithProxyURL("http://proxy.company.com:8080"),
        settings.WithHTTPHeaders(map[string]string{
            // Proxy-specific headers
            "Proxy-Authorization": "Basic dXNlcm5hbWU6cGFzc3dvcmQ=",
            "X-Proxy-Client": "Conductor-Go-SDK/1.0",
            "X-Proxy-Request-ID": "req-12345",
            
            // General HTTP headers
            "User-Agent": "MyApp/1.0",
            "X-Custom-Header": "custom-value",
        }),
    )
    
    apiClient := client.NewAPIClientFromSettings(clientSettings)
}
```

## Troubleshooting

### Common Proxy Issues

1. **Critical Error: "failed to load proxy from environment"**
   - **Cause**: `CONDUCTOR_PROXY` contains invalid URL format
   - **Solution**: Fix the proxy URL format (e.g., `http://proxy.company.com:8080`)
   - **Example**: `export CONDUCTOR_PROXY="http://proxy.company.com:8080"` (not `proxy.company.com:8080`)

2. **System Proxy Not Working**
   - Verify system environment variables (`HTTP_PROXY`, `HTTPS_PROXY`)
   - Check if `NO_PROXY` is blocking the target host
   - Ensure system proxy variables are set correctly