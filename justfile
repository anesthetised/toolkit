set shell := ["sh", "-c"]

default:
    @just --list

[no-cd]
test:
    go test $(go list ./...) -v

[no-cd]
lint:
    golangci-lint run

[working-directory: 'net/http/gen/example']
build-genhttp-example:
    go build -trimpath

[no-cd]
update-deps:
    go get -u ./... && go mod tidy
