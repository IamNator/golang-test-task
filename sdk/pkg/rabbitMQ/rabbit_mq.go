package rabbitMQ

import "github.com/streadway/amqp"

type (
	MQ struct {
		Conn *amqp.Connection
	}
)

func NewConnection(url string) (*MQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	return &MQ{Conn: conn}, nil
}

func (mq MQ) Close() error {
	return mq.Conn.Close()
}
