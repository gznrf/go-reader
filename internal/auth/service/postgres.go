package auth_service

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	postgres_proto "github.com/gznrf/go-reader/pkg/api/proto/go/postgres"
	"github.com/gznrf/go-reader/pkg/utils/encryption"
)

type PostgresService struct {
	postgresClient postgres_proto.PostgresClient
}

func (ps *PostgresService) SetPostgresClient(addr string) error {
	const op = "auth.service.postgres.set_postgres_client"
	cc, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error - %s at %s", err, op)
		return err
	}
	ps.postgresClient = postgres_proto.NewPostgresClient(cc)
	return nil
}
func (ps *PostgresService) Register(email, name, password string) (int64, error) {
	const op = "auth.service.postgres.register"

	passwordHash, err := encryption.HashPassword(password)
	if err != nil {
		log.Printf("error - %s at %s", err, op)
		return 0, err
	}

	res, err := ps.postgresClient.CreateUser(context.Background(), &postgres_proto.CreateUserRequest{
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
	})
	if err != nil {
		log.Printf("error - %s at %s", err, op)
		return 0, err
	}

	return res.UserId, nil
}

func (ps *PostgresService) Login(email, password string) (string, error) {
	const op = "auth.service.postgres.Login"

	passwordHash, err := encryption.HashPassword(password)
	if err != nil {
		log.Printf("error - %s at %s", err, op)
		return "", err
	}

	res, err := ps.postgresClient.GetUser(context.Background(), &postgres_proto.GetUserRequest{
		Email:        email,
		PasswordHash: passwordHash,
	})
	if err != nil {
		log.Printf("error - %s at %s", err, op)
		return "", err
	}

	return fmt.Sprint(res.UserId), nil
}
