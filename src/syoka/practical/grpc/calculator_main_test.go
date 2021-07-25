package grpc

import (
	"google.golang.org/grpc"
	"net"
	"testing"
)

func TestCalculatorGrpc(t *testing.T) {

	server, err := net.Listen("tcp", ":9999")
	if err == nil {
		cal := Calculator{}
		grpcServer := grpc.NewServer()
		RegisterCalculatorServer(grpcServer, &cal)

		grpcServer.Serve(server)
	}
}
