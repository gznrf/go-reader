package postgres_server

import (
	"log"

	"google.golang.org/grpc"

	postgres_service "github.com/gznrf/go-reader/internal/postgres_service/service"
	postgres_proto "github.com/gznrf/go-reader/pkg/api/proto/go/postgres"
)

type PostgresServer struct {
	postgres_proto.UnimplementedPostgresServer
	services *postgres_service.Service
}

func NewPostgresServer(s *postgres_service.Service) *PostgresServer {
	return &PostgresServer{
		services: s,
	}
}

func Register(gRPC *grpc.Server, postgresServer PostgresServer, addr string) {
	const op = "postgres.server"
	postgres_proto.RegisterPostgresServer(gRPC, &postgresServer)
	log.Printf("server started at %s on port %s", op, addr)
}
