package httpserver

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(port string, handler http.Handler) error {
	const op = "pkg.httpserver.run"
	s.httpServer = &http.Server{
		Addr:           port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	log.Println("server start at " + s.httpServer.Addr + " " + op)
	return s.httpServer.ListenAndServe()
}
