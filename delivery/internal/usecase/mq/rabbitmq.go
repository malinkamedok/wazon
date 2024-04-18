package mq

import (
	"context"
	"delivery/internal/usecase"
	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	channel *amqp091.Channel
	queue   amqp091.Queue
}

func (r RabbitMQ) SendOrderUpdateMessage(ctx context.Context, orderUUID uuid.UUID, status string) error {
	//TODO implement me
	panic("implement me")
}

var _ usecase.DeliveryMQ = (*RabbitMQ)(nil)

func NewRabbitMQProducer(ch *amqp091.Channel, q amqp091.Queue) *RabbitMQ {
	return &RabbitMQ{channel: ch, queue: q}
}
