#!/usr/bin/env bash
set -euox pipefail
IFS=$'\n\t'

# This script pushes a tagged artifact to a new release branch.
# It is intended to be run by the CI system after a successful build.

# Usage: commit-artifact.sh <version> <artifact-path>/<artifact-name>
# Example: commit-artifact.sh v0.1.0 dist/status-writer-action_v0.1.0-dev_darwin_amd64.tar.gz

echo "Committing artifact '${2}' to branch 'release-${1}'"

# Configure git to push to the current repository
git config user.name "${GITHUB_ACTOR}"
git config user.email "${GITHUB_ACTOR}@users.noreply.github.com"
git remote set-url origin "https://x-access-token:${GITHUB_TOKEN}@github.com/${GITHUB_REPOSITORY}.git"
git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

# The version is based on the tag that triggered the script
readonly version="release-${1}"
# Strip the path, version and extension from the filename
readonly filenamewithextension="${2##*/}"
readonly filename="${filenamewithextension%.tar.gz}"

# Create a new branch if it doesn't exist, otherwise just checkout
git checkout "${version}" 2>/dev/null || git checkout -b "${version}"
tar -xzf "${2}"

mkdir -p bin
readonly movedfilename="bin/${filename}"
mv status-writer-action "${movedfilename}"
git add "${movedfilename}"
git commit -m "Add ${filename}"
git push origin "${version}"
