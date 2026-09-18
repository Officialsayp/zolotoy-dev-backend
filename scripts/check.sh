#!/usr/bin/env bash

set -euo pipefail

cd "$(git rev-parse --show-toplevel)/services/order-service"

echo "==> Formatting"
unformatted="$(gofmt -l .)"

if [ -n "$unformatted" ]; then
    echo "Unformatted files:"
    echo "$unformatted"
    exit 1
fi

echo "==> Vet"
go vet ./...

echo "==> Build"
go build ./...

echo "==> Test"
go test ./...

echo
echo "All checks passed."
