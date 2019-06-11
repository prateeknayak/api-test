#!/bin/bash
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin

if [[ -z "${TRAVIS_TAG}" ]]; then
    docker push pnyak/api-test:latest
else
    echo "travis tag set"
    docker push pnyak/api-test:${TRAVIS_TAG}
fi
