# TLS/SSL Configuration Guide

This guide explains how to configure TLS/SSL settings for the Conductor Go SDK, including support for self-signed certificates and mutual TLS authentication.

## Table of Contents

- [Quick Start](#quick-start)
- [Configuration Methods](#configuration-methods)
  - [Programmatic Configuration](#programmatic-configuration)
  - [Environment Variables](#environment-variables)
- [Common Use Cases](#common-use-cases)
- [Security Considerations](#security-considerations)
- [Troubleshooting](#troubleshooting)
- [API Reference](#api-reference)

## Quick Start

### Trust a Self-Signed Certificate

The most common use case - connecting to a server with a self-signed certificate:

```go
import (
    "github.com/conductor-sdk/conductor-go/sdk/client"
    "github.com/conductor-sdk/conductor-go/sdk/settings"
)

func main() {
    apiClient := client.NewAPIClient(
        settings.NewAuthenticationSettings("your-key", "your-secret"),
        settings.NewHttpSettings("https://conductor.internal.company.com/api"),
        settings.WithCACertFromFile("/etc/ssl/certs/company-ca.pem"),
    )
    
    // Use the client...
}
```

### Using Environment Variables

Set environment variables and let the SDK configure automatically:

```bash
export CONDUCTOR_SERVER_URL="https://conductor.internal.company.com/api"
export CONDUCTOR_AUTH_KEY="your-key"
export CONDUCTOR_AUTH_SECRET="your-secret"
export CONDUCTOR_TLS_CA_CERT="/etc/ssl/certs/company-ca.pem"
```

```go
func main() {
    // Automatically loads all settings from environment
    apiClient := client.NewAPIClientFromEnv()
    
    // Use the client...
}
```

## Configuration Methods

### Programmatic Configuration

#### 1. Disable Certificate Verification

**⚠️ WARNING**: Only use this for local development/testing!

```go
apiClient := client.NewAPIClient(
    authSettings,
    httpSettings,
    settings.WithInsecureSkipVerify(true),
)
```

#### 2. Trust Custom CA Certificate

**Recommended for production with self-signed certificates:**

```go
// From file
apiClient := client.NewAPIClient(
    authSettings,
    httpSettings,
    settings.WithCACertFromFile("/path/to/ca-cert.pem"),
)

// From PEM bytes (e.g., from secrets manager)
certData, _ := os.ReadFile("ca-cert.pem")
apiClient := client.NewAPIClient(
    authSettings,
    httpSettings,
    settings.WithCACertFromPEM(certData),
)
```

#### 3. Hybrid Mode (System + Custom CAs)

Trust both public CAs and your custom CA:

```go
apiClient := client.NewAPIClient(
    authSettings,
    httpSettings,
    settings.WithSelfSignedCert("/path/to/internal-ca.pem"),
)
```

This allows connecting to both:
- Public services with valid SSL certificates
- Internal services with self-signed certificates

#### 4. Mutual TLS (Client Certificates)

For servers requiring client authentication:

```go
apiClient := client.NewAPIClient(
    authSettings,
    httpSettings,
    settings.WithCACertFromFile("/path/to/server-ca.pem"),
    settings.WithClientCert(
        "/path/to/client-cert.pem",
        "/path/to/client-key.pem",
    ),
)
```


### Environment Variables

#### Conductor-Specific Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `CONDUCTOR_TLS_INSECURE_SKIP_VERIFY` | Disable certificate verification (⚠️ insecure) | `true` |
| `CONDUCTOR_TLS_CA_CERT` | Path to CA certificate (PEM) | `/etc/ssl/certs/ca.pem` |
| `CONDUCTOR_TLS_CLIENT_CERT` | Path to client certificate for mTLS | `/etc/ssl/certs/client.pem` |
| `CONDUCTOR_TLS_CLIENT_KEY` | Path to client private key for mTLS | `/etc/ssl/private/key.pem` |

#### Configuration Priority

1. Programmatic configuration (highest priority)
2. `CONDUCTOR_TLS_*` environment variables
3. System default certificates (lowest priority)

## Common Use Cases

### Use Case 1: Self-Signed Certificate

**Scenario**: Connecting to a server with a self-signed certificate that is not trusted by the system certificate store

**Option A: Skip Certificate Verification**

This option disables certificate verification entirely. The client will accept any certificate presented by the server.

```bash
export CONDUCTOR_SERVER_URL="https://localhost:8443/api"
export CONDUCTOR_TLS_INSECURE_SKIP_VERIFY=true
```

```go
apiClient := client.NewAPIClientFromEnv()
```

**Option B: Trust the Self-Signed Certificate**

This option explicitly trusts the self-signed certificate by providing it as a CA certificate. The client will verify the server's certificate against the provided CA.

```bash
# Generate self-signed cert (one-time)
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes

export CONDUCTOR_SERVER_URL="https://localhost:8443/api"
export CONDUCTOR_TLS_CA_CERT="./cert.pem"
```

```go
apiClient := client.NewAPIClientFromEnv()
```

### Use Case 2: Custom CA Certificate

**Scenario**: Connecting to a server whose certificate is signed by a custom Certificate Authority that is not in the system certificate store

```bash
export CONDUCTOR_SERVER_URL="https://conductor.internal.company.com/api"
export CONDUCTOR_AUTH_KEY="your-key"
export CONDUCTOR_AUTH_SECRET="your-secret"
export CONDUCTOR_TLS_CA_CERT="/etc/ssl/certs/corporate-ca.pem"
```

```go
func main() {
    apiClient := client.NewAPIClientFromEnv()
    
    // All TLS configuration is handled automatically
    // Client verifies server certificate against the provided CA
}
```

### Use Case 3: Secrets Manager Integration

**Scenario**: Load certificates from a secrets management service

```go
import (
    "github.com/conductor-sdk/conductor-go/sdk/client"
    "github.com/conductor-sdk/conductor-go/sdk/settings"
)

func main() {
    // Load CA certificate from secrets manager (for server verification)
    caCertPEM, err := getFromSecretsManager("conductor/ca-cert")
    if err != nil {
        log.Fatal(err)
    }
    
    // Create client with CA certificate loaded from memory
    apiClient := client.NewAPIClient(
        authSettings,
        httpSettings,
        settings.WithCACertFromPEM(caCertPEM), // ✅ Loads CA cert from memory
    )
}

func getFromSecretsManager(path string) ([]byte, error) {
    // Your secrets manager integration code
    return nil, nil
}
```

**For mTLS (client certificates) from secrets manager:**

```go
func mainWithMTLS() {
    // Get certificates from secrets manager
    caCertPEM, _ := getFromSecretsManager("conductor/ca-cert")
    clientCertPEM, _ := getFromSecretsManager("conductor/client-cert")
    clientKeyPEM, _ := getFromSecretsManager("conductor/client-key")
    
    // Simple and clean - all certificates loaded from memory
    apiClient := client.NewAPIClient(
        authSettings,
        httpSettings,
        settings.WithCACertFromPEM(caCertPEM),
        settings.WithClientCertFromPEM(clientCertPEM, clientKeyPEM),
    )
}
```

### Use Case 4: Mutual TLS (mTLS)

**Scenario**: Conductor server requires client certificate authentication

```bash
export CONDUCTOR_SERVER_URL="https://secure.conductor.com/api"
export CONDUCTOR_TLS_CA_CERT="/etc/ssl/certs/server-ca.pem"
export CONDUCTOR_TLS_CLIENT_CERT="/etc/ssl/certs/client.pem"
export CONDUCTOR_TLS_CLIENT_KEY="/etc/ssl/private/client-key.pem"
```

```go
// Automatically configured from environment
apiClient := client.NewAPIClientFromEnv()

// Or programmatically:
apiClient := client.NewAPIClient(
    settings.NewAuthenticationSettings("", ""), // May not need API keys with mTLS
    settings.NewHttpSettings("https://secure.conductor.com/api"),
    settings.WithCACertFromFile("/etc/ssl/certs/server-ca.pem"),
    settings.WithClientCert(
        "/etc/ssl/certs/client.pem",
        "/etc/ssl/private/client-key.pem",
    ),
)
```

## Security Considerations

### TLS Version

The SDK enforces **TLS 1.2** as the minimum version.

## Troubleshooting

### Common Errors

| Error | Solution |
|-------|----------|
| `certificate signed by unknown authority` | Use `WithCACertFromFile()` or set `CONDUCTOR_TLS_CA_CERT` |
| `certificate is valid for X, not Y` | Use `WithTLSServerName()` to override hostname |
| `certificate has expired` | Renew certificate or use `WithInsecureSkipVerify` (testing only) |
| Client certificate not sent | Set both `CONDUCTOR_TLS_CLIENT_CERT` and `CONDUCTOR_TLS_CLIENT_KEY` |

### Testing TLS Configuration

Test your TLS setup with `openssl`:

```bash
# Test server certificate
openssl s_client -connect conductor.example.com:443 -showcerts

# Test with custom CA
openssl s_client -connect conductor.example.com:443 -CAfile /path/to/ca.pem
```

## API Reference

### Available Options

| Option | Description |
|--------|-------------|
| `WithCACertFromFile(path)` | Trust custom CA from file |
| `WithCACertFromPEM(pemCert)` | Trust custom CA from PEM bytes |
| `WithSelfSignedCert(path)` | Trust custom CA + system CAs |
| `WithClientCert(certPath, keyPath)` | Load client cert from files (mTLS) |
| `WithClientCertFromPEM(certPEM, keyPEM)` | Load client cert from memory (mTLS) |
| `WithTLSServerName(serverName)` | Override server name for SNI |
| `WithInsecureSkipVerify(skip)` | ⚠️ Disable cert verification (insecure) |
| `WithTLSSettings(tlsSettings)` | Set complete TLS settings |

