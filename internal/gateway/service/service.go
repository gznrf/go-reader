package gateway_service

type Authorization interface {
	Register(email, name, password string) (int64, error)
	Login(email, password string) (string, error)
}

type BookService interface {
}

type Service struct {
	Authorization
	BookService
}

func NewService() *Service {
	return &Service{
		Authorization: &AuthService{},
		BookService:   nil,
	}
}
