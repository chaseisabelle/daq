package main

import (
	"flag"
	"github.com/chaseisabelle/daq/src/server"
	"log"
)

func main() {
	adr := flag.String("address", ":6969", "the address to listen on")

	flag.Parse()

	srv, err := server.New(*adr)

	if err != nil {
		log.Fatalln(err)
	}

	err = srv.Serve()

	if err != nil {
		log.Fatalln(err)
	}
}
