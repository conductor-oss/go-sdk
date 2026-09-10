#!/usr/bin/env bash
#
# Run a local Conductor server for the ai_e2e tests, with its LLM recorder in
# record or replay mode. See ../README.md, "Recording and replaying".
#
#   conductor-server.sh record <recordings-dir>   real LLM; every response is saved
#   conductor-server.sh replay <recordings-dir>   no LLM; responses come from the dir
#   conductor-server.sh plain                     neither; an ordinary server
#   conductor-server.sh stop
#
# CONDUCTOR_SERVER_JAR must point at a conductor-server boot jar built from a
# branch that has the recorder (conductor.ai.record-mode). Record mode also
# needs the provider key, e.g. OPENAI_API_KEY, in this shell: the server reads
# it from its own environment.
set -euo pipefail

port="${CONDUCTOR_SERVER_PORT:-8080}"
state="${TMPDIR:-/tmp}/conductor-e2e-server"
mode="${1:-}"

if [[ "$mode" == "stop" ]]; then
  if [[ -f "$state/pid" ]] && kill -0 "$(cat "$state/pid")" 2>/dev/null; then
    kill "$(cat "$state/pid")"
    echo "stopped server pid $(cat "$state/pid")"
  else
    echo "no server running from $state"
  fi
  rm -f "$state/pid"
  exit 0
fi

jar="${CONDUCTOR_SERVER_JAR:?set CONDUCTOR_SERVER_JAR to a conductor-server boot jar with the LLM recorder}"
props=(--server.port="$port" --conductor.integrations.ai.enabled=true)
case "$mode" in
  record)
    dir="${2:?record needs a recordings directory}"
    [[ -n "${OPENAI_API_KEY:-}" ]] || echo "warning: OPENAI_API_KEY is not set; the server cannot reach the provider" >&2
    props+=(--conductor.ai.record-mode=true --conductor.ai.enable-llm-mocks=false)
    ;;
  replay)
    dir="${2:?replay needs a recordings directory}"
    props+=(--conductor.ai.record-mode=false --conductor.ai.enable-llm-mocks=true)
    ;;
  plain)
    dir=""
    ;;
  *)
    echo "usage: $0 record|replay <recordings-dir> | plain | stop" >&2
    exit 2
    ;;
esac
if [[ -n "$dir" ]]; then
  mkdir -p "$dir"
  dir="$(cd "$dir" && pwd)"
  props+=(--conductor.ai.recordings-directory="$dir")
fi

if lsof -iTCP:"$port" -sTCP:LISTEN >/dev/null 2>&1; then
  echo "port $port is already in use; stop that server first (two local servers share a datastore)" >&2
  exit 1
fi

mkdir -p "$state"
# Run from the state dir: the OSS server writes its SQLite datastore (c123.db)
# into the working directory, and that must not land in the repository.
(cd "$state" && exec nohup java -jar "$jar" "${props[@]}" >"$state/server.log" 2>&1) &
pid=$!
echo "$pid" >"$state/pid"
echo "starting conductor ($mode) pid $pid, log $state/server.log"

for _ in $(seq 1 120); do
  if ! kill -0 "$pid" 2>/dev/null; then
    echo "server exited during startup; last log lines:" >&2
    tail -20 "$state/server.log" >&2
    exit 1
  fi
  if curl -sf "http://localhost:$port/health" >/dev/null 2>&1; then
    echo "ready at http://localhost:$port/api${dir:+, recordings in $dir}"
    exit 0
  fi
  sleep 1
done
echo "server did not become healthy in 120s; see $state/server.log" >&2
exit 1
