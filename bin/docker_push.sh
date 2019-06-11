#!/bin/bash
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push pnyak/api-test:latest
docker push pnyak/api-test:${TRAVIS_BUILD_ID}.${TRAVIS_BUILD_NUMBER}
