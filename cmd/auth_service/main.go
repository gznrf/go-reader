package main

import (
	"log"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	auth_server "github.com/gznrf/go-reader/internal/auth/server"
	auth_service "github.com/gznrf/go-reader/internal/auth/service"
	"github.com/gznrf/go-reader/pkg/utils"
)

func main() {
	// Read config
	if err := utils.InitConfig("auth_config"); err != nil {
		log.Panicf("error with init gateway config - %s", err)
	}

	serviceAddr := viper.GetString("service_addr")
	l, err := net.Listen("tcp", serviceAddr)
	if err != nil {
		panic("Error with starting auth service")
	}

	//Init service and set postgres client
	as := auth_service.NewService()
	as.SetPostgresClient(viper.GetString("service_postgres_client_addr"))

	grpcServer := grpc.NewServer()
	authServer := auth_server.NewAuthServer(as)
	auth_server.Register(grpcServer, *authServer, serviceAddr)
	err = grpcServer.Serve(l)
	if err != nil {
		panic("Error with starting auth service")
	}
}
