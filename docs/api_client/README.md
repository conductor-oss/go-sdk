# API Client Configuration

The Conductor Go SDK provides flexible API client configuration with support for authentication, HTTP settings, proxy configuration, and environment variable management.

## Table of Contents

- [Quick Start](#quick-start)
- [Client Creation Methods](#client-creation-methods)
- [Environment Variables](#environment-variables)
- [Configuration Options](#configuration-options)
- [Advanced Configuration](#advanced-configuration)
- [Examples](#examples)
- [Related Documentation](#related-documentation)

## Quick Start

### Basic Client Creation

```go
import (
    "github.com/conductor-sdk/conductor-go/sdk/client"
    "github.com/conductor-sdk/conductor-go/sdk/settings"
)

// Create client from environment variables
apiClient := client.NewAPIClientFromEnv()

// Create client from environment with additional options
apiClient := client.NewAPIClientFromEnv(
    settings.WithServerURL("https://custom-server.com/api"),
    settings.WithHTTPTimeout(60*time.Second),
)

// Create client with centralized settings
clientSettings := settings.NewClientSettings(
    settings.WithServerURL("https://your-server.com/api"),
    settings.WithAuthCredentials("your_key", "your_secret"),
)
apiClient := client.NewAPIClientFromSettings(clientSettings)
```
## Client Creation Methods

> **All client creation methods support options functions** that allow you to modify individual client properties. This provides maximum flexibility for configuration.
> 
> **Recommended approach**: Use `NewAPIClientFromEnv()` with options for most cases, or `NewAPIClientFromSettings()` for complex configurations.

## Environment Variables

The SDK supports the following environment variables:

| Variable | Description | Default | Example |
|----------|-------------|---------|---------|
| `CONDUCTOR_SERVER_URL`* | Conductor server API URL | `http://localhost:8080/api` | `https://api.orkes.io` or `https://api.orkes.io/api` |
| `CONDUCTOR_AUTH_KEY` | Authentication key | - | `your_auth_key` |
| `CONDUCTOR_AUTH_SECRET` | Authentication secret | - | `your_auth_secret` |
| `CONDUCTOR_CLIENT_HTTP_TIMEOUT` | HTTP timeout in seconds | `30` (default) | `60` |
| `CONDUCTOR_PROXY` | Proxy URL | - | `http://proxy.company.com:8080` |

\* You can provide the server URL with or without the `/api` suffix. The SDK will normalize it to end with `/api`.

### Environment Variable Examples

```bash
# Basic configuration
export CONDUCTOR_SERVER_URL="https://api.orkes.io/api"
export CONDUCTOR_AUTH_KEY="your_key"
export CONDUCTOR_AUTH_SECRET="your_secret"

# With timeout (default is 30 seconds)
export CONDUCTOR_CLIENT_HTTP_TIMEOUT="60"

# With proxy
export CONDUCTOR_PROXY="http://proxy.company.com:8080"
```

## Configuration Options

### Authentication Options

```go
// Set authentication credentials
clientSettings := settings.NewClientSettings(
    settings.WithAuthCredentials("your_key", "your_secret"),
)
```

### HTTP Options

```go
// Set server URL
clientSettings := settings.NewClientSettings(
    settings.WithServerURL("https://api.orkes.io/api"),
)

// Set HTTP timeout (default is 30 seconds)
clientSettings := settings.NewClientSettings(
    settings.WithHTTPTimeout(60*time.Second),
)

// Set custom headers
headers := map[string]string{
    "User-Agent": "MyApp/1.0",
    "X-Custom-Header": "value",
}
clientSettings := settings.NewClientSettings(
    settings.WithHTTPHeaders(headers),
)
```

### Proxy Options

```go
// Set proxy URL
clientSettings := settings.NewClientSettings(
    settings.WithProxyURL("http://proxy.company.com:8080"),
)

// Set proxy with credentials
clientSettings := settings.NewClientSettings(
    settings.WithProxyURL("http://username:password@proxy.company.com:8080"),
)

// Set proxy from parsed URL
proxyURL, _ := url.Parse("http://proxy.company.com:8080")
clientSettings := settings.NewClientSettings(
    settings.WithProxy(proxyURL),
)

// Set proxy credentials separately
clientSettings := settings.NewClientSettings(
    settings.WithProxyURL("http://proxy.company.com:8080"),
    settings.WithProxyCredentials("username", "password"),
)
```

### Token Options

```go
import "github.com/conductor-sdk/conductor-go/sdk/authentication"

// Set custom token expiration settings
tokenExpiration := authentication.NewTokenExpiration(
    45*time.Minute,  // default expiration
    3*time.Hour,     // cleanup interval
)
clientSettings := settings.NewClientSettings(
    settings.WithTokenExpiration(tokenExpiration),
)

// Set custom token manager
tokenManager := authentication.NewTokenManager(authSettings, tokenExpiration)
clientSettings := settings.NewClientSettings(
    settings.WithTokenManager(tokenManager),
)
```

## Advanced Configuration

### Complete Configuration Example

```go
package main

import (
    "log"
    "time"
    
    "github.com/conductor-sdk/conductor-go/sdk/authentication"
    "github.com/conductor-sdk/conductor-go/sdk/client"
    "github.com/conductor-sdk/conductor-go/sdk/settings"
)

func main() {
    // Method 1: Using option functions (recommended)
    clientSettings := settings.NewClientSettings(
        // Server configuration
        settings.WithServerURL("https://api.orkes.io/api"),
        settings.WithHTTPTimeout(120*time.Second),
        settings.WithHTTPHeaders(map[string]string{
            "User-Agent": "MyProductionApp/1.0",
        }),
        
        // Authentication
        settings.WithAuthCredentials("prod_key", "prod_secret"),
        
        // Proxy configuration
        settings.WithProxyURL("http://corporate-proxy.com:8080"),
        settings.WithProxyCredentials("proxy_user", "proxy_pass"),
        
        // Token management
        settings.WithTokenExpiration(authentication.NewTokenExpiration(
            45*time.Minute,  // default expiration
            3*time.Hour,     // cleanup interval
        )),
        settings.WithTokenManager(authentication.NewTokenManager(
            settings.NewAuthenticationSettings("prod_key", "prod_secret"),
            authentication.NewTokenExpiration(45*time.Minute, 3*time.Hour),
        )),
    )
    
    apiClient := client.NewAPIClientFromSettings(clientSettings)
    
    // Method 2: Using environment variables with overrides
    apiClient = client.NewAPIClientFromEnv(
        settings.WithServerURL("https://api.orkes.io/api"),
        settings.WithHTTPTimeout(120*time.Second),
        settings.WithProxyURL("http://corporate-proxy.com:8080"),
    )
    
    // Method 3: Manual settings construction
    authSettings := settings.NewAuthenticationSettings("key", "secret")
    httpSettings := settings.NewHttpSettings("https://api.orkes.io/api")
    proxySettings := settings.NewProxySettingsFromEnv()
    
    tokenExpiration := authentication.NewTokenExpiration(
        45*time.Minute,  // default expiration
        3*time.Hour,     // cleanup interval
    )
    tokenManager := authentication.NewTokenManager(authSettings, tokenExpiration)
    
    clientSettings = &settings.ClientSettings{
        Authentication:  authSettings,
        HTTP:           httpSettings,
        Proxy:          proxySettings,
        TokenExpiration: tokenExpiration,
        TokenManager:   tokenManager,
    }
    
    apiClient = client.NewAPIClientFromSettings(clientSettings)
    
    // Use the client...
    log.Println("API Client configured successfully")
}
```

## Related Documentation

- [Proxy Configuration](proxy_configuration.md) - Detailed proxy setup guide
