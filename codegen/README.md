# Code Generation for Conductor Go SDK

This directory contains scripts and tools for generating Go SDK code from the OpenAPI specification.

## Prerequisites

- Docker installed and running
- Python 3.x
- Internet connection (for Docker image download on first run)

## Usage

```bash
cd codegen
./generate.sh <client_version>
```

### Examples

```bash
# Generate Orkes Cloud client
./generate.sh orkes

# Generate Community version client  
./generate.sh conductor
```

## What it does

1. **Fixes OpenAPI specification:**
   - Removes `/api` prefix from all paths
   - Fixes path parameters
   - Resolves circular references
   - Fixes array schemas

2. **Generates Go SDK using Docker:**
   - Uses OpenAPI Generator v7.1.0
   - Generates models, APIs, and supporting files

3. **Output location:**
   ```
   sdk/generated/http/<client_version>/
   ├── api_*.go          # API client methods
   ├── model_*.go        # Data models/DTOs
   ├── client.go         # HTTP client
   ├── configuration.go  # Client configuration
   └── utils.go          # Utility functions
   ```
   
   For example:
   - `./generate.sh orkes` → `sdk/generated/http/orkes/`
   - `./generate.sh community` → `sdk/generated/http/conductor/`

## Configuration

Edit `generate.sh` to modify:
- `DOCKER_IMAGE`: OpenAPI Generator Docker image version (default: v7.1.0)
- `INPUT_SPEC`: Path to OpenAPI specification (default: `codegen/api-docs.json`)

The package name and output path are automatically determined from the client version parameter.

## Troubleshooting

- **Docker not running:** Start Docker Desktop
- **Permission denied:** Run `chmod +x generate.sh`
- **Python errors:** Ensure Python 3 is installed

## Notes

- The generated code is placed in `sdk/generated/http/<client_version>/`
- Client version is required and determines the output directory and package name
- Original spec must be in `codegen/api-docs.json`
- Temporary files are cleaned up automatically
- Multiple client versions can coexist in different directories
