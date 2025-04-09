package main

import (
	"log"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	postgres_server "github.com/gznrf/go-reader/internal/postgres_service/server"
	postgres_service "github.com/gznrf/go-reader/internal/postgres_service/service"
	"github.com/gznrf/go-reader/pkg/utils"
)

func main() {
	// Read config
	if err := utils.InitConfig("postgres_config"); err != nil {
		log.Panicf("error with init gateway config - %s", err)
	}

	serviceAddr := viper.GetString("service_addr")
	l, err := net.Listen("tcp", serviceAddr)
	if err != nil {
		panic("Error with starting auth service")
	}

	// Init service and set postgres client
	as := postgres_service.NewService()

	grpcServer := grpc.NewServer()
	authServer := postgres_server.NewPostgresServer(as)
	postgres_server.Register(grpcServer, *authServer, serviceAddr)
	err = grpcServer.Serve(l)
	if err != nil {
		panic("Error with starting postgres service")
	}
}
