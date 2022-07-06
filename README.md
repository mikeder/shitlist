# shitlist service

Shitlist is a project service that makes use of several technologies I've wanted to try out.

## Getting Started

* Install dependencies

```bash
$ go install github.com/bufbuild/buf/cmd/buf@latest
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
```

* Generate code

```bash
$ buf lint
$ buf generate
```

## Building the server

```bash
$ go build -o ./.bin/shitlist cmd/server/main.go
```

## Running the server

```bash
$ go run cmd/server/main.go
```