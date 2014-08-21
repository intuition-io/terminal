#!/bin/bash
# Wrote by hashicorp (https://github.com/hashicorp/terraform/blob/master/scripts/build.sh)
# This script builds the application from source for only this platform.
set -e

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd $DIR
# Get the application name from root directory
# TODO Overwrite it if $1 is supplied ?
APP=${PWD##*/}

# Get the git commit
GIT_COMMIT=$(git rev-parse HEAD)
GIT_DIRTY=$(test -n "`git status --porcelain`" && echo "+CHANGES" || true)

# If we're building on Windows, specify an extension
EXTENSION=""
if [ "$(go env GOOS)" = "windows" ]; then
  EXTENSION=".exe"
fi

GOPATHSINGLE=${GOPATH%%:*}
if [ "$(go env GOOS)" = "windows" ]; then
  GOPATHSINGLE=${GOPATH%%;*}
fi

# Install dependencies
echo "--> Getting dependencies..."
go get ./...

# Delete the old dir
echo "--> Removing old directory..."
rm -f bin/*

# Build!
echo "--> Building..."
gox \
  -os="$(go env GOOS)" \
  -arch="$(go env GOARCH)" \
  -ldflags "-X main.GitCommit ${GIT_COMMIT}${GIT_DIRTY}" \
  -output "bin/${APP}-{{.Dir}}" \
  ./...
mv bin/${APP}-${APP}${EXTENSION} bin/${APP}${EXTENSION}
cp bin/${APP}* ${GOPATHSINGLE}/bin

if [ "${TF_DEV}x" = "x" ]; then
  # Zip and copy to the dist dir
  echo "--> Packaging..."
  mkdir -p pkg
  cd bin/
  zip ../pkg/$(go env GOOS)_$(go env GOARCH).zip ./*
  cd $DIR
fi

# Done!
echo
echo "--> Results:"
ls -hl bin/
