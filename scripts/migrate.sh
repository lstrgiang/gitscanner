#!/usr/bin/env bash

# This is shortcut for running `migrate` against the default database as
# configured in config/database.json

# Source the common.sh script
# shellcheck source=./common.sh
. "$(git rev-parse --show-toplevel || echo ".")/scripts/common.sh"

# Use default postgres port
if [[ -z "${DB_URL}" ]]; then
    LOCAL_DB_URL="postgres://postgres:postgres@localhost:5432/gitscan?sslmode=disable"
    echo_info "Using default db url: $DB_URL"
else
    echo_info "Using db url: $LOCAL_DB_URL"
fi

# Go to project root dir to make sure that we can call other scripts correctly
cd "$PROJECT_DIR"

cmd=$1
shift
echo_info "Run migrate command: $cmd"
./bin/migrate -verbose -database "$LOCAL_DB_URL" -path ./db/migrations/ $cmd $*
cd $WORKING_DIR
