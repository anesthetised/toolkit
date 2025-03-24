set shell := ["bash", "-c"]

default:
  just --list

test:
  go test $(go list ./...) -v

lint:
  golangci-lint run

[working-directory: 'net/http/gen/example']
build-genhttp-example:
    go build
