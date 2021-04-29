#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

echo "Installing svu..."
brew install caarlos0/tap/svu
echo "Install complete."

next_version="$(svu n)"

if [[ "${dry_run:=1}" == "1" ]]; then
    goreleaser --snapshot --skip-publish --rm-dist
else
    echo "Tagging repo with ${next_version}"
    git tag "$(svu n)"

    echo "Pushing tags"
    git push --tags
    
    echo "Tagging complete."
    echo ""

    goreleaser release --rm-dist
fi

echo "Release complete."
