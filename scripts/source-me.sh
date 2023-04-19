#!/bin/bash
# Useful variables. Source from the root of the project

# Shockingly hard to get the sourced script's directory in a portable way
if [[ "${0}" == "bash" || "${0}" == "sh" ]]; then
    script_name="${BASH_SOURCE[0]}"
else
    script_name="${0}"
fi

export GO111MODULE=on
export GOBIN="$PWD/bin"
export GOPATH="$HOME/go"
export PATH="$PATH:$PWD/bin:$PWD/tools/protoc-3.6.1/bin"
export DOCKER_BUILDKIT=1

echo "=> Environment Variables Loaded"