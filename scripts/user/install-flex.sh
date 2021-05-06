#!/usr/bin/env bash

set -o errexit
set -o nounset

version_to_install="${1:-latest}"
skip_download=${skip_download:=0}
download_folder_path="${download_folder_path:=$(realpath dist)}"
install_folder_name='.devx-workflows'
install_path="${install_path:=$(realpath ${install_folder_name})}"
user_scripts_install_path="${install_path}/scripts/user"

echo "Installing flex version $version_to_install!"

# Generate the platform specific file name to download.
os=$(uname | tr '[:upper:]' '[:lower:]')
file_name="flex_${os}_amd64.tar.gz"
base_url='https://github.com/fp-mt-test-org/devx-workflows/releases'
if [[ "${version_to_install}" == "latest" ]]; then
    url="${base_url}/latest/download/${file_name}"
else
    url="${base_url}/download/v${version_to_install}/${file_name}"
fi
     
mkdir -p "${install_path}"
mkdir -p "${download_folder_path}"

download_file_path="${download_folder_path}/${file_name}"

if [ "${skip_download}" -ne "1" ]; then
    echo "Downloading ${url} to ${download_file_path}"
    curl -L "${url}" --output "${download_file_path}"
fi

echo "Extracting ${download_file_path} to ${install_path}"
tar -xvf "${download_file_path}" -C "${install_path}"

git_ignore_file='.gitignore'

if ! grep -qs "${install_folder_name}" "${git_ignore_file}"; then
    echo "Updating ${git_ignore_file} to ignore the ${install_path} install_path..."
    echo "${install_folder_name}" >> "${git_ignore_file}"
fi

echo "Configuring the local host..."
"${user_scripts_install_path}/configure-localhost.sh"

if [ "${auto_clean:=1}" == "1" ]; then
    echo "Cleaning up ${download_file_path}"
    rm -fdr "${download_file_path}"
fi

echo "Installation complete!"
echo ""
