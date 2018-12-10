BAZEL_ROOT=$(shell pwd)/bazel-tmp

build:
	bazel --output_user_root=$(BAZEL_ROOT) run //:gazelle
	bazel --output_user_root=$(BAZEL_ROOT) run //cmd/nekomata:nekomata
	bazel --output_user_root=$(BAZEL_ROOT) run //cmd/nekomatad:nekomatad

protocol:
	protoc -I $(shell pwd)/../pkg/protocol $(shell pwd)/../pkg/protocol/*.proto --go_out=plugins=grpc:$(shell pwd)/../pkg/protocol

clean:
	bazel --output_user_root=$(BAZEL_ROOT) clean --expunge

.PHONY: build protocol clean
