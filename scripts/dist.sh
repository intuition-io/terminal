#!/bin/bash
# Wrote by hashicorp (https://github.com/hashicorp/terraform/blob/master/scripts/dist.sh)
set -e

if [ -z $BINTRAY_USER ]; then
    echo "Please set BINTRAY_USER."
    exit 1
fi
if [ -z $BINTRAY_REPO ]; then
    echo "Please set BINTRAY_REPO"
    exit 1
fi
BINTRAY_URL="https://api.bintray.com/content/${BINTRAY_USER}"

# Get the version from the command line
VERSION=$1
if [ -z $VERSION ]; then
    echo "Please specify a version."
    exit 1
fi

# Make sure we have a bintray API key
if [ -z $BINTRAY_API_KEY ]; then
    echo "Please set your bintray API key in the BINTRAY_API_KEY env var."
    exit 1
fi

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that dir because we expect that
cd $DIR
# Get the application name from root directory
# TODO Overwrite it if $1 is supplied ?
APP=${PWD##*/}

# Zip all the files
rm -rf ./pkg/dist
mkdir -p ./pkg/dist
for FILENAME in $(find ./pkg -mindepth 1 -maxdepth 1 -type f); do
    FILENAME=$(basename $FILENAME)
    cp ./pkg/${FILENAME} ./pkg/dist/${APP}_${VERSION}_${FILENAME}
done

# Make the checksums
pushd ./pkg/dist
shasum -a256 * > ./${APP}_${VERSION}_SHA256SUMS
popd

# Upload
for ARCHIVE in ./pkg/dist/*; do
    ARCHIVE_NAME=$(basename ${ARCHIVE})

    echo Uploading: $ARCHIVE_NAME
    curl \
        -T ${ARCHIVE} \
        -u${BINTRAY_USER}:${BINTRAY_API_KEY} \
        "${BINTRAY_URL}/${BINTRAY_REPO}/${APP}/${VERSION}/${ARCHIVE_NAME}"
done
# TODO Automatic publishing ?

exit 0
