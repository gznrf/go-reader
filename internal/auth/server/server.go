package auth_server

import (
	"log"

	"google.golang.org/grpc"

	auth_service "github.com/gznrf/go-reader/internal/auth/service"
	auth_proto "github.com/gznrf/go-reader/pkg/api/proto/go/auth"
)

type AuthServer struct {
	auth_proto.UnimplementedAuthServer
	services *auth_service.Service
}

func NewAuthServer(s *auth_service.Service) *AuthServer {
	return &AuthServer{
		services: s,
	}
}

func Register(gRPC *grpc.Server, authServer AuthServer, addr string) {
	const op = "auth.server"
	auth_proto.RegisterAuthServer(gRPC, &authServer)
	log.Printf("server started at %s on port %s", op, addr)
}
