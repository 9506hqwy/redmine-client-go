#!/bin/bash
set -euo pipefail

BASE_DIR=$(dirname "$(dirname "$0")")
SOURCE_FILE="/tmp/openapi-orig.yml"
DEST_FILE="/tmp/openapi.yml"

curl -fsSL \
    -o "${SOURCE_FILE}" \
    "https://github.com/9506hqwy/openapi-spec-redmine/raw/refs/heads/main/openapi.yml"

# shellcheck disable=SC2016
cat "${SOURCE_FILE}" | \
    # description
    yq '.. |= select(has("parameters")) |= with(.parameters; map(select(has("description")) | .x-oapi-codegen-extra-tags.jsonschema = "description=" + .description | .))' | \
    yq '.. |= select(has("properties")) |= with(.properties; map_values(select(has("description")) | .x-oapi-codegen-extra-tags.jsonschema += "description=" + .description | .))' | \
    yq '.. |= select(has("x-oapi-codegen-extra-tags")) |= .x-oapi-codegen-extra-tags |= .jsonschema |= sub(",", "\\\\,")' | \
    yq '.. |= select(has("x-oapi-codegen-extra-tags")) |= .x-oapi-codegen-extra-tags |= .jsonschema |= trim' | \
    yq '.. |= select(has("x-oapi-codegen-extra-tags")) |= .x-oapi-codegen-extra-tags |= .jsonschema |= sub("\n", " ")' | \
    yq '.. |= select(has("x-oapi-codegen-extra-tags")) |= .x-oapi-codegen-extra-tags |= .jsonschema |= sub("\\\<", "<")' | \
    yq '.. |= select(has("x-oapi-codegen-extra-tags")) |= .x-oapi-codegen-extra-tags |= .jsonschema |= sub("\"", "\\\"")' | \
    yq '.. |= select(has("x-oapi-codegen-extra-tags")) |= .x-oapi-codegen-extra-tags |= .jsonschema |= sub("`", "\\\"")' | \
    # default
    yq '.. |= select(has("parameters")) |= .parameters |= map(. as $a | select(has("default")) | $a.x-oapi-codegen-extra-tags.jsonschema = ",default=" + $a.default | $a)' | \
    yq '.. |= select(has("properties")) |= .properties |= map_values(. as $a | select(has("default")) | $a.x-oapi-codegen-extra-tags.jsonschema += ",default=" + $a.default | $a)' | \
    # minLength
    yq '.. |= select(has("parameters")) |= .parameters |= map(. as $a | select(has("minLength")) | $a.x-oapi-codegen-extra-tags.jsonschema = ",minLength=" + $a.minLength | $a)' | \
    yq '.. |= select(has("properties")) |= .properties |= map_values(. as $a | select(has("minLength")) | $a.x-oapi-codegen-extra-tags.jsonschema += ",minLength=" + $a.minLength | $a)' | \
    # maxLengt
    yq '.. |= select(has("parameters")) |= .parameters |= map(. as $a | select(has("maxLengt")) | $a.x-oapi-codegen-extra-tags.jsonschema = ",maxLengt=" + $a.maxLengt | $a)' | \
    yq '.. |= select(has("properties")) |= .properties |= map_values(. as $a | select(has("maxLengt")) | $a.x-oapi-codegen-extra-tags.jsonschema += ",maxLengt=" + $a.maxLengt | $a)' | \
    # pattern
    yq '.. |= select(has("parameters")) |= .parameters |= map(. as $a | select(has("pattern")) | $a.x-oapi-codegen-extra-tags.jsonschema = ",pattern=" + $a.pattern | $a)' | \
    yq '.. |= select(has("properties")) |= .properties |= map_values(. as $a | select(has("pattern")) | $a.x-oapi-codegen-extra-tags.jsonschema += ",pattern=" + $a.pattern | $a)' | \
    # format
    yq '.. |= select(has("parameters")) |= .parameters |= map(. as $a | select(has("format")) | $a.x-oapi-codegen-extra-tags.jsonschema = ",format=" + $a.format | $a)' | \
    yq '.. |= select(has("properties")) |= .properties |= map_values(. as $a | select(has("format")) | $a.x-oapi-codegen-extra-tags.jsonschema += ",format=" + $a.format | $a)' | \
    # example
    yq '.. |= select(has("parameters")) |= .parameters |= map(. as $a | (.schema.examples[], .schema.items.examples[]) as $e ireduce({}; $a.x-oapi-codegen-extra-tags.jsonschema += ",example=" + $e) | $a)' | \
    yq '.. |= select(has("properties")) |= .properties |= map_values(. as $a | .examples[] as $e ireduce({}; $a.x-oapi-codegen-extra-tags.jsonschema += ",example=" + $e) | $a)' | \
    # enum
    yq '.. |= select(has("parameters")) |= .parameters |= map(. as $a | (.schema.enum[], .schema.items.enum[]) as $e ireduce({}; $a.x-oapi-codegen-extra-tags.jsonschema += ",enum=" + $e) | $a)' | \
    yq '.. |= select(has("properties")) |= .properties |= map_values(. as $a | .enum[] as $e ireduce({}; $a.x-oapi-codegen-extra-tags.jsonschema += ",enum=" + $e) | $a)' | \
    # minimum
    yq '.. |= select(has("parameters")) |= .parameters |= map(. as $a | select(has("minimum")) | $a.x-oapi-codegen-extra-tags.jsonschema = ",minimum=" + $a.minimum | $a)' | \
    yq '.. |= select(has("properties")) |= .properties |= map_values(. as $a | select(has("minimum")) | $a.x-oapi-codegen-extra-tags.jsonschema += ",minimum=" + $a.minimum | $a)' | \
    # maximum
    yq '.. |= select(has("parameters")) |= .parameters |= map(. as $a | select(has("maximum")) | $a.x-oapi-codegen-extra-tags.jsonschema = ",maximum=" + $a.maximum | $a)' | \
    yq '.. |= select(has("properties")) |= .properties |= map_values(. as $a | select(has("maximum")) | $a.x-oapi-codegen-extra-tags.jsonschema += ",maximum=" + $a.maximum | $a)' | \
    # exclusiveMaximum
    yq '.. |= select(has("parameters")) |= .parameters |= map(. as $a | select(has("exclusiveMaximum")) | $a.x-oapi-codegen-extra-tags.jsonschema = ",exclusiveMaximum=" + $a.exclusiveMaximum | $a)' | \
    yq '.. |= select(has("properties")) |= .properties |= map_values(. as $a | select(has("exclusiveMaximum")) | $a.x-oapi-codegen-extra-tags.jsonschema += ",exclusiveMaximum=" + $a.exclusiveMaximum | $a)' | \
    # exclusiveMinimum
    yq '.. |= select(has("parameters")) |= .parameters |= map(. as $a | select(has("exclusiveMinimum")) | $a.x-oapi-codegen-extra-tags.jsonschema = ",exclusiveMinimum=" + $a.exclusiveMinimum | $a)' | \
    yq '.. |= select(has("properties")) |= .properties |= map_values(. as $a | select(has("exclusiveMinimum")) | $a.x-oapi-codegen-extra-tags.jsonschema += ",exclusiveMinimum=" + $a.exclusiveMinimum | $a)' | \
    # minItems
    yq '.. |= select(has("parameters")) |= .parameters |= map(. as $a | select(has("minItems")) | $a.x-oapi-codegen-extra-tags.jsonschema = ",minItems=" + $a.minItems | $a)' | \
    yq '.. |= select(has("properties")) |= .properties |= map_values(. as $a | select(has("minItems")) | $a.x-oapi-codegen-extra-tags.jsonschema += ",minItems=" + $a.minItems | $a)' | \
    # maxItems
    yq '.. |= select(has("parameters")) |= .parameters |= map(. as $a | select(has("maxItems")) | $a.x-oapi-codegen-extra-tags.jsonschema = ",maxItems=" + $a.maxItems | $a)' | \
    yq '.. |= select(has("properties")) |= .properties |= map_values(. as $a | select(has("maxItems")) | $a.x-oapi-codegen-extra-tags.jsonschema += ",maxItems=" + $a.maxItems | $a)' | \
    # remove enum
    yq 'del(.. | select(. | key == "enum"))' \
    > "${DEST_FILE}"

go tool oapi-codegen -config "${BASE_DIR}/cfg.yml" "${DEST_FILE}"
