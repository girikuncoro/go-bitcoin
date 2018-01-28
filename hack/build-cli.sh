#!/bin/bash

# This builds the application from source for multiple platforms.
set -eo pipefail

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR"

BINARY_NAME="coin"

# Determine the arch/os combos we're building for
XC_ARCH=${XC_ARCH:-"amd64 arm"}
XC_OS=${XC_OS:-"darwin linux windows"}

# Build!
echo "==> Building..."
gox \
  -verbose \
  -os="${XC_OS}" \
  -arch="${XC_ARCH}" \
  -osarch="!darwin/arm" \
  -ldflags "${LDFLAGS}" \
  -output "dist/bin/${BINARY_NAME}-{{.OS}}-{{.Arch}}/${BINARY_NAME}" \
  ./cmd/...

# Done!
echo "==> Results:"
echo "==>./dist/bin"
cp "dist/bin/${BINARY_NAME}-$(go env GOOS)-$(go env GOARCH)/${BINARY_NAME}" dist/bin/
ls -hlR dist/bin/*
