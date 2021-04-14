#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

devx_workflow_scripts_folder=".devx-workflows"
os="darwin" && [[ -n "${CI:-}" ]] && os="linux"

# Build and output the workflow binaries.
goreleaser --snapshot --skip-publish --rm-dist
mkdir -p ${devx_workflow_scripts_folder} && tar -xzvf dist/$(ls dist | grep -e "SNAPSHOT.*${os}") -C ${devx_workflow_scripts_folder}
