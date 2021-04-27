#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

current_path=$(realpath .)
dist_folder_name='dist'
dist_folder_path="${current_path}/${dist_folder_name}"
dist_user_scripts_path="${dist_folder_path}/scripts/user"

install_folder_name='.devx-workflows'
helloworld_repo_name='devx-workflows-test-empty-repo'
helloworld_repo_folder_path="${current_path}/${helloworld_repo_name}"
helloworld_repo_install_path="${helloworld_repo_folder_path}/${install_folder_name}"

if [ -d "${helloworld_repo_folder_path}" ]; then
    echo "Pre-test Cleanup: Clearing out the test repo if left over from previous test..."
    rm -rdf "${helloworld_repo_folder_path}"
fi

echo ""
echo "TEST: Install/init flow"
echo "Cloning a blank repo..."
git clone https://github.com/fp-mt-test-org/${helloworld_repo_name}.git
echo "Clone complete."
echo ""

cd "$helloworld_repo_folder_path"

echo "Installing flex from ${dist_folder_path} into ${helloworld_repo_folder_path}"
skip_download=1 download_folder_path="${dist_folder_path}" "${dist_user_scripts_path}/install-flex.sh"

build_cmd="hello"

echo "Executing init workflow..."
{ echo "echo ${build_cmd}"; sleep 1; echo "done"; sleep 1; echo "helloworld-service"; } | "${helloworld_repo_install_path}/flex" init
echo "Init complete, executing build..."

build_output=$("${helloworld_repo_install_path}/flex" build)

echo "-- Build Output --"
echo "${build_output}"
echo "-- End Build Output --"

echo "Assert build output is as expected:"
if [ "${build_output}" != "${build_cmd}" ]; then
    echo "Fail: Build output doesn't contain ${build_cmd}"
    exit 1
fi
echo "Pass!"

gitignore_file_name='.gitignore'
gitignore_file_path="${helloworld_repo_folder_path}/${gitignore_file_name}"
echo "Assert ${gitignore_file_name} updated:"
if ! grep -q "${install_folder_name}" "${gitignore_file_path}"; then
    echo "Fail: ${install_folder_name} is missing from ${gitignore_file_path}"
    exit 1
fi
echo "Pass!"

cd ..

if [ -d "${helloworld_repo_folder_path}" ]; then
    echo "Post-test Cleanup: Clearing out the test repo if left over from previous test..."
    rm -rdf "${helloworld_repo_folder_path}"
fi
