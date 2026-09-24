#!/usr/bin/env bash
# Run every agent example against a playback-enabled Conductor server built
# from a conductor-oss/conductor checkout (branch main).
#
#   scripts/run-agents-playback.sh /path/to/conductor
#
# The server must already be running; .github/workflows/agents-playback.yml
# starts it with the checkout's start-playback action. The HTTP and MCP
# fixtures are started here unless CONDUCTOR_PLAYBACK_SERVICES_STARTED=true.
set -euo pipefail
repo_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
conductor_dir=$(cd "${1:?usage: run-agents-playback.sh CONDUCTOR_CHECKOUT}" && pwd)
cd "$repo_dir"
export CONDUCTOR_SERVER_URL=${CONDUCTOR_SERVER_URL:-http://localhost:8080/api}
export CONDUCTOR_AGENT_LLM_MODEL=mock/mockLLM
export CONDUCTOR_AGENTS_PLAYBACK=true
export GITHUB_REPOS_URL='http://localhost:3002/users/Conductor/repos?per_page=5&sort=updated'
services_script="$conductor_dir/.github/actions/start-playback-services/start-services.sh"
mkdir -p tmp
playback_dir=${CONDUCTOR_PLAYBACK_WORK_DIR:-$(mktemp -d "$repo_dir/tmp/agent-playback.XXXXXX")}
mkdir -p "$playback_dir"
playback_dir=$(cd "$playback_dir" && pwd)
[[ -d "$conductor_dir/llm-recordings" ]]
curl --fail --silent --show-error --max-time 10 "${CONDUCTOR_SERVER_URL%/api}/health" > /dev/null

pids=()
cleanup() {
  for pid in "${pids[@]}"; do kill "$pid" 2>/dev/null || true; done
  for pid in "${pids[@]}"; do wait "$pid" 2>/dev/null || true; done
}
trap cleanup EXIT
trap 'exit 130' INT
trap 'exit 143' TERM

if [[ "${CONDUCTOR_PLAYBACK_SERVICES_STARTED:-false}" == true ]]; then
  bash "$services_script" --check
else
  CONDUCTOR_PLAYBACK_WORK_DIR="$playback_dir" bash "$services_script"
  pids+=("$(cat "$playback_dir/http.pid")" "$(cat "$playback_dir/mcp.pid")")
fi

cd examples
go build ./agents/...
go run ./agents/cmd external-workers > "$playback_dir/workers.log" 2>&1 &
pids+=("$!")

echo "Running every agent example against $CONDUCTOR_SERVER_URL"
echo "Logs: $playback_dir"
go test -count=1 -v -timeout 60m -run TestExamplesPlayback ./agents/ 2>&1 | tee "$playback_dir/tests.log"
