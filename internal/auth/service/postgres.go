package auth_service

import (
	"log"

	postgres_proto "github.com/gznrf/go-reader/pkg/api/proto/go/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PostgresService struct {
	postgresClient postgres_proto.PostgresClient
}

func (ps *PostgresService) SetPostgresClient(addr string) error {
	const op = "gateway.service.auth.set_auth_client"
	cc, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error - %s at %s", err, op)
		return err
	}
	ps.postgresClient = postgres_proto.NewPostgresClient(cc)
	return nil
}
func (ps *PostgresService) Register(email, name, password string) (int64, error) {
	return 0, nil
}
func (ps *PostgresService) Login(email, password string) (string, error) {
	return "", nil
}
