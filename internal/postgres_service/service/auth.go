package postgres_service

import "log"

type AuthService struct {
}

func (as *AuthService) CreateUser(email, name, password string) (int64, error) {
	const op = "postgres.service.auth.createUser"
	log.Printf("User created at %s", op)
	return 0, nil
}
func (as *AuthService) GetUser(email, password string) (int64, error) {
	const op = "postgres.service.auth.getUser"
	log.Printf("User got at %s", op)
	return 0, nil
}
