package mqProducer

import (
	"delivery/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func Connect() (config.MQConfig, error) {
	var mqConfig config.MQConfig

	conn, err := amqp.Dial("amqp://rmuser:rmpassword@rabbitmq:5672/")
	if err != nil {
		for i := 2; i < 10; i++ {
			time.Sleep(time.Duration(2) * time.Second)
			log.Printf("connection to RabbitMQ failed. try %d of 10", i)
			conn, err = amqp.Dial("amqp://rmuser:rmpassword@rabbitmq:5672/")
			if err == nil {
				break
			}
			if i == 9 {
				log.Panicf("Failed to connect to RabbitMQ: %s", err)
				return mqConfig, err
			}
		}
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("Failed to open a channel: %s", err)
		return mqConfig, err
	}

	q, err := ch.QueueDeclare(
		"deliveryUpdate", // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		log.Panicf("Failed to declare a queue: %s", err)
		return mqConfig, err
	}

	mqConfig.Connection = conn
	mqConfig.Channel = ch
	mqConfig.Queue = q
	return mqConfig, nil
}
