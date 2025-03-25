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

[no-cd]
update-deps:
  go get -u ./... && go mod tidy
