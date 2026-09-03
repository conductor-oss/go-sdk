#!/usr/bin/env bash
#
# Spin up a local Conductor OSS stack and run the integration suite against
# it, similar to the `integration-tests-oss` job in
# .github/workflows/integration-tests-sm.yml -- except this runs plain `go
# test`, so there is no gotestsum JUnit report and no --rerun-fails=3 flake
# retry.
#
# Tests covering functionality plain OSS Conductor doesn't support are skipped
# via testdata.SkipIfOSS, and individual assertions via
# testdata.OSSGapSkippedReason; every gap and the evidence for it is listed in
# test/integration_tests/oss_gaps_test.go.
#
# The stack (Conductor OSS + Postgres + httpbin) is defined in
# scripts/docker-compose-oss.yaml and is torn down automatically on exit. The
# image is always pulled before starting, since `latest` (the local default)
# is a mutable tag and a cached copy would otherwise go stale silently.
#
# Usage:
#   scripts/run-integration-oss.sh [--keep-up] [--version <tag>] [--include-gated] [-- go test args]
# Examples:
#   scripts/run-integration-oss.sh                  # run against `latest`
#   scripts/run-integration-oss.sh --version 3.32.0-rc18
#   scripts/run-integration-oss.sh --keep-up        # leave the stack running afterwards
#   scripts/run-integration-oss.sh --include-gated  # also run tests normally skipped as Orkes-only
#   scripts/run-integration-oss.sh -- -run TestExecuteWorkflow
set -euo pipefail

KEEP_UP=0
INCLUDE_GATED=0
extra=()

while [[ $# -gt 0 ]]; do
  case "$1" in
    --keep-up) KEEP_UP=1; shift ;;
    --version) OSS_CONDUCTOR_VERSION="${2:?--version needs a tag}"; shift 2 ;;
    --include-gated) INCLUDE_GATED=1; shift ;;
    -h|--help)
      echo "Usage: $0 [--keep-up] [--version <tag>] [--include-gated] [-- go test args]"
      exit 0
      ;;
    --) shift; extra=("$@"); break ;;
    *) echo "Unknown argument: $1" >&2; exit 1 ;;
  esac
done

export OSS_CONDUCTOR_VERSION="${OSS_CONDUCTOR_VERSION:-latest}"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"
COMPOSE_FILE="${SCRIPT_DIR}/docker-compose-oss.yaml"
cd "${REPO_ROOT}"

compose() { docker compose -f "${COMPOSE_FILE}" "$@"; }

cleanup() {
  if [[ "${KEEP_UP}" == "1" ]]; then
    echo "--keep-up set: leaving the OSS stack running. Tear down with:"
    echo "  docker compose -f ${COMPOSE_FILE} down -v"
    return
  fi
  echo "Tearing down Conductor OSS stack..."
  compose down -v || true
}
trap cleanup EXIT

echo "Using conductoross/conductor:${OSS_CONDUCTOR_VERSION}"

# `docker compose up` only pulls an image when it is missing locally, so a
# previously-cached `latest` (or any other mutable tag) would silently be
# reused instead of getting the current version. Pull unconditionally so the
# stack always reflects the tag we just printed.
echo "Pulling conductoross/conductor:${OSS_CONDUCTOR_VERSION} to ensure it's current..."
compose pull conductor-server

echo "Starting Conductor OSS stack..."
compose up -d

echo "Waiting for Conductor to be healthy..."
HEALTH_TIMEOUT="${HEALTH_TIMEOUT:-180}"
deadline=$(( SECONDS + HEALTH_TIMEOUT ))
until curl -sf http://localhost:8080/health >/dev/null 2>&1; do
  if (( SECONDS >= deadline )); then
    echo "Error: Conductor did not become healthy within ${HEALTH_TIMEOUT}s." >&2
    compose logs conductor-server || true
    exit 1
  fi
  sleep 5
done
echo "Conductor is up."

export CONDUCTOR_SERVER_URL="http://localhost:8080/api"
# Must stay "oss" even with --include-gated: testdata.RequireAtLeast depends
# on it too, to bypass comparing OSS's own (unrelated, numerically lower)
# version numbering against Orkes Enterprise version thresholds. Unsetting it
# would cause nearly the entire suite to skip for an unrelated reason.
export CONDUCTOR_SERVER_TYPE="oss"
unset CONDUCTOR_AUTH_KEY CONDUCTOR_AUTH_SECRET

if [[ "${INCLUDE_GATED}" == "1" ]]; then
  echo "--include-gated set: tests normally skipped as Orkes-only will run against OSS too."
  export CONDUCTOR_INCLUDE_GATED_TESTS="true"
fi

go test -v -count=1 ./test/integration_tests/... "${extra[@]+"${extra[@]}"}"
