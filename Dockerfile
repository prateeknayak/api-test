ARG app_version='latest'
FROM golang:alpine AS build

ARG app_version
ENV VERSION=$app_version

RUN apk update --no-cache && apk add git
WORKDIR /app

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

ADD ./ /app

RUN COMMIT_SHA=$(git rev-list -1 HEAD) && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.Ver=${VERSION} -X main.SHA=${COMMIT_SHA} -w -extldflags '-static'" -o api-test

FROM scratch
WORKDIR /app
COPY --from=build /app/api-test /app

ENTRYPOINT ["/app/api-test"]

EXPOSE 8080
