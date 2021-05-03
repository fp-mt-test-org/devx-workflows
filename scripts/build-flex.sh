#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

echo "Building Flex!"
echo ""

# Execute a release in "dry run" mode to create a local build that can be tested.
dry_run=1 ./scripts/release-flex.sh

current_path=$(realpath .)
scripts_folder_name='scripts/user'
distribution_folder_name='dist'
distribution_folder_path="${current_path}/${distribution_folder_name}"

echo ""
echo "current_path:             ${current_path}"
echo "distribution_folder_path: ${distribution_folder_path}"
echo ""

user_scripts_source_path="${current_path}/${scripts_folder_name}"
user_scripts_dist_path="${distribution_folder_path}/${scripts_folder_name}"

echo "Copy user scripts..."
mkdir -p "${user_scripts_dist_path}"
cp -vR "${user_scripts_source_path}/." "${user_scripts_dist_path}"
echo "Copy completed."
echo "Build completed!"

if [ "${auto_install:=0}" == "1" ]; then
    echo ""
    echo "Auto installing..."
    skip_download=1 download_folder_path="${distribution_folder_path}" "${user_scripts_dist_path}/install-flex.sh"
fi
