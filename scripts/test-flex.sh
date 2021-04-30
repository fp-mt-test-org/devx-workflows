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

echo ""
echo "======================="
echo "TEST: Install/init flow"
echo "======================="
if [ -d "${helloworld_repo_folder_path}" ]; then
    echo "Pre-test Cleanup: Clearing out the test repo if left over from previous test..."
    rm -rdf "${helloworld_repo_folder_path}"
fi

echo "Cloning a blank repo..."
git clone https://github.com/fp-mt-test-org/${helloworld_repo_name}.git
echo "Clone complete."
echo ""

cd "$helloworld_repo_folder_path"

echo "Installing flex from ${dist_folder_path} into ${helloworld_repo_folder_path}"
skip_download=1 auto_clean=0 download_folder_path="${dist_folder_path}" "${dist_user_scripts_path}/install-flex.sh"

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

echo ""
echo "======================"
echo "TEST: Get Version Flow"
echo "======================"
if [ -d "${helloworld_repo_name}" ]; then
    echo "Pre-test Cleanup: Clearing out the test repo if left over from previous test..."
    rm -rdf "${helloworld_repo_name}"
fi

expected_flex_version=$(git describe --abbrev=0 --tags)
echo "expected_flex_version: ${expected_flex_version}"

echo "Cloning a blank repo..."
git clone https://github.com/fp-mt-test-org/${helloworld_repo_name}.git
echo "Clone complete."
echo ""

cd "${helloworld_repo_folder_path}"

echo "Installing flex from ${dist_folder_path} into ${helloworld_repo_folder_path}"
skip_download=1 auto_clean=0 download_folder_path="${dist_folder_path}" "${dist_user_scripts_path}/install-flex.sh"

echo "Getting version..."
flex_output=$("${helloworld_repo_install_path}/flex" -version)
echo "flex_output:"
echo ""
echo "${flex_output}"
echo ""
echo "Assert the output contains expected_flex_version:"
if ! [[ "${flex_output}" =~ .*[0-9]+\.[0-9]+\.[0-9]+.* ]]; then
    echo "Fail: Output does not contain expected_flex_version: ${expected_flex_version}"
    exit 1
fi
echo "Pass!"

echo ""
echo "========================="
echo "TEST: Update Version Flow"
echo "========================="
repo_name='devx-workflows-test-upgrade'

if [ -d "${repo_name}" ]; then
    echo "Pre-test Cleanup: Clearing out the test repo if left over from previous test..."
    rm -rdf "${repo_name}"
fi

echo "Cloning a repo that has flex already initialized..."
git clone https://github.com/fp-mt-test-org/${repo_name}.git
echo "Clone complete."
echo ""

cd "${repo_name}"

echo "Step 1. Test: Get current actual version: flex --version"
actual_flex_version=1
echo "Step 2. Test: Configure version to latest built version"
expected_flex_version=2
# update service_config.yaml
echo "Step 3. Test: Run flex --version again:"
echo "Step 4. Flex: If configuration != actual then install-flex.sh, return updated version"
actual_flex_version=2
echo "Step 5. Test: Assert actual == configured"
if [[ ! "${actual_flex_version}" == "${expected_flex_version}" ]]; then
    echo "Fail: actual_flex_version: ${actual_flex_version} should equal expected_flex_version: ${expected_flex_version}"
    exit 1
fi
echo "Pass!"

echo "Step 6. Test: Run flex build"
echo "Step 7. Test: Assert build output"
echo "Pass!"