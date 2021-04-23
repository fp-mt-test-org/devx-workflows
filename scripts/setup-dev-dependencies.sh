#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

./scripts/user/configure-localhost.sh

# Local install of goreleaser for build script
brew install goreleaser
