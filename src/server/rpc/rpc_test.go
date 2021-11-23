package rpc

import (
	"context"
	"fmt"
	"github.com/chaseisabelle/daq/pkg"
	"github.com/chaseisabelle/daq/src/service"
	"github.com/phayes/freeport"
	"google.golang.org/grpc"
	"sync"
	"testing"
)

func TestRPC(t *testing.T) {
	max := 100 //<< number of messages to queue
	prt, err := freeport.GetFreePort()

	if err != nil {
		t.Error(err)
	}

	adr := fmt.Sprintf("0.0.0.0:%d", prt)

	ser, err := service.New()

	if err != nil {
		t.Error(err)
	}

	srv, err := New(adr, ser)

	if err != nil {
		t.Error(err)
	}

	go func() {
		err := srv.Serve()

		if err != nil {
			t.Error(err)
		}
	}()

	con, err := grpc.Dial(adr, grpc.WithInsecure())

	if err != nil {
		t.Error(err)
	}

	defer func() {
		err := con.Close()

		if err != nil {
			t.Error(err)
		}
	}()

	wg := sync.WaitGroup{}
	cli := daqpb.NewServiceClient(con)

	wg.Add(1)

	go func() {
		defer wg.Done()

		for i := 0; i < max; i++ {
			_, err := cli.Enqueue(context.TODO(), &daqpb.EnqueueRequest{
				Body: fmt.Sprintf("message %d", i),
			})

			if err != nil {
				t.Error(err)
			}
		}
	}()

	wg.Add(1)

	go func() {
		wg.Done()

		for i := 0; i < max; i++ {
			_, err := cli.Dequeue(context.TODO(), &daqpb.DequeueRequest{})

			if err != nil {
				t.Error(err)
			}
		}
	}()

	wg.Wait()
}
