#!/bin/bash
# Useful variables. Source from the root of the project

# Shockingly hard to get the sourced script's directory in a portable way
set_env() {
  if [ "$(uname -s)" == "MINGW"* ]; then
    export "$1"="$2"
  else
    echo "export $1=$2"
  fi
}

if [[ "${0}" == "bash" || "${0}" == "sh" ]]; then
    script_name="${BASH_SOURCE[0]}"
else
    script_name="${0}"
fi
dir_path="$( cd "$(dirname "$script_name")" >/dev/null 2>&1 ; pwd -P )"
secrets_path="${dir_path}/../secret"
test ! -d $secrets_path && echo "ERR: ../secret dir missing!" && return 1

echo "secrets_path"
echo "$secrets_path"
echo "dir_path"
echo "$dir_path"

set_env "GO111MODULE" "on"
set_env "GOOS" "windows"
set_env "GOARCH" "amd64" # for 64-bit Windows
set_env "GOBIN" "$PWD/bin"
set_env "GOPATH" "$HOME/go"
set_env "PATH" "$PATH:$PWD/bin:$PWD/tools/protoc-3.6.1/bin"
set_env "DOCKER_BUILDKIT" "1"
set_env "OPENAI_API_KEY" "$(cat ${secrets_path}/openai_api_key)"
set_env "PINECONE_API_KEY" "$(cat ${secrets_path}/pinecone_api_key)"
set_env "PINECONE_API_ENDPOINT" "$(cat ${secrets_path}/pinecone_api_endpoint)"

echo "=> Environment Variables Loaded"
