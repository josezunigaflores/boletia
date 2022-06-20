# Boletia Project

<p align="center">
  <a href="">
    <img alt="boletia" src="https://home-statics.boletia.com/assets/logo/icon-820b253c8aae8a65e9cf244bc6060e27.png" width="100px" height="100px" />
  </a>
</p>

<h1 align="center">
   Hexagonal Architecture & CQRS in Golang
</h1>


<p align="center">
  Example of a <strong>Golang application using Command Query Responsibility Segregation
  (CQRS) principles</strong> keeping the code as simple as possible.
  <br />
  <a href=".">View Demo</a>
  ·

</p>

## 🚀 Environment Setup

### 🐳 Needed tools

1. [Install Docker](https://www.docker.com/get-started)
2. Clone this project: `git clone git@github.com:josezunigaflores/boletia.git`
3. Move to the project folder: `cd boletia`
4. [Install task](https://taskfile.dev/installation/)
5. go install github.com/jstemmer/go-junit-report@latest

### 🛠️ Environment configuration

1. Create a local environment file (`config.json`) if you want to modify any parameter

### 🔥 Application execution

1. Install all the dependencies and bring up the project with Docker executing: `go mod tidy`

### ✅ Tests execution

1. Install the dependencies if you haven't done it previously: `go mod tidy`
2. Execute native test you should use: `task test` or to get coverage `task coverage`

### 🎯 Hexagonal Architecture

This repository follows the Hexagonal Architecture pattern.

```scala
$ tree -L 4 -I recursos
.
|-- README.md
|-- Taskfile.yml
|-- c.out
|-- cmd
|   `-- api
|       |-- bootstrap
|       |   |-- bootstrap.go // call all artifacts to will create the server.
|       |   `-- bootstrap_test.go
|       `-- main.go // main file
|-- config.json
|-- coverage.html // view about the coverage of our code.
|-- coverage.out
|-- docker // all related with docker imagen
|   `-- app
|       |-- Dockerfile
|       |-- reflex.conf
|       `-- start.sh
|-- docker-compose.yml
|-- docs // documentation
|   |-- docs.go
|   |-- init.go
|   |-- swagger.json
|   `-- swagger.yaml
|-- etc
|   `-- swag.sh
|-- go.mod
|-- go.sum
|-- golangci.yml
|-- internal
|   |-- calls.go
|   |-- config
|   |   `-- config.go
|   |-- currency
|   |   |-- command.go
|   |   |-- response.go
|   |   `-- service.go
|   |-- currency.go
|   |-- currency_test.go
|   |-- events.go
|   |-- events_test.go
|   |-- mocks // mocks are files make for mockery
|   |   |-- repository_calls.go
|   |   |-- repository_currency.go
|   |   |-- repository_currency_find.go
|   |   `-- repository_http.go
|   |-- plataform
|   |   |-- bus // implement of bus command
|   |   |   `-- inmemory
|   |   |-- http // calls to http
|   |   |   |-- currency.go
|   |   |   |-- currency_test.go
|   |   |   |-- example.json
|   |   |   |-- repository.go
|   |   |   `-- repository_test.go
|   |   |-- server // contain hander functions and engine of gin.
|   |   |   |-- handler
|   |   |   `-- server.go
|   |   `-- storage // calls to db
|   |       `-- postgres
|   |-- schedule // service of schedule
|   |   |-- event.go
|   |   |-- event_test.go
|   |   |-- service.go
|   |   `-- service_test.go
|   `-- utils // utils common request and response.
|       |-- badRequest.go
|       |-- internalError.go
|       |-- response.go
|       `-- tools.go
|-- kit // source about the events  and command bus
|   |-- command
|   |   |-- command.go
|   |   `-- commandmocks
|   |       `-- bus.go
|   `-- event
|       |-- event.go
|       `-- eventmocks
|           |-- bus.go
|           `-- event.go
|-- report.json
`-- test
    `-- currency.http

27 directories, 55 files

```
