all: protocol nekomata nekomata-server

nekomata:
	go build -o bin/nekomata cmd/nekomata/main.go

nekomata-server: bin/nekomata-server
	go build -o bin/nekomata-server cmd/nekomata-server/main.go

protocol:
	env GO111MODULE=off go get -u -v github.com/golang/protobuf/proto
	env GO111MODULE=off go get -u -v github.com/golang/protobuf/protoc-gen-go
	protoc -I $(shell pwd)/pkg/protocol $(shell pwd)/pkg/protocol/*.proto --go_out=plugins=grpc:$(shell pwd)/pkg/protocol

.PHONY: all nekomata nekomata-server protocol
