package main

import (
	gateway_handler "github.com/gznrf/go-reader/internal/gateway/handler"
	gateway_service "github.com/gznrf/go-reader/internal/gateway/service"
	"github.com/gznrf/go-reader/pkg/httpserver"
)

func main() {
	//Gateway start
	gws := gateway_service.NewService()
	h := gateway_handler.NewHandler(gws)

	serv := httpserver.NewServer()
	if err := serv.Run(":8080", h.InitRoutes()); err != nil {
		panic("Error with starting server")
	}
}
