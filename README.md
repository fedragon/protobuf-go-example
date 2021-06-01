# protobuf-go-example

## Setup

- Download the latest release of the Protobuf compiler [here](https://github.com/protocolbuffers/protobuf/releases)
- Uncompress it and add its `bin` directory to your `PATH`
- Run `go install google.golang.org/protobuf/cmd/protoc-gen-go`
- Add `$(go env GOPATH)/bin` to your `PATH`

## Run

    $ make gen-proto build
    $ ./example