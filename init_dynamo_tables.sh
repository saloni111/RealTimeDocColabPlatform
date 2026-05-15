#!/usr/bin/env bash
# Create DynamoDB Local tables needed by DocHub (safe to run multiple times).
set -euo pipefail

ENDPOINT="${DYNAMODB_ENDPOINT:-http://localhost:9000}"

if ! command -v aws >/dev/null 2>&1; then
  echo "Install AWS CLI, or create manually: docs (HASH document_id), doc_users (HASH user_id) at ${ENDPOINT}."
  exit 1
fi

aws dynamodb create-table \
  --endpoint-url "$ENDPOINT" \
  --region us-east-1 \
  --table-name docs \
  --attribute-definitions AttributeName=document_id,AttributeType=S \
  --key-schema AttributeName=document_id,KeyType=HASH \
  --billing-mode PAY_PER_REQUEST \
  >/dev/null 2>&1 || true

aws dynamodb create-table \
  --endpoint-url "$ENDPOINT" \
  --region us-east-1 \
  --table-name doc_users \
  --attribute-definitions AttributeName=user_id,AttributeType=S \
  --key-schema AttributeName=user_id,KeyType=HASH \
  --billing-mode PAY_PER_REQUEST \
  >/dev/null 2>&1 || true

echo "DynamoDB tables docs and doc_users ensured at ${ENDPOINT}."
