package main

import (
	"flag"
	"fmt"
	"github.com/chaseisabelle/daq/src/server"
	"github.com/chaseisabelle/daq/src/server/rest"
	"github.com/chaseisabelle/daq/src/server/rpc"
	"github.com/chaseisabelle/daq/src/service"
	"log"
	"strings"
)

var daq server.Server

func main() {
	api := flag.String("api", "grpc", "the api to use (grpc or rest)")
	adr := flag.String("address", ":6969", "the address to listen on")

	flag.Parse()

	srv, err := service.New()

	if err != nil {
		log.Fatalln(err)
	}

	switch strings.ToLower(*api) {
	case "rest":
		daq, err = rest.New(*adr, srv)
	case "grpc":
		daq, err = rpc.New(*adr, srv)
	default:
		err = fmt.Errorf("invalid api: %s", *api)
	}

	if err != nil {
		log.Fatalln(err)
	}

	err = daq.Serve()

	if err != nil {
		log.Fatalln(err)
	}
}
