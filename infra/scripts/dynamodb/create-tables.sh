#!/usr/bin/env bash
set -euo pipefail

TABLE_NAME=${TABLE_NAME:-jobwatch_main}
ENDPOINT_URL=${DDB_ENDPOINT_URL:-http://localhost:18000}

aws dynamodb create-table \
    --table-name "$TABLE_NAME" \
    --attribute-definitions AttributeName=pk,AttributeType=S \
    --key-schema AttributeName=pk,KeyType=HASH \
    --billing-mode PAY_PER_REQUEST \
    --endpoint-url "$ENDPOINT_URL"
