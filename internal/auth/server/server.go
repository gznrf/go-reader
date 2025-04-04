package auth_server

import (
	"context"
	"log"

	auth_proto "github.com/gznrf/go-reader/pkg/api/proto/go/auth"
	"google.golang.org/grpc"
)

type AuthServer struct {
	auth_proto.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	const op = "auth.server"
	auth_proto.RegisterAuthServer(gRPC, &AuthServer{})
	log.Printf("server started at %s", op)
}

func (as *AuthServer) Register(ctx context.Context, req *auth_proto.RegisterRequest) (*auth_proto.RegisterResponse, error) {

	return &auth_proto.RegisterResponse{
		UserId: 200,
	}, nil
}

func (as *AuthServer) Login(ctx context.Context, req *auth_proto.LoginRequest) (*auth_proto.LoginResponse, error) {
	return &auth_proto.LoginResponse{
		Token: "Login toched",
	}, nil
}
