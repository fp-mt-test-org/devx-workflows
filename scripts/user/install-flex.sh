#!/usr/bin/env bash

set -o errexit
set -o nounset

skip_download="${skip_download:=false}"
download_dir='dist'
install_dir='.devx-workflows'
os=$(uname | tr '[:upper:]' '[:lower:]')
file_name="flex_${os}_amd64.tar.gz"
url="https://github.com/fp-mt-test-org/devx-workflows/releases/latest/download/${file_name}"
download_file_path="${download_dir}/${file_name}"

mkdir -p "${install_dir}"
mkdir -p "${download_dir}"

if [ "$skip_download" == "false" ]; then
    echo "Downloading ${url} to ${download_dir}"
    curl -L "${url}" --output "${download_file_path}"
fi

echo "Extracting ${download_file_path} to ${install_dir}"
tar -xf "${download_file_path}" -C "${install_dir}"

echo "Cleaning up ${download_dir}"
rm -fdr "${download_dir}"

echo "Installation complete!"
echo ""
