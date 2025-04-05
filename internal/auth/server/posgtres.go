package auth_server

import (
	"context"

	auth_proto "github.com/gznrf/go-reader/pkg/api/proto/go/auth"
)

func (as *AuthServer) Register(ctx context.Context, req *auth_proto.RegisterRequest) (*auth_proto.RegisterResponse, error) {
	var userID int64

	userID, err := as.services.Postgres.Register(req.Email, req.Name, req.Password)
	if err != nil {
		return nil, err
	}

	return &auth_proto.RegisterResponse{
		UserId: userID,
	}, nil
}

func (as *AuthServer) Login(ctx context.Context, req *auth_proto.LoginRequest) (*auth_proto.LoginResponse, error) {
	var token string

	token, err := as.services.Postgres.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &auth_proto.LoginResponse{
		Token: token,
	}, nil
}
