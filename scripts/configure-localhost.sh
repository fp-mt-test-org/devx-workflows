#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

echo $SHELL
cat ~/.bashrc

profile_path=~/.zshrc
flex_relative_path="alias flex=\"./.devx-workflows/flex\""

if ! grep -q "${flex_relative_path}" "${profile_path}"; then
    # Save it to the profile so it's execute for each shell session.
    echo "${flex_relative_path}" >> "${profile_path}"

    echo "Added ${flex_relative_path} to your ${profile_path}."

    # Start a new shell session so that the
    # user can use the alias immediately.
    zsh
fi

# Local install of goreleaser for build script
brew install goreleaser
