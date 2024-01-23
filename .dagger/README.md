# Agate CI

Agate use [Dagger](https://dagger.io/) to handle its CI and local development.

## Table of Contents

- [Requirements](#requirements)
- [Dagger commands](#dagger-commands)
    - [Open a shell inside a container with the project loaded](#open-a-shell-inside-a-container-with-the-project-loaded)
    - [Install Agate binary from local source](#install-agate-binary-from-local-source)
    - [Start a PostgreSQL database](#start-a-postgresql-database)
    - [Execute CI locally](#execute-ci-locally)

## Requirements

- [Dagger CLI installed in your host](https://docs.dagger.io/cli/465058/install)

## Dagger commands

💡 You can list all functions using `dagger functions`

### Open a shell inside a container with the project loaded

```shell
dagger shell project --entrypoint /bin/sh
```

Use it to debug flacky tests or test things inside the CI or container environment.

### Install Agate binary from local source

```shell
dagger export binary
```

You should see `agate` binary in your current host directory.

### Start a PostgreSQL database

```shell
dagger up --native database
```

You can either configure the database with flags in the CLI or using your environment.

💡 Execute `dagger call database --help` to see the flags.

### Execute CI locally

```shell
dagger call ci run
```

This will build, lint and run unit tests in parallel.

💡 Execute `dagger call ci --help` to see additional subcommands.