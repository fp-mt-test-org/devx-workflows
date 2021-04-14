#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

devx_scripts_dir=".devx-workflows"
helloworld_dir="devx-workflows-test-empty-repo"
build_cmd="hello"

trap "rm -rf ${helloworld_dir} ../${helloworld_dir}" ERR EXIT

echo "--- Cloning helloworld service repository"
git clone https://github.com/fp-mt-test-org/${helloworld_dir}.git
cp -R ${devx_scripts_dir} ${helloworld_dir}
cd ${helloworld_dir}

echo "--- Testing init script"
{ echo "echo hello"; sleep 1; echo "done"; sleep 1; echo "helloworld-service"; } | ./${devx_scripts_dir}/flex init

echo "--- Running build command"
if [ $(./${devx_scripts_dir}/flex build) != "hello" ]; then
    echo "Build command was not initialized correctly!"
    exit 1
fi
echo "Build command successful!"
