package main

import (
	"gitlab.com/expected.sh/agent/pkg/function"
	"gitlab.com/expected.sh/agent/pkg/protocol"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	log.SetFlags(log.Ldate|log.Ltime|log.Llongfile)
	lis, err := net.Listen("tcp", ":9809")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	protocol.RegisterFunctionServiceServer(server, &function.Server{})
	server.Serve(lis)
}
