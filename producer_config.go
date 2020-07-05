package rabbitmq

type ProducerConfig struct {
	Url          string `mapstructure:"url"`
	ExchangeName string `mapstructure:"exchange_name"`
	ExchangeKind string `mapstructure:"exchange_kind"`
	Key          string `mapstructure:"key"`
	AutoDelete   bool   `mapstructure:"auto_delete"`
	ContentType  string `mapstructure:"content_type"`
}
