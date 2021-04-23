#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

case "$SHELL" in
 "/bin/zsh") profile_path=~/.zshrc ;;
 *) profile_path=~/.bashrc ;;
esac

flex_alias="alias flex=\"./.devx-workflows/flex\""

if ! grep -q "${flex_alias}" "${profile_path}"; then
    profile_content=$(cat ${profile_path})
    # Save it to the profile so it's execute for each shell session.
    echo -e "${flex_alias}\n${profile_content}" > "${profile_path}"

    echo "Added ${flex_alias} to your ${profile_path}."

    # Start a new shell session so that the
    # user can use the alias immediately.
    $SHELL
fi

