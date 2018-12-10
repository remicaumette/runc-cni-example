package function

import (
	"context"
	"gitlab.com/expected.sh/agent/pkg/protocol"
)

type Server struct {}

func (server *Server) Create(ctx context.Context, req *protocol.CreateRequest) (res *protocol.CreateResponse, err error) {
	return &protocol.CreateResponse{Function:nil}, nil
}

func (server *Server) List(ctx context.Context, req *protocol.ListRequest) (res *protocol.ListResponse, err error) {
	return &protocol.ListResponse{Function:nil}, nil
}
