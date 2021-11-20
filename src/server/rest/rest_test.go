package rest

import (
	daqpb "github.com/chaseisabelle/daq/pkg"
	"testing"
)

func TestREST(t *testing.T) {
	cli := daqpb.NewServiceClient()
	srv := New()
}
