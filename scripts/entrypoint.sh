#!/bin/sh
set -e

# Run migrations
go run infra/goose/migrations.go up

# Install reflex to watch for code changes and restart server
go install github.com/cespare/reflex@latest

# Start the app server
reflex go run main.go --start-service

exec "$@"