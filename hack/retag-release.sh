#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# This script re-tags a release branch after the binaries have been committed.
# It is intended to be run by the CI system after a successful release

# Usage: retag-release.sh <version> <number of commits back from HEAD>

echo "Re-tagging ${1}"

# Use -f as the tag will already exist
git tag -a -m "${1}" "${1}" -f
git push origin "${1}" -f

# Tag on main so the next release can calculate release notes.
git tag -a -m "Will be used by the next release to calculate the release notes" "${1}-rn" HEAD~"${2}"
git push origin "${1}-rn"
