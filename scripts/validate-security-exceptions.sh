#!/usr/bin/env bash

set -euo pipefail

script_dir=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
repo_root=$(cd -- "${script_dir}/.." && pwd)
exceptions_file="${repo_root}/.github/security-exceptions.json"

if [[ ! -e "$exceptions_file" ]]; then
  echo "no security exceptions file configured"
  exit 0
fi

command -v jq >/dev/null 2>&1 || {
  echo "jq is required to validate security exceptions" >&2
  exit 1
}

today=$(date -u +%F)
jq --arg today "$today" -e '
  .version == 1 and
  (.exceptions | type == "array") and
  all(.exceptions[];
    type == "object" and
    (.scanner | type == "string" and test("^[a-z0-9-]+$")) and
    (.id | type == "string" and length > 0 and (contains("*") | not)) and
    (.component | type == "string" and length > 0 and (contains("*") | not)) and
    (.reason | type == "string" and length > 0) and
    (.owner | type == "string" and length > 0) and
    (.issue | type == "string" and test("^https://github\\.com/")) and
    (.expires_on | type == "string" and test("^[0-9]{4}-[0-9]{2}-[0-9]{2}$") and . >= $today)
  )
' "$exceptions_file" >/dev/null || {
  echo "security exceptions are invalid, expired, or contain wildcards" >&2
  exit 1
}

echo "security exceptions are valid"
