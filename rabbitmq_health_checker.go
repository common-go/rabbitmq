package rabbitmq

import "context"

type RabbitMQHealthChecker struct {
	url  string
	name string
}

func NewRabbitMQHealthChecker(url string, options ...string) *RabbitMQHealthChecker {
	var name string
	if len(options) >= 1 {
		name = options[0]
	} else {
		name = "rabbitmq"
	}
	return &RabbitMQHealthChecker{url, name}
}

func (s *RabbitMQHealthChecker) Name() string {
	return s.name
}

func (s *RabbitMQHealthChecker) Check(ctx context.Context) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	channel, er1 := NewChannel(s.url)
	if er1 != nil {
		return res, er1
	}
	er2 := channel.Close()
	return res, er2
}

func (s *RabbitMQHealthChecker) Build(ctx context.Context, data map[string]interface{}, err error) map[string]interface{} {
	if err == nil {
		return data
	}
	data["error"] = err.Error()
	return data
}
