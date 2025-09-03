#!/usr/bin/env bash

sqlc generate
infisical run --env=dev -- go run ./cmd/server/main.go
