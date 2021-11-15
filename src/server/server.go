package server

import (
	"errors"
	"github.com/chaseisabelle/daq/src/server/rest"
	"github.com/chaseisabelle/daq/src/server/rpc"
	"github.com/chaseisabelle/daq/src/service"
)

type Server struct {
	service *service.Service
	rpc     *rpc.RPC
	rest    *rest.REST
}

func New(ser *service.Service) (*Server, error) {
	return &Server{
		service: ser,
	}, nil
}

func (s *Server) RPC(adr string) (*rpc.RPC, error) {
	rpc, err := rpc.New(adr, s.service)

	if err != nil {
		return nil, err
	}

	s.rpc = rpc

	return rpc, nil
}

func (s *Server) REST(adr string) (*rest.REST, error) {
	rst, err := rest.New(adr, s.service)

	if err != nil {
		return nil, err
	}

	s.rest = rst

	return rst, nil
}

func (s *Server) Serve() error {
	if s.rpc == nil && s.rest == nil {
		return errors.New("no server implemented")
	}

	ech := make(chan error)

	if s.rpc != nil {
		err := s.rpc.Serve()

		if err != nil {
			return err
		}
	}

	if s.rest != nil {
		err := s.rest.Serve()

		if err != nil {
			return err
		}
	}


}
