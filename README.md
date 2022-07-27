# shitlist service

Shitlist is a project service that makes use of several technologies I've wanted to try out.

## Getting Started

### Install dependencies

```bash
$ go install github.com/bufbuild/buf/cmd/buf@latest
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
```

### Add Go install directories to PATH
```bash
$ [ -n "$(go env GOBIN)" ] && export PATH="$(go env GOBIN):${PATH}"
$ [ -n "$(go env GOPATH)" ] && export PATH="$(go env GOPATH)/bin:${PATH}"
```

### Generate code

```bash
$ buf lint
$ buf generate
```

### Add

## Building the server

```bash
$ go build -o ./.bin/shitlist cmd/server/main.go
```

## Configuration

```bash
This application is configured via the environment. The following environment
variables can be used:

KEY                           TYPE       DEFAULT                                        REQUIRED    DESCRIPTION
DATABASE_HOST                 String                                                    true        
DATABASE_PORT                 Integer    5432                                                       
DATABASE_USER                 String                                                    true        
DATABASE_PASSWORD             String                                                    true        
DATABASE_SCHEMA_NAME          String                                                    true        
GITHUB_OAUTH_CLIENT_ID        String                                                    true        
GITHUB_OAUTH_CLIENT_SECRET    String                                                    true        
GITHUB_OAUTH_REDIRECT_URL     String     http://localhost:10000/auth/github/callback                
GOOGLE_OAUTH_CLIENT_ID        String                                                    true        
GOOGLE_OAUTH_CLIENT_SECRET    String                                                    true        
GOOGLE_OAUTH_REDIRECT_URL     String     http://localhost:10000/auth/google/callback                
TEMPLATES_DIRECTORY           String     ../../templates                                            
SERVER_LISTEN_ADDRESS         String     :10000
```

## Running the server

```bash
$ go run cmd/server/main.go
```

