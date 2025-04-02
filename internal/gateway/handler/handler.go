package gateway_handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	gateway_service "github.com/gznrf/go-reader/internal/gateway/service"
)

type Handler struct {
	services *gateway_service.Service
}

func NewHandler(services *gateway_service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	const op = "internal.gateway.handler.handler"
	router := mux.NewRouter()

	//Auth
	auth := router.PathPrefix("/auth").Subrouter()
	{
		log.Println("auth routes inited at " + op)
		auth.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) { fmt.Println("Registration echo") }).Methods("POST")
		auth.HandleFunc("/authorization", func(w http.ResponseWriter, r *http.Request) { fmt.Println("Authorization echo") }).Methods("POST")
	}

	//Books

	return router
}
