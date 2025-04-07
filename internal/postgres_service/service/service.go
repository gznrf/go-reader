package postgres_service

type Auth interface {
	Register(email, name, password string) (int64, error)
	Login(email, password string) (string, error)
}

type Service struct {
	Auth
}

func NewService() *Service {
	return &Service{
		Auth: nil,
	}
}
