package server

import (
	"context"
	"gitlab.com/expected.sh/agent/pkg/protocol"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	Storage    string
	Runtime    string
	CNIPlugins string
	CNIConfig  string
}

func (server *Server) Listen(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	protocol.RegisterFunctionServiceServer(s, server)
	return s.Serve(listener)
}

func (server *Server) Create(ctx context.Context, req *protocol.CreateRequest) (*protocol.CreateResponse, error) {
	container, err := server.CreateContainer(req.Name)
	if err != nil {
		return nil, err
	}
	return &protocol.CreateResponse{
		Function: container.ToFunction(),
	}, nil
}

func (server *Server) List(ctx context.Context, req *protocol.ListRequest) (res *protocol.ListResponse, err error) {
	panic("implement me")
}
