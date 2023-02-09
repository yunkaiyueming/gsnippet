.PHONY: dep test check build run clean

BINARY_NAME="pprof"

MAIN_GO=trac1.go

dep:
	go version

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

check:
	go fmt ./...
	go vet ./...

env:
	go env

build: dep env clean
	go build -o ${BINARY_NAME} ${MAIN_GO}

run:
	./${BINARY_NAME}

clean:
	go clean
	rm -f ${BINARY_NAME}-arm64-darwin
	rm -f ${BINARY_NAME}-amd64-darwin
	rm -f ${BINARY_NAME}-amd64-linux
	rm -f ${BINARY_NAME}-amd64-windows
