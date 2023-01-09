package worker

import (
	"fmt"
	"twitch_chat_analysis/sdk/pkg/redis"

	"twitch_chat_analysis/sdk/pkg/rabbitMQ"
)

type (
	App struct {
		mq          *rabbitMQ.MQ
		redisClient *redis.Redis
	}
)

func NewApp(mq *rabbitMQ.MQ, redisClient *redis.Redis) *App {
	return &App{mq: mq, redisClient: redisClient}
}

func (w App) SubscribeToQueue(queueName string) error {

	// Open a channel
	ch, err := w.mq.Conn.Channel()
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

	// Connect to Redis

	// Subscribe to messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	// Process messages
	go func() {
		for msg := range msgs {
			// Save the message to Redis
			err := w.redisClient.Client.Set("message_"+msg.AppId, string(msg.Body), 0).Err()
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	return nil
}
