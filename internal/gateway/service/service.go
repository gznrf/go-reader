package gateway_service

type Authorization interface {
}

type BookService interface {
}

type Service struct {
	Authorization
	BookService
}

func NewService() *Service {
	return &Service{
		Authorization: nil,
		BookService:   nil,
	}
}
