package auth_server

import (
	"log"

	auth_service "github.com/gznrf/go-reader/internal/auth/service"
	auth_proto "github.com/gznrf/go-reader/pkg/api/proto/go/auth"
	"google.golang.org/grpc"
)

type AuthServer struct {
	Services *auth_service.Service
	auth_proto.UnimplementedAuthServer
}

func NewAuthServer() AuthServer {
	return AuthServer{}
}

func Register(gRPC *grpc.Server) {
	const op = "auth.server"
	auth_proto.RegisterAuthServer(gRPC, &AuthServer{})
	log.Printf("server started at %s", op)
}
