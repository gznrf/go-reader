package gateway_service

import "io"

type Authorization interface {
	SetAuthClient(addr string) error
	Register(email, name, password string) (int64, error)
	Login(email, password string) (string, error)
}

type BookService interface {
	SetBookClient(addr string) error
	UploadBook(file io.Reader) (int64, error)
	GetAllBooks(id int64) error
	GetBookById(id int64) error
	DeleteBookById(id int64) error
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
