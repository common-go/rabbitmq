package rabbitmq

import "context"

type RabbitMQHealthService struct {
	url  string
	name string
}

func NewDefaultRabbitMQHealthService(url string) *RabbitMQHealthService {
	return NewRabbitMQHealthService(url, "rabbitmq")
}

func NewRabbitMQHealthService(url string, name string) *RabbitMQHealthService {
	return &RabbitMQHealthService{url, name}
}

func (s *RabbitMQHealthService) Name() string {
	return s.name
}

func (s *RabbitMQHealthService) Check(ctx context.Context) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	channel, er1 := NewChannel(s.url)
	if er1 != nil {
		return res, er1
	}
	er2 := channel.Close()
	return res, er2
}

func (s *RabbitMQHealthService) Build(ctx context.Context, data map[string]interface{}, err error) map[string]interface{} {
	if err == nil {
		return data
	}
	data["error"] = err.Error()
	return data
}
