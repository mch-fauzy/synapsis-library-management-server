#!/bin/sh

# Load docker environment variables from .env.docker file
source .env.docker

# Wait for the database container to become available
dockerize -wait tcp://${DOCKER_DB_POSTGRESQL_HOST}:${DOCKER_DB_POSTGRESQL_PORT} -timeout 1m

# Run database migration
OPTIONS="-config=dbconfig.yml -env development"
sql-migrate up $OPTIONS

# Start the API server
./main
