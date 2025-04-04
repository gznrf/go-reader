package gateway_service

import (
	"context"
	"log"

	auth_proto "github.com/gznrf/go-reader/pkg/api/proto/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthService struct {
	authClient auth_proto.AuthClient
}

func NewAuthService(addr string) (*AuthService, error) {
	const op = "gateway.service.auth.new_auth_client"
	cc, err := grpc.NewClient(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &AuthService{
		authClient: auth_proto.NewAuthClient(cc),
	}, nil

}

func (as *AuthService) Register(email, name, password string) (int64, error) {
	const op = "gateway.service.register"
	var userID int64

	res, err := as.authClient.Register(context.Background(), &auth_proto.RegisterRequest{Name: name, Email: email, Password: password})
	if err != nil {
		log.Printf("error - %s at %s", err, op)
		return 0, err
	}
	userID = res.UserId
	return userID, nil
}
func (as *AuthService) Login(email, password string) (string, error) {
	const op = "gateway.service.login"
	var token string

	cc, err := grpc.NewClient(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error - %s at %s", err, op)
		return 0, err
	}

	cli := auth_proto.NewAuthClient(cc)

	res, err := cli.Register(context.Background(), &auth_proto.RegisterRequest{Name: name, Email: email, Password: password})
	if err != nil {
		log.Printf("error - %s at %s", err, op)
		return 0, err
	}
	userID = res.UserId
	return userID, nil

	return op, nil
}
