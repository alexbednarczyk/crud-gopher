# crud-gopher

**[Table of Contents]**

- [Synopsis](#synopsis)
- [Motivation](#motivation)
- [Code Example](#code-example)
- [Endpoint Documentation](#endpoint-documentation)
- [Getting Started](#getting-started)
  - [Configuration and/or Environment Variables](#configuration-andor-environment-variables)
  - [Prerequisites](#prerequisites)
  - [Installing](#installing)
- [Running the tests](#running-the-tests)
  - [Break down into unit tests](#break-down-into-unit-tests)
- [Built With](#built-with)
- [License](#license)
- [Links](#links)

This is an example or starter providing guidance in creating APIs using Golang.

## Synopsis

This repo provides a central location of guidance for best practices when creating APIs using Golang

## Motivation

The goal of this repo is to create a location for folks to get information they need to create an API using Golang. Hunting for best practices, and missing unknown unknowns results in poor performing APIs that have missing pieces.

In the end this repo will provide guidance for the following:

- API Documentation
- Basic and Advanced API examples
- Continuous integration and continuous deployment
- Developer workflow
- Integration, Performance, and Unit tests

## Endpoint Documentation

[Docs folder](docs) or when the service is running locally go to http://localhost:8080/swagger/index.html

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Configuration and/or Environment Variables

The location of the configuration and/or environment variables for to get this project up and running.

#### .env File

Place this file in the root directory of the repo and add custom values for each env variable.

```
DEPLOYED=false
LOG_LEVEL=trace
POSTGRES_HOST=
POSTGRES_PORT=
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_SSLMODE=
POSTGRES_DATABASE=
# Optional
# POSTGRES_LAZY_CONNECT=
# POSTGRES_POOL_MAX_CONNECTION_LIFETIME=
# POSTGRES_POOL_MAX_CONNECTION_IDLE_TIME=
# POSTGRES_POOL_MAX_CONNECTIONS=
# POSTGRES_POOL_MIN_CONNECTIONS=
# POSTGRES_POOL_HEALTH_CHECK_PERIOD=
```

#### Available Git Hooks

##### pre-commit

1. Check lint using golangci-lint from https://github.com/golangci/golangci-lint
2. Run unit tests.
3. Go mod updates.

##### prepare-commit-msg

Check that branch name meets naming requirement.

##### commit-msg

Check commit message meets style requirement.

#### Prerequisites

What things you need to install the software and how to install them

- [golang](https://golang.org/doc/install)
- [golangci-lint](https://github.com/golangci/golangci-lint#install)
- [swaggo](https://github.com/swaggo/swag#getting-started)

### Installing

A step by step series of examples that tell you have to get a development env running

Setup Git Hooks

```
./scripts/setup_git_hooks.sh
```

Create Swagger Docs

```
swag init
```

Start API Service

```
go run main.go
```

Check API Service is running

```
http://localhost:8080/v0alpha/ping
```

Should result in "pong" being displayed in the browser

Check API Docs are available

```
http://localhost:8080/swagger/index.html
```

## Running the tests

Explain how to run the automated tests for this system

### Break down into unit tests

Basic unit test for the v0alpha/ping endpoint

Test command:

```
go test -v -covermode=count -coverprofile=coverage.out ./...
```

Results in:

```
=== RUN   TestPingRoute
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /v0alpha/ping             --> github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/dal.GetPong.func1 (3 handlers)
[GIN-debug] GET    /v0alpha/examples/attribute --> github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/dal.Attribute.func1 (3 handlers)
[GIN-debug] GET    /v0alpha/examples/calc    --> github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/dal.CalcSumOfTwo.func1 (3 handlers)
[GIN-debug] GET    /v0alpha/examples/groups/:group_id/accounts/:account_id --> github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/dal.GetAccountIDInGroup.func1 (3 handlers)
[GIN-debug] GET    /v0alpha/examples/header  --> github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/dal.UseHeader.func1 (3 handlers)
[GIN-debug] GET    /v0alpha/examples/securities --> github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/dal.Securities.func1 (3 handlers)
[GIN] 2020/04/26 - 11:17:07 | 200 |      22.323µs |                 | GET      "/v0alpha/ping"
--- PASS: TestPingRoute (0.00s)
PASS
coverage: 0.0% of statements
ok      github.com/alexbednarczyk/crud-gopher   1.285s  coverage: 0.0% of statements
?       github.com/alexbednarczyk/crud-gopher/docs      [no test files]
?       github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/dal  [no test files]
?       github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/routes       [no test files]
=== RUN   TestSetLogLevel
--- PASS: TestSetLogLevel (0.00s)
=== RUN   TestStandardFormatLogs
--- PASS: TestStandardFormatLogs (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/alexbednarczyk/crud-gopher/internal/common   0.773s  coverage: 100.0% of statements
=== RUN   TestEnvironmentVariables
time="2020-04-26T11:17:05-05:00" level=info msg="All Environment Variables" DEPLOYED=false LOG_LEVEL=trace
--- PASS: TestEnvironmentVariables (0.00s)
PASS
coverage: 66.7% of statements
ok      github.com/alexbednarczyk/crud-gopher/internal/environment/variables    1.049s  coverage: 66.7% of statements
```

Coverage information command:

```
go tool cover -func coverage.out
```

Results in:

```
github.com/alexbednarczyk/crud-gopher/internal/common/log.go:11:                        SetLogLevel                     100.0%
github.com/alexbednarczyk/crud-gopher/internal/common/log.go:31:                        StandardFormatLogs              100.0%
github.com/alexbednarczyk/crud-gopher/internal/environment/variables/variables.go:17:   FetchAllEnvironmentVariables    66.7%
github.com/alexbednarczyk/crud-gopher/main.go:25:                                       main                            0.0%
total:                                                                                  (statements)                    68.4%

```

Visual coverage command:

```
go tool cover -html=coverage.out
```

## Build Commands

`docker build -t cruggolang:latest -f build/package/Dockerfile .`
`docker run -p 8080:8080 cruggolang:latest`

## Built With

- [Golang](https://golang.org/) - Go, the open source programming language
- [Gin](https://github.com/gin-gonic/gin) - Gin is a HTTP web framework written in Go
- [golang-migrate](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library.
- [pgx](https://github.com/jackc/pgx) - PostgreSQL driver and toolkit for Go
- [Swagger](https://swagger.io/) - API tooling
- [Docker](https://www.docker.com/) - Docker

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Links/References

- [Project-Layout](https://github.com/golang-standards/project-layout)
- [Dockerfile](https://github.com/chemidy/smallest-secured-golang-docker-image/blob/master/Dockerfile)
- [CustomizingGitGitHooks](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks)
- [ConventionalCommits](https://www.conventionalcommits.org/en/v1.0.0-beta.2/)
