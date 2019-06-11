#!/bin/bash

echo "travis tag set: ${TRAVIS_TAG}"
if [ -z ${TRAVIS_TAG+x} ]; then
    echo "travis tag"
    docker build . -t pnyak/api-test:${TRAVIS_TAG} --build-arg app_version=${TRAVIS_TAG}
fi

docker build . -t pnyak/api-test:latest --build-arg app_version=latest

