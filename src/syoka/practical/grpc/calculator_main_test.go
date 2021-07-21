package grpc

import (
	"google.golang.org/grpc"
	"net"
	"testing"
)

func TestCalculatorGrpc(t *testing.T) {

	server, err := net.Listen("tcp", ":9999")
	if err == nil {
		calculator := GoCalculator{}
		gNewServer := grpc.NewServer()
		RegisterCalculatorServer(gNewServer, &calculator)

		grpcServer := grpc.NewServer()
		grpcServer.Serve(server)
	}
}
