package auth_service

type Postgres interface {
	SetPostgresClient(addr string) error
	Register(email, name, password string) (int64, error)
	Login(email, password string) (string, error)
}

type Mongo interface {
}

type Redis interface {
}

type Service struct {
	Postgres
	Mongo
	Redis
}

func NewService() *Service {
	return &Service{
		Postgres: &PostgresService{},
		Mongo:    nil,
		Redis:    nil,
	}
}
