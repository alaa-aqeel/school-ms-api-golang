#!/bin/bash

# Set the absolute path for the migrations directory
ABSOLUTE_PATH="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MIGRATIONS_PATH="${ABSOLUTE_PATH}/database/migrations"

case "$1" in
    "make")
        # docker run \
        #     -v "${MIGRATIONS_PATH}":/migrations \
        #     --network fastapi_db-network \
        #     migrate/migrate \
        #     create -ext sql -dir /migrations "$2"
        migrate create -ext sql -dir "${MIGRATIONS_PATH}" -seq "$2"

        # Check if the create command was successful
        if [ $? -eq 0 ]; then
            echo "Migration file created successfully."
        else
            echo "Error: Failed to create migration file."
            exit 1
        fi
        ;;
    "up")
        # Run migrate up to apply the migrations
        docker run \
            -v "${MIGRATIONS_PATH}":/migrations \
            --network school-ms-api-golang_pgdb-net migrate/migrate \
            -path /migrations \
            -database postgres://db_username:db_pass@postgres-db:5432/app_db?sslmode=disable up

        # Check if the up command was successful
        if [ $? -eq 0 ]; then
            echo "Migrations applied successfully."
        else
            echo "Error: Failed to apply migrations."
            exit 1
        fi
        ;;
    "down")
        # Run migrate down to undo the last migration
        docker run \
            -v "${MIGRATIONS_PATH}":/migrations \
            --network school-ms-api-golang_pgdb-net migrate/migrate \
            -path /migrations \
            -database postgres://db_username:db_pass@postgres-db:5432/app_db?sslmode=disable down $2

        # Check if the down command was successful
        if [ $? -eq 0 ]; then
            echo "Last migration undone successfully."
        else
            echo "Error: Failed to undo last migration."
            exit 1
        fi
        ;;
    *)
        echo "Usage: $0 {make|up|down} [migration_name]"
        exit 1
        ;;
esac