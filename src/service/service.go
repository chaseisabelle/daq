package service

import (
	"context"
	"github.com/chaseisabelle/daq/pkg"
	"github.com/chaseisabelle/daq/src/message"
	"github.com/chaseisabelle/daq/src/queue"
)

type Service struct {
	queue *queue.Queue
}

func New() (*Service, error) {
	q, err := queue.New()

	return &Service{
		queue: q,
	}, err
}

func (s *Service) Enqueue(ctx context.Context, req *daqpb.EnqueueRequest) (*daqpb.EnqueueResponse, error) {
	msg, err := message.New(req.Body)

	if err != nil {
		return nil, err
	}

	return &daqpb.EnqueueResponse{}, s.queue.Enqueue(msg)
}

func (s *Service) Dequeue(ctx context.Context, req *daqpb.DequeueRequest) (*daqpb.DequeueResponse, error) {
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

func (s *Service) Requeue(ctx context.Context, req *daqpb.RequeueRequest) (*daqpb.RequeueResponse, error) {
	msg, err := message.New(req.Body)

	if err != nil {
		return nil, err
	}

	return &daqpb.RequeueResponse{}, s.queue.Requeue(msg)
}
