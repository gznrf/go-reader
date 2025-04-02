package gateway_service

type Authorization interface {
	RegisterUser()
}

type BookService interface {
}

type PostgreService interface {
}

type MongoService interface {
}

type RedisService interface {
}

type Service struct {
	Authorization
	BookService
	PostgreService
	MongoService
	RedisService
}

func NewService() *Service {
	return &Service{
		Authorization:  nil,
		BookService:    nil,
		PostgreService: nil,
		MongoService:   nil,
		RedisService:   nil,
	}
}
