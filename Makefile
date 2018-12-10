protocol:
	protoc -I $(shell pwd)/pkg/protocol $(shell pwd)/pkg/protocol/*.proto --go_out=plugins=grpc:$(shell pwd)/pkg/protocol

.PHONY: build protocol clean
