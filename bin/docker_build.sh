#!/bin/bash

echo "travis tag ${TRAVIS_TAG}:"
if [[ -z "${TRAVIS_TAG}" ]]; then
    docker build . -t pnyak/api-test:latest --build-arg app_version=latest
else
    echo "tavis tag set"
    docker build . -t pnyak/api-test:${TRAVIS_TAG} --build-arg app_version=${TRAVIS_TAG}
fi


