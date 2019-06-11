# Test 2 API

[![Build Status](https://travis-ci.org/prateeknayak/api-test.svg?branch=master)](https://travis-ci.org/prateeknayak/api-test)

Summary:
- Simple api written in Golang
- Developed using `github.com/gorilla/mux` library
- Travis-ci is used for creating a CI pipeline that gets triggered on each push, pull request, tag etc.
- The application is intended to run as a docker container and the CI pipeline builds a docker image

### Table of Contents

- Pre-requisites
- How to run the app?
- How to build the app?
- CI pipeline and Versioning
- Risks and Limitations
- Future Improvements


### 1 Pre-requisites

- Docker engine
- git
- golang 1.12
- Set GO11MODULE=on in your bashrc

 `export GO111MODULE=on`

### 2 How to run the app?

#### 2.1 Docker

The travis pipeline in this project publishes a docker image to docker hub.

##### 2.1.1 Latest

Execute the following command to run the latest version of this app (not recommended for production, use a specific version)

```
docker run -it -p 8080:8080 pnyak/api-test:latest
```

The above command should result in a log line like

```
2019/06/11 17:20:57 starting the api
```
this means the API has started successfully and now you can access it on `http://localhost:8080/`

##### 2.1.2 v1.0.3

Execute the following command to run the v1.0.3 version of this app
```
docker run -it -p 8080:8080 pnyak/api-test:v1.0.3
```

The above command should result in a log line like
```
2019/06/11 22:54:06 starting the api
```
this means the API has started successfully and now you can access it on `http://localhost:8080/`


#### 2.2 Golang


*NOTE:* For this you will need to have `git` and `golang` installed and configured on your machine.

Clone the repo

```
git clone https://github.com/prateeknayak/api-test.git
```

From the root of the project
```
go run main.go
```

The above command should result in log line like
```
2019/06/12 07:27:56 starting the api
```

This means the API has started successfully and now you can access it on `http://localhost:8080/`

### 3 How to build the app?

*NOTE:* For this you will need to have `git` and `golang` installed and configured on your machine.

Clone the repo

```
git clone https://github.com/prateeknayak/api-test.git
```

Run the tests

```
# navigate to the root of the project
go test -v ./...
```

Build the project

```
# navigate to the root of the project
go build -o api-test
```

Run the app from artefact

```
./api-test
```

### 4 CI Pipeline and Versioning

This project has a travis CI pipeline enabled to build the application and publish the docker image to dockerhub. The API follows a loose versioning scheme which would need to be evolved per the use case of the team that works on this. the current versioning scheme and ci strategy is

- Push docker image when merged / pused to master branch. Use `latest` as version for both app and docker image.
- Push docker image when a tag is pushed. Use the tag as version for both app and docker image.
- CI builds on push on every single branch and pull request but will not push any image.
- CI is leveraging travis's `deploy` to target branches rather then logic in shell script.

### 5 Risks and Limitations

- The API does not have logging setup which means it is hard to debug the individual API calls in the application.
- The API does not have any monitoring configured.
- Currently the API is served over http which is a big risk in terms of production deploy.
- It does not have any authentication / authorisation support.
- It does not support deploying to any orchestration engine.
- Uses environment vars on travis for dockerhub auth.


### 6 Future Improvements

- Add observability support to the application by configuring logging, monitoring, tracing.
- Enable ssl to serve traffic over https
- Design OIDC or JWT flow for authentication and authorisation support
- Secure way of providing docker credential. i.e use `travis encrypt` functionality.
- Make this application deployable to a container orchestration engine like kubernetes.
- Design a app / image versioning strategy based on the teams requirements.



