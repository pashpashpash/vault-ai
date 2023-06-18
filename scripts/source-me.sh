#!/bin/bash
# Useful variables. Source from the root of the project

# Shockingly hard to get the sourced script's directory in a portable way
if [[ "${0}" == "bash" || "${0}" == "sh" ]]; then
    script_name="${BASH_SOURCE[0]}"
else
    script_name="${0}"
fi
dir_path="$( cd "$(dirname "$script_name")" >/dev/null 2>&1 ; pwd -P )"
secrets_path="${dir_path}/../secret"
test ! -d $secrets_path && echo "ERR: ../secret dir missing!" && return 1

export GO111MODULE=on
export GOBIN="$PWD/bin"
export GOPATH="$HOME/go"
export PATH="$PATH:$PWD/bin:$PWD/tools/protoc-3.6.1/bin"
export DOCKER_BUILDKIT=1
export OPENAI_API_KEY="$(cat ${secrets_path}/openai_api_key)"
export PINECONE_API_KEY="$(cat ${secrets_path}/pinecone_api_key 2>/dev/null)"
export PINECONE_API_ENDPOINT="$(cat ${secrets_path}/pinecone_api_endpoint 2>/dev/null)"

echo "=> Environment Variables Loaded"