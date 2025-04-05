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

func (as *AuthService) SetAuthClient(addr string) error {
	const op = "gateway.service.auth.set_auth_client"
	cc, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error - %s at %s", err, op)
		return err
	}
	as.authClient = auth_proto.NewAuthClient(cc)
	return nil
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

	res, err := as.authClient.Login(context.Background(), &auth_proto.LoginRequest{Email: email, Password: password})
	if err != nil {
		log.Printf("error - %s at %s", err, op)
		return "", err
	}
	token = res.Token
	return token, nil
}
