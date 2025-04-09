package postgres_service

type Auth interface {
	CreateUser(email, name, password string) (int64, error)
	GetUser(email, password string) (int64, error)
}

type Service struct {
	Auth
}

func NewService() *Service {
	return &Service{
		Auth: nil,
	}
}
