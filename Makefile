build:
	bazel run //:gazelle
	bazel run //cmd/nekomata:nekomata
	bazel run //cmd/nekomatad:nekomatad

protocol:
	protoc -I $(shell pwd)/../pkg/protocol $(shell pwd)/../pkg/protocol/*.proto --go_out=plugins=grpc:$(shell pwd)/../pkg/protocol

clean:
	bazel clean --expunge

.PHONY: build protocol clean
