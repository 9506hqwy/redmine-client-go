#!/bin/bash
set -euo pipefail

BASE_DIR=$(dirname "$(dirname "$0")")
SOURCE_FILE="/tmp/openapi.yml"
DEST_FILE="/tmp/openapi_slim.yml"

curl -fsSL \
    -o "${SOURCE_FILE}" \
    "https://github.com/9506hqwy/openapi-spec-redmine/raw/refs/heads/main/openapi.yml"

yq 'del(.. | select(. | key == "enum"))' "${SOURCE_FILE}" > "${DEST_FILE}"

go tool oapi-codegen -config "${BASE_DIR}/cfg.yml" "${DEST_FILE}"
