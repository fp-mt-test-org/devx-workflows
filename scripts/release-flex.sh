#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset

if [[ "${dry_run:=true}" == "true" ]]; then
    goreleaser --snapshot --skip-publish --rm-dist
else
    echo "Installing svu..."
    brew install caarlos0/tap/svu
    echo "Install complete."

    next_version="$(svu n)"

    echo "Tagging repo with ${next_version}"
    git tag "$(svu n)"

    echo "Pushing tags"
    git push --tags
    
    echo "Tagging complete."
    echo ""

    goreleaser release --rm-dist
fi

echo "Release complete."
