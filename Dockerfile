FROM golang:alpine AS build

RUN apk update --no-cache && apk add git
WORKDIR /app

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

ADD ./ /app

RUN go build -o api-test

FROM alpine
WORKDIR /app
COPY --from=build /app/api-test /app

ENTRYPOINT ["/app/api-test"]

EXPOSE 8080
