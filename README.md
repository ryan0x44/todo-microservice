# Todo Microservice experiments

This repository contains experiments of building a Todo Microservice.

Each implementation of Todo Microservice should support a multi-transport,
multi-version API using HTTP JSON and gRPC/Protobufs.

The goal is that by doing so, we can demonstrate the strength and weaknesses of
various tools/frameworks/libraries/patterns.

## Domain Model

We keep it simple:

* Projects have a name
* Tasks have a description
* Projects can have many tasks assigned

# Docker

We have a `docker-compose.yml` file which enables us to quickly build and
run each service with only Docker and Docker Compose as dependencies.

e.g. `docker-compose up -d --build` (run without `--build` to cache builds)

To correctly shutdown containers and clean-up the database volume:

e.g. `docker-compose down --volumes`

Our set-up includes MySQL with the DB initialized from `database/todo.sql`.

## HTTP Contract Tests

We use [newman](https://github.com/postmanlabs/newman) to run contract tests.

e.g. `docker run -t postman/newman_ubuntu1404 run contract-tests`

## Implementations

__Go Kit__

Uses <http://gokit.io/>
