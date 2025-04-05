package main

import (
	"log"
	"net"

	auth_server "github.com/gznrf/go-reader/internal/auth/server"
	"github.com/gznrf/go-reader/pkg/utils"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	// Read config
	if err := utils.InitConfig("auth_config"); err != nil {
		log.Panicf("error with init gateway config - %s", err)
	}
	gRPCServer := grpc.NewServer()

	auth_server.Register(gRPCServer)

	l, err := net.Listen("tcp", viper.GetString("service_addr"))
	if err != nil {
		panic("Error with starting auth service")
	}
	err = gRPCServer.Serve(l)
	if err != nil {
		panic("Error with starting auth service")
	}
}
