package config

import amqp "github.com/rabbitmq/amqp091-go"

type MQConfig struct {
	Channel    *amqp.Channel
	Queue      amqp.Queue
	Connection *amqp.Connection
}
