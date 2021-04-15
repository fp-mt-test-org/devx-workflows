#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

target_dir='.devx-workflows'
os=$(uname | tr '[:upper:]' '[:lower:]')
url="https://github.com/fp-mt-test-org/devx-workflows/releases/latest/download/flex_${os}_amd64.tar.gz"

echo "Fetching ${url}"

mkdir -p "${target_dir}" && curl -LX GET "${url}" | tar -xvz -C "${target_dir}"