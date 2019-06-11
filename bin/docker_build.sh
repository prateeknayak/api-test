#!/bin/bash

if [ -z ${TRAVIS_TAG+x} ]; then
    docker build . -t pnyak/api-test:latest --build-arg app_version=latest
else
    echo "travis tag"
    docker build . -t pnyak/api-test:${TRAVIS_TAG} --build-arg app_version=${TRAVIS_TAG}
fi


