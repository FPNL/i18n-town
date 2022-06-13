package service

type IPingService interface {
	Pong() (string, error)
}
type pingService struct{}

var singlePing = pingService{}

func Ping() IPingService {
	return &singlePing
}

func (service pingService) Pong() (string, error) {
	return "pong", nil
}
