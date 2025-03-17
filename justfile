set shell := ["bash", "-c"]

default:
  just --list

test:
  go test $(go list ./...) -v

[working-directory: 'net/http/gen/example']
build-genhttp-example:
    go build
