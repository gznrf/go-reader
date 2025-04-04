package gateway_handler

import (
	"log"

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
		auth.HandleFunc("/registration", h.Register).Methods("POST")
		auth.HandleFunc("/authorization", h.Login).Methods("POST")
	}

	//Books

	return router
}
