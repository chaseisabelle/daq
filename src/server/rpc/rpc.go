package rpc

import (
	daqpb "github.com/chaseisabelle/daq/pkg"
	"github.com/chaseisabelle/daq/src/service"
	"google.golang.org/grpc"
	"net"
)

type RPC struct {
	server   *grpc.Server
	listener net.Listener
}

func New(adr string, ser *service.Service) (*RPC, error) {
	lis, err := net.Listen("tcp", adr)

	if err != nil {
		return nil, err
	}

	srv := grpc.NewServer()

	daqpb.RegisterServiceServer(srv, ser)

	return &RPC{
		server:   srv,
		listener: lis,
	}, nil
}

func (r *RPC) Serve() error {
	return r.server.Serve(r.listener)
}
