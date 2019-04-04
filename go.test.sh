#!/usr/bin/env bash

set -e
echo "" > coverage.txt

go test -race -coverprofile=profile.out -covermode=atomic github.com/tarcisio/gotm1
if [ -f profile.out ]; then
    cat profile.out >> coverage.txt
    rm profile.out
fi

go test -race -coverprofile=profile.out -covermode=atomic github.com/tarcisio/gotm1/http
if [ -f profile.out ]; then
    cat profile.out >> coverage.txt
    rm profile.out
fi
