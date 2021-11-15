package rest

import (
	"github.com/chaseisabelle/daq/src/service"
	"net/http"
)

type REST struct {
	server  *http.Server
	service *service.Service
}

func New(adr string, ser *service.Service) (*REST, error) {
	srv := &REST{
		service: ser,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/enqueue", srv.enqueue)
	mux.HandleFunc("/dequeue", srv.dequeue)
	mux.HandleFunc("/requeue", srv.requeue)

	srv.server = &http.Server{
		Addr:    adr,
		Handler: mux,
	}

	return srv, nil
}

func (r *REST) Serve() error {
	return r.server.ListenAndServe()
}

func (r *REST) enqueue(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPut {
		http.Error(res, "method not allowed", http.StatusMethodNotAllowed)

		return
	}

	// convert request/response
}

func (r *REST) dequeue(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "method not allowed", http.StatusMethodNotAllowed)

		return
	}
}

func (r *REST) requeue(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "method not allowed", http.StatusMethodNotAllowed)

		return
	}
}
