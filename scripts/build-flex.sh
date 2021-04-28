#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

echo "PWD: $(pwd)"

# Tag if in CI and on main branch
trunk_branch=$(git branch | grep -o -m1 "\b\(master\|main\)\b")
current_branch=$(git branch --show-current)

echo ""
echo "Trunk Branch:   ${trunk_branch}"
echo "Current Branch: ${current_branch}"
echo ""

if [[ -n "${CI:-}" ]]; then
    echo "Detected running in CI, checking branches for tagging..."

    if [[ "${trunk_branch}" == "${current_branch}" ]]; then
        echo "This build is running on the trunk branch ${trunk_branch}, tagging..."
        
        echo "Installing svu..."
        brew install caarlos0/tap/svu
        echo "Install complete."
        next_version="$(svu n)"
        echo "Tagging repo with ${next_version}"
        git tag "$(svu n)"
        echo "Pushing tags"
        git push --tags
        echo "Tagging complete."
        echo ""
    fi
fi

# Build and output the workflow binaries.
goreleaser --snapshot --skip-publish --rm-dist

current_path=$(realpath .)
echo "current_path: ${current_path}"

scripts_folder_name='scripts/user'
distribution_folder_name='dist'

distribution_folder_path="${current_path}/${distribution_folder_name}"

echo "distribution_folder_path: ${distribution_folder_path}"

user_scripts_source_path="${current_path}/${scripts_folder_name}"
user_scripts_dist_path="${distribution_folder_path}/${scripts_folder_name}"

# Copy user scripts
mkdir -p "${user_scripts_dist_path}"
cp -r "${user_scripts_source_path}/" "${user_scripts_dist_path}/"

echo "Build completed!"

if [ "${auto_install:=false}" == "true" ]; then
    echo "Auto installing..."
    skip_download=1 download_folder_path="${distribution_folder_path}" "${user_scripts_dist_path}/install-flex.sh"
fi
