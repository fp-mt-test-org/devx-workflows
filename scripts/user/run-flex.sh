#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

auto_update="${auto_update:-0}"
devx_workflows_path='./.devx-workflows'
flex_path="${devx_workflows_path}/flex"
flex_version_command="${flex_path} -version"
initial_flex_version=$(${flex_version_command})
service_config_path='./service_config.yml'

echo "Running Flex ${initial_flex_version}"

# Check the service_config, if it exists (i.e. is not first run of flex)
if [[ "${auto_update}" == "1" ]] && [[ -f "${service_config_path}" ]]; then
    service_config=$(cat ${service_config_path})

    if [[ "${service_config}" =~ [0-9]+.[0-9]+.[0-9]+ ]]; then
        configured_flex_version="${BASH_REMATCH[0]}"
        echo "service_config: flex: version: ${configured_flex_version}"

        # Regex for matching snapshot versions such as v0.8.3-SNAPSHOT-27afad4
        configured_flex_version_regex=".*${configured_flex_version}.*"

        if ! [[ "${initial_flex_version}" =~ ${configured_flex_version_regex} ]]; then
            echo "Current version is different than configured, upgrading..."
            "${devx_workflows_path}/scripts/user/install-flex.sh" "${configured_flex_version}"
            echo "Current version is now:"
            ${flex_version_command}
            echo "Upgrade complete."
        fi
    fi
fi

"${flex_path}" "$@"
