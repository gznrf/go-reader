package gateway_handler

import gateway_service "github.com/gznrf/go-reader/internal/gateway/service"

type Handler struct {
	services *gateway_service.Service
}

func NewHandler(services *gateway_service.Service) *Handler {
	return &Handler{services: services}
}
