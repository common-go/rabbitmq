package rabbitmq

type ConsumerConfig struct {
	Url          string `mapstructure:"url"`
	ExchangeName string `mapstructure:"exchange_name"`
	ExchangeKind string `mapstructure:"exchange_kind"`
	QueueName    string `mapstructure:"queue_name"`
	AutoDelete   bool   `mapstructure:"auto_delete"`
}
