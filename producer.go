package rabbitmq

import (
	"context"
	"time"

	"github.com/streadway/amqp"
)

type Producer struct {
	Channel      *amqp.Channel
	ExchangeName string
	Key          string
	ContentType  string
}

func NewProducer(channel *amqp.Channel, exchangeName string, key string, contentType string) (*Producer, error) {
	if len(key) == 0 {
		key = "info"
	}
	if len(contentType) == 0 {
		contentType = "text/plain"
	}
	return &Producer{Channel: channel, ExchangeName: exchangeName, Key: key, ContentType: contentType}, nil
}

func NewProducerByConfig(config ProducerConfig) (*Producer, error) {
	channel, er1 := NewChannel(config.Url)
	if er1 != nil {
		return nil, er1
	}
	er2 := channel.ExchangeDeclare(config.ExchangeName, config.ExchangeKind, true, config.AutoDelete, false, false, nil)
	if er2 != nil {
		return nil, er2
	}
	return NewProducer(channel, config.ExchangeName, config.Key, config.ContentType)
}

func (p *Producer) Produce(ctx context.Context, data []byte, messageAttributes map[string]string) (string, error) {
	opts := MapToTable(messageAttributes)
	msg := amqp.Publishing{
		Headers:      opts,
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  p.ContentType,
		Body:         data,
	}
	err := p.Channel.Publish(p.ExchangeName, p.Key, false, false, msg)
	return "", err
}

func MapToTable(messageAttributes map[string]string) amqp.Table {
	opts := amqp.Table{}
	if messageAttributes != nil {
		for k, v := range messageAttributes {
			opts[k] = v
		}
	}
	return opts
}
