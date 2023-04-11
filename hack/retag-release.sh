#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# This script re-tags a release branch with a stable version after the binaries have been committed.
# It is intended to be run by the CI system after a successful release.

# Usage: retag-release.sh <version>
# Example: retag-release.sh v1

echo "Re-tagging ${1}"

# Use -f as the tag will already exist
git tag -a -m "${1}" "${1}" -f
git push origin "${1}" -f
