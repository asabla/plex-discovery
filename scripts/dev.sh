#!/usr/bin/env bash
set -euo pipefail

[ -d web/node_modules ] || npm --prefix web install

npm --prefix web run dev -- --host 0.0.0.0 --port 5173 &
VITE_PID=$!

cleanup() {
    if kill -0 "$VITE_PID" 2>/dev/null; then
        kill "$VITE_PID"
        wait "$VITE_PID" 2>/dev/null || true
    fi
}

trap cleanup EXIT

air
