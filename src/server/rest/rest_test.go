package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	daqpb "github.com/chaseisabelle/daq/pkg"
	"github.com/chaseisabelle/daq/src/service"
	"github.com/phayes/freeport"
	"net/http"
	"sync"
	"testing"
)

func TestREST(t *testing.T) {
	max := 100 //<< number of messages to queue
	ser, err := service.New()
	ehm := make(map[string]bool)
	rhm := make(map[string]bool)

	if err != nil {
		t.Error(err)
	}

	prt, err := freeport.GetFreePort()

	if err != nil {
		t.Error(err)
	}

	srv, err := New(fmt.Sprintf("0.0.0.0:%d", prt), ser)

	if err != nil {
		t.Error(err)
	}

	go func() {
		err := srv.Serve()

		if err != nil {
			t.Error(err)
		}
	}()

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()

		url := fmt.Sprintf("http://%s/enqueue", srv.server.Addr)

		for i := 0; i < max; i++ {
			msg := fmt.Sprintf("enqueue %d", i)
			_, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(msg)))

			if err != nil {
				t.Error(err)
			} else {
				ehm[msg] = false
			}
		}
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()

		url := fmt.Sprintf("http://%s/dequeue", srv.server.Addr)

		for i := 0; i < max; i++ {
			res, err := http.Post(url, "application/json", bytes.NewBuffer([]byte("")))

			if err != nil {
				t.Error(err)
			}

			raw := daqpb.DequeueResponse{}
			err = json.NewDecoder(res.Body).Decode(&raw)

			if err != nil {
				t.Error(err)

				continue
			}

			switch raw.Type {
			case "OK":
				_, ok := raw.Body
			case "EMPTY":

			default:
				t.Errorf("unexpected type %+v", raw.Type)
			}
		}
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()

		url := fmt.Sprintf("http://%s/requeue", srv.server.Addr)

		for i := 0; i < max; i++ {
			msg := fmt.Sprintf("requeue %d", i)
			_, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(msg)))

			if err != nil {
				t.Error(err)
			} else {
				rhm[msg] = false
			}
		}
	}()

	wg.Wait()
}
