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
	const op = "gateway.handler.handler.initRoutes"
	router := mux.NewRouter()

	//Auth
	auth := router.PathPrefix("/auth").Subrouter()
	{
		log.Println("auth routes inited at " + op)
		auth.HandleFunc("/registration", h.Register).Methods("POST") // Create user
		auth.HandleFunc("/authorization", h.Login).Methods("POST")   //Auth user
	}

	//Books
	books := router.PathPrefix("/books").Subrouter()
	{
		log.Println("books routes inited at " + op)
		books.HandleFunc("/", h.UploadBook).Methods("POST")           //Upload book
		books.HandleFunc("/", h.GetAllBooks).Methods("GET")           //All books
		books.HandleFunc("/{id}", h.GetBookById).Methods("GET")       // Get book by id
		books.HandleFunc("/{id}", h.DeleteBookById).Methods("DELETE") // Delete book by id
	}

	return router
}
