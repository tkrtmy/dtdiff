#!/bin/bash

REPO_ROOT=$(git rev-parse --show-toplevel)

echo "Start release"

gox -os="darwin linux windows" -arch="amd64 386" -output="$REPO_ROOT/dist/{{.Dir}}_{{.OS}}_{{.Arch}}" ./cmd/dtdiff
ghr -u tkrtmy -t $GITHUB_TOKEN --replace `grep 'Version =' ./cmd/dtdiff/version.go | sed -E 's/.*"(.+)"$$/\1/'` dist/

echo "Completed"
