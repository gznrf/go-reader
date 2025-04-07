package postgres_service

import (
	"context"

	postgres_proto "github.com/gznrf/go-reader/pkg/api/proto/go/postgres"
)

type AuthService struct {
}

func (as *AuthService) CreateUser(context.Context, *postgres_proto.CreateUserRequest) (*postgres_proto.CreateUserResponse, error) {
	return nil, nil
}
func (as *AuthService) GetUser(context.Context, *postgres_proto.GetUserRequest) (*postgres_proto.GetUserResponse, error) {
	return nil, nil
}
