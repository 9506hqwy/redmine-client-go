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
    yq '.. |= select(has("parameters")) |= with(.parameters; . as $a | select(tag == "!!seq") | map(select(has("description")) | .x-oapi-codegen-extra-tags.jsonschema = "description=" + .description) | $a)' | \
    yq '.. |= select(has("properties")) |= with(.properties; . as $a | select(tag == "!!map") | map_values(select(has("description")) | .x-oapi-codegen-extra-tags.jsonschema = "description=" + .description) | $a)' | \
    yq '.. |= select(has("x-oapi-codegen-extra-tags")) |= .x-oapi-codegen-extra-tags |= .jsonschema |= sub("\s+", " ")' | \
    yq '.. |= select(has("x-oapi-codegen-extra-tags")) |= .x-oapi-codegen-extra-tags |= .jsonschema |= trim' | \
    yq '.. |= select(has("x-oapi-codegen-extra-tags")) |= .x-oapi-codegen-extra-tags |= .jsonschema |= sub(",", "\\\\,")' | \
    yq '.. |= select(has("x-oapi-codegen-extra-tags")) |= .x-oapi-codegen-extra-tags |= .jsonschema |= sub("\\\\<", "<")' | \
    yq '.. |= select(has("x-oapi-codegen-extra-tags")) |= .x-oapi-codegen-extra-tags |= .jsonschema |= sub("\\\\ ", " ")' | \
    yq '.. |= select(has("x-oapi-codegen-extra-tags")) |= .x-oapi-codegen-extra-tags |= .jsonschema |= sub("\"", "\\\"")' | \
    yq '.. |= select(has("x-oapi-codegen-extra-tags")) |= .x-oapi-codegen-extra-tags |= .jsonschema |= sub("`", "\\\"")' | \
    # default
    yq '.. |= select(has("parameters")) |= with(.parameters; . as $a | select(tag == "!!seq") | map(select(.schema.default != null) | .x-oapi-codegen-extra-tags.jsonschema += ",default=" + .schema.default) | $a)' | \
    yq '.. |= select(has("properties")) |= with(.properties; . as $a | select(tag == "!!map") | map_values(select(has("default")) | .x-oapi-codegen-extra-tags.jsonschema += ",default=" + .default) | $a)' | \
    # minLength
    yq '.. |= select(has("parameters")) |= with(.parameters; . as $a | select(tag == "!!seq") | map(select(.schema.minLength != null) | .x-oapi-codegen-extra-tags.jsonschema += ",minLength=" + .schema.minLength) | $a)' | \
    yq '.. |= select(has("properties")) |= with(.properties; . as $a | select(tag == "!!map") | map_values(select(has("minLength")) | .x-oapi-codegen-extra-tags.jsonschema += ",minLength=" + .minLength) | $a)' | \
    # maxLengt
    yq '.. |= select(has("parameters")) |= with(.parameters; . as $a | select(tag == "!!seq") | map(select(.schema.maxLengt != null) | .x-oapi-codegen-extra-tags.jsonschema += ",maxLengt=" + .schema.maxLengt) | $a)' | \
    yq '.. |= select(has("properties")) |= with(.properties; . as $a | select(tag == "!!map") | map_values(select(has("maxLengt")) | .x-oapi-codegen-extra-tags.jsonschema += ",maxLengt=" + .maxLengt) | $a)' | \
    # pattern
    yq '.. |= select(has("parameters")) |= with(.parameters; . as $a | select(tag == "!!seq") | map(select(.schema.pattern != null) | .x-oapi-codegen-extra-tags.jsonschema += ",pattern=" + .schema.pattern) | $a)' | \
    yq '.. |= select(has("properties")) |= with(.properties; . as $a | select(tag == "!!map") | map_values(select(has("pattern")) | .x-oapi-codegen-extra-tags.jsonschema += ",pattern=" + .pattern) | $a)' | \
    # format
    yq '.. |= select(has("parameters")) |= with(.parameters; . as $a | select(tag == "!!seq") | map(select(.schema.format != null) | .x-oapi-codegen-extra-tags.jsonschema += ",format=" + .schema.format) | $a)' | \
    yq '.. |= select(has("properties")) |= with(.properties; . as $a | select(tag == "!!map") | map_values(select(has("format")) | .x-oapi-codegen-extra-tags.jsonschema += ",format=" + .format) | $a)' | \
    # example
    yq '.. |= select(has("parameters")) |= .parameters |= select(tag == "!!seq") |= map(. as $a | select(.schema.examples != null) | .schema.examples[] as $e ireduce({}; $a.x-oapi-codegen-extra-tags.jsonschema += ",example=" + $e) | $a)' | \
    yq '.. |= select(has("parameters")) |= .parameters |= select(tag == "!!seq") |= map(. as $a | select(.schema.items.examples != null) | .schema.items.examples[] as $e ireduce({}; $a.x-oapi-codegen-extra-tags.jsonschema += ",example=" + $e) | $a)' | \
    yq '.. |= select(has("properties")) |= .properties |= select(tag == "!!map") |= map_values(. as $a | select(has("examples")) | .examples[] as $e ireduce({}; $a.x-oapi-codegen-extra-tags.jsonschema += ",example=" + $e) | $a)' | \
    # enum
    yq '.. |= select(has("parameters")) |= .parameters |= select(tag == "!!seq") |= map(. as $a | select(.schema.enum != null) | .schema.enum[] as $e ireduce({}; $a.x-oapi-codegen-extra-tags.jsonschema += ",enum=" + $e) | $a)' | \
    yq '.. |= select(has("parameters")) |= .parameters |= select(tag == "!!seq") |= map(. as $a | select(.schema.items.enum != null) | .schema.items.enum[] as $e ireduce({}; $a.x-oapi-codegen-extra-tags.jsonschema += ",enum=" + $e) | $a)' | \
    yq '.. |= select(has("properties")) |= .properties |= select(tag == "!!map") |= map_values(. as $a | select(has("enum")) | .enum[] as $e ireduce({}; $a.x-oapi-codegen-extra-tags.jsonschema += ",enum=" + $e) | $a)' | \
    # minimum
    yq '.. |= select(has("parameters")) |= with(.parameters; . as $a | select(tag == "!!seq") | map(select(.schema.minimum != null) | .x-oapi-codegen-extra-tags.jsonschema += ",minimum=" + .schema.minimum) | $a)' | \
    yq '.. |= select(has("properties")) |= with(.properties; . as $a | select(tag == "!!map") | map_values(select(has("minimum")) | .x-oapi-codegen-extra-tags.jsonschema += ",minimum=" + .minimum) | $a)' | \
    # maximum
    yq '.. |= select(has("parameters")) |= with(.parameters; . as $a | select(tag == "!!seq") | map(select(.schema.maximum != null) | .x-oapi-codegen-extra-tags.jsonschema += ",maximum=" + .schema.maximum) | $a)' | \
    yq '.. |= select(has("properties")) |= with(.properties; . as $a | select(tag == "!!map") | map_values(select(has("maximum")) | .x-oapi-codegen-extra-tags.jsonschema += ",maximum=" + .maximum) | $a)' | \
    # exclusiveMaximum
    yq '.. |= select(has("parameters")) |= with(.parameters; . as $a | select(tag == "!!seq") | map(select(.schema.exclusiveMaximum != null) | .x-oapi-codegen-extra-tags.jsonschema += ",exclusiveMaximum=" + .schema.exclusiveMaximum) | $a)' | \
    yq '.. |= select(has("properties")) |= with(.properties; . as $a | select(tag == "!!map") | map_values(select(has("exclusiveMaximum")) | .x-oapi-codegen-extra-tags.jsonschema += ",exclusiveMaximum=" + .exclusiveMaximum) | $a)' | \
    # exclusiveMinimum
    yq '.. |= select(has("parameters")) |= with(.parameters; . as $a | select(tag == "!!seq") | map(select(.schema.exclusiveMinimum != null) | .x-oapi-codegen-extra-tags.jsonschema += ",exclusiveMinimum=" + .schema.exclusiveMinimum) | $a)' | \
    yq '.. |= select(has("properties")) |= with(.properties; . as $a | select(tag == "!!map") | map_values(select(has("exclusiveMinimum")) | .x-oapi-codegen-extra-tags.jsonschema += ",exclusiveMinimum=" + .exclusiveMinimum) | $a)' | \
    # minItems
    yq '.. |= select(has("parameters")) |= with(.parameters; . as $a | select(tag == "!!seq") | map(select(.schema.minItems != null) | .x-oapi-codegen-extra-tags.jsonschema += ",minItems=" + .schema.minItems) | $a)' | \
    yq '.. |= select(has("properties")) |= with(.properties; . as $a | select(tag == "!!map") | map_values(select(has("minItems")) | .x-oapi-codegen-extra-tags.jsonschema += ",minItems=" + .minItems) | $a)' | \
    # maxItems
    yq '.. |= select(has("parameters")) |= with(.parameters; . as $a | select(tag == "!!seq") | map(select(.schema.maxItems != null) | .x-oapi-codegen-extra-tags.jsonschema += ",maxItems=" + .schema.maxItems) | $a)' | \
    yq '.. |= select(has("properties")) |= with(.properties; . as $a | select(tag == "!!map") | map_values(select(has("maxItems")) | .x-oapi-codegen-extra-tags.jsonschema += ",maxItems=" + .maxItems) | $a)' | \
    # remove enum
    yq 'del(.. | select(. | key == "enum"))' \
    > "${DEST_FILE}"

go tool oapi-codegen -config "${BASE_DIR}/cfg.yml" "${DEST_FILE}"
