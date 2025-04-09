package main

import (
	"log"

	"github.com/spf13/viper"

	gateway_handler "github.com/gznrf/go-reader/internal/gateway/handler"
	gateway_service "github.com/gznrf/go-reader/internal/gateway/service"
	"github.com/gznrf/go-reader/pkg/httpserver"
	"github.com/gznrf/go-reader/pkg/utils"
)

func main() {
	// Read config
	if err := utils.InitConfig("gateway_config"); err != nil {
		log.Panicf("error with init gateway config - %s", err)
	}
	//Gateway start
	gws := gateway_service.NewService()
	err := gws.Authorization.SetAuthClient(viper.GetString("service_auth_client_addr"))
	if err != nil {
		log.Printf("error - %s with setting auth client in gateway", err)
	}
	h := gateway_handler.NewHandler(gws)
	serv := httpserver.NewServer()
	if err := serv.Run(viper.GetString("service_addr"), h.InitRoutes()); err != nil {
		panic("Error with starting gateway service")
	}
}
