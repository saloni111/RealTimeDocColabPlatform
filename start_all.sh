#!/bin/bash

set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$ROOT"

# Optional: DynamoDB Local on port 9000 (host maps to container 8000 in compose).
if docker info >/dev/null 2>&1; then
  if docker compose version >/dev/null 2>&1; then
    docker compose up -d dynamodb || true
  elif command -v docker-compose >/dev/null 2>&1; then
    docker-compose up -d dynamodb || true
  fi
  echo "DynamoDB Local should be reachable at http://localhost:9000 (run ./init_dynamo_tables.sh once if tables are missing)."
else
  echo "Docker not running: using existing DynamoDB endpoint (default http://localhost:9000)."
fi

sleep 2

# Start user-service
cd "${ROOT}/user-service"
nohup go run main.go >"${ROOT}/user-service.log" 2>&1 &
echo "Started user-service"

cd "${ROOT}/document-service"
nohup go run main.go >"${ROOT}/document-service.log" 2>&1 &
echo "Started document-service"

cd "${ROOT}/collaboration-service"
nohup go run main.go >"${ROOT}/collaboration-service.log" 2>&1 &
echo "Started collaboration-service"

cd "${ROOT}/api-gateway"
nohup go run main.go >"${ROOT}/api-gateway.log" 2>&1 &
echo "Started api-gateway"

cd "$ROOT"
echo "All services started. Check *.log files for output."
