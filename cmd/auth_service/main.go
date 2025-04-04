package main

import (
	"net"

	auth_server "github.com/gznrf/go-reader/internal/auth/server"
	"google.golang.org/grpc"
)

func main() {
	gRPCServer := grpc.NewServer()

	auth_server.Register(gRPCServer)

	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic("Error with starting auth service")
	}
	gRPCServer.Serve(l)
}
