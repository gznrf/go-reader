package main

import (
	"log"
	"net"

	auth_server "github.com/gznrf/go-reader/internal/auth/server"
	auth_service "github.com/gznrf/go-reader/internal/auth/service"
	"github.com/gznrf/go-reader/pkg/utils"
	"github.com/spf13/viper"
)

func main() {
	// Read config
	if err := utils.InitConfig("auth_config"); err != nil {
		log.Panicf("error with init gateway config - %s", err)
	}

	l, err := net.Listen("tcp", viper.GetString("service_addr"))
	if err != nil {
		panic("Error with starting auth service")
	}

	as := auth_service.NewService()
	as.SetPostgresClient(viper.GetString("service_postgres_client_addr"))

	gRPCServer := auth_server.NewAuthServer()
	auth_server.Register(gRPCServer)
	err = gRPCServer.Serve(l)
	if err != nil {
		panic("Error with starting auth service")
	}
}
