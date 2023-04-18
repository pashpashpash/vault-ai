#!/bin/bash

# MUST use . and not `source`
# . scripts/source-me.sh

pretty_echo() {
    echo -e "\033[1;35m->\033[0m" "$@"
}

# What to compile...
TARGET="$1"
if [ "$TARGET" == "" ]; then
    echo " Usage: $0 <go package name>"
    exit 1
fi

# Install direct code dependencies
pretty_echo "Installing '$TARGET' dependencies"

go get -v "$TARGET"
RESULT="$?"
if [ "$RESULT" != "0" ]; then
    echo "   ... error"
    exit $RESULT
fi

# Compile / Install the server
pretty_echo " Compiling '$TARGET'"

go install -v "$TARGET"
RESULT="$?"
if [ "$RESULT" == "0" ]; then
    echo "   ... done"
    exit 0
else
    echo "   ... error"
    exit $RESULT
fi
