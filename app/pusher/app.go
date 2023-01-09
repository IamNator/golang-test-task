package pusher

import (
	"github.com/streadway/amqp"
	"twitch_chat_analysis/sdk/pkg/rabbitMQ"
)

type (
	App struct {
		mq *rabbitMQ.MQ
	}
)

func NewApp(mq *rabbitMQ.MQ) *App {
	return &App{mq: mq}
}

func (p App) PushToQueue(queueName string, message []byte) error {
	// Connect to the RabbitMQ server

	// Open a channel
	ch, err := p.mq.Conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Declare the queue
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	// Publish the message
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		return err
	}

	return nil
}
