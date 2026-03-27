# Kubernetes Manifests

This directory contains Kubernetes manifests for deploying the Go SDK harness worker to the certification clusters.

## Prerequisites

**Set your namespace environment variable:**
```bash
export NS=your-namespace-here
```

All kubectl commands below use `-n $NS` to specify the namespace. The manifests intentionally do not include hardcoded namespaces.

**Ensure the GHCR image pull secret exists in the namespace** (see step 1 below).

## Files

| File | Description |
|---|---|
| `deployment.yaml` | Deployment (single file, works on all clusters) |
| `configmap-aws.yaml` | Conductor URL + auth key for certification-aws |
| `configmap-azure.yaml` | Conductor URL + auth key for certification-az |
| `configmap-gcp.yaml` | Conductor URL + auth key for certification-gcp |
| `secret-conductor.yaml` | Conductor auth secret (placeholder template) |
| `secret-ghcr-pull.yaml` | GHCR image pull secret (instructions only) |

## Quick Start

### 1. Create the Image Pull Secret

The deployment pulls from a private GHCR package. The orkes-io org's `regcredorkesgit` secret does not have access to this OSS repo's packages, so a dedicated pull secret is needed.

```bash
kubectl create secret docker-registry ghcr-oss-sdk \
  --docker-server=ghcr.io \
  --docker-username=YOUR_GITHUB_USERNAME \
  --docker-password=YOUR_GITHUB_PAT \
  -n $NS
```

The PAT must be a GitHub PAT (classic) with `read:packages` scope. Verify it works first with `docker login ghcr.io`.

See `secret-ghcr-pull.yaml` for more details.

### 2. Create the Conductor Auth Secret

The `CONDUCTOR_AUTH_SECRET` must be created as a Kubernetes secret before deploying.

```bash
kubectl create secret generic conductor-credentials \
  --from-literal=auth-secret=YOUR_AUTH_SECRET \
  -n $NS
```

If the `conductor-credentials` secret already exists in the namespace (e.g. from the e2e-testrunner-worker), it can be reused as-is.

See `secret-conductor.yaml` for more details.

### 3. Apply the ConfigMap for Your Cluster

```bash
# AWS
kubectl apply -f manifests/configmap-aws.yaml -n $NS

# Azure
kubectl apply -f manifests/configmap-azure.yaml -n $NS

# GCP
kubectl apply -f manifests/configmap-gcp.yaml -n $NS
```

### 4. Deploy

```bash
kubectl apply -f manifests/deployment.yaml -n $NS
```

### 5. Verify

```bash
# Check pod status
kubectl get pods -n $NS -l app=go-sdk-harness-worker

# Watch logs
kubectl logs -n $NS -l app=go-sdk-harness-worker -f
```

## Building and Pushing the Image

From the repository root:

```bash
# Build the harness target and push to GHCR (multiplatform)
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  --target harness \
  -t ghcr.io/conductor-oss/go-sdk/harness-worker:latest \
  --push .
```

After pushing a new image with the same tag, restart the deployment to pull it:

```bash
kubectl rollout restart deployment/go-sdk-harness-worker -n $NS
kubectl rollout status deployment/go-sdk-harness-worker -n $NS
```

## Tuning

The harness worker accepts these optional environment variables (set in `deployment.yaml`):

| Variable | Default | Description |
|---|---|---|
| `HARNESS_WORKFLOWS_PER_SEC` | 2 | Workflows to start per second |
| `HARNESS_BATCH_SIZE` | 20 | Tasks each worker polls per batch |
| `HARNESS_POLL_INTERVAL_MS` | 100 | Milliseconds between poll cycles |

Edit `deployment.yaml` to change these, then re-apply:

```bash
kubectl apply -f manifests/deployment.yaml -n $NS
```

## Troubleshooting

### Pod not starting

```bash
kubectl describe pod -n $NS -l app=go-sdk-harness-worker
kubectl logs -n $NS -l app=go-sdk-harness-worker --tail=100
```

### Secret not found

```bash
kubectl get secret conductor-credentials -n $NS
```

### Image pull errors

Verify the `ghcr-oss-sdk` image pull secret exists and is correctly formatted:

```bash
# Check it exists
kubectl get secret ghcr-oss-sdk -n $NS

# Verify type is kubernetes.io/dockerconfigjson
kubectl get secret ghcr-oss-sdk -n $NS -o jsonpath='{.type}'

# Decode and inspect the docker config
kubectl get secret ghcr-oss-sdk -n $NS \
  -o jsonpath='{.data.\.dockerconfigjson}' | base64 -d
```

## Resource Limits

Default resource allocation:
- **Memory**: 256Mi (request) / 512Mi (limit)
- **CPU**: 100m (request) / 500m (limit)

Adjust in `deployment.yaml` based on workload. Higher `HARNESS_WORKFLOWS_PER_SEC` values may need more CPU/memory.

## Service

The harness worker does **not** need a Service or Ingress. It connects to Conductor via outbound HTTP polling. All communication is outbound.
