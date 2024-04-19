package mq

import (
	"context"
	"delivery/internal/entity"
	"delivery/internal/usecase"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

type RabbitMQ struct {
	channel *amqp091.Channel
	queue   amqp091.Queue
}

var _ usecase.DeliveryMQ = (*RabbitMQ)(nil)

func NewRabbitMQProducer(ch *amqp091.Channel, q amqp091.Queue) *RabbitMQ {
	return &RabbitMQ{channel: ch, queue: q}
}

func (r RabbitMQ) SendOrderUpdateMessage(ctx context.Context, orderUUID uuid.UUID, status string) error {
	var order entity.OrderList
	order.UUID = orderUUID
	order.Status = entity.OrderStatus(status)
	orderJson, err := json.Marshal(order)
	if err != nil {
		log.Println(err)
		return err
	}

	err = r.channel.PublishWithContext(ctx, "", r.queue.Name, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        orderJson,
	})
	if err != nil {
		log.Println("failed to publish to mq: ", err)
		return err
	}

	return nil
}
