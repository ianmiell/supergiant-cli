#!/bin/bash
VERSION=$(tail -1 version.go | awk '{print $5}' | sed 's/"//g')
STABLE=$(tail -1 version.go | awk '{print $5}' | sed 's/"//g' | cut -f1 -d"-")
if [[ ! -z "$TRAVIS_TAG" && "$STABLE" != "$VERSION" ]]; then
  echo "Releasing supergiant-cli version: ${VERSION}, pre-release"
  ghr --username supergiant --token $GITHUB_TOKEN --replace --prerelease --debug unstable-$VERSION dist/
  exit 0
elif [ ! -z "$TRAVIS_TAG" ]; then
  echo "Releasing supergiant-cli version: ${TRAVIS_TAG}, as latest release."
  ghr --username supergiant --token $GITHUB_TOKEN --replace --debug $TRAVIS_TAG dist/
  exit 0
fi
echo "Unable to determine tag."
exit 5
