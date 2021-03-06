## Golang Codebase

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com) 
[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/contains-technical-debt.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/check-it-out.svg)](https://forthebadge.com)

Your Go Codebase Built with :heart:

### Requirements
- Golang version 1.12+ (https://golang.org/)
- Docker (https://docs.docker.com/)
- Docker compose (https://docs.docker.com/compose/)

### Building and running tests
The software is designed to be able to run both on the machine on top of docker. You can see Makefile to see what kind of targets supported.

#### Docker compose
open <a href="https://gitlab.com/Wuriyanto/go-codebase/blob/development/docker-compose.yml">docker-compose.yml</a> for details
```shell
$ docker-compose up
```

#### Code identation should be consistent (Make sure run `make format` before pushing your code)
```shell
$ make format
```

### Running Test, Coverage and Linter

#### Linter

Get golangci-lint binary
```shell
$ make lint-prepare
```

Run Linter
```shell
$ make lint
```

#### Unit Test

This test run without environment variable set.

```
$ make test
```

#### Coverage

Running Coverage and produce html output

```
$ make cover-html
```

Then open `coverages/index.html`
```shell
$ open coverages/index.html
```

#### Copyright (c) 2019 PT Telekomunikasi Indonesia Tbk.