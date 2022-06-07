#!/usr/bin/env bash

. "$(git rev-parse --show-toplevel || echo ".")/scripts/common.sh"

echo_info "Download migration"
make setup

echo_info "Build docker containers"
make docker/build

echo_info "Start docker containers"
make docker/up

echo_info "Wait 5 seconds for database container to start"
sleep 5

echo_info "Begin migrations"

while ! make db/up
do 
    echo_info "Re-run migration"
done

