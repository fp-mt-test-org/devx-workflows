#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

skip_download="${skip_download:=false}"
download_dir='dist'
install_dir='.devx-workflows'
os=$(uname | tr '[:upper:]' '[:lower:]')
file_name="flex_${os}_amd64.tar.gz"
url="https://github.com/fp-mt-test-org/devx-workflows/releases/latest/download/${file_name}"

mkdir -p "${install_dir}"
mkdir -p "${download_dir}"

cd "${download_dir}"

if [ "$skip_download" == "false" ]; then
    echo "Downloading ${url} to ${download_dir}"
    curl -L "${url}"
fi

echo "Extracting ${file_name} to ${install_dir}"
tar -xf "${file_name}" -C "../${install_dir}"

echo "Installation complete!"
echo ""
