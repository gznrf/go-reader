package postgres_server

import (
	"context"
	"errors"
	"log"

	postgres_proto "github.com/gznrf/go-reader/pkg/api/proto/go/postgres"
)

const zeroValue = 0

func (ps *PostgresServer) CreateUser(ctx context.Context, req *postgres_proto.CreateUserRequest) (*postgres_proto.CreateUserResponse, error) {
	const op = "postgres.server.auth.createUser"

	userId, err := ps.services.Auth.CreateUser(req.Email, req.Name, req.PasswordHash)
	if err != nil {
		log.Printf("error - %s at %s", err, op)
		return nil, err
	}

	return &postgres_proto.CreateUserResponse{
		UserId: userId,
	}, nil
}
func (ps *PostgresServer) GetUser(ctx context.Context, req *postgres_proto.GetUserRequest) (*postgres_proto.GetUserResponse, error) {
	const op = "postgres.server.auth.getUser"

	userId, err := ps.services.Auth.GetUser(req.Email, req.PasswordHash)
	if err != nil {
		log.Printf("error - %s at %s", err, op)
		return nil, err
	}

	// if userId didnt returned it means that he does not exist
	if userId == zeroValue {
		log.Println("user does not exist or incorrect password")
		return nil, errors.New("user does not exist or incorrect password")
	}

	// else auth him
	return &postgres_proto.GetUserResponse{
		UserId: userId,
	}, nil
}
