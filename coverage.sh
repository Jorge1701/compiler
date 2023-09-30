#!/bin/bash

if [ $# -eq 2 ]; then
    echo "Getting coverage for package $1 running tests $2"
    go test -coverprofile cover.out "$1" -run "$2"
else
    echo "Getting coverage for all packages"
    go test -coverprofile cover.out ./...
fi

go tool cover -html=cover.out -o cover.html
