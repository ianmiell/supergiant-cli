#!/bin/sh
VERSION=$(tail -1 version.go | awk '{print $5}' | sed 's/"//g')
if [ -z "$TRAVIS_TAG" ] && [ "$TRAVIS_BRANCH" == "master" ]; then
  echo "Releasing supergiant-cli version: ${VERSION}, pre-release"
  ghr --username spankenstein --token $GITHUB_TOKEN --replace --prerelease --debug $VERSION dist/
else
  echo "Releasing supergiant-cli version: ${TRAVIS_TAG}, as latest release."
  ghr --username spankenstein --token $GITHUB_TOKEN --replace --debug $TRAVIS_TAG dist/
fi