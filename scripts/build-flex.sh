#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

# Tag if in CI and on main branch
trunk_branch=$(git branch | grep -o -m1 "\b\(master\|main\)\b")
current_branch=$(git branch --show-current)

echo "Trunk Branch: ${trunk_branch}"
echo "Current Branch: ${current_branch}"

if [[ -n "${CI:-}" ]]; then
    echo "Detected running in CI, checking braches for tagging..."
    git tag

    if [[ "${trunk_branch}" == ${current_branch} ]]; then
        brew install caarlos0/tap/svu
        git tag $(svu n)
        git push --tags
    fi
fi

devx_workflow_scripts_folder=".devx-workflows"
os="darwin" && [[ -n "${CI:-}" ]] && os="linux"

# Build and output the workflow binaries.
goreleaser --snapshot --skip-publish --rm-dist

mkdir -p ${devx_workflow_scripts_folder}
tar -xzvf dist/$(ls dist | grep -e "SNAPSHOT.*${os}") -C ${devx_workflow_scripts_folder}

