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
- Architectural patterns for maintainable and scalable APIs (e.g., Repository Pattern)
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

#### .golangci.yml

This file configures `golangci-lint`, a popular Go linter, to ensure code quality and consistency across the project.

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
- [postman](https://www.postman.com/)
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

### Break down into integration tests

Happy path integration tests for v0alpha/crud/users.

#### Executing the integration tests

Start the service and ensure it runs without errors, follow the instructions above in [installing](#installing) for instructions to get the service running.

1. Open Postman and **import** the Postman collection located in this repo at `tests/v0alpha-crud-users.postman_collection.json`
2. Once imported, there will be a folder under Collections named `v0alpha/crud/users`.
3. **Hover over** the folder and a play button will appear on the right, **click** the play button.
4. A Run button with a blue background will appear, **click** the Run button.
5. After clicking the Run button a new window named Collection Runner will appear. This tool enables the user to execute the integration tests with ease.
6. **Set** the iterations to 1, can try more later.
7. **Set** the delay to 1000 ms.
8. **Press** the blue button near the bottom named `Run v0alpha/crud...`

The Collection Runner will now execute the happy path integration tests. There should be activity in the logs of the service that show the requests being made.

#### Test explanation and Results

At a high level these integration tests execute the happy path to create a user, modify the user, delete the user, with confirmation steps between to ensure the actions are being taken and executed properly.

These tests are a WIP, sometimes these tests will fail based on database and service times. A longer delay between in the Collection Runner results in fewer failures.

### Break down into unit tests

The project now utilizes the Repository pattern, abstracting database interactions behind an interface. This allows for effective unit testing of HTTP handlers by injecting mock implementations of the `UserRepository`.

**Handler Unit Tests:**
Unit tests for API handlers (e.g., for user CRUD operations) are located alongside the handlers themselves. These tests leverage `github.com/stretchr/testify/mock` to simulate repository behavior, ensuring handlers function correctly in isolation.

Test command for all unit tests:

```
go test -v ./...
```

Example Results (may vary):

```
=== RUN   TestGetPong
--- PASS: TestGetPong (0.00s)
=== RUN   TestGetAllUsers
--- PASS: TestGetAllUsers (0.00s)
=== RUN   TestCreateUser
--- PASS: TestCreateUser (0.00s)
=== RUN   TestGetSpecificUser_NotFound
--- PASS: TestGetSpecificUser_NotFound (0.00s)
PASS
ok  	github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/handlers	0.XXXs
... (other test outputs)
```

Coverage information command:

```
go tool cover -func coverage.out
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
- [testify](https://github.com/stretchr/testify) - A Go testing toolkit, including mock generation
- [google/uuid](https://github.com/google/uuid) - UUID package for Go

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Links/References

- [Project-Layout](https://github.com/golang-standards/project-layout)
- [Dockerfile](https://github.com/chemidy/smallest-secured-golang-docker-image/blob/master/Dockerfile)
- [CustomizingGitGitHooks](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks)
- [ConventionalCommits](https://www.conventionalcommits.org/en/v1.0.0-beta.2/)
