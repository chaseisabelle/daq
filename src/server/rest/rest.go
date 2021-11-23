package rest

import (
	"context"
	"encoding/json"
	"github.com/chaseisabelle/daq/pkg"
	"github.com/chaseisabelle/daq/src/service"
	"log"
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
	if req.Method != http.MethodPost {
		http.Error(res, "method not allowed", http.StatusMethodNotAllowed)

		return
	}

	enr := &daqpb.EnqueueRequest{}
	err := json.NewDecoder(req.Body).Decode(enr)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

	eqr, err := r.service.Enqueue(context.TODO(), enr)

	res.WriteHeader(http.StatusCreated)

	_, err = res.Write([]byte(eqr.String()))

	if err != nil {
		log.Println(err)
	}
}

func (r *REST) dequeue(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "method not allowed", http.StatusMethodNotAllowed)

		return
	}

	der := &daqpb.DequeueRequest{}
	err := json.NewDecoder(req.Body).Decode(der)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

	dqr, err := r.service.Dequeue(context.TODO(), der)

	res.WriteHeader(http.StatusCreated)

	_, err = res.Write([]byte(dqr.String()))

	if err != nil {
		log.Println(err)
	}
}

func (r *REST) requeue(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "method not allowed", http.StatusMethodNotAllowed)

		return
	}

	rer := &daqpb.RequeueRequest{}
	err := json.NewDecoder(req.Body).Decode(rer)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)

		return
	}

	rqr, err := r.service.Requeue(context.TODO(), rer)

	res.WriteHeader(http.StatusCreated)

	_, err = res.Write([]byte(rqr.String()))

	if err != nil {
		log.Println(err)
	}
}
