package service

type Interface_pingService interface {
	Pong() (string, error)
}

type struct_pingService struct{}

var pingService = struct_pingService{}

func PingService() Interface_pingService {
	return &pingService
}

func (service struct_pingService) Pong() (string, error) {
	return "pong", nil
}
