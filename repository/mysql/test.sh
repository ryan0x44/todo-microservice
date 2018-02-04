#!/bin/bash -e

# Deps
go get github.com/go-sql-driver/mysql
go get github.com/NYTimes/gizmo/config/mysql
go get github.com/kelseyhightower/envconfig

# Load env vars from file
ENV_FILE=../../dev.env
if [ ! -f $ENV_FILE ]; then
    echo "Env file $ENV_FILE not found"
    exit 1
fi
export $(grep -v "^#" $ENV_FILE | xargs)

# Check database is running
nc -z 127.0.0.1 3306 &> /dev/null || \
    (echo "DB error. Start the DB container?" && exit 1)

# Override mysql hostname (since we're not in the docker network)
MYSQL_HOST_NAME=127.0.0.1

# Run go tests
go test -mysql=true
