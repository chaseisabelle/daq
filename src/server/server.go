package server

import (
	"context"
	"github.com/chaseisabelle/daq/pkg"
	"github.com/chaseisabelle/daq/src/message"
	"github.com/chaseisabelle/daq/src/queue"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	server   *grpc.Server
	listener *net.Listener
	queue    *queue.Queue
}

func New(adr string) (*Server, error) {
	lis, err := net.Listen("tcp", adr)

	if err != nil {
		return nil, err
	}

	srv := grpc.NewServer()
	q, err := queue.New()

	if err != nil {
		return nil, err
	}

	ins := &Server{
		server:   srv,
		listener: &lis,
		queue:    q,
	}

	daqpb.RegisterServiceServer(srv, ins)

	return ins, nil
}

func (s *Server) Serve() error {
	return s.server.Serve(*s.listener)
}

func (s *Server) Enqueue(ctx context.Context, req *daqpb.EnqueueRequest) (*daqpb.EnqueueResponse, error) {
	msg, err := message.New(req.Body)

	if err != nil {
		return nil, err
	}

	return &daqpb.EnqueueResponse{}, s.queue.Enqueue(msg)
}

func (s *Server) Dequeue(ctx context.Context, req *daqpb.DequeueRequest) (*daqpb.DequeueResponse, error) {
	msg, err := s.queue.Dequeue()

	if err != nil {
		return nil, err
	}

	res := &daqpb.DequeueResponse{
		Body: "",
		Type: "OK",
	}

	if msg != nil {
		res.Body = msg.Body
	} else {
		res.Type = "EMPTY"
	}

	return res, nil
}

func (s *Server) Requeue(ctx context.Context, req *daqpb.RequeueRequest) (*daqpb.RequeueResponse, error) {
	msg, err := message.New(req.Body)

	if err != nil {
		return nil, err
	}

	return &daqpb.RequeueResponse{}, s.queue.Requeue(msg)
}
