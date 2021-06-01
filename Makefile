.DEFAULT_GOAL := default
.PHONY: default

default: gen-proto build

gen-proto:
	@rm -rf semver/
	@rm -rf types/
	@protoc --go_out=paths=source_relative:. -I./api types/v1/messages.proto semver/v1/messages.proto

build: gen-proto
	@go build -o example main.go
