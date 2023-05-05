#!/bin/bash

pretty_echo() {
    echo -e "\033[1;35m->\033[0m" "$@"
}

# What to compile...
TARGET="$1"
if [ "$TARGET" == "" ]; then
    echo " Usage: $0 <go package name>"
    exit 1
fi

# Save the current directory and change to the target directory
original_dir=$(pwd)
cd "$TARGET"

# Install direct code dependencies
pretty_echo "Installing '$TARGET' dependencies"

go get -v
RESULT="$?"
if [ "$RESULT" != "0" ]; then
    echo "   ... error"
    exit $RESULT
fi

# Compile the server
pretty_echo " Compiling '$TARGET'"

env GOOS=windows GOARCH=amd64 go build -o ../bin/vault-web-server.exe -v
RESULT="$?"
if [ "$RESULT" == "0" ]; then
    echo "   ... done"
else
    echo "   ... error"
fi

# Change back to the original directory
cd "$original_dir"

exit $RESULT
