#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

devx_scripts_dir=".devx-workflows"
helloworld_dir="devx-workflows-test-empty-repo"
build_cmd="hello"

if [ -d "${helloworld_dir}" ]; then
    echo "Pre-test Cleanup: Clearing out the test repo if left over from previous test..."
    rm -rdf "${helloworld_dir}"
fi

echo ""
echo "TEST: Install/init workflow"
git clone https://github.com/fp-mt-test-org/${helloworld_dir}.git
echo "Clone complete, installing flex from ${devx_scripts_dir} to ${helloworld_dir}"
cp -R ${devx_scripts_dir} ${helloworld_dir}
cd ${helloworld_dir}
echo "Install complete, executing init..."
{ echo "echo ${build_cmd}"; sleep 1; echo "done"; sleep 1; echo "helloworld-service"; } | ./${devx_scripts_dir}/flex init
echo "Init complete, executing build..."
build_output="$(./${devx_scripts_dir}/flex build)"

echo "-- Build Output --"
echo "${build_output}"
echo "-- End Build Output --"

if [ "${build_output}" != "${build_cmd}" ]; then
    echo "Build command was not initialized correctly!"
    exit 1
fi
echo "Build command successful!"

cd ..

if [ -d "${helloworld_dir}" ]; then
    echo "Post-test Cleanup: Clearing out the test repo if left over from previous test..."
    rm -rdf "${helloworld_dir}"
fi
