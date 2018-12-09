package main

import (
	"context"
	"gitlab.com/expected.sh/agent/pkg/protocol"
	"google.golang.org/grpc"
	"log"
	"net"
)

type functionServer struct {}

func (server *functionServer) Create(ctx context.Context, req *protocol.CreateRequest) (res *protocol.CreateResponse, err error) {
	return &protocol.CreateResponse{Function:nil}, nil
}

func (server *functionServer) List(ctx context.Context, req *protocol.ListRequest) (res *protocol.ListResponse, err error) {
	return &protocol.ListResponse{Function:nil}, nil
}

func main() {
	log.SetFlags(log.Ldate|log.Ltime|log.Llongfile)
	lis, err := net.Listen("tcp", ":9809")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	protocol.RegisterFunctionServiceServer(server, &functionServer{})
	server.Serve(lis)
}
